package event

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/external/database"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
	tempdata "github.com/mhmdiamd/go-social-service/temp_data"
	"github.com/stretchr/testify/require"
)

var (
	svc       service
	currentId string
)

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	// Connect to database
	db, err := database.ConnectPostgres(config.Cfg.Db)
	if err != nil {
		panic(err)
	}

	repo := NewRepository(db)
	svc = NewService(repo)
}

func Test_CreateEvent(t *testing.T) {
	type mock struct {
		Status  string
		Payload CreateEventRequestPayload
		Err     error
	}

	currentId = uuid.NewString()

	dataMock := map[string]mock{
		"create new event": {
			Status: "success",
			Payload: CreateEventRequestPayload{
				UserPublicId:        currentId,
				EventDemographicsId: 19,
				Name:                "Tes event 1",
				Thumbnail:           "default.jpg",
				StartDate:           time.Now(),
				EndDate:             time.Now(),
			},
			Err: nil,
		},

		"name id is required": {
			Status: "fail",
			Payload: CreateEventRequestPayload{
				UserPublicId:        uuid.NewString(),
				EventDemographicsId: 19,
				Thumbnail:           "default.jpg",
				StartDate:           time.Now(),
				EndDate:             time.Now(),
			},
			Err: response.ErrNameRequired,
		},
	}

	for i, e := range dataMock {
		t.Run(fmt.Sprintf("%s, %s", e.Status, i), func(t *testing.T) {
			err := svc.Create(context.Background(), e.Payload)

			if e.Err != nil {
				require.NotNil(t, err)
				require.Equal(t, e.Err, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func Test_UpdateEventById(t *testing.T) {
	type mock struct {
		Status  string
		Payload UpdateEventRequestPayload
		Err     error
	}

	dataMock := map[string]mock{
		"update new event": {
			Status: "success",
			Payload: UpdateEventRequestPayload{
				PublicId:            tempdata.TempCurrentEventPublicId,
				EventDemographicsId: 19,
				Name:                "Tes event 1",
				Thumbnail:           "default.jpg",
				StartDate:           time.Now(),
				EndDate:             time.Now(),
			},
			Err: nil,
		},

		"update new event with different event demographics": {
			Status: "success",
			Payload: UpdateEventRequestPayload{
				PublicId:            tempdata.TempCurrentEventPublicId,
				EventDemographicsId: 52,
				Name:                "Tes event 1",
				Thumbnail:           "default.jpg",
				StartDate:           time.Now(),
				EndDate:             time.Now(),
			},
			Err: nil,
		},

		"name id is required": {
			Status: "fail",
			Payload: UpdateEventRequestPayload{
				PublicId:  currentId,
				Thumbnail: "default.jpg",
				StartDate: time.Now(),
				EndDate:   time.Now(),
			},
			Err: response.ErrNameRequired,
		},
	}

	for i, e := range dataMock {
		t.Run(fmt.Sprintf("%s, %s", e.Status, i), func(t *testing.T) {
			err := svc.UpdateById(context.Background(), e.Payload)

			if e.Err != nil {
				require.NotNil(t, err)
				require.Equal(t, e.Err, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func Test_GetListWithPaginationEvent(t *testing.T) {
	t.Run("success, get list with pagination", func(t *testing.T) {
		req := ListEventRequestPayload{
			Size:   20,
			Cursor: 30,
		}

		events, err := svc.GetAllWithPagination(context.Background(), req)
		require.Nil(t, err)
		require.NotNil(t, events)
	})
}

func Test_GetDetailEventById(t *testing.T) {
	t.Run("success, get detail by id", func(t *testing.T) {
		event, err := svc.GetDetailById(context.Background(), tempdata.TempCurrentEventPublicId)
		require.NotNil(t, event)
		require.Nil(t, err)
	})

	t.Run("fail, event not found", func(t *testing.T) {
		_, err := svc.GetDetailById(context.Background(), uuid.NewString())
		require.NotNil(t, err)
		require.Equal(t, response.ErrNotFound, err)
	})
}

func Test_DeleteById(t *testing.T) {
	t.Run("success, delete by id", func(t *testing.T) {
		err := svc.DeleteById(context.Background(), tempdata.TempCurrentEventPublicId)
		require.Nil(t, err)
	})

	t.Run("fail, id not found", func(t *testing.T) {
		err := svc.DeleteById(context.Background(), tempdata.TempCurrentEventPublicId)
		require.NotNil(t, err)
	})
}
