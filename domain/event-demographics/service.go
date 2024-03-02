package eventdemographics

import (
	"context"
	"fmt"
)

type Repository interface {
  Create(ctx context.Context, entity EventDemographicsEntity) (err error)
  UpdateById(ctx context.Context, model EventDemographicsEntity) (entity EventDemographicsEntity, err error)
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

func (s service) UpdateById(ctx context.Context, req UpdateEventDemographicsRequestPayload) (model EventDemographicsEntity, err error) {

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
  model, err = s.repo.UpdateById(ctx, eventDemographicsEntity)
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
  fmt.Println(res.Graduation)

  return 
}


