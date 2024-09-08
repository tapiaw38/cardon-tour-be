package profilesite

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
)

type (
	Repository interface {
		Create(context.Context, string, string) (domain.ProfileSite, error)
		Delete(context.Context, string, string) (domain.ProfileSite, error)
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
