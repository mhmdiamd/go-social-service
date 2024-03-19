package event

import (
	"testing"
	"time"

	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/stretchr/testify/require"
)

func Test_ValidateEventEntity_Create(t *testing.T) {

  t.Run("success", func(t *testing.T) {

      req := CreateEventRequestPayload{
        Name: "Muhamad Ilham",
        Description: "example description",
        Address: "example address",
        Thumbnail: "test.jpg",
        StartDate: time.Now(),
        EndDate: time.Now(),
      }

      entity := NewEventFromCreate(req)

      err := entity.Validate()
      require.Nil(t, err)
  })

  t.Run("fail", func(t *testing.T) {

    t.Run("validate name", func(t *testing.T) {
      req := CreateEventRequestPayload{
        Name: "",
        Description: "example description",
        Address: "example address",
        Thumbnail: "test.jpg",
        StartDate: time.Now(),
        EndDate: time.Now(),
      }

      entity := NewEventFromCreate(req)

      err := entity.Validate()
      require.NotNil(t, err)
      require.Equal(t, response.ErrNameRequired, err)
    })

    t.Run("validate description", func(t *testing.T) {
      req := CreateEventRequestPayload{
        Name: "Muhamad Ilham",
        Description: "",
        Address: "example address",
        Thumbnail: "test.jpg",
        StartDate: time.Now(),
        EndDate: time.Now(),
      }

      entity := NewEventFromCreate(req)

      err := entity.Validate()
      require.NotNil(t, err)
      require.Equal(t, response.ErrDescriptionRequired, err)
    })

    t.Run("validate address", func(t *testing.T) {
      req := CreateEventRequestPayload{
        Name: "Muhamad Ilham",
        Description: "example description",
        Address: "",
        Thumbnail: "test.jpg",
        StartDate: time.Now(),
        EndDate: time.Now(),
      }

      entity := NewEventFromCreate(req)

      err := entity.Validate()
      require.NotNil(t, err)
      require.Equal(t, response.ErrAddressRequired, err)
    })

    t.Run("validate date", func(t *testing.T) {

      t.Run("start date required", func(t *testing.T) {
        req := CreateEventRequestPayload{
          Name: "Muhamad Ilham",
          Description: "example description",
          Address: "example address",
          Thumbnail: "test.jpg",
          EndDate: time.Now(),
        }

        entity := NewEventFromCreate(req)

        err := entity.Validate()
        require.NotNil(t, err)
        require.Equal(t, response.ErrStartDateRequired, err)
      })

      t.Run("end date required", func(t *testing.T) {
        req := CreateEventRequestPayload{
          Name: "Muhamad Ilham",
          Description: "example description",
          Address: "example address",
          Thumbnail: "test.jpg",
          StartDate: time.Now(),
        }

        entity := NewEventFromCreate(req)

        err := entity.Validate()
        require.NotNil(t, err)
        require.Equal(t, response.ErrEndDateRequired, err)
      })

      t.Run("start date greater than end date", func(t *testing.T) {
        req := CreateEventRequestPayload{
          Name: "Muhamad Ilham",
          Description: "example description",
          Address: "example address",
          Thumbnail: "test.jpg",
          StartDate: time.Now().Add(time.Minute * time.Duration(30)),
          EndDate: time.Now(),
        }

        entity := NewEventFromCreate(req)

        err := entity.Validate()
        require.NotNil(t, err)
        require.Equal(t, response.ErrStartDateGreaterThanEndDate, err)
      })

    })

  })
}

func Test_ValidateEventEntity_Update(t *testing.T) {
  t.Run("success", func(t *testing.T) {
    req := UpdateEventRequestPayload{
      Name: "Muhamad Ilham",
      Description: "example description",
      Address: "example address",
      Thumbnail: "test.jpg",
      StartDate: time.Now(),
      EndDate: time.Now(),
    }

    entity := NewEventFromUpdate(req)

    err := entity.Validate()
    require.Nil(t, err)
  })
}
