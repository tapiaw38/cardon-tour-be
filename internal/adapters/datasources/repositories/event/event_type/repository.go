package eventtype

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
)

type (
	Repository interface {
		Create(ctx context.Context, eventType domain.EventType) (string, error)
		Get(ctx context.Context, eventTypeID string) (domain.EventType, error)
		List(ctx context.Context) ([]domain.EventType, error)
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
