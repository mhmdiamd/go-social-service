package eventcommittee

import "github.com/google/uuid"

type CreateEventCommitteeRequestPayload struct {
	UserPublicID  uuid.UUID              `json:"user_public_id"`
	EventPublicID uuid.UUID              `json:"event_public_id"`
	Position      EventCommitteePosition `json:"position"`
}

type UpdateEventCommitteeRequestPayload struct {
	Id       int                    `json:"id"`
	Position EventCommitteePosition `json:"position"`
}

type EventDemographicsRequestQuery struct {
	Id int `query:"id"`
}
