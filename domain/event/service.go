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
  Update(ctx context.Context, public_id uuid.UUID, entity Event) (err error)
  DeleteById(ctx context.Context, public_id uuid.UUID) (err error)

  GetAllWithPagination(ctx context.Context, pagination ListEventRequestPayload) (events []Event, err error)
  GetById(ctx context.Context, public_id uuid.UUID) (events Event, err error)
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
    EventPublicId: event.PublicId.String(),
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

func (e *Event) Update(ctx context.Context, req UpdateEventRequestPayload) (err error) {
  return 
}

func (e *Event) DeleteById(ctx context.Context, UserPublicId string) (err error) {
  return
}






