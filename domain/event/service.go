package event

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
  Create(ctx context.Context, model Event, pagination ListEventRequestPayload) (err error)
  Update(ctx context.Context, public_id uuid.UUID, model Event) (err error)
  DeleteById(ctx context.Context, public_id uuid.UUID) (err error)

  GetAllWithPagination(ctx context.Context, pagination ListEventRequestPayload) (events []Event, err error)
  GetById(ctx context.Context, public_id uuid.UUID) (events Event, err error)
}

type service struct {
  repo Repository
}

func newService(r Repository) service {
  return service{
    repo: r,
  }
}

