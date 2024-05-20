package event

import "github.com/mhmdiamd/go-social-service/infra/response"

type EventDemographics struct {
	Id         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Gender     string `db:"gender" json:"gender"`
	Graduation string `db:"graduation" json:"graduation"`
	StartAge   int    `db:"start_age" json:"start_age"`
	EndAge     int    `db:"end_age" json:"end_age"`
}

func (ed *EventDemographics) IsExists() error {
	if ed.Id == 0 {
		return response.ErrNotFound
	}

	return nil
}
