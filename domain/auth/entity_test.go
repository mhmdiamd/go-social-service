package auth

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/stretchr/testify/require"
)

func Test_ValidateName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		request := RegisterRequestPayload{
			Name:                 "Muhamad Ilham",
			Password:             "password",
			PasswordConfirmation: "password",
			PublicIdUserOtp:      uuid.New(),
		}

		authData := NewAuthEntityFromRegister(request)

		err := authData.Validate()
		require.Nil(t, err)
	})

	t.Run("error, required name", func(t *testing.T) {
		request := RegisterRequestPayload{
			Name:                 "",
			Password:             "password",
			PasswordConfirmation: "password",
			PublicIdUserOtp:      uuid.New(),
		}

		authData := NewAuthEntityFromRegister(request)

		err := authData.Validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrNameRequired, err)
	})

	t.Run("error, name invalid", func(t *testing.T) {
		request := RegisterRequestPayload{
			Name:                 "asd",
			Password:             "password",
			PasswordConfirmation: "password",
			PublicIdUserOtp:      uuid.New(),
		}

		authData := NewAuthEntityFromRegister(request)

		err := authData.Validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrNameInvalid, err)
	})
}

func Test_ValidatePassword(t *testing.T) {
	t.Run("error, password required", func(t *testing.T) {
		request := RegisterRequestPayload{
			Name:                 "Muhamad Ilham",
			Password:             "",
			PasswordConfirmation: "password",
			PublicIdUserOtp:      uuid.New(),
		}

		authData := NewAuthEntityFromRegister(request)

		err := authData.Validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordRequired, err)
	})

	t.Run("error, password invalid", func(t *testing.T) {
		request := RegisterRequestPayload{
			Name:                 "Muhamad Ilham",
			Password:             "passw",
			PasswordConfirmation: "password",
			PublicIdUserOtp:      uuid.New(),
		}

		authData := NewAuthEntityFromRegister(request)

		err := authData.Validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordInvalid, err)
	})

	t.Run("error, password confirmation required", func(t *testing.T) {
		request := RegisterRequestPayload{
			Name:                 "Muhamad Ilham",
			Password:             "password",
			PasswordConfirmation: "",
			PublicIdUserOtp:      uuid.New(),
		}

		authData := NewAuthEntityFromRegister(request)

		err := authData.Validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordConfirmationRequired, err)
	})

	t.Run("error, password confirmation invalid", func(t *testing.T) {
		request := RegisterRequestPayload{
			Name:                 "Muhamad Ilham",
			Password:             "password",
			PasswordConfirmation: "passw",
			PublicIdUserOtp:      uuid.New(),
		}

		authData := NewAuthEntityFromRegister(request)

		err := authData.Validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordConfirmationInvalid, err)
	})

	t.Run("error, password confirmation not match", func(t *testing.T) {
		request := RegisterRequestPayload{
			Name:                 "Muhamad Ilham",
			Password:             "password",
			PasswordConfirmation: "password23123",
			PublicIdUserOtp:      uuid.New(),
		}

		authData := NewAuthEntityFromRegister(request)

		err := authData.Validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordConfirmationNotMatch, err)
	})
}

func Test_ValidatePublicIdUserOtp(t *testing.T) {
	t.Run("error, public id user otp is required", func(t *testing.T) {
		request := RegisterRequestPayload{
			Name:                 "Muhamad Ilham",
			Password:             "password",
			PasswordConfirmation: "password",
		}

		authData := NewAuthEntityFromRegister(request)

		err := authData.Validate()
		require.NotNil(t, response.ErrPublicIdUserOtpRequired, err)
	})
}
