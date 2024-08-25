package site

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

type (
	Repository interface {
		List(ctx context.Context) ([]domain.Site, error)
		GetByID(ctx context.Context, id string) (*domain.Site, error)
		GetBySlug(ctx context.Context, slug string) (*domain.Site, error)
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
