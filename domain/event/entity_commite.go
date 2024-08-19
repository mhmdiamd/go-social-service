package event

import (
	"time"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/domain/auth"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type Position string

const (
	EventCommitteePosition_Director Position = "director"
	EventPosition_Admin             Position = "admin"
	EventPosition_Staff             Position = "staff"
	EventCommitteePosition_Member   Position = "member"
)

type EventCommittee struct {
	Id            int       `db:"-" json:"id"`
	UserPublicId  uuid.UUID `db:"user_public_id" json:"user_public_id"`
	EventPublicId uuid.UUID `db:"event_public_id" json:"event_public_id"`
	Position      Position  `db:"position" json:"position"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

func NewEventCommitteeFromCreate(req CreateEventCommitteeRequestPayload) EventCommittee {
	ec := EventCommittee{
		UserPublicId:  req.UserPublicId,
		EventPublicId: req.EventPublicId,
		Position:      EventCommitteePosition_Member,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if req.Position != "" {
		ec.Position = req.Position
	}

	return ec
}

func (ec *EventCommittee) Validate() (err error) {
	if err = ec.ValidateUserPublicId(); err != nil {
		return
	}

	if err = ec.ValidateEventPublicId(); err != nil {
		return
	}

	return
}

func (ec *EventCommittee) ValidateUserPublicId() (err error) {
	if ec.UserPublicId.String() == "" {
		return response.ErrUserPublicIdRequired
	}

	return
}

func (ec *EventCommittee) ValidateEventPublicId() (err error) {
	if ec.EventPublicId.String() == "" {
		return response.ErrEventPublicIdRequired
	}

	return
}

func (ec *EventCommittee) IsMember() bool {
	return ec.Position == EventCommitteePosition_Member
}

func (e *EventCommittee) NewEventCommitteeRepsonse(user auth.AuthEntity) EventCommitteeResponse {
	return EventCommitteeResponse{
		Id:            e.Id,
		UserPublicId:  e.UserPublicId,
		EventPublicId: e.EventPublicId,
		Position:      e.Position,
		User: EventUserEntityResponse{
			Name:  user.Name,
			Email: user.Email,
		},
	}
}
