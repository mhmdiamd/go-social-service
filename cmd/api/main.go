package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mhmdiamd/go-social-service/domain/auth"
	"github.com/mhmdiamd/go-social-service/external/database"
	"github.com/mhmdiamd/go-social-service/internal/config"
)

func main() {
  filename := "./cmd/api/config.yaml"
  if err := config.LoadConfig(filename); err != nil {
    panic(err)
  }

  db, err := database.ConnectPostgres(config.Cfg.Db);
  if err != nil {
    panic(err)
  }

  if db != nil {
    log.Print("db connected")
  }

  router := fiber.New(fiber.Config{
    Prefork: true,
    AppName: config.Cfg.App.Name,
  })

  // router.Use(infrafiber.)

  auth.Init(router, db)

  router.Listen(config.Cfg.App.Port)
}
