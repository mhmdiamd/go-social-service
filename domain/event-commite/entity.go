package eventcommite

import (
	"time"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type EventCommitteePosition string

const (
	EVENT_DIRECTOR  EventCommitteePosition = "EVENT_DIRECTOR"
	EVENT_SECERTARY EventCommitteePosition = "SECERTARY"
)

type EventCommitteeEntity struct {
	Id            int                    `db:"id"`
	UserPublicID  uuid.UUID              `db:"user_public_id"`
	EventPublicID uuid.UUID              `db:"event_public_id"`
	Position      EventCommitteePosition `db:"position"`
	CreatedAt     time.Time              `db:"created_at"`
	UpdatedAt     time.Time              `db:"updated_at"`
}

func NewEventCommitteeEntityFromUpdate(req UpdateEventCommitteeRequestPayload) EventCommitteeEntity {
	entity := EventCommitteeEntity{
		Id:       req.Id,
		Position: req.Position,
	}

	return entity
}

func NewEventCommitteeEntity(req CreateEventCommitteeRequestPayload) EventCommitteeEntity {
	entity := EventCommitteeEntity{
		UserPublicID:  req.UserPublicID,
		EventPublicID: req.EventPublicID,
		Position:      req.Position,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return entity
}

func (e EventCommitteeEntity) Validate() (err error) {
	if err = e.ValidateUserPublicID(); err != nil {
		return
	}

	if err = e.ValidateEventPublicID(); err != nil {
		return
	}

	return
}

func (e EventCommitteeEntity) ValidateUserPublicID() (err error) {
	if e.UserPublicID.String() == "" {
		return response.ErrNameRequired
	}

	return
}

func (e EventCommitteeEntity) ValidateEventPublicID() (err error) {
	if e.EventPublicID.String() == "" {
		return response.ErrNameRequired
	}

	return
}

func (e EventCommitteeEntity) ValidateId() (err error) {
	if e.Id == 0 {
		return response.ErrIdRequired
	}

	return
}

func (e EventCommitteeEntity) ValidatePosition() (err error) {
	if e.Position == "" {
		return response.ErrIdRequired
	}

	return
}
