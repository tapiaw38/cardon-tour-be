package repositories

import (
	"database/sql"

	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/profile"
)

type Repositories struct {
	Profile profile.Repository
}

func CreateRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Profile: profile.NewRepository(db),
	}
}
