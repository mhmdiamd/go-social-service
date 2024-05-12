package event

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type repository struct {
  Db *sqlx.DB
}

func (r repository) Create(ctx context.Context, tx *sqlx.Tx, event Event) (err error) {

  query := `INSERT INTO events (
    id, name, thumbnail, start_at, end_at, created_at, end_at
  ) VALUES (
    :id, :name, :thumbnail, :start_at, :end_at, :created_at, end_at
  )`

  stmt, err := tx.PrepareNamedContext(ctx, query)
  if err != nil {
    return 
  }

  _, err = stmt.ExecContext(ctx, event)
  if err != nil {
    return
  }

  defer stmt.Close()

  return
}

func (r repository) UpdateById(ctx context.Context, tx *sqlx.Tx, event Event) (err error) {
  
  query := `
    UPDATE events 
    SET (
      name=:name, description=:description, address=:address, thummbnail=:thumbnail, event_demographics_id=:event_demographics_id, start_at=:start_at, end_at=:end_at
    )
  `

  stmt, err := tx.PrepareNamedContext(ctx, query)
  if err != nil {
    return
  }

  _, err = stmt.ExecContext(ctx, event)
  if err != nil {
    return
  }

  defer stmt.Close()

  return
}

func (r repository) DeleteById(ctx context.Context, public_id string) (err error) {

  query := `DELETE FROM events WHERE id=$1`
  _, err = r.Db.QueryContext(ctx, query, public_id)
  if err != nil {
    return
  }

  return
}

func (r repository) GetDetailById(ctx context.Context, public_id string) (event Event, err error) {



  return
}

func (r repository) GetAll(ctx context.Context, pagination EventPagination) (events []Event, err error) {


  return
}

