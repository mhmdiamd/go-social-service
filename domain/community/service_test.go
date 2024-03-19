package community

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/external/database"
	"github.com/mhmdiamd/go-social-service/external/google"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
	tempdata "github.com/mhmdiamd/go-social-service/temp_data"
	"github.com/stretchr/testify/require"
)

var svc service
var userPublicId = uuid.MustParse("7eaf37cb-1408-459c-ba13-94d3dbcff0aa")

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

  con, err := google.ConnectServiceGoogleDrive()
  if err != nil {
    panic(err)
  }

  googleDriveService := google.NewGoogleDriveService(con)

  repo := newRepository(db)
  svc = newService(repo, googleDriveService)
}

func Test_CreateCommunity_Success(t *testing.T) {
  req := CreateCommunityRequestPayload{
    Name: "name from test",
    CategoryCommunityID: 1,
    UserPublicId: userPublicId,
  }

  err := svc.CreateCommunity(context.Background(), req)
  require.Nil(t, err)
}

func Test_CreateCommunity_Fail(t *testing.T) {
  t.Run("fail, name is required", func(t *testing.T) {
    req := CreateCommunityRequestPayload{
      Name: "",
      CategoryCommunityID: 1,
      UserPublicId: userPublicId,
    }

    err := svc.CreateCommunity(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrNameRequired, err)
  })

  t.Run("fail, category_community_id is required", func(t *testing.T) {
    req := CreateCommunityRequestPayload{
      Name: "test",
      UserPublicId: userPublicId,
    }

    err := svc.CreateCommunity(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrCommunityCategoryIdRequired, err)
  })
}

func Test_UpdateCommunity(t *testing.T) {
  req := UpdateCommunityRequestPayload{
    Id: tempdata.LastCommunityID,
    Name: "renamed name",
    CategoryCommunityID: 1,
  }

  model, err := svc.GetById(context.Background(), tempdata.LastCommunityID)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(tempdata.LastCommunityID)

  err = svc.UpdateById(context.Background(), userPublicId, req)
  require.Nil(t, err)
  require.NotNil(t, model)
}

func Test_GetByIdCommunity(t *testing.T) {
  t.Run("success", func (t *testing.T) {
    community, err := svc.GetById(context.Background(), tempdata.LastCommunityID)
    require.Nil(t, err)
    require.NotNil(t, community)
  })

  t.Run("fail, community not found", func (t *testing.T) {
    _, err := svc.GetById(context.Background(), 9999)
    require.NotNil(t, err)
    require.Equal(t, response.ErrNotFound, err)
  })
}

func Test_GetAllCommunity(t *testing.T) {
  t.Run("success", func (t *testing.T) {
    payloadPagination := ListCommunityRequestPayload{}
    community, err := svc.GetAll(context.Background(), payloadPagination.GenerateDefaultValue())
    require.Nil(t, err)
    require.NotNil(t, community)
  })
}

func Test_DeleteByIdCommunity(t *testing.T) {
  t.Run("success", func (t *testing.T) {
    err := svc.DeleteById(context.Background(), tempdata.LastCommunityID)
    require.Nil(t, err)
  })

  t.Run("fail, community not found", func (t *testing.T) {
    err := svc.DeleteById(context.Background(), 9999)
    require.NotNil(t, err)
    require.Equal(t, response.ErrNotFound, err)
  })
}

