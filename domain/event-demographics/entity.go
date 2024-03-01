package eventdemographics

import (
	"fmt"
	"strings"
	"time"

	"github.com/mhmdiamd/go-social-service/infra/response"
)

type EventDemographicsEntity struct {
  Name string `db:"name"`
  Gender string `db:"gender"`
  Graduation string `db:"graduation"`
  StartAge int `db:"start_age"`
  EndAge int `db:"end_age"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}

func NewEventDemographicsEntity(req CreateEventDemographicsRequestPayload) EventDemographicsEntity {
  entity := EventDemographicsEntity{
    Name : req.Name,
    Gender : "All",
    StartAge: req.StartAge,
    EndAge: req.EndAge,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }

  if req.Gender != "" {
    entity.Gender = req.Gender
  }

  if len(req.Graduation) != 0 {
    entity.ConvertGraduationToString(req.Graduation)
  }

  return entity
}

func (e EventDemographicsEntity) Validate() (err error) {
  if err = e.ValidateName(); err != nil {
    return
  }

  return 
}

func (e EventDemographicsEntity) ValidateName() (err error) {
  if e.Name == "" {
    return response.ErrNameRequired
  }

  return
}

func (e EventDemographicsEntity) ValidateAge() (err error) {
  if e.StartAge == 0 {
    return response.ErrStartAgeRequired
  }

  if e.StartAge < 1 {
    return response.ErrStartAgeToMin
  }

  if e.StartAge > e.EndAge {
    return response.ErrStartAgeGreaterThenEndAge
  }

  if e.StartAge > 99 {
    return response.ErrStartAgeToMax
  }

  if e.EndAge == 0 {
    return response.ErrEndAgeRequired
  }

  if e.EndAge < 0 {
    return response.ErrEndAgeToMin
  }

  if e.EndAge > 99 {
    return response.ErrEndAgeToMax
  }

  if e.EndAge < e.StartAge {
    return response.ErrEndAgeLowerThenStartAge 
  }

  return 

}

func (e *EventDemographicsEntity) ConvertGraduationToString(graduations []string) (err error) {
  e.Graduation = fmt.Sprint(graduations)
  return
}

func (e EventDemographicsEntity) ConvertGraduationToSlice() (graduations []string, err error) {
  arrGraduation := strings.Split(e.Graduation, ",")
  return arrGraduation, err
}


