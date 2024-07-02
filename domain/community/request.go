package community

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type CreateCommunityRequestPayload struct {
	UserPublicId        uuid.UUID `json:"user_public_id"`
	Name                string    `json:"name"`
	ExternalCategories  []string  `json:"external_categories"`
	CategoryCommunityID int       `json:"category_community_id"`
}

type UpdateCommunityRequestPayload struct {
	Id                  int                   `json:"id"`
	Name                string                `json:"name"`
	Description         string                `json:"description"`
	Logo                *multipart.FileHeader `json:"logo"`
	ExternalCategories  string                `json:"external_categories"`
	CategoryCommunityID int                   `json:"category_community_id"`
}

type CreateCommunityMembersRequestPayload struct {
	Nik      string              `json:"nik"`
	PhotoKTP string              `json:"photo_ktp"`
	Role     CommunityMemberRole `json:"role"`
	IsActive int                 `json:"is_active"`

	// user public id foreign key
	UserPublicId uuid.UUID `json:"user_public_id"`
	// Community Id
	CommunityId int `json:"community_id"`
}

type ListCommunityRequestPayload struct {
	Cursor int `query:"cursor" json:"cursor"`
	Size   int `query:"size" json:"size"`
}

func (l ListCommunityRequestPayload) GenerateDefaultValue() ListCommunityRequestPayload {
	if l.Cursor < 0 {
		l.Cursor = 1
	}

	if l.Size <= 0 {
		l.Size = 10
	}

	return l
}
