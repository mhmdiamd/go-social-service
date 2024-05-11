package event

import (
	"time"

	"github.com/mhmdiamd/go-social-service/domain/auth"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type Position string

const (
  EventPosition_Admin Position = "admin"
  EventPosition_Staff Position = "staff"
  EventPosition_Member Position = "member"
)

type EventCommite struct {
  Id int `db:"id"`
  UserPublicId string `db:"user_public_id"`
  EventPublicId string `db:"event_id"`
  Position Position `db:"position"`
  CreatedAt time.Time `db:"created_at" json:"created_at"`
  UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func NewEventCommiteFromCreate(req CreateEventCommiteRequestPayload) EventCommite {
  ec := EventCommite{
    UserPublicId : req.UserPublicId,
    EventPublicId : req.EventPublicId,
    Position: EventPosition_Member,
    CreatedAt : time.Now(),
    UpdatedAt : time.Now(),
  }

  if req.Position != "" {
    ec.Position = req.Position
  }

  return ec
}

func (ec *EventCommite) Validate() (err error) {

  if err = ec.ValidateUserPublicId(); err != nil {
    return
  }

  if err = ec.ValidateEventPublicId(); err != nil {
    return
  }

  return 
}

func (ec *EventCommite) ValidateUserPublicId() (err error) {

  if ec.UserPublicId == "" {
    return response.ErrUserPublicIdRequired
  }

  return 
}

func (ec *EventCommite) ValidateEventPublicId() (err error) {
  if ec.EventPublicId == "" {
    return response.ErrEventPublicIdRequired
  }

  return 
}

func (ec *EventCommite) IsMember() bool {
  return ec.Position == EventPosition_Member
}

func (e *EventCommite) NewEventCommiteRepsonse(user auth.AuthEntity) EventCommiteResponse {
  return EventCommiteResponse{
    Id: e.Id,
    UserPublicId: e.UserPublicId,
    EventPublicId: e.EventPublicId,
    Position: e.Position,
    User: EventUserEntityResponse {
      Name : user.Name,
      Email : user.Email,
    },
  }

}













