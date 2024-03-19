package event

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type repository struct {
  db *sqlx.DB
}

func (r repository) Create(ctx context.Context, req Event) (err error) {

  return
}

func (r repository) UpdateById(ctx context.Context, public_id uuid.UUID, req Event) (err error) {


  return
}

func (r repository) DeleteById(ctx context.Context, public_id uuid.UUID) (err error) {


  return
}

func (r repository) GetById(ctx context.Context, public_id uuid.UUID) (event Event, err error) {


  return
}

func (r repository) GetAll(ctx context.Context, pagination EventPagination) (events []Event, err error) {


  return
}

