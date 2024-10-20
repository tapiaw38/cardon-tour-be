package event_type

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"

type EventTypeInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func toEventTypeInput(eventTypeInput EventTypeInput) domain.EventType {
	return domain.EventType{
		ID:   eventTypeInput.ID,
		Name: eventTypeInput.Name,
	}
}
