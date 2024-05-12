package event

import (
	"time"

)

type CreateEventCommiteRequestPayload struct {
  UserPublicId string 
  EventPublicId string
  Position Position
}
 
type CreateEventRequestPayload struct {
  Name string `json:"name"`
  UserPublicId string `json:"user_public_id"` 
  Description string `json:"description"`
  Address string `json:"address"`
  Thumbnail string `json:"thumbnail"`
  StartDate time.Time `json:"start_date"`
  EndDate time.Time `json:"end_date"`
}


type UpdateEventRequestPayload struct {
  PublicId string `json:"event_public_id"`
  Name string `json:"name"`
  Description string `json:"description"`
  Address string `json:"address"`
  Thumbnail string `json:"thumbnail"`
  StartDate time.Time `json:"start_date"`
  EndDate time.Time `json:"end_date"`
}

type ListEventRequestPayload struct {
  Cursor int `query:"cursor" json:"cursor"`
  Size int `query:"size" json:"size"`
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

