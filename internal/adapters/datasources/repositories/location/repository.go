package location

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/location"
)

type (
	Repository interface {
		CreateCountry(context.Context, domain.Country) (string, error)
		DeleteCountry(context.Context, string) error
		CreateProvince(context.Context, domain.Province, string) (string, error)
		GetProvinceBySlug(context.Context, string) (*domain.Province, error)
		ListProvince(context.Context) ([]domain.Province, error)
		DeleteProvince(context.Context, string) error
		CreateCity(context.Context, domain.City, string) (string, error)
		DeleteCity(context.Context, string) error
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
