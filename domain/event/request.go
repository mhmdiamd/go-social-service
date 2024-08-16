package event

import (
	"time"

	"github.com/google/uuid"
)

type CreateEventCommitteeRequestPayload struct {
	UserPublicId  uuid.UUID
	EventPublicId uuid.UUID
	Position      Position
}

type CreateEventRequestPayload struct {
	Name                string    `json:"name"`
	UserPublicId        uuid.UUID `json:"user_public_id"`
	EventDemographicsId int       `json:"event_demographic_id"`
	Description         *string   `json:"description"`
	Address             *string   `json:"address"`
	Thumbnail           string    `json:"thumbnail"`
	StartDate           time.Time `json:"start_at"`
	EndDate             time.Time `json:"end_at"`
}

type UpdateEventRequestPayload struct {
	PublicId            uuid.UUID `json:"event_public_id"`
	EventDemographicsId int       `json:"event_demographic_id"`
	Name                string    `json:"name"`
	Description         *string   `json:"description"`
	Address             *string   `json:"address"`
	Thumbnail           string    `json:"thumbnail"`
	StartDate           time.Time `json:"start_at"`
	EndDate             time.Time `json:"start_at"`
}

type ListEventRequestPayload struct {
	Cursor int `query:"cursor" json:"cursor"`
	Size   int `query:"size" json:"size"`
}

func (l ListEventRequestPayload) GenerateDefaultValue() ListEventRequestPayload {
	if l.Cursor <= 0 {
		l.Cursor = 0
	}

	if l.Size <= 0 {
		l.Size = 10
	}

	return l
}
