package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := newHandler(svc)

	authRoute := router.Group("auth")
  {
		authRoute.Post("send-otp", handler.SendOtp)
		authRoute.Post("verify-otp", handler.VerifyOtp)
		authRoute.Post("register", handler.Register)
		authRoute.Post("login", handler.login)
	}
}
