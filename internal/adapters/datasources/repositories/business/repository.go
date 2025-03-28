package business

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

type (
	Repository interface {
		Create(ctx context.Context, business domain.Business) (string, error)
		Get(context.Context, string) (domain.Business, error)
		List(context.Context, ListFilterOptions) ([]domain.Business, error)
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
