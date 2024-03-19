package community

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-social-service/external/google"
	infrafiber "github.com/mhmdiamd/go-social-service/infra/fiber"
)

func Init(router fiber.Router, db *sqlx.DB){

  con, err := google.ConnectServiceGoogleDrive()
  if err != nil {
    panic(err)
  }

  googleDriveService := google.NewGoogleDriveService(con)

  repo := newRepository(db)
  svc := newService(repo, googleDriveService)
  handler := newHandler(svc)


  communityRoute := router.Group("community")
  {

    communityRoute.Get("", handler.GetAll)

    // Authorization middleware
    communityRoute.Use(infrafiber.CheckAuth())

    communityRoute.Get("/:id", handler.GetById)
    communityRoute.Post("", handler.Create)
    communityRoute.Put("/:id", handler.UpdateById)
    communityRoute.Delete("/:id", handler.DeleteById)
  }

} 
