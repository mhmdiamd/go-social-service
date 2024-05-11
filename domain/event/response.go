package event

import (
	"time"

	"github.com/google/uuid"
)

type EventUserEntityResponse struct {
  Name string `json:"name"`
  Email string `json:"email"`
}

type EventResponse struct {
  Id string `json:"id"`
  PublicId uuid.UUID `json:"public_id"`
  Name string `json:"name"`
  Description string `json:"description"`
  Address string `json:"address"`
  Thumbnail string `json:"thumbnail"`
  StartDate time.Time `json:"start_date"`
  EndDate time.Time `json:"end_date"`
}

type EventCommiteResponse struct {
  Id int `json:"id"`
  UserPublicId string `json:"user_public_id"`
  EventPublicId string `json:"event_public_id"`
  Position Position `json:"position"`
  User EventUserEntityResponse `json:"user"`
}

func ConvertToEventResponseList(events []Event) (lists []EventResponse) {
  for _, e := range events {
    er := e.ConvertToEventResponse()

    lists = append(lists, er)
  }

  return  
}
