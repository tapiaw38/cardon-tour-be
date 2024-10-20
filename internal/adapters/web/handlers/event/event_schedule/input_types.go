package event_schedule

import (
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
	"time"
)

type EventScheduleInput struct {
	EventID string    `json:"event_id"`
	Active  *bool     `json:"active"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}

func toEventScheduleInput(input EventScheduleInput) domain.EventSchedule {
	return domain.EventSchedule{
		EventID: input.EventID,
		Active:  input.Active,
		StartAt: input.StartAt,
		EndAt:   input.EndAt,
	}
}
