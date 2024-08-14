package event

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	eventRoute := router.Group("event")
	{
		eventRoute.Get("", handler.GetAll)
		eventRoute.Get("/:id", handler.GetById)
		eventRoute.Post("", handler.Create)
		eventRoute.Put("/:id", handler.UpdateById)
		eventRoute.Delete("/:id", handler.DeleteById)
	}
}
