package auth

import (
	"context"

	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
)

type Repository interface {
  GetAuthByEmail(ctx context.Context, email string) (model AuthEntity, err error)
  CreateUser(ctx context.Context, mode AuthEntity) (err error) 
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
  if err = authEntity.EncryptedPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
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
  return s.repo.CreateUser(ctx, model)
}

// func (s service) login(ctx context.Context, req LoginRequestPayload) (model AuthEntity, err error) {\
//
//
//   return
// }

