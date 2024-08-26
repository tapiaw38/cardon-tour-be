package profile

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
)

type (
	Repository interface {
		Create(context.Context, domain.Profile) (string, error)
		Get(context.Context, string) (*domain.Profile, error)
		GetByUserID(context.Context, string) (*domain.Profile, error)
		Update(context.Context, string, domain.Profile) error
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
