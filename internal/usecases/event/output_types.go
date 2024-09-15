package event

import (
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
	"time"
)

type (
	EventOutputData struct {
		ID          string                    `json:"id"`
		SiteID      string                    `json:"site_id"`
		Name        string                    `json:"name"`
		Description string                    `json:"description"`
		EventType   EventTypeOutputData       `json:"event_type"`
		Schedule    []EventScheduleOutputData `json:"schedule"`
		CreatedAt   string                    `json:"created_at"`
		CreatedBy   string                    `json:"created_by"`
	}

	EventTypeOutputData struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	EventScheduleOutputData struct {
		ID        string `json:"id"`
		Active    bool   `json:"active"`
		StartAt   string `json:"start_at"`
		EndAt     string `json:"end_at"`
		CreatedAt string `json:"created_at"`
	}
)

func toEventOutputData(event domain.Event) EventOutputData {
	var eventType EventTypeOutputData
	if event.EventType != nil {
		eventType = EventTypeOutputData{
			ID:   event.EventType.ID,
			Name: event.EventType.Name,
		}
	}
	var schedule []EventScheduleOutputData
	for _, s := range event.Schedule {
		schedule = append(schedule, EventScheduleOutputData{
			ID:        s.ID,
			Active:    s.Active,
			StartAt:   s.StartAt.Format(time.RFC3339),
			EndAt:     s.EndAt.Format(time.RFC3339),
			CreatedAt: s.CreatedAt.Format(time.RFC3339),
		})
	}

	return EventOutputData{
		ID:          event.ID,
		SiteID:      event.SiteID,
		Name:        event.Name,
		Description: event.Description,
		EventType:   eventType,
		Schedule:    schedule,
		CreatedAt:   event.CreatedAt.Format(time.RFC3339),
		CreatedBy:   event.CreatedBy,
	}
}
