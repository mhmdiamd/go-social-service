package eventcommite

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/stretchr/testify/require"
)

func Test_EventDemographicsValidate_Success(t *testing.T) {
	payload := CreateEventCommitteeRequestPayload{
		Name: "Cakupan Anak SMK",
	}

	entity := NewEventCommitteeEntity(payload)

	err := entity.Validate()
	require.Nil(t, err)
}

func Test_EventDemographicsValidate_Fail(t *testing.T) {
	t.Run("fail, user public id required", func(t *testing.T) {
		payload := CreateEventCommitteeRequestPayload{
			EventPublicID: uuid.MustParse("be36dc6f-f886-4320-8627-882de71c7eas"),
			Position:      EVENT_DIRECTOR,
		}

		entity := NewEventCommitteeEntity(payload)

		err := entity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrNameRequired, err)
	})

	t.Run("fail, event public id required", func(t *testing.T) {
		payload := CreateEventCommitteeRequestPayload{
			UserPublicID: uuid.MustParse("be36dc6f-f886-4320-8627-882de71c7eas"),
			Position:     EVENT_DIRECTOR,
		}

		entity := NewEventCommitteeEntity(payload)

		err := entity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrStartAgeRequired, err)
	})

	t.Run("fail, start age invalid, to min", func(t *testing.T) {
		payload := CreateEventCommitteeRequestPayload{
			UserPublicID:  uuid.MustParse("be36dc6f-f886-4320-8627-882de71c7eas"),
			EventPublicID: uuid.MustParse("be36dc6f-f886-4320-8627-882de71c7eas"),
		}

		entity := NewEventCommitteeEntity(payload)

		err := entity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrStartAgeToMin, err)
	})
}
