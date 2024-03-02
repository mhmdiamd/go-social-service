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

  return
}

func (r repository) UpdateById(ctx context.Context, model EventDemographicsEntity) (entity EventDemographicsEntity, err error) {

  query := `
    UPDATE event_demographics 
      SET name=$1, gender=$2, start_age=$3, end_age=$4, graduation=$5
    WHERE id=$6
  `
 
  err = r.db.SelectContext(ctx, &entity, query, model.Name, model.Gender, model.StartAge, model.EndAge, model.Graduation, model.Id)
  if err != nil {
    return
  } 

  return
}

func (r repository) GetAll(ctx context.Context) (models []EventDemographicsEntity, err error) {
  
  query := `
    SELECT * FROM event_demographics LIMIT 10
  `

  err = r.db.GetContext(ctx, &models, query)
  if err != nil {
    if err == sql.ErrNoRows {
      return []EventDemographicsEntity{}, err
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
  }

  return
}

func (r repository) DeleteById(ctx context.Context, id int) (err error){

  query := `
    DELETE FROM event_demographics WHERE id=$1
  `

  stmt, err := r.db.PrepareNamedContext(ctx, query)
  if err != nil {
    return
  }

  _, err = stmt.ExecContext(ctx, id)
  if err != nil {
    return
  }

  return
}


