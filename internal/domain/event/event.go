package domain

import (
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
	"time"
)

type (
	EventType struct {
		ID   string
		Name string
	}

	EventSchedule struct {
		ID        string
		EventID   string
		Active    *bool
		StartAt   time.Time
		EndAt     time.Time
		CreatedAt time.Time
	}

	Event struct {
		ID          string
		EventTypeID string
		SiteID      string
		Site        *domain.Site
		EventType   *EventType
		Name        string
		Description string
		Schedule    []EventSchedule
		CreatedAt   time.Time
		CreatedBy   string
	}
)
