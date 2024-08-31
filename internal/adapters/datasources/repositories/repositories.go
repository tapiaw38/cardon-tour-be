package repositories

import (
	"database/sql"

	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/business"
	businesstype "github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/business/business_type"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/profile"
	profilesite "github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/profile/profile_site"
	profiletype "github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/profile/profile_type"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/site"
)

type Repositories struct {
	Profile      profile.Repository
	ProfileType  profiletype.Repository
	ProfileSite  profilesite.Repository
	Site         site.Repository
	Business     business.Repository
	BusinessType businesstype.Repository
}

func CreateRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Profile:      profile.NewRepository(db),
		ProfileType:  profiletype.NewRepository(db),
		ProfileSite:  profilesite.NewRepository(db),
		Site:         site.NewRepository(db),
		Business:     business.NewRepository(db),
		BusinessType: businesstype.NewRepository(db),
	}
}
