package event

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
)

type (
	Repository interface {
		Get(context.Context, string) (*domain.Event, error)
		List(context.Context, ListFilterOptions) ([]*domain.Event, error)
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
