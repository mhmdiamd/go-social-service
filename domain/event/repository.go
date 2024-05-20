package event

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository {
	return repository{
		Db: db,
	}
}

func (r repository) Begin(ctx context.Context) (tx *sqlx.Tx, err error) {
	return r.Db.BeginTxx(ctx, &sql.TxOptions{})
}

func (r repository) Commit(ctx context.Context, tx *sqlx.Tx) (err error) {
	return tx.Commit()
}

func (r repository) Rollback(ctx context.Context, tx *sqlx.Tx) (err error) {
	return tx.Rollback()
}

func (r repository) Create(ctx context.Context, tx *sqlx.Tx, event Event) (err error) {
	query := `INSERT INTO events (
    public_id, event_demographics_id, name, thumbnail, start_at, end_at, event_demographics_snapshot, created_at, updated_at
  ) VALUES (
    :public_id, :event_demographics_id, :name, :thumbnail, :start_at, :end_at, :event_demographics_snapshot, :created_at, :updated_at
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
    SET name=:name, description=:description, address=:address, thumbnail=:thumbnail, event_demographics_id=:event_demographics_id, event_demographics_snapshot=:event_demographics_snapshot, start_at=:start_at, end_at=:end_at
    WHERE public_id=:public_id
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

	query := `DELETE FROM events WHERE public_id=$1`
	_, err = r.Db.QueryContext(ctx, query, public_id)
	if err != nil {
		return
	}

	return
}

func (r repository) GetDetailById(ctx context.Context, public_id string) (event Event, err error) {

	query := `
    SELECT
      e.public_id AS "public_id", 
      e.event_demographics_id AS "event_demographics_id", 
      e.name AS "name", 
      e.description AS "description", 
      e.address AS "address", 
      e.start_at AS "start_at", 
      e.end_at AS "end_at", 
      e.created_at AS "created_at", 
      e.updated_at AS "updated_at",
      ed.id AS "event_demographics.id",
      ed.name AS "event_demographics.name",
      ed.gender AS "event_demographics.gender",
      ed.graduation AS "event_demographics.graduation",
      ed.start_age AS "event_demographics.start_age",
      ed.end_age AS "event_demographics.end_age"
    FROM events e
    LEFT JOIN event_demographics ed ON ed.id = e.event_demographics_id
    WHERE e.public_id=$1
  `

	err = r.Db.GetContext(ctx, &event, query, public_id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
		}
	}

	return
}

func (r repository) GetAllWithPagination(ctx context.Context, pagination ListEventRequestPayload) (events []Event, err error) {

	query := `
    SELECT 
      e.public_id AS "public_id", 
      e.event_demographics_id AS "event_demographics_id", 
      e.name AS "name", 
      e.description AS "description", 
      e.address AS "address", 
      e.start_at AS "start_at", 
      e.end_at AS "end_at", 
      e.created_at AS "created_at", 
      e.updated_at AS "updated_at",
      ed.id AS "event_demographics.id",
      ed.name AS "event_demographics.name",
      ed.gender AS "event_demographics.gender",
      ed.graduation AS "event_demographics.graduation",
      ed.start_age AS "event_demographics.start_age",
      ed.end_age AS "event_demographics.end_age"
    FROM events e
    LEFT JOIN event_demographics ed ON ed.id = e.event_demographics_id
    WHERE e.id>$1
    ORDER BY e.id ASC
    LIMIT $2
  `

	err = r.Db.SelectContext(ctx, &events, query, pagination.Cursor, pagination.Size)
	if err != nil {
		return
	}

	return
}

func (r repository) GetEventDemographicsById(ctx context.Context, eventDemographicsId int) (ed EventDemographics, err error) {

  query := `
    SELECT 
      id, name, gender, graduation, start_age, end_age
    FROM event_demographics
    WHERE id=$1
  `

  err = r.Db.GetContext(ctx, &ed, query, eventDemographicsId)
  if err != nil {
    if err == sql.ErrNoRows {
      err = response.ErrNotFound
    }
  }

  return
}

func (r repository) CreateEventCommite(ctx context.Context, tx *sqlx.Tx, ec EventCommite) (err error) {

	query := `
    INSERT INTO event_commite (
      user_public_id, event_public_id, position, created_at, updated_at
    ) VALUES (
      :user_public_id, :event_public_id, :position, :created_at, :updated_at
    )
  `

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, ec)
	if err != nil {
		return
	}

	defer stmt.Close()

	return
}
