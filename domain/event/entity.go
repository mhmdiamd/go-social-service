package event

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type Event struct {
	Id                int             `db:"-"`
	PublicId          string          `db:"public_id" json:"public_id"`
	Name              string          `db:"name" json:"name"`
	Description       *string          `db:"description" json:"description"`
	Address           *string          `db:"address" json:"address"`
	Thumbnail         string          `db:"thumbnail" json:"thumbnail"`
	EventDemographicsJSON json.RawMessage `db:"event_demographics_snapshot" json:"event_demographics_snapshot"`
	StartDate         time.Time       `db:"start_at" json:"start_at"`
	EndDate           time.Time       `db:"end_at" json:"end_at"`

	EventDemographicsId int `db:"event_demographics_id" json:"event_demographics_id"`
	EventDemographics *EventDemographics `db:"event_demographics"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func NewEventFromCreate(req CreateEventRequestPayload) Event {
	entity := Event{
		PublicId:            uuid.NewString(),
		EventDemographicsId: req.EventDemographicsId,
		Name:                req.Name,
		Thumbnail:           "default.jpg",
		StartDate:           req.StartDate,
		EndDate:             req.EndDate,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	if req.Thumbnail != "" {
		entity.Thumbnail = req.Thumbnail
	}

	return entity
}

func NewEventFromUpdate(req UpdateEventRequestPayload) Event {
	var event = Event{
		PublicId:    req.PublicId,
    EventDemographicsId : req.EventDemographicsId,
		Description: req.Description,
		Address:     req.Address,
		Name:        req.Name,
		Thumbnail:   req.Thumbnail,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	return event
}

func (e Event) Validate() (err error) {

	if err = e.ValidateName(); err != nil {
		return
	}

	if err = e.ValidateDate(); err != nil {
		return
	}

	return
}

func (e Event) ValidateName() (err error) {
	if e.Name == "" {
		return response.ErrNameRequired
	}

	return
}

func (e Event) ValidateDate() (err error) {

	if e.StartDate.IsZero() {
		return response.ErrStartDateRequired
	}

	if e.EndDate.IsZero() {
		return response.ErrEndDateRequired
	}

	if e.StartDate.After(e.EndDate) {
		return response.ErrStartDateGreaterThanEndDate
	}

	return
}

func (e Event) IsEventDemographicSame(newIdEventDemographic int) bool {
  return e.EventDemographicsId == newIdEventDemographic
}

func (e *Event) SetEventDemographicsJSON(ed EventDemographics) (err error) {
  edJson, err := json.Marshal(ed)
  if err != nil {
    return 
  }

  e.EventDemographicsJSON = edJson

  return
}

func (e *Event) ConvertToEventResponse() EventResponse {
	return EventResponse{
		Id:          e.Id,
		PublicId:    e.PublicId,
    EventDemographicsId: e.EventDemographicsId,
		Name:        e.Name,
		Description: e.Description,
		Address:     e.Address,
		Thumbnail:   e.Thumbnail,
    EventDemographics: e.EventDemographics,
		StartDate:   e.StartDate,
		EndDate:     e.EndDate,
	}
}
