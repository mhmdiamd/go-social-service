package communitymember

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type Repository interface {
  CommunityMemberTransactionRepository
  CommunityMemberRepository
  CommunityRepository
  AuthRepository
}

type CommunityMemberTransactionRepository interface {
  Begin(ctx context.Context) (tx *sqlx.Tx, err error)
  Commit(ctx context.Context, tx *sqlx.Tx) (err error)
  Rollback(ctx context.Context, tx *sqlx.Tx) (err error)
}

type CommunityMemberRepository interface {
  Create(ctx context.Context, cm CommunityMember) (err error)
  Update(ctx context.Context, cm CommunityMember) (err error)
  Delete(ctx context.Context, userId string, communityId int) (err error)
  UpdateRole(ctx context.Context, memberId string, role CommunityMemberRole) (err error)

  GetAllByCommunityId(ctx context.Context, communityId int, pagination CommunityMemberPagination) (members CommunityMembers, err error) 
  GetDetailMember(ctx context.Context, userId string, communityId int) (member CommunityMember, err error)

  UpdateMemberStatus(ctx context.Context, communityId int, memberId string, status int) (err error)
}

type CommunityRepository interface {
  GetCommunityById(ctx context.Context, communityId int) (community Community, err error)
}

type AuthRepository interface {
  IsUserAuthExist(ctx context.Context, userId string) (err error)
}

type service struct {
  repo Repository
}

func newService(repo Repository) service {
  return service{
    repo: repo,
  }
}

func (s service) GetAllMemberByCommunityId(ctx context.Context, communityId int, paginatonPayload CommunityMemberListRequestPayload) (membersConverted []CommunityMemberResponse, err error) {

  // Get Community for checking
  _, err = s.repo.GetCommunityById(ctx, communityId);
  if err != nil {
    return
  }

  // Create new pagination
  pagination := NewCommunityMemberPagination(paginatonPayload)

  // Get All Member from Community
  members, err := s.repo.GetAllByCommunityId(ctx, communityId, pagination)
  if err != nil {
    return
  }

  // Converted member to response
  membersConverted = NewCommunityMemberListFromEntity(members)

  return
}

func (s service) AddMember(ctx context.Context, req AddCommunityMemberRequestPayload) (err error) {
  cm := NewCommunityMemberFromAdd(req)

  if err = cm.Validate(); err != nil {
    return
  }

  // Chech is user alredy exists
  if err = s.repo.IsUserAuthExist(ctx, cm.UserPublicId); err != nil {
    return
  }

  if err = s.repo.Create(ctx, cm); err != nil {
    return
  }

  return
}

func (s service) Update(ctx context.Context, req UpdateCommunityMemberRequestPayload) (err error) {
  cm := NewCommunityMemberFromUpdate(req)

  if err = cm.Validate(); err != nil {
    return
  }

  // Start Transaction
  tx, err := s.repo.Begin(ctx)
  if err != nil {
    return
  }

  defer tx.Rollback()

  // Check The entity first
  _, err = s.repo.GetDetailMember(ctx, req.UserPublicId, req.CommunityId)
  if err != nil {
    return
  }

  // Update the data
  if err = s.repo.Update(ctx, cm); err != nil {
    return
  }

  // Commit the query
  if err = s.repo.Commit(ctx, tx); err != nil {
    return
  }

  return
}

func(s service) KickMember(ctx context.Context, editorId, userId string, communityId int) (err error) {

  if userId == "" {
    return response.ErrIdRequired
  }

  editor, err := s.repo.GetDetailMember(ctx, editorId, communityId); 
  if err != nil {
    return
  }

  // Check editor is admin or owner
  if !editor.IsOwner() || !editor.IsAdmin() {
    return response.ErrCommunityMemberRoleNotPermitted
  }

  // Get user
  member, err := s.repo.GetDetailMember(ctx, userId, communityId); 
  if err != nil {
    return
  }

  member.IsActive = 0

  if err = s.repo.Update(ctx, member); err != nil {
    return
  }

  return
}

func (s service) DeleteCommunityMember(ctx context.Context, userId string, communityId int) (err error){

  // Get user for checking
  _, err = s.repo.GetDetailMember(ctx, userId, communityId); 
  if err != nil {
    return
  }

  // Delete by id
  if err = s.repo.Delete(ctx, userId, communityId); err != nil {
    return
  }

  return
}









