package event

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
  EventRepository
  EventCommiteRepository
  EventTransactionRepository
}

type EventTransactionRepository interface {
  Begin(ctx context.Context) (tx *sqlx.Tx, err error)
  Commite(ctx context.Context, tx *sqlx.Tx) error
  Rollback(ctx context.Context, tx *sqlx.Tx) error
}

type EventRepository interface {
  Create(ctx context.Context, tx *sqlx.Tx, entity Event) (err error)
  Update(ctx context.Context, tx *sqlx.Tx, entity Event) (err error)
  DeleteById(ctx context.Context, eventPublicId string) (err error)

  GetAllWithPagination(ctx context.Context, pagination ListEventRequestPayload) (events []Event, err error)
  GetDetailById(ctx context.Context, event_public_id string) (event Event, err error)
}

type EventCommiteRepository interface {
  CreateCommite(ctx context.Context, tx *sqlx.Tx, entity EventCommite) (err error)
}

type service struct {
  repo Repository
}

func newService(r Repository) service {
  return service{
    repo: r,
  }
}

func (s *service) GetAllWithPagination(ctx context.Context, req ListEventRequestPayload) (events []EventResponse, err error) {
  entities, err := s.repo.GetAllWithPagination(ctx, req)
  if err != nil {
    return 
  }

  return ConvertToEventResponseList(entities), nil
}

func (s *service) GetDetailById(ctx context.Context, event_id string) (event Event, err error) {
  event, err = s.repo.GetDetailById(ctx, event_id); 
  if err != nil {
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
    return 
  }

  // Create Event First
  if err = s.repo.Create(ctx, tx, event); err != nil {
    return
  }

  commite := EventCommite{
    UserPublicId: req.UserPublicId,
    EventPublicId: event.PublicId,
    Position: EventPosition_Admin,
  }
  // Create Event Commite admin
  if err = s.repo.CreateCommite(ctx, tx, commite); err != nil {
    return
  }

  if err = s.repo.Commite(ctx, tx); err != nil {
    return
  }

  defer s.repo.Rollback(ctx, tx)
  return 
}

func (s *service) Update(ctx context.Context, req UpdateEventRequestPayload) (err error) {

  newEvent := NewEventFromUpdate(req)

  tx, err := s.repo.Begin(ctx)
  if err != nil {
    return
  }

  // make sure that event id is exist
  _, err = s.repo.GetDetailById(ctx, newEvent.PublicId)
  if err != nil {
    return
  }

  // Update the event by event_id
  if err = s.repo.Update(ctx, tx, newEvent); err != nil {
    return 
  }

  // Commite the transactions
  if err = s.repo.Commite(ctx, tx); err != nil {
    return
  }

  defer s.repo.Rollback(ctx, tx)
  return 
}

func (s *service) DeleteById(ctx context.Context, eventPublicId string) (err error) {

  // Check is the event exist
  _, err = s.repo.GetDetailById(ctx, eventPublicId); 
  if err != nil {
    return
  }

  // Delete the Event
  if err = s.repo.DeleteById(ctx, eventPublicId); err != nil {
    return
  }

  return
}






