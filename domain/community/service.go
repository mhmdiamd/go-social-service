package community

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	communitymember "github.com/mhmdiamd/go-social-service/domain/community_member"
	"github.com/mhmdiamd/go-social-service/external/google"
	tempdata "github.com/mhmdiamd/go-social-service/temp_data"
)

type Repository interface {
	CommunityRepository
	CommunityMemberRepository
	CommunityTransactionDBRepository
}

type CommunityTransactionDBRepository interface {
	Begin(ctx context.Context) (tx *sqlx.Tx, err error)
	Commit(ctx context.Context, tx *sqlx.Tx) (err error)
	Rollback(ctx context.Context, tx *sqlx.Tx) (err error)
}

type CommunityRepository interface {
	Create(ctx context.Context, entity Community) (id int, err error)
	UpdateById(ctx context.Context, entity Community) (err error)
	DeleteById(ctx context.Context, id int) (err error)

	GetAll(ctx context.Context, communityPagination CommunityPagination) (communities []Community, err error)
	GetById(ctx context.Context, id int) (model Community, err error)
}

type CommunityMemberRepository interface {
	GetAllByUserPublicId(ctx context.Context, publicId string) (communities []Community, err error)
	DeleteCommunityMemberByIdCommunity(ctx context.Context, communityId int) (err error)
	CreateCommunityMember(ctx context.Context, entity CommunityMember) (err error)
}

type service struct {
	repo          Repository
	googleService *google.GoogleDrive
}

func newService(repo Repository, gs *google.GoogleDrive) service {
	return service{
		repo:          repo,
		googleService: gs,
	}
}

func (s service) CreateCommunity(ctx context.Context, req CreateCommunityRequestPayload) (err error) {

	communityEntity := NewCommunityFromCreate(req)

	// Validation payload first
	if err = communityEntity.Validate(); err != nil {
		return
	}

	// Start transaction
	tx, err := s.repo.Begin(ctx)
	if err != nil {
		return
	}

	// defer rollback if any error on transaction
	defer tx.Rollback()

	// Create new Community
	communityId, err := s.repo.Create(ctx, communityEntity)
	if err != nil {
		return
	}

	tempdata.LastCommunityID = communityId

	// Create new community member as an owner
	reqCommunityMember := CreateCommunityMembersRequestPayload{
		CommunityId:  communityId,
		UserPublicId: req.UserPublicId,
		Role:         communitymember.CommunityMemberRole_owner,
	}

	communityMemberEntity := NewCommunityMembersFromCreate(reqCommunityMember)
	if err = s.repo.CreateCommunityMember(ctx, communityMemberEntity); err != nil {
		return
	}

	// Commit the transaction
	if err = s.repo.Commit(ctx, tx); err != nil {
		return
	}

	return
}

func (s service) UpdateById(ctx context.Context, publicUserId uuid.UUID, req UpdateCommunityRequestPayload) (err error) {

	entity := NewCommunityFromUpdate(req)

	// validation check
	if err = entity.Validate(); err != nil {
		return
	}

	// Check Image first
	community, err := s.repo.GetById(ctx, req.Id)
	if err != nil {
		return
	}

	if community.IsImageExist() {

		// Handle if photo not exist
		if req.Logo != nil {
			err = s.googleService.UpdateFileById(ctx, community.FileIdGdrive, req.Logo)
			if err != nil {
				return
			}
		}

		entity.FileIdGdrive = community.FileIdGdrive
		entity.Logo = community.Logo

		err = s.repo.UpdateById(ctx, entity)
		if err != nil {
			return
		}

	} else {

		if req.Logo != nil {
			fileRes, err := s.googleService.UploadFile(ctx, publicUserId, req.Logo)
			if err != nil {
				return err
			}

			entity.FileIdGdrive = fileRes.FileId
			entity.Logo = fileRes.FileUrl
		}

		err = s.repo.UpdateById(ctx, entity)
		if err != nil {
			return err
		}
	}

	return
}

func (s service) DeleteById(ctx context.Context, id int) (err error) {

	// Check the entity first
	_, err = s.GetById(ctx, id)
	if err != nil {
		return
	}

	// Create transaction
	tx, err := s.repo.Begin(ctx)
	if err != nil {
		return
	}

	defer tx.Rollback()

	// Delete Community By Id
	err = s.repo.DeleteById(ctx, id)
	if err != nil {
		return
	}

	// Delete Community Member By Community Id
	err = s.repo.DeleteCommunityMemberByIdCommunity(ctx, id)
	if err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func (s service) GetAll(ctx context.Context, cq ListCommunityRequestPayload) ([]CommunityListResponse, error) {
	pagination := NewCommunityPaginationFromListCommunityRequest(cq)

	communities, err := s.repo.GetAll(ctx, pagination)
	if err != nil {
		return []CommunityListResponse{}, err
	}

	// Convert to response
	communityListResponse := NewCommunityListResponseFromEntity(communities)
	return communityListResponse, nil
}

func (s service) GetById(ctx context.Context, id int) (communityRes CommunityDetailResponse, err error) {
	model, err := s.repo.GetById(ctx, id)

	communityRes = model.ToCommunityResponse()
	if err != nil {
		return
	}

	return
}
