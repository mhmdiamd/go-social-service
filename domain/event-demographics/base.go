package eventdemographics

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB){

  repo := newRepository(db)
  svc := newService(repo)
  handler := newHandler(svc)


  eventDemographicsRoute := router.Group("event-demographics")
  {
    eventDemographicsRoute.Get("", handler.GetAll)
    eventDemographicsRoute.Get("/:id", handler.GetById)
    eventDemographicsRoute.Post("", handler.Create)
    eventDemographicsRoute.Put("/:id", handler.UpdateById)
    eventDemographicsRoute.Delete("/:id", handler.DeleteById)
  }

} 
