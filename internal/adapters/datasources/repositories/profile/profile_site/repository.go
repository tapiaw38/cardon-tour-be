package profilesite

import (
	"context"
	"database/sql"
)

type (
	Repository interface {
		Create(context.Context, string, string) error
		Delete(context.Context, string, string) error
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
