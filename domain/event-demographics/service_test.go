package eventdemographics

import (
	"context"
	"testing"

	"github.com/mhmdiamd/go-social-service/external/database"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
	"github.com/stretchr/testify/require"
)

var svc service
var CURRENT_ID_EVENT_DEMOGRAPHICS int

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

func Test_CreateEventDemographics(t *testing.T) {
  t.Run("success", func(t *testing.T) {

    req := CreateEventDemographicsRequestPayload{
      Name: "event name from test",
      Gender: "male",
      StartAge: 1,
      EndAge: 50,
      Graduation : []string{"SMK", "SMP"},
    }

    err := svc.AddEventDemographics(context.Background(), req)
    require.Nil(t, err)

  })

  t.Run("fail, name is required", func(t *testing.T) {
    req := CreateEventDemographicsRequestPayload{
      Name: "",
      Gender: "male",
      StartAge: 1,
      EndAge: 50,
      Graduation : []string{"SMK"},
    }

    err := svc.AddEventDemographics(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrNameRequired, err)
  })

  t.Run("fail, gender is invalid", func(t *testing.T) {
    req := CreateEventDemographicsRequestPayload{
      Name: "example from test",
      Gender: "asdasda",
      StartAge: 1,
      EndAge: 50,
      Graduation : []string{"SMK"},
    }

    err := svc.AddEventDemographics(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrGenderInvalid, err)
  })
}

func Test_GetById_Success(t *testing.T) {
  model, err := svc.GetById(context.Background(), 21)

  require.Nil(t, err)
  require.NotNil(t, model, err)
}

func Test_UpatedEventDemographics(t *testing.T) {
  t.Run("success", func(t *testing.T) {

    req := UpdateEventDemographicsRequestPayload{
      Id: 21,
      Name: "event name from test",
      Gender: "male",
      StartAge: 1,
      EndAge: 50,
      Graduation : []string{"SMK", "SMP", "S1"},
    }

    model, err := svc.UpdateById(context.Background(), req)
    require.Nil(t, err)
    require.NotNil(t, model)

  })

  t.Run("fail, name is required", func(t *testing.T) {
    req := UpdateEventDemographicsRequestPayload{
      Name: "",
      Gender: "male",
      StartAge: 1,
      EndAge: 50,
      Graduation : []string{"SMK"},
    }

    _, err := svc.UpdateById(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrNameRequired, err)
  })

  t.Run("fail, gender is invalid", func(t *testing.T) {
    req := UpdateEventDemographicsRequestPayload{
      Name: "example from test",
      Gender: "asdasda",
      StartAge: 1,
      EndAge: 50,
      Graduation : []string{"SMK"},
    }

    _, err := svc.UpdateById(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrGenderInvalid, err)
  })
}
