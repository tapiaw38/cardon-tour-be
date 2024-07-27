package profiletype

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
)

type (
	Repository interface {
		Create(context.Context, domain.ProfileType) error
		Delete(context.Context, string) error
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
