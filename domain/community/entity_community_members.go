package community

import (
	"time"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type CommunityMemberRole string

const (
	CommunityMemberRole_member CommunityMemberRole = "member"
	CommunityMemberRole_admin  CommunityMemberRole = "admin"
	CommunityMemberRole_owner  CommunityMemberRole = "owner"
)

var (
	CommunityMemberMapping = map[string]CommunityMemberRole{
		"member": CommunityMemberRole_member,
		"admin":  CommunityMemberRole_admin,
		"owner":  CommunityMemberRole_owner,
	}
)

type CommunityMember struct {
	Id       int                 `db:"id"`
	Role     CommunityMemberRole `db:"role"`
	Nik      string              `db:"nik"`
	IsActive int                 `db:"is_active"`
	PhotoKTP int                 `db:"photoktp"`

	// user public id foreign key
	UserPublicId uuid.UUID `db:"user_public_id"`

	// id community
	CommunityId int `db:"community_id"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func NewCommunityMembersFromCreate(req CreateCommunityMembersRequestPayload) CommunityMember {
	entity := CommunityMember{
		CommunityId:  req.CommunityId,
		UserPublicId: req.UserPublicId,
		Role:         CommunityMemberRole_member,
		IsActive:     1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if req.Role != "" {
		entity.Role = req.Role
	}

	return entity
}

func (c CommunityMember) Validate() (err error) {

	if err = c.ValidateCommunityId(); err != nil {
		return
	}

	return
}

func (c CommunityMember) ValidateCommunityId() (err error) {

	if c.CommunityId == 0 {
		return response.ErrIdRequired
	}

	return
}
