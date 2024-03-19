package eventdemographics

import (
	"testing"

	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/stretchr/testify/require"
)

func Test_EventDemographicsValidate_Success(t *testing.T) {

    payload := CreateEventDemographicsRequestPayload{
      Name: "Cakupan Anak SMK",
      Gender : "",
      Graduation: []string{"SMK"},
      StartAge: 1,
      EndAge: 50,
    }

    entity := NewEventDemographicsEntity(payload)

    err := entity.Validate()
    require.Nil(t, err)
}

func Test_EventDemographicsValidate_Fail(t *testing.T) {
  t.Run("fail, name required", func(t *testing.T) {
    payload := CreateEventDemographicsRequestPayload{
      Name: "",
      Graduation: []string{"SMK"},
    }

    entity := NewEventDemographicsEntity(payload)

    err := entity.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrNameRequired, err)
  })

  t.Run("fail, start age required", func(t *testing.T) {
    payload := CreateEventDemographicsRequestPayload{
      Name: "Ilham",
      Graduation: []string{"SMK"},
      StartAge: 0,
    }

    entity := NewEventDemographicsEntity(payload)

    err := entity.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrStartAgeRequired, err)
  })

  t.Run("fail, start age invalid, to min", func(t *testing.T) {
    payload := CreateEventDemographicsRequestPayload{
      Name: "Ilham",
      Graduation: []string{"SMK"},
      StartAge : -1,
    }

    entity := NewEventDemographicsEntity(payload)

    err := entity.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrStartAgeToMin, err)
  })

  t.Run("fail, start age invalid, to max", func(t *testing.T) {
    payload := CreateEventDemographicsRequestPayload{
      Name: "Ilham",
      Graduation: []string{"SMK"},
      StartAge : 100,
    }

    entity := NewEventDemographicsEntity(payload)

    err := entity.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrStartAgeToMax, err)
  })

  t.Run("fail, start age greater then end age", func(t *testing.T) {
    payload := CreateEventDemographicsRequestPayload{
      Name: "Ilham",
      Graduation: []string{"SMK"},
      StartAge : 80,
      EndAge : 50,
    }

    entity := NewEventDemographicsEntity(payload)

    err := entity.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrStartAgeGreaterThenEndAge, err)
  })

  t.Run("fail, end age invalid, to max", func(t *testing.T) {
    payload := CreateEventDemographicsRequestPayload{
      Name: "Ilham",
      Graduation: []string{"SMK"},
      StartAge : 1,
      EndAge: 100,
    }

    entity := NewEventDemographicsEntity(payload)

    err := entity.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrEndAgeToMax, err)
  })
}
