package business_type

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

type (
	Repository interface {
		Create(ctx context.Context, businessType domain.BusinessType) (string, error)
		Get(ctx context.Context, id string) (domain.BusinessType, error)
		Update(ctx context.Context, id string, businessType domain.BusinessType) (string, error)
		Delete(ctx context.Context, id string) (string, error)
		List(ctx context.Context) ([]domain.BusinessType, error)
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
