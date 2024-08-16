package event

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	infrafiber "github.com/mhmdiamd/go-social-service/infra/fiber"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	eventRoute := router.Group("event")
	{
		eventRoute.Get("", handler.GetAll)
		eventRoute.Get("/:id", handler.GetById)
		eventRoute.Post("", infrafiber.CheckAuth(), handler.Create)
		eventRoute.Put("/:id", infrafiber.CheckAuth(), handler.UpdateById)
		eventRoute.Delete("/:id", infrafiber.CheckAuth(), handler.DeleteById)
	}
}
