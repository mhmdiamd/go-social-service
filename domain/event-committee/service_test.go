package eventcommittee

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/external/database"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
	"github.com/stretchr/testify/require"
)

var (
	svc                        service
	CURRENT_ID_EVENT_COMMITTEE int64
)

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
		req := CreateEventCommitteeRequestPayload{
			UserPublicID:  uuid.MustParse("be36dc6f-f886-4320-8627-882de71c7eas"),
			EventPublicID: uuid.MustParse("be36dc6f-f886-4320-8627-882de71c7eas"),
			Position:      EVENT_DIRECTOR,
		}

		err := svc.AddEventCommittee(context.Background(), req)
		require.Nil(t, err)
	})

	t.Run("fail, user public id is required", func(t *testing.T) {
		req := CreateEventCommitteeRequestPayload{
			EventPublicID: uuid.MustParse("be36dc6f-f886-4320-8627-882de71c7eas"),
			Position:      EVENT_DIRECTOR,
		}

		err := svc.AddEventCommittee(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrNameRequired, err)
	})

	t.Run("fail, event public id is invalid", func(t *testing.T) {
		req := CreateEventCommitteeRequestPayload{
			UserPublicID: uuid.MustParse("be36dc6f-f886-4320-8627-882de71c7eas"),
			Position:     EVENT_DIRECTOR,
		}

		err := svc.AddEventCommittee(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrGenderInvalid, err)
	})

	t.Run("fail, event public id is invalid", func(t *testing.T) {
		req := CreateEventCommitteeRequestPayload{
			UserPublicID:  uuid.MustParse("be36dc6f-f886-4320-8627-882de71c7eas"),
			EventPublicID: uuid.MustParse("be36dc6f-f886-4320-8627-882de71c7eas"),
		}

		err := svc.AddEventCommittee(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrGenderInvalid, err)
	})
}

func Test_GetById_Success(t *testing.T) {
	model, err := svc.GetById(context.Background(), int(CURRENT_ID_EVENT_COMMITTEE))

	require.Nil(t, err)
	require.NotNil(t, model, err)
}

func Test_GetById_Fail(t *testing.T) {
	t.Run("fail, not found", func(t *testing.T) {
		_, err := svc.GetById(context.Background(), 99999)

		require.NotNil(t, err)
		require.Equal(t, response.ErrNotFound, err)
	})
}

func Test_GetAll_Success(t *testing.T) {
	models, err := svc.GetListEventCommittee(context.Background())

	require.Nil(t, err)
	require.NotNil(t, models, err)

	for _, model := range models {
		fmt.Println(model)
	}
}

func Test_UpatedEventDemographics(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := UpdateEventCommitteeRequestPayload{
			Id:       int(CURRENT_ID_EVENT_COMMITTEE),
			Position: EVENT_SECERTARY,
		}

		err := svc.UpdateById(context.Background(), req)
		require.Nil(t, err)
	})

	t.Run("fail, id is required", func(t *testing.T) {
		req := UpdateEventCommitteeRequestPayload{
			Position: EVENT_SECERTARY,
		}

		err := svc.UpdateById(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrNameRequired, err)
	})

	t.Run("fail, position is required", func(t *testing.T) {
		req := UpdateEventCommitteeRequestPayload{
			Id: int(CURRENT_ID_EVENT_COMMITTEE),
		}

		err := svc.UpdateById(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrGenderInvalid, err)
	})
}
