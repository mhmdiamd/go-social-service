package eventcommittee

import (
	"context"

	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/log"
)

type Repository interface {
	Create(ctx context.Context, entity EventCommitteeEntity) (err error)
	UpdateById(ctx context.Context, model EventCommitteeEntity) (err error)
	GetById(ctx context.Context, id int) (entity EventCommitteeEntity, err error)
	GetAll(ctx context.Context) (models []EventCommitteeEntity, err error)
	DeleteById(ctx context.Context, id int) (err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) AddEventCommittee(ctx context.Context, req CreateEventCommitteeRequestPayload) (err error) {
	// Create event demographics entity
	ec := NewEventCommitteeEntity(req)

	// Entity Validation
	if err = ec.Validate(); err != nil {
		log.Log.Errorf(ctx, "[Create, Validate] with error detail %s", err.Error())
		return
	}

	// hit service to create event demographics entity
	err = s.repo.Create(ctx, ec)
	if err != nil {
		log.Log.Errorf(ctx, "[Create, AddEventCommittee] with error detail %s", err.Error())
		return
	}

	return
}

func (s service) UpdateById(ctx context.Context, req UpdateEventCommitteeRequestPayload) (err error) {
	ec := NewEventCommitteeEntityFromUpdate(req)
	if err = ec.Validate(); err != nil {
		log.Log.Errorf(ctx, "[Create, Validate] with error detail %s", err.Error())
		return
	}

	if err = ec.ValidateId(); err != nil {
		return
	}

	// Check is event demographis alredy Exists firs
	_, err = s.repo.GetById(ctx, ec.Id)
	if err != nil {
		log.Log.Errorf(ctx, "[Create, GetById] with error detail %s", err.Error())
		return
	}

	// hit service to create event demographics entity
	err = s.repo.UpdateById(ctx, ec)
	if err != nil {
		log.Log.Errorf(ctx, "[Create, UpdateById] with error detail %s", err.Error())
		return
	}

	return
}

func (s service) DeleteById(ctx context.Context, id int) (err error) {
	_, err = s.repo.GetById(ctx, id)
	if err != nil {
		log.Log.Errorf(ctx, "[DeleteById, GetById] with error detail %s", err.Error())
		return
	}

	err = s.repo.DeleteById(ctx, id)
	if err != nil {
		log.Log.Errorf(ctx, "[DeleteById, DeleteById] with error detail %s", err.Error())
		return
	}

	return
}

func (s service) GetById(ctx context.Context, id int) (ecr EventCommitteeEntityResponse, err error) {
	model, err := s.repo.GetById(ctx, id)
	if err != nil {
		return
	}

	// Convert graduation to []string
	ecr = NewEventCommitteeEntityResponse(model)

	return
}

func (s service) GetListEventCommittee(ctx context.Context) (eventDemographies []EventCommitteeEntityResponse, err error) {
	ecs, err := s.repo.GetAll(ctx)
	if err != nil {
		if err == response.ErrNotFound {
			return []EventCommitteeEntityResponse{}, err
		}

		return
	}

	if len(ecs) == 0 {
		return []EventCommitteeEntityResponse{}, err
	}

	eventDemographies = NewListEventCommitteeEntityResponse(ecs)

	return
}
