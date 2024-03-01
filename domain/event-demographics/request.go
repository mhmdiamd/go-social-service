package eventdemographics

type CreateEventDemographicsRequestPayload struct {
  Name string `json:"name"`
  Gender string `json:"gender"`
  Graduation []string `json:"graduation"`
  StartAge int `json:"start_age"`
  EndAge int `json:"end_age"`
}



