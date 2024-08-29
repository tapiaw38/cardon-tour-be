package business

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

type (
	ListFilterOptions struct {
		SiteSlug         string
		BusinessTypeSlug string
	}

	Repository interface {
		List(ctx context.Context, filter ListFilterOptions) ([]domain.Business, error)
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
