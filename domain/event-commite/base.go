package eventcommite

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)
	kafkaReader := NewEventReaderEventCommittee("event-committee", svc)
	go func() {
		if err := kafkaReader.Init(context.Background()); err != nil {
			fmt.Println(err)
		}
	}()

	eventRoute := router.Group("event")

	eventcommiteRoute := eventRoute.Group("event-committee")
	{
		eventcommiteRoute.Get("", handler.GetAll)
		eventcommiteRoute.Get("/:id", handler.GetById)
		eventcommiteRoute.Post("", handler.Create)
		eventcommiteRoute.Put("/:id", handler.UpdateById)
		eventcommiteRoute.Delete("/:id", handler.DeleteById)
	}
}
