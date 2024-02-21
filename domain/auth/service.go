package auth

import (
	"context"

	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
)

type Repository interface {
  GetAuthByEmail(ctx context.Context, email string) (model AuthEntity, err error)
  CreateAuth(ctx context.Context, model AuthEntity) (err error) 
}

type service struct {
  repo Repository
}

func newService(repo Repository) service {
  return service{
    repo: repo,
  }
}

func (s service) register(ctx context.Context, req RegisterRequestPayload) (err error) {

  authEntity := NewAuthEntityFromRegister(req)

  // Validation
  if err = authEntity.Validate(); err != nil {
    return 
  }

  // Password encryption 
  if err = authEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
    return
  }

  // Check is account exists
  model, err := s.repo.GetAuthByEmail(ctx, authEntity.Email)
  if err != nil {
    return
  }

  if model.IsExists() {
    return response.ErrEmailAlreadyUsed
  }

  // Execute Serivce
  return s.repo.CreateAuth(ctx, authEntity)
}

func (s service) login(ctx context.Context, req LoginRequestPayload) (token string, err error) {
  authEntity := NewAuthEntityFromLogin(req)

  // Validation
  if err = authEntity.Validate(); err != nil {
    return 
  }

  // Check is account exists
  model, err := s.repo.GetAuthByEmail(ctx, req.Email)
  if err != nil {
    return
  }

  // Validation Password
  if err = authEntity.VerifyPasswordFromPlain(model.Password); err != nil {
    err = response.ErrPasswordNotMatch
    return
  }

  token, err = model.GenerateToken(config.Cfg.App.Encryption.JWTSecret)

  return 
}

