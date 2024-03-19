package eventdemographics

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type repository struct {
  db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
  return repository{
    db: db,
  }
}

func (r repository) Create(ctx context.Context, entity EventDemographicsEntity) (err error) {

  query := `
    INSERT INTO event_demographics(
      name, gender, start_age, end_age, graduation, created_at, updated_at
    ) VALUES (
      :name, :gender, :start_age, :end_age, :graduation, :created_at, :updated_at
    )
  `

  stmt, err := r.db.PrepareNamedContext(ctx, query)
  if err != nil {
    return err
  }

  _, err = stmt.ExecContext(ctx, entity)
  if err != nil {
    return err
  }

  defer stmt.Close()

  return
}

func (r repository) UpdateById(ctx context.Context, model EventDemographicsEntity) (err error) {

  query := `
    UPDATE event_demographics 
    SET name=:name, gender=:gender, start_age=:start_age, end_age=:end_age, graduation=:graduation
    WHERE id=:id
  `
 
  stmt, err := r.db.PrepareNamedContext(ctx, query)
  if err != nil {
    return
  } 

  _, err = stmt.ExecContext(ctx, model)
  if err != nil {
    return
  }

  defer stmt.Close()

  return
}

func (r repository) GetAll(ctx context.Context) (models []EventDemographicsEntity, err error) {
  
  query := `
    SELECT 
      id, name, gender, graduation, start_age, end_age, created_at, updated_at
    FROM event_demographics 
  `

  err = r.db.SelectContext(ctx, &models, query)

  if err != nil {
    if err == sql.ErrNoRows {
      return nil, err
    }

    return 
  }

  return
}

func (r repository) GetById(ctx context.Context, id int) (model EventDemographicsEntity, err error){

  query := `
    SELECT 
      id, name, gender, graduation, start_age, end_age, created_at, updated_at
    FROM event_demographics
    WHERE id=$1
  `

  err = r.db.GetContext(ctx, &model, query, id)
  if err != nil {
    if err == sql.ErrNoRows {
      err = response.ErrNotFound
    }

    return
  }

  return
}

func (r repository) DeleteById(ctx context.Context, id int) (err error){

  query := `DELETE FROM event_demographics WHERE id=$1`

  _, err = r.db.QueryContext(ctx, query, id)
  if err != nil {
    return
  }


  return
}


