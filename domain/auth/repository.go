package auth

import (
	"context"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type repository struct {
  db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
  return repository{
    db : db,
  }
}

func (r repository) CreateAuth(ctx context.Context, model AuthEntity) (err error){

  query := `
    INSERT INTO auth (
      public_id, email, password, created_at, updated_at
    ) VALUES (
      :public_id, :email, :password, :created_at, :updated_at
    )
  `

  stmt, err := r.db.PrepareNamedContext(ctx, query)
  if err != nil {
    return
  }

  _, err = stmt.ExecContext(ctx, model)

  defer stmt.Close()

  return
}

func (r repository) GetAuthByEmail(ctx context.Context, email string) (model AuthEntity, err error) {

  query := `
    SELECT 
      id, public_id, email, password, name, gender, no_tlp, address
    FROM auth
    WHERE email=$1
  `

  err = r.db.GetContext(ctx, &model, query, email)
  if err != nil {
    if err == sql.ErrNoRows {
      err = response.ErrNotFound
      return
    }
  }

  return
}

func (r repository) CreateOTP(ctx context.Context, model OtpEntity) (err error) {

  query := `
    INSERT INTO user_otp (
      otp, email, password, is_active, expired_at, updated_at, created_at
    ) VALUES (
      :otp, :email, :password, :is_active, :expired_at, :updated_at, :created_at
    )
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

func (r repository)GetOtpByEmail(ctx context.Context, email string) (models []OtpEntity, err error) {

  log.Println(email)

  query := `SELECT * FROM user_otp WHERE email=$1`

  var otpEntities []OtpEntity

  err = r.db.SelectContext(ctx, &otpEntities, query, email)

  if err != nil {
    if err == sql.ErrNoRows {
      err = response.ErrNotFound
      return nil, err
    }
  }

  return 
}

