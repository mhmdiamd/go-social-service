package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
  repo := newRepository(db)
  svc := newService(repo)
  handler := newHandler(svc)

  authRoute := router.Group("auth")
  {
    authRoute.Post("send-otp", handler.sendOtp)
    authRoute.Post("verify-otp", handler.verifyOtp)
    authRoute.Post("register", handler.register)
    authRoute.Post("login", handler.login)
  }
}
