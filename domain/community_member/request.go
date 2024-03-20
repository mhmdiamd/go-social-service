package communitymember

type CommunityMemberListRequestPayload struct {
	Cursor int `query:"cursor" json:"cursor"`
	Size   int `query:"size" json:"size"`
}

type AddCommunityMemberRequestPayload struct {
	UserPublicId string              `json:"user_id"`
	CommunityId  int                 `json:"community_id"`
	Role         CommunityMemberRole `json:"role"`
}

type UpdateCommunityMemberRequestPayload struct {
	Role     CommunityMemberRole `json:"role"`
	Nik      string              `json:"nik"`
	IsActive int                 `json:"is_active"`
	PhotoKTP string              `json:"photo_ktp"`

	// user public id foreign key
	UserPublicId string `json:"user_public_id"`

	CommunityId int `json:"community_id"`
}

func (c CommunityMemberListRequestPayload) GenerateDefaultValue() CommunityMemberListRequestPayload {

	if c.Cursor <= 0 {
		c.Cursor = 1
	}

	if c.Size <= 0 {
		c.Size = 10
	}

	return c
}
