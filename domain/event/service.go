package event

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-social-service/internal/log"
	tempdata "github.com/mhmdiamd/go-social-service/temp_data"
)

type Repository interface {
	EventRepository
  EventDemographicsRepository
	EventCommiteRepository
	EventTransactionRepository
}

type EventTransactionRepository interface {
	Begin(ctx context.Context) (tx *sqlx.Tx, err error)
	Commit(ctx context.Context, tx *sqlx.Tx) error
	Rollback(ctx context.Context, tx *sqlx.Tx) error
}

type EventRepository interface {
	Create(ctx context.Context, tx *sqlx.Tx, entity Event) (err error)
	UpdateById(ctx context.Context, tx *sqlx.Tx, entity Event) (err error)
	DeleteById(ctx context.Context, eventPublicId string) (err error)

	GetAllWithPagination(ctx context.Context, pagination ListEventRequestPayload) (events []Event, err error)
	GetDetailById(ctx context.Context, eventPublicId string) (event Event, err error)
}

type EventDemographicsRepository interface {
  GetEventDemographicsById(ctx context.Context, eventDemographicsId int) (ed EventDemographics, err error)
}

type EventCommiteRepository interface {
	CreateEventCommite(ctx context.Context, tx *sqlx.Tx, ec EventCommite) (err error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) service {
	return service{
		repo: r,
	}
}

func (s *service) GetAllWithPagination(ctx context.Context, req ListEventRequestPayload) (events []EventResponse, err error) {
	pagination := req.GenerateDefaultValue()

	entities, err := s.repo.GetAllWithPagination(ctx, pagination)
	if err != nil {
    log.Log.Errorf(ctx, "[GetAllWithPagination, GetAllWithPagination] with error detail %s", err.Error())
		return
	}

	return ConvertToEventResponseList(entities), nil
}

func (s *service) GetDetailById(ctx context.Context, event_public_id string) (event Event, err error) {
	event, err = s.repo.GetDetailById(ctx, event_public_id)
	if err != nil {
    log.Log.Errorf(ctx, "[GetDetailById, GetDetailById] with error detail %s", err.Error())
		return
	}

	return
}

func (s *service) Create(ctx context.Context, req CreateEventRequestPayload) (err error) {

	event := NewEventFromCreate(req)

	if err = event.Validate(); err != nil {
		return
	}

	tx, err := s.repo.Begin(ctx)

	if err != nil {
    log.Log.Errorf(ctx, "[Create, Validate] with error detail %s", err.Error())
		return
	}

  // Get event demographics firs
  ed, err := s.repo.GetEventDemographicsById(ctx, event.EventDemographicsId)
  if err != nil {
    log.Log.Errorf(ctx, "[Create, GetEventDemographicsById] with error detail %s", err.Error())
    return
  }

  // set Into json
  event.SetEventDemographicsJSON(ed)

	// Create Event First
	if err = s.repo.Create(ctx, tx, event); err != nil {
    log.Log.Errorf(ctx, "[Create, Create] with error detail %s", err.Error())
		return
	}

	commite := CreateEventCommiteRequestPayload{
		UserPublicId:  req.UserPublicId,
		EventPublicId: event.PublicId,
		Position:      EventPosition_Admin,
	}

  newCommite := NewEventCommiteFromCreate(commite)
  tempdata.TempCurrentEventPublicId = event.PublicId
	// Create Event Commite admin
	if err = s.repo.CreateEventCommite(ctx, tx, newCommite); err != nil {
    log.Log.Errorf(ctx, "[Create, CreateEventCommite] with error detail %s", err.Error())
		return
	}

	if err = s.repo.Commit(ctx, tx); err != nil {
    log.Log.Errorf(ctx, "[Create, Commit] with error detail %s", err.Error())
		return
	}

	defer s.repo.Rollback(ctx, tx)
	return
}

func (s *service) UpdateById(ctx context.Context, req UpdateEventRequestPayload) (err error) {

	newEvent := NewEventFromUpdate(req) 

  // Validation
  if err = newEvent.Validate(); err != nil {
    return
  }

	tx, err := s.repo.Begin(ctx)
	if err != nil {
    log.Log.Errorf(ctx, "[Update, Begin] with error detail %s", err.Error())
		return
	}

	// make sure that event is exist
  _, err = s.repo.GetDetailById(ctx, newEvent.PublicId)
	if err != nil {
    log.Log.Errorf(ctx, "[Update, GetDetailById] with error detail %s", err.Error())
		return
	}

  ed, err := s.repo.GetEventDemographicsById(ctx, newEvent.EventDemographicsId)
  if err != nil {
    log.Log.Errorf(ctx, "[Update, GetEventDemographicsById] with error detail %s", err.Error())
    return
  }

  // if event demographics has a different data, set new data as new event demographics
  newEvent.SetEventDemographicsJSON(ed)

	// Update the event by event_id
	if err = s.repo.UpdateById(ctx, tx, newEvent); err != nil {
    log.Log.Errorf(ctx, "[Update, UpdateById] with error detail %s", err.Error())
		return
	}

	// Commite the transactions
	if err = s.repo.Commit(ctx, tx); err != nil {
    log.Log.Errorf(ctx, "[Update, Commit] with error detail %s", err.Error())
		return
	}

	defer s.repo.Rollback(ctx, tx)
	return
}

func (s *service) DeleteById(ctx context.Context, eventPublicId string) (err error) {

	// Check is the event exist
	_, err = s.repo.GetDetailById(ctx, eventPublicId)
	if err != nil {
    log.Log.Errorf(ctx, "[DeleteById, GetDetailById] with error detail %s", err.Error())
		return
	}

	// Delete the Event
	if err = s.repo.DeleteById(ctx, eventPublicId); err != nil {
    log.Log.Errorf(ctx, "[DeleteById, DeleteById] with error detail %s", err.Error())
		return
	}

	return
}
