package community

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-social-service/external/google"
	infrafiber "github.com/mhmdiamd/go-social-service/infra/fiber"
)

func Init(router fiber.Router, db *sqlx.DB) {
	con, err := google.ConnectServiceGoogleDrive()
	if err != nil {
		panic(err)
	}

	googleDriveService := google.NewGoogleDriveService(con)

	repo := NewRepository(db)
	svc := NewService(repo, googleDriveService)
	handler := newHandler(svc)

	communityRoute := router.Group("community")
	{

		communityRoute.Get("", handler.GetAll)

		// Authorization middleware
		// communityRoute.Use(infrafiber.CheckAuth())

		communityRoute.Get("/:id", infrafiber.CheckAuth(), handler.GetById)
		communityRoute.Post("", infrafiber.CheckAuth(), handler.Create)
		communityRoute.Put("/:id", infrafiber.CheckAuth(), handler.UpdateById)
		communityRoute.Delete("/:id", infrafiber.CheckAuth(), handler.DeleteById)
	}
}
