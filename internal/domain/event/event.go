package domain

import "time"

type (
	EventType struct {
		ID   string
		Name string
	}

	EventSchedule struct {
		ID        string
		Active    bool
		StartAt   time.Time
		EndAt     time.Time
		CreatedAt time.Time
	}

	Event struct {
		ID          string
		EventTypeID string
		EventType   *EventType
		Name        string
		Description string
		Schedule    *EventSchedule
		CreatedAt   time.Time
		CreatedBy   string
	}
)
