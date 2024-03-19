package eventdemographics

type EventDemographicsEntityResponse struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Gender string `json:"gender"`
  Graduation []string `json:"graduations"`
  StartAge int `json:"start_age"`
  EndAge int `json:"end_age"`
}

func NewEventDemographicsEntityResponse(entity EventDemographicsEntity) EventDemographicsEntityResponse {

  response := EventDemographicsEntityResponse {
    Id : entity.Id,
    Name : entity.Name,
    Gender : entity.Gender,
    Graduation : []string{},
    StartAge : entity.StartAge,
    EndAge : entity.EndAge,
  }

  if entity.Graduation != "" {
    response.Graduation = entity.ConvertGraduationToSlice()
  }

  return response
}

func NewListEventDemographicsEntityResponse(events []EventDemographicsEntity) []EventDemographicsEntityResponse {
  var newEvents []EventDemographicsEntityResponse

  for _, event := range events {
    response := NewEventDemographicsEntityResponse(event)

    if event.Graduation != "" {
      response.Graduation = event.ConvertGraduationToSlice()
    }

    newEvents = append(newEvents, response)
  }

  return newEvents 
}
