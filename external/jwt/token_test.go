package jwt

import (
	"fmt"
	"testing"

	"github.com/mhmdiamd/go-social-service/domain/auth"
	"github.com/mhmdiamd/go-social-service/internal/config"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		model := auth.RegisterRequestPayload{
			Name:     "am@gmail.com",
			Password: "password",
		}
		authEntity := auth.NewAuthEntityFromRegister(model)

		err := authEntity.EncryptPassword(model.Password, int(config.Cfg.App.Encryption.Salt))

		require.Nil(t, err)
		fmt.Println(authEntity.Password)
	})
}
