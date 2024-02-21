package auth

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/external/database"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
	"github.com/stretchr/testify/require"
)

var svc service

func init() {
  filename := "../../cmd/api/config.yaml"
  err := config.LoadConfig(filename)

  db, err := database.ConnectPostgres(config.Cfg.Db)
  if err != nil {
    panic(err)
  }

  repo := newRepository(db)
  svc = newService(repo)
}

func TestCreateAuth_Success(t *testing.T) {
  t.Run("success", func (t *testing.T) {
    req := RegisterRequestPayload{
      Email : fmt.Sprintf("%vam@gmail.com", uuid.NewString()),
      Password : "password",
    }

    err := svc.register(context.Background(), req)
    require.Nil(t, err)
  })
}

func TestCreateAuth_Fail(t *testing.T) {
  t.Run("error email is already use", func(t *testing.T) {
    email := fmt.Sprintf("%vam@gmail.com", uuid.NewString())
    req := RegisterRequestPayload{
      Email : email,
      Password : "password",
    }

    err := svc.register(context.Background(), req)
    require.Nil(t, err)

    err = svc.register(context.Background(), req)
    require.Equal(t, response.ErrEmailAlreadyUsed, err)
  })
}

func TestLogin(t *testing.T) {
  email := fmt.Sprintf("%vam@gmail.com", uuid.NewString())

  t.Run("success", func(t *testing.T) {
    req := RegisterRequestPayload{
      Email : email,
      Password : "password",
    }

    err := svc.register(context.Background(), req)
    require.Nil(t, err)

    reqLogin := LoginRequestPayload{
      Email: email,
      Password : "password",
    }

    token, err := svc.login(context.Background(), reqLogin)
    require.Nil(t, err)
    require.NotEmpty(t, token)
  })

  t.Run("error password not match", func(t *testing.T) {
    reqLogin := LoginRequestPayload{
      Email: email,
      Password : "password123",
    }

    _, err := svc.login(context.Background(), reqLogin)
    require.NotNil(t, err)
    require.Equal(t, response.ErrPasswordNotMatch, err)
  })
}


