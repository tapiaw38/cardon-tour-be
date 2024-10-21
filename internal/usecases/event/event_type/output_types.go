package event_type

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"

type EventTypeOutputData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func toEventTypeOutputData(eventType domain.EventType) EventTypeOutputData {
	return EventTypeOutputData{
		ID:   eventType.ID,
		Name: eventType.Name,
	}
}
