package eventdemographics

type EventDemographicsEntityResponse struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Gender string `json:"gender"`
  Graduation []string `json:"graduation"`
  StartAge int `json:"start_age"`
  EndAge int `json:"end_age"`
}

func NewEventDemographicsEntityResponse(entity EventDemographicsEntity) EventDemographicsEntityResponse {
  return EventDemographicsEntityResponse {
    Name : entity.Name,
    Gender : entity.Gender,
    Graduation : entity.ConvertGraduationToSlice(),
    StartAge : entity.StartAge,
    EndAge : entity.EndAge,
  }
}
