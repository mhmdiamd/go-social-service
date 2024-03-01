package categorycomunity

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
  repo := newRepository(db)
  svc := newService(repo)
  handler := newHandler(svc)

  authRoute := router.Group("category-community")
  {
    authRoute.Post("", handler.CreateCategoryCommunity)
  }
}
