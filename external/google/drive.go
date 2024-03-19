package google

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/helper"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type ResponseUploadFile struct {
	FileId  string
	FileUrl string
}

type GoogleCredentials struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `yaml:"client_x509_cert_url"`
	UniverseDomain          string `yaml:"universe_domain"`
}

func NewGoogleDriveServiceAccount() ([]byte, error) {
	credentials := GoogleCredentials{
		Type:                    config.Cfg.App.External.Google.Drive.Type,
		ProjectId:               config.Cfg.App.External.Google.Drive.ProjectId,
		PrivateKeyId:            config.Cfg.App.External.Google.Drive.PrivateKeyId,
		PrivateKey:              config.Cfg.App.External.Google.Drive.PrivateKey,
		ClientEmail:             config.Cfg.App.External.Google.Drive.ClientEmail,
		ClientId:                config.Cfg.App.External.Google.Drive.ClientId,
		AuthUri:                 config.Cfg.App.External.Google.Drive.AuthUri,
		TokenUri:                config.Cfg.App.External.Google.Drive.TokenUri,
		AuthProviderX509CertUrl: config.Cfg.App.External.Google.Drive.AuthProviderX509CertUrl,
		ClientX509CertUrl:       config.Cfg.App.External.Google.Drive.ClientX509CertUrl,
		UniverseDomain:          config.Cfg.App.External.Google.Drive.UniverseDomain,
	}

	googleCredentialsJson, err := json.Marshal(credentials)

	if err != nil {
		return nil, fmt.Errorf("Google Drive credentials error: %s", err)
	}
	return googleCredentialsJson, nil
}

func ConnectServiceGoogleDrive() (*drive.Service, error) {
	gDriveAccount, err := NewGoogleDriveServiceAccount()
	if err != nil {
		panic(err)
	}

	service, err := drive.NewService(context.Background(), option.WithCredentialsJSON(gDriveAccount), option.WithScopes(
		drive.DriveScope,
		drive.DriveFileScope,
		drive.DriveReadonlyScope,
	))

	if err != nil {
		log.Printf("Error: Unable to create drive Client %v", err)
		return nil, err
	}

	return service, nil
}

type GoogleDrive struct {
	svc            *drive.Service
	ImageExtention map[string]string
	MaxSize        int
	FileExtention  map[string]string
}

func NewGoogleDriveService(svc *drive.Service) *GoogleDrive {
	return &GoogleDrive{
		svc:            svc,
		ImageExtention: map[string]string{"png": "png", "jpg": "jpg", "jpeg": "jpeg"},
		FileExtention:  map[string]string{"pdf": "key"},
		MaxSize:        1_000_000,
	}
}

func (g GoogleDrive) GenerateLink(fileId string) string {
	return fmt.Sprintf("https://drive.google.com/uc?id=%v", fileId)
}

func (g GoogleDrive) GenerateNewName(PublicUserId uuid.UUID, filename string) string {
	return fmt.Sprintf("%s-u-%d-%s", PublicUserId, uint64(time.Now().UnixNano()), filename)
}

func (g GoogleDrive) IsExtentionAvailable(filename string) bool {
	extention := helper.GetFileExtention(filename)
	_, ok := g.ImageExtention[extention]
	return ok
}

func (g GoogleDrive) IsImageSizeAvailable(size int) bool {
	return size < g.MaxSize
}

func (g GoogleDrive) UpdateFileById(ctx context.Context, fileIdGdrive string, file *multipart.FileHeader) (err error) {
	newFile, err := file.Open()

	defer newFile.Close()

	// Check Extention
	if !g.IsExtentionAvailable(file.Filename) {
		err = response.ErrImageTypeNotCompatible
		return
	}

	// Check Size
	if !g.IsImageSizeAvailable(int(file.Size)) {
		err = response.ErrImageOversize
		return
	}

	// Create File metadata
	f := &drive.File{
	}

	// Create and upload the file
	_, err = g.svc.Files.
		Update(fileIdGdrive, f).
		Media(newFile).
		Context(ctx). // file, fileInf.Size(), baseMimeType).
		ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).
		Do()

	if err != nil {
		return
	}

	return
}

func (g GoogleDrive) UploadFile(ctx context.Context, PublicUserId uuid.UUID, file *multipart.FileHeader) (newRes ResponseUploadFile, err error) {
	newFile, err := file.Open()

	defer newFile.Close()

	// Check Extention
	if !g.IsExtentionAvailable(file.Filename) {
		err = response.ErrImageTypeNotCompatible
		return
	}

	// Check Size
	if !g.IsImageSizeAvailable(int(file.Size)) {
		err = response.ErrImageOversize
		return
	}

	// Create File metadata
	f := &drive.File{
		Name:    g.GenerateNewName(PublicUserId, file.Filename),
		Parents: []string{"1JG1-i9jswFvugmsFa0YDegBbPCrAWT_-"},
	}

	// Create and upload the file
	res, err := g.svc.Files.
		Create(f).
		Media(newFile).
		Context(ctx). // file, fileInf.Size(), baseMimeType).
		ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).
		Do()

	if err != nil {
		return
	}

	newRes.FileId = res.Id
	newRes.FileUrl = g.GenerateLink(res.Id)

	return
}
