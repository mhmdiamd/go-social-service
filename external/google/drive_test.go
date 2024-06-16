package google

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
	"github.com/stretchr/testify/require"
)

var svc *GoogleDrive

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	con, err := ConnectServiceGoogleDrive()
	if err != nil {
		panic(err)
	}

	svc = NewGoogleDriveService(con)
}

func Test_GoogleDriveConnect_Success(t *testing.T) {
	id, err := svc.UploadFile(context.Background(), uuid.New(), "meme-fb.jpg")
	require.Nil(t, err)
	require.NotNil(t, id)
}

func Test_GoogleDriveUpload(t *testing.T) {
	t.Run("fail, image oversize", func(t *testing.T) {
	})

	t.Run("fail, image type not compatible", func(t *testing.T) {
		file, err := svc.UploadFile(context.Background(), uuid.New(), "coverage.out")
		require.NotNil(t, err)
		require.NotNil(t, file)
		require.Equal(t, response.ErrImageTypeNotCompatible, err)
	})
}
