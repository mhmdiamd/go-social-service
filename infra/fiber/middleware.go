package infrafiber

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mhmdiamd/go-social-service/external/jwt"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
)

func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("authorization")

		if authorization == "" {
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		bearer := strings.Split(authorization, "Bearer ")
		if len(bearer) != 2 {
			log.Println("Invalid Token")
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		// Decrypt token
		token := bearer[1]
		publicId, err := jwt.ValidateToken(token, config.Cfg.App.Encryption.JWTSecret)
		if err != nil {
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		c.Locals("PUBLIC_ID", publicId)

		return c.Next()
	}
}
