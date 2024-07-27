package repositories

import (
	"database/sql"

	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/profile"
	profiletype "github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/profile/profile_type"
)

type Repositories struct {
	Profile     profile.Repository
	ProfileType profiletype.Repository
}

func CreateRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Profile:     profile.NewRepository(db),
		ProfileType: profiletype.NewRepository(db),
	}
}
