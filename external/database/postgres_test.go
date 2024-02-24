package database

import (
	"testing"

	"github.com/mhmdiamd/go-social-service/internal/config"
	"github.com/stretchr/testify/require"


)

func init() {
  filename := "../../cmd/api/config.yaml"
  err := config.LoadConfig(filename)

  if err != nil {
    panic(err)
  }
}

func TestConnectionPostgres(t *testing.T) {
  t.Run("success", func(t *testing.T) {
    db, err := ConnectPostgres(config.Cfg.Db)

    require.NotNil(t, db)
    require.Nil(t, err)
  })
}
