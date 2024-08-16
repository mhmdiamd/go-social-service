package eventcommittee

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

	EventCommitteeRoute := router.Group("event/event-committee")
	{
		EventCommitteeRoute.Get("", handler.GetAll)
		EventCommitteeRoute.Get("/:id", handler.GetById)
		EventCommitteeRoute.Post("", handler.Create)
		EventCommitteeRoute.Put("/:id", handler.UpdateById)
		EventCommitteeRoute.Delete("/:id", handler.DeleteById)
	}
}
