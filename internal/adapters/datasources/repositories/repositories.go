package repositories

import (
	"database/sql"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/event"
	eventschedule "github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/event/event_schedule"
	eventtype "github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/event/event_type"
	sitebusinesstype "github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/site/business_type"

	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/business"
	businessimage "github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/business/business_image"
	businesstype "github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/business/business_type"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/location"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/profile"
	profilesite "github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/profile/profile_site"
	profiletype "github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/profile/profile_type"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/site"
)

type Repositories struct {
	Profile          profile.Repository
	ProfileType      profiletype.Repository
	ProfileSite      profilesite.Repository
	Location         location.Repository
	Site             site.Repository
	SiteBusinessType sitebusinesstype.Repository
	Business         business.Repository
	BusinessType     businesstype.Repository
	BusinessImage    businessimage.Repository
	Event            event.Repository
	EventType        eventtype.Repository
	EventSchedule    eventschedule.Repository
}

func CreateRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Profile:          profile.NewRepository(db),
		ProfileType:      profiletype.NewRepository(db),
		ProfileSite:      profilesite.NewRepository(db),
		Location:         location.NewRepository(db),
		Site:             site.NewRepository(db),
		SiteBusinessType: sitebusinesstype.NewRepository(db),
		Business:         business.NewRepository(db),
		BusinessType:     businesstype.NewRepository(db),
		BusinessImage:    businessimage.NewRepository(db),
		Event:            event.NewRepository(db),
		EventType:        eventtype.NewRepository(db),
		EventSchedule:    eventschedule.NewRepository(db),
	}
}
