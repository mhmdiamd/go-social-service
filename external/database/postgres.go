package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-social-service/internal/config"

  _ "github.com/lib/pq"
)

func ConnectPostgres(cfg config.DBConfig) (db *sqlx.DB, err error) {
  dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
      cfg.Host,
      cfg.Port,
      cfg.User,
      cfg.Password,
      cfg.Name,
  )

  db, err = sqlx.Open("postgres", dsn)
  if err != nil {
    return
  }

  if err = db.Ping(); err != nil {
    return
  }

  return 
}
