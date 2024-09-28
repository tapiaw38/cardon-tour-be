package event_schedule

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
)

type (
	Repository interface {
		Create(context.Context, domain.EventSchedule) (string, error)
		Get(context.Context, string) (*domain.EventSchedule, error)
		List(context.Context, string) ([]domain.EventSchedule, error)
		Update(context.Context, domain.EventSchedule) (string, error)
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
