package site_business_type

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

type (
	Repository interface {
		Create(context.Context, string, string) (domain.SiteBusinessType, error)
		Delete(context.Context, string, string) (domain.SiteBusinessType, error)
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
