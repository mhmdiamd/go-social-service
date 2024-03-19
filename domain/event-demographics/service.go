package eventdemographics

import (
	"context"

	"github.com/mhmdiamd/go-social-service/infra/response"
)

type Repository interface {
  Create(ctx context.Context, entity EventDemographicsEntity) (err error)
  UpdateById(ctx context.Context, model EventDemographicsEntity) (err error)
  GetById(ctx context.Context, id int) (entity EventDemographicsEntity, err error)
  GetAll(ctx context.Context) (models []EventDemographicsEntity, err error)
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

func (s service) AddEventDemographics(ctx context.Context, req CreateEventDemographicsRequestPayload) (err error) {

  // Create event demographics entity
  eventDemographicsEntity := NewEventDemographicsEntity(req)
 
   // Entity Validation
  if err = eventDemographicsEntity.Validate(); err != nil {
    return
  }

  // Convert graduation from slice string to string
  eventDemographicsEntity.ConvertGraduationToString(req.Graduation)

  // hit service to create event demographics entity
  err = s.repo.Create(ctx, eventDemographicsEntity)
  if err != nil {
    return 
  }

  return
}

func (s service) UpdateById(ctx context.Context, req UpdateEventDemographicsRequestPayload) (err error) {

  eventDemographicsEntity := NewEventDemographicsEntityFromUpdate(req)
  if err = eventDemographicsEntity.Validate(); err != nil {
    return
  }

  if err = eventDemographicsEntity.ValidateId(); err != nil {
    return
  }

  // Check is event demographis alredy Exists firs
  _, err = s.repo.GetById(ctx, eventDemographicsEntity.Id)
  if err != nil {
    return
  }

  // hit service to create event demographics entity
  err = s.repo.UpdateById(ctx, eventDemographicsEntity)
  if err != nil {
    return 
  }

  return 
}

func (s service) DeleteById(ctx context.Context, id int) (err error) {

  _, err = s.repo.GetById(ctx, id)
  if err != nil {
    return
  }

  err = s.repo.DeleteById(ctx, id)
  if err != nil {
    return
  }

  return 
}

func (s service) GetById(ctx context.Context, id int) (res EventDemographicsEntityResponse, err error) {

  model, err := s.repo.GetById(ctx, id)
  if err != nil {
    return
  }

  // Convert graduation to []string
  res = NewEventDemographicsEntityResponse(model)

  return 
}

func (s service) GetListEventDemographics(ctx context.Context) (eventDemographies []EventDemographicsEntityResponse, err error) {
  
  events, err := s.repo.GetAll(ctx)

  if err != nil {
    if err == response.ErrNotFound {
      return []EventDemographicsEntityResponse{}, err
    }

    return
  }

  if len(events) == 0 {
    return []EventDemographicsEntityResponse{}, err
  }

  eventDemographies = NewListEventDemographicsEntityResponse(events)

  return
}
