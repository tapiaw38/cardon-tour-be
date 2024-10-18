package site

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

type (
	Repository interface {
		Create(context.Context, domain.Site) (string, error)
		List(context.Context, ListFilterOptions) ([]domain.Site, error)
		Get(context.Context, string) (*domain.Site, error)
		GetBySlug(context.Context, string) (*domain.Site, error)
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
