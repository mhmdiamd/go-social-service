package community

import (
	"strings"
	"time"

	"github.com/mhmdiamd/go-social-service/infra/response"
)

type Community struct {
	Id                  int    `db:"id"`
	Name                string `db:"name"`
	Description         string `db:"description"`
	Logo                string `db:"logo"`
	ExternalCategories  string `db:"external_categories"`
	CategoryCommunityID int    `db:"category_community_id"`
	FileIdGdrive        string `db:"file_id_gdrive"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CommunityPagination struct {
	Cursor int `db:"cursor" json:"cursor"`
	Size   int `db:"size" json:"size"`
}

func NewCommunityFromCreate(req CreateCommunityRequestPayload) Community {
	entity := Community{
		Name:                req.Name,
		Logo:                "default.jpg",
		CategoryCommunityID: req.CategoryCommunityID,
		FileIdGdrive:        "",
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	if len(req.ExternalCategories) > 0 {
		entity.ConvertExternalCategoriesToString(req.ExternalCategories)
	}

	return entity
}

func NewCommunityFromUpdate(req UpdateCommunityRequestPayload) Community {
	entity := Community{
		Id:                  req.Id,
		Name:                req.Name,
		Description:         req.Description,
		CategoryCommunityID: req.CategoryCommunityID,
	}

	if req.Logo == nil {
		entity.Logo = "https://drive.google.com/uc?id=default.jpg"
	}

	return entity
}

func NewCommunityPaginationFromListCommunityRequest(req ListCommunityRequestPayload) CommunityPagination {
	req = req.GenerateDefaultValue()

	return CommunityPagination{
		Cursor: req.Cursor,
		Size:   req.Size,
	}
}

func (c Community) Validate() (err error) {
	if err = c.ValidateName(); err != nil {
		return
	}

	if err = c.ValidateCategory(); err != nil {
		return
	}

	return
}

func (c Community) ValidateName() (err error) {
	if c.Name == "" {
		return response.ErrNameRequired
	}

	return
}

func (c Community) ValidateCategory() (err error) {
	if c.CategoryCommunityID == 0 {
		if len(strings.Split(c.ExternalCategories, ",")) <= 1 && strings.Split(c.ExternalCategories, ",")[0] == "" {
			return response.ErrCommunityCategoryIdRequired
		}
	}

	return
}

func (c *Community) ConvertExternalCategoriesToString(externalCategories []string) {
	newString := ""
	for i, category := range externalCategories {
		if i == 0 {
			newString = newString + category
		} else {
			newString = newString + "," + category
		}
	}

	c.ExternalCategories = newString
}

func (c Community) IsImageExist() bool {
	return c.FileIdGdrive != ""
}

func (c Community) ToCommunityListResponse() CommunityListResponse {
	externalCategories := strings.Split(c.ExternalCategories, ",")

	return CommunityListResponse{
		Id:                  c.Id,
		Name:                c.Name,
		Description:         c.Description,
		Logo:                c.Logo,
		ExternalCategories:  externalCategories,
		CategoryCommunityId: c.CategoryCommunityID,
	}
}

func (c Community) ToCommunityResponse() CommunityDetailResponse {
	externalCategories := strings.Split(c.ExternalCategories, ",")
	return CommunityDetailResponse{
		Id:                  c.Id,
		Name:                c.Name,
		Description:         c.Description,
		Logo:                c.Logo,
		ExternalCategories:  externalCategories,
		CategoryCommunityId: c.CategoryCommunityID,
		CreatedAt:           c.CreatedAt,
		UpdatedAt:           c.UpdatedAt,
	}
}
