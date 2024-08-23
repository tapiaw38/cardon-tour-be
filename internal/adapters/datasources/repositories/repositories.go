package repositories

import (
	"database/sql"

	businesstype "github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/business/business_type"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/profile"
	profiletype "github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/profile/profile_type"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/site"
)

type Repositories struct {
	Profile      profile.Repository
	ProfileType  profiletype.Repository
	Site         site.Repository
	BusinessType businesstype.Repository
}

func CreateRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Profile:      profile.NewRepository(db),
		ProfileType:  profiletype.NewRepository(db),
		Site:         site.NewRepository(db),
		BusinessType: businesstype.NewRepository(db),
	}
}
