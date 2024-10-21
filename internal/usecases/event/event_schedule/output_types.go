package event_schedule

import (
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
	"time"
)

type EventScheduleOutputData struct {
	ID        string `json:"id"`
	Active    *bool  `json:"active"`
	StartAt   string `json:"start_at"`
	EndAt     string `json:"end_at"`
	CreatedAt string `json:"created_at"`
}

func toEventScheduleOutputData(eventSchedule *domain.EventSchedule) EventScheduleOutputData {
	return EventScheduleOutputData{
		ID:        eventSchedule.ID,
		Active:    eventSchedule.Active,
		StartAt:   eventSchedule.StartAt.Format(time.RFC3339),
		EndAt:     eventSchedule.EndAt.Format(time.RFC3339),
		CreatedAt: eventSchedule.CreatedAt.Format(time.RFC3339),
	}
}
