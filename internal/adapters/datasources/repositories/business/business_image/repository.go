package business_image

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

type (
	Repository interface {
		Create(ctx context.Context, businessImage domain.BusinessImage) (string, error)
		Get(context.Context, string) (domain.BusinessImage, error)
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
