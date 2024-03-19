package communitymember

import (
	"time"

	"github.com/mhmdiamd/go-social-service/infra/response"
)

type CommunityMemberRole string

const (
  CommunityMemberRole_member CommunityMemberRole = "member"
  CommunityMemberRole_admin CommunityMemberRole = "admin"
  CommunityMemberRole_owner CommunityMemberRole = "owner"
)

var (
  CommunityMemberMapping = map[string]CommunityMemberRole{
    "member" : CommunityMemberRole_member,
    "admin" : CommunityMemberRole_admin,
    "owner" : CommunityMemberRole_owner,
  }
)

type CommunityMemberPagination struct {
  Cursor int `db:"cursor"`
  Size int `db:"size"`
}

func NewCommunityMemberPagination(req CommunityMemberListRequestPayload) CommunityMemberPagination {
  req = req.GenerateDefaultValue()

  return CommunityMemberPagination{
    Cursor: req.Cursor,
    Size: req.Size,
  }
}

type CommunityMember struct {
  Id int `db:"id"`
  Role CommunityMemberRole `db:"role"`
  Nik string `db:"nik"`
  IsActive int `db:"is_active"`
  PhotoKTP int `db:"photo_ktp"`

  // user public id foreign key
  UserPublicId string `db:"user_public_id"`

  // id community
  CommunityId int `db:"community_id"`

  CreatedAt time.Time `db:"created_at" json:"created_at"`
  UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CommunityMembers []CommunityMember

func NewCommunityMemberFromAdd(req AddCommunityMemberRequestPayload) CommunityMember {
  cm := CommunityMember{
    UserPublicId: req.UserPublicId,
    Role: CommunityMemberRole_member,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }

  registredRole, ok := CommunityMemberMapping[string(req.Role)];
  if ok {
    cm.Role = registredRole 
  }

  return cm
}

func NewCommunityMemberFromUpdate(req UpdateCommunityMemberRequestPayload) CommunityMember {
  cm := CommunityMember{
    UserPublicId: req.UserPublicId,
    Role: CommunityMemberRole_member,
    PhotoKTP: req.PhotoKTP ,
    Nik: req.Nik,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }

  registredRole, ok := CommunityMemberMapping[string(req.Role)];
  if ok {
    cm.Role = registredRole 
  }

  return cm
}

func (c CommunityMember) Validate() (err error) {

  if err = c.ValidateUserId(); err != nil {
    return
  }

  return
}

func (c CommunityMember) ValidateUserId() (err error) {
  if c.UserPublicId == "" {
    return response.ErrIdRequired
  }

  return
}

func (c CommunityMember) IsAdmin() bool {
  return c.Role == CommunityMemberRole_admin
}

func (c CommunityMember) IsOwner() bool {
  return c.Role == CommunityMemberRole_owner
}

func (c CommunityMember) ToCommunityMemberResponse() CommunityMemberResponse {
  return CommunityMemberResponse{
    Role : c.Role,
    Nik : c.Nik,
    IsActive : c.IsActive,
    PhotoKTP : c.PhotoKTP,
    UserPublicId : c.UserPublicId,
    CommunityId : c.CommunityId,
    CreatedAt : c.CreatedAt,
    UpdatedAt : c.UpdatedAt,
  }
}


