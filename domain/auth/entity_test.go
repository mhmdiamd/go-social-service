package auth

import (
	"testing"

	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/stretchr/testify/require"
)

func TestValidateEmail(t *testing.T) {
  t.Run("success", func(t *testing.T) {
    request := RegisterRequestPayload{
      Email : "am@gmail.com",
      Password: "123456",
    }

    authData := NewAuthEntityFromRegister(request)
    err := authData.Validate()

    require.Nil(t, err)
  })

  t.Run("error required email", func(t *testing.T) {
    request := RegisterRequestPayload{
      Password: "123456",
    }
    authData := NewAuthEntityFromRegister(request)

    err := authData.Validate()

    require.NotNil(t, err)
    require.Equal(t, response.ErrEmailRequired, err)
  })

  t.Run("error invalid email", func(t *testing.T) {
    request := RegisterRequestPayload{
      Email : "amgmail.com",
      Password: "123456",
    }
    authData := NewAuthEntityFromRegister(request)

    err := authData.Validate()

    require.NotNil(t, err)
    require.Equal(t, response.ErrEmailInvalid, err)
  })
}

func TestValidatePassword(t *testing.T) {
  t.Run("success", func(t *testing.T) {
    request := RegisterRequestPayload{
      Email : "am@gmail.com",
      Password: "123456",
    }
    authData := NewAuthEntityFromRegister(request)

    err := authData.Validate()
    require.Nil(t, err)
  })

  t.Run("error required password", func(t *testing.T) {
    request := RegisterRequestPayload{
      Email : "am@gmail.com",
    }
    authData := NewAuthEntityFromRegister(request)

    err := authData.Validate()

    require.NotNil(t, err)
    require.Equal(t, response.ErrPasswordRequired, err)
  })

  t.Run("error invalid password", func(t *testing.T) {
    request := RegisterRequestPayload{
      Email : "am@gmail.com",
      Password: "12345",
    }
    authData := NewAuthEntityFromRegister(request)

    err := authData.Validate()

    require.NotNil(t, err)
    require.Equal(t, response.ErrPasswordInvalid, err)
  })
}
