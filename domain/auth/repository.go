package auth

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
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
      public_id, name, email, password, user_otp_public_id, created_at, updated_at
    ) VALUES (
      :public_id, :name, :email, :password, :user_otp_public_id, :created_at, :updated_at
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
     public_id, otp, email, is_active, expired_at, updated_at, created_at
    ) VALUES (
      :public_id, :otp, :email, :is_active, :expired_at, :updated_at, :created_at
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

func (r repository) GetOtpByEmail(ctx context.Context, email string) (models []OtpEntity, err error) {

  query := `
    SELECT 
      id, public_id, otp, email, is_active, expired_at, created_at, updated_at
    FROM user_otp WHERE email=$1
  `

  err = r.db.SelectContext(ctx, &models, query, email)

  if err != nil {
    if err == sql.ErrNoRows {
      err = response.ErrNotFound
      return nil, err
    }
  }

  return 
}

func (r repository) GetDetailOtp(ctx context.Context, public_id_user_otp uuid.UUID) (model OtpEntity, err error) {

  query := `
    SELECT 
      id, public_id, otp, email, is_active, expired_at, created_at, updated_at
    FROM user_otp
    WHERE public_id=$1
  `

  err = r.db.GetContext(ctx, &model, query, public_id_user_otp)

  if err != nil {
    if err == sql.ErrNoRows {
      err = response.ErrNotFound
      return
    }

    return
  }

  return
}

func (r repository) GetDetailOtpByEmailAndOtp(ctx context.Context, req VerifyOtpRequestPayload) (model OtpEntity, err error) {

  query := `
    SELECT 
      id, public_id, otp, email, is_active, expired_at, created_at, updated_at
    FROM user_otp
    WHERE otp=$1 AND email=$2
  `

  err = r.db.GetContext(ctx, &model, query, req.Otp, req.Email)

  if err != nil {
    if err == sql.ErrNoRows {
      err = response.ErrNotFound
      return
    }

    return
  }

  return
}

func (r repository) DeleteAuthByEmail(ctx context.Context, email string) (err error) {
  
  query := `
    DELETE FROM auth WHERE email=$1
  `

  _, err = r.db.ExecContext(ctx, query, email)

  if err != nil {
    return
  }


  return 

}

func (r repository) UnactiveOtp(ctx context.Context, email string) (err error) {
  
  query := `
    UPDATE user_otp SET is_active=0 WHERE email=$1
  `

  _, err = r.db.ExecContext(ctx, query, email)

  if err != nil {
    return
  }


  return 

}

