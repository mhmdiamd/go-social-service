package eventdemographics

type CreateEventDemographicsRequestPayload struct {
  Name string `json:"name" form:"name"`
  Gender string `json:"gender" form:"gender"`
  Graduation []string `json:"graduation" form:"graduation"`
  StartAge int `json:"start_age" form:"start_age"`
  EndAge int `json:"end_age" form:"end_age"`
}

type UpdateEventDemographicsRequestPayload struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Gender string `json:"gender"`
  Graduation []string `json:"graduation"`
  StartAge int `json:"start_age"`
  EndAge int `json:"end_age"`
}

type EventDemographicsRequestQuery struct {
  Id int `query:"id"`
}



