package communitymember

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
  repo := newRepository(db)
  svc := newService(repo)
  _ = newHandler(svc)

  // communityMemberRouter := router.Use("/community-member", handler)
  {
  }
}
