package event

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/stretchr/testify/require"
)

func Test_EventCommitteeValidate(t *testing.T) {

	type mockPayload struct {
		Status  string
		Payload EventCommittee
		Err     error
	}

	testMock := map[string]mockPayload{
		"validate user_public_id": {
			Status: "success",
			Payload: EventCommittee{
				UserPublicId:  uuid.New(),
				EventPublicId: uuid.New(),
				Position:      EventPosition_Admin,
			},
			Err: nil,
		},

		"user_public_id is required": {
			Status: "fail",
			Payload: EventCommittee{
				EventPublicId: uuid.New(),
				Position:      EventPosition_Admin,
			},
			Err: response.ErrUserPublicIdRequired,
		},

		"validate event public id": {
			Status: "success",
			Payload: EventCommittee{
				UserPublicId:  uuid.New(),
				EventPublicId: uuid.New(),
				Position:      EventPosition_Admin,
			},
			Err: nil,
		},

		"event_public_id is required": {
			Status: "fail",
			Payload: EventCommittee{
				UserPublicId: uuid.New(),
				Position:     EventPosition_Admin,
			},
			Err: response.ErrEventPublicIdRequired,
		},
	}

	for i, e := range testMock {
		t.Run(fmt.Sprintf("%s, %s", e.Status, i), func(t *testing.T) {
			event := e.Payload
			err := event.Validate()

			if e.Err != nil {
				require.NotNil(t, err)
				require.Equal(t, e.Err, err)
			} else {
				require.Nil(t, err)
			}
		})
	}

}
