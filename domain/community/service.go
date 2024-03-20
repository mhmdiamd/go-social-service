package community

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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

type Service struct {
	Repo          Repository
	GoogleService *google.GoogleDrive
}

func NewService(repo Repository, gs *google.GoogleDrive) Service {
	return Service{
		Repo:          repo,
		GoogleService: gs,
	}
}

func (s Service) CreateCommunity(ctx context.Context, req CreateCommunityRequestPayload) (err error) {

	communityEntity := NewCommunityFromCreate(req)

	// Validation payload first
	if err = communityEntity.Validate(); err != nil {
		return
	}

	// Start transaction
	tx, err := s.Repo.Begin(ctx)
	if err != nil {
		return
	}

	// defer rollback if any error on transaction
	defer tx.Rollback()

	// Create new Community
	communityId, err := s.Repo.Create(ctx, communityEntity)
	if err != nil {
		return
	}

	tempdata.LastCommunityID = communityId

	// Create new community member as an owner
	reqCommunityMember := CreateCommunityMembersRequestPayload{
		CommunityId:  communityId,
		UserPublicId: req.UserPublicId,
		Role:         CommunityMemberRole_owner,
	}

	communityMemberEntity := NewCommunityMembersFromCreate(reqCommunityMember)
	if err = s.Repo.CreateCommunityMember(ctx, communityMemberEntity); err != nil {
		return
	}

	// Commit the transaction
	if err = s.Repo.Commit(ctx, tx); err != nil {
		return
	}

	return
}

func (s Service) UpdateById(ctx context.Context, publicUserId uuid.UUID, req UpdateCommunityRequestPayload) (err error) {

	entity := NewCommunityFromUpdate(req)

	// validation check
	if err = entity.Validate(); err != nil {
		return
	}

	// Check Image first
	community, err := s.Repo.GetById(ctx, req.Id)
	if err != nil {
		return
	}

	if community.IsImageExist() {

		// Handle if photo not exist
		if req.Logo != nil {
			err = s.GoogleService.UpdateFileById(ctx, community.FileIdGdrive, req.Logo)
			if err != nil {
				return
			}
		}

		entity.FileIdGdrive = community.FileIdGdrive
		entity.Logo = community.Logo

		err = s.Repo.UpdateById(ctx, entity)
		if err != nil {
			return
		}

	} else {

		if req.Logo != nil {
			fileRes, err := s.GoogleService.UploadFile(ctx, publicUserId, req.Logo)
			if err != nil {
				return err
			}

			entity.FileIdGdrive = fileRes.FileId
			entity.Logo = fileRes.FileUrl
		}

		err = s.Repo.UpdateById(ctx, entity)
		if err != nil {
			return err
		}
	}

	return
}

func (s Service) DeleteById(ctx context.Context, id int) (err error) {

	// Check the entity first
	_, err = s.GetById(ctx, id)
	if err != nil {
		return
	}

	// Create transaction
	tx, err := s.Repo.Begin(ctx)
	if err != nil {
		return
	}

	defer tx.Rollback()

	// Delete Community By Id
	err = s.Repo.DeleteById(ctx, id)
	if err != nil {
		return
	}

	// Delete Community Member By Community Id
	err = s.Repo.DeleteCommunityMemberByIdCommunity(ctx, id)
	if err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func (s Service) GetAll(ctx context.Context, cq ListCommunityRequestPayload) ([]CommunityListResponse, error) {
	pagination := NewCommunityPaginationFromListCommunityRequest(cq)

	communities, err := s.Repo.GetAll(ctx, pagination)
	if err != nil {
		return []CommunityListResponse{}, err
	}

	// Convert to response
	communityListResponse := NewCommunityListResponseFromEntity(communities)
	return communityListResponse, nil
}

func (s Service) GetById(ctx context.Context, id int) (communityRes CommunityDetailResponse, err error) {
	model, err := s.Repo.GetById(ctx, id)

	communityRes = model.ToCommunityResponse()
	if err != nil {
		return
	}

	return
}
