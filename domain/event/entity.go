package event

import (
	"time"

	"github.com/mhmdiamd/go-social-service/infra/response"
)

type Event struct {
  Id string `db:"-"`
  PublicId string `db:"public_id" json:"public_id"`
  Name string `db:"name" json:"name"`
  Description string `db:"description" json:"description"`
  Address string `db:"address" json:"address"`
  Thumbnail string `db:"thumbnail" json:"thumbnail"`
  StartDate time.Time `db:"start_date" json:"start_date"`
  EndDate time.Time `db:"end_date" json:"end_date"`

  CreatedAt time.Time `db:"created_at" json:"created_at"`
  UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type EventPagination struct {
  Cursor int `json:"cursor"`
  Size int `json:"size"`
}

func NewEventFromCreate(req CreateEventRequestPayload) Event {
  entity := Event{
    Name: req.Name,
    Description: req.Description,
    Address: req.Address,
    Thumbnail: "default.jpg",
    StartDate: req.StartDate,
    EndDate: req.EndDate,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }

  if req.Thumbnail != "" {
    entity.Thumbnail = req.Thumbnail
  }

  return entity
}

func NewEventFromUpdate(req UpdateEventRequestPayload) Event {
  return Event{
    PublicId: req.PublicId,
    Name: req.Name,
    Description: req.Description,
    Address: req.Address,
    Thumbnail: req.Thumbnail,
    StartDate: req.StartDate,
    EndDate: req.EndDate,
  }
}

func (e Event) Validate() (err error) {

  if err = e.ValidateName(); err != nil {
    return
  }

  if err = e.ValidateDescription(); err != nil {
    return
  }

  if err = e.ValidateAddress(); err != nil {
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

func (e Event) ValidateDescription() (err error) {
  if e.Description == "" {
    return response.ErrDescriptionRequired
  }

  return
}

func (e Event) ValidateAddress() (err error) {
  if e.Address == "" {
    return response.ErrAddressRequired
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

func (e *Event) ConvertToEventResponse() EventResponse {
  return EventResponse{
    Id: e.Id ,
    PublicId : e.PublicId ,
    Name : e.Name,
    Description : e.Description,
    Address : e.Address, 
    Thumbnail : e.Thumbnail,
    StartDate : e.StartDate, 
    EndDate : e.EndDate,
  }
}

