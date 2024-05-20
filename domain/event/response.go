package event

import (
	"time"
)

type EventUserEntityResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type EventResponse struct {
	Id          int       `json:"id"`
	PublicId    string    `json:"public_id"`
	EventDemographicsId    int    `json:"event_demographics_id"`
	Name        string    `json:"name"`
	Description *string    `json:"description"`
	Address     *string    `json:"address"`
	Thumbnail   string    `json:"thumbnail"`
  EventDemographics *EventDemographics `json:"event_demographic"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type EventCommiteResponse struct {
	Id            int                     `json:"id"`
	UserPublicId  string                  `json:"user_public_id"`
	EventPublicId string                  `json:"event_public_id"`
	Position      Position                `json:"position"`
	User          EventUserEntityResponse `json:"user"`
}

func ConvertToEventResponseList(events []Event) []EventResponse {

  var el []EventResponse

	for _, e := range events {
		er := e.ConvertToEventResponse()
		el = append(el, er)
	}

	return el
}
