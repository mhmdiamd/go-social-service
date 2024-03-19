package categorycomunity

import (
	"context"
	"testing"

	"github.com/mhmdiamd/go-social-service/external/database"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
	"github.com/stretchr/testify/require"
)

var svc service

func init() {
  filename := "../../cmd/api/config.yaml"
  err := config.LoadConfig(filename)
  if err != nil {
    panic(err)
  }

  db, err := database.ConnectPostgres(config.Cfg.Db)
  if err != nil {
    panic(err)
  }

  repo := newRepository(db)
  svc = newService(repo)
}

func Test_CreateCategoryCommunity(t *testing.T) {
  t.Run("success", func(t *testing.T) {
    req := CreateCategoryCommunityRequestPayload{
      Name: "Technology",
    }

    err := svc.AddCategoryCommunity(context.Background(), req); 
    require.Nil(t, err)
  })

  t.Run("fail, name required", func(t *testing.T) {
    req := CreateCategoryCommunityRequestPayload{
      Name: "",
    }

    err := svc.AddCategoryCommunity(context.Background(), req); 
    require.NotNil(t, err)
    require.Equal(t, response.ErrNameRequired, err)
  })
}
