package eventdemographics

import (
	"strings"
	"time"

	"github.com/mhmdiamd/go-social-service/infra/response"
)

var GENDER_MAPPING = map[string]string {
    "male" : "male",
    "female" : "female",
    "all" : "all",
  }


type EventDemographicsEntity struct {
  Id int `db:"id"`
  Name string `db:"name"`
  Gender string `db:"gender"`
  Graduation string `db:"graduation"`
  StartAge int `db:"start_age"`
  EndAge int `db:"end_age"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}

func NewEventDemographicsEntityFromUpdate(req UpdateEventDemographicsRequestPayload) EventDemographicsEntity {
  entity := EventDemographicsEntity {
    Id : req.Id,
    Name : req.Name,
    Gender : req.Gender,
    StartAge: req.StartAge,
    EndAge: req.EndAge,
  }
  
  if len(req.Graduation) != 0 {
    entity.ConvertGraduationToString(req.Graduation)
  }

  return entity 
}

func NewEventDemographicsEntity(req CreateEventDemographicsRequestPayload) EventDemographicsEntity {
  entity := EventDemographicsEntity{
    Name : req.Name,
    Gender : "all",
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

  if err = e.ValidateAge(); err != nil {
    return
  }

  if err = e.ValidateGender(); err != nil {
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

func (e EventDemographicsEntity) ValidateGender() (err error) {
  gender := strings.ToLower(e.Gender)

  if _, ok := GENDER_MAPPING[gender]; !ok {
    return response.ErrGenderInvalid
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
  
  if e.StartAge > 99 {
    return response.ErrStartAgeToMax
  }

  if e.StartAge > e.EndAge {
    return response.ErrStartAgeGreaterThenEndAge
  }

  if e.EndAge == 0 {
    return response.ErrEndAgeRequired
  }

  if e.EndAge > 99 {
    return response.ErrEndAgeToMax
  }

  return 
}

func (e EventDemographicsEntity) ValidateId() (err error) { 
  if e.Id == 0 {
    return response.ErrIdRequired
  }

  return
}


func (e *EventDemographicsEntity) ConvertGraduationToString(graduations []string) {

  newString := ""
  for i, graduation := range graduations {
    if i == 0 {
      newString = newString + graduation
    }else {
      newString = newString + "," + graduation
    }
  }

  e.Graduation = newString
}

func (e EventDemographicsEntity) ConvertGraduationToSlice() (graduations []string) {
  return strings.Split(e.Graduation, ",")
}


