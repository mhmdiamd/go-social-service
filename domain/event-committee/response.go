package eventcommittee

import "github.com/google/uuid"

type EventCommitteeEntityResponse struct {
	Id            int                    `json:"id"`
	UserPublicID  uuid.UUID              `json:"user_public_id"`
	EventPublicID uuid.UUID              `json:"event_public_id"`
	Position      EventCommitteePosition `json:"position"`
}

func NewEventCommitteeEntityResponse(entity EventCommitteeEntity) EventCommitteeEntityResponse {
	response := EventCommitteeEntityResponse{
		Id:            entity.Id,
		UserPublicID:  entity.UserPublicID,
		EventPublicID: entity.EventPublicID,
		Position:      entity.Position,
	}

	return response
}

func NewListEventCommitteeEntityResponse(events []EventCommitteeEntity) []EventCommitteeEntityResponse {
	var newEvents []EventCommitteeEntityResponse

	for _, event := range events {
		response := NewEventCommitteeEntityResponse(event)

		newEvents = append(newEvents, response)
	}

	return newEvents
}
