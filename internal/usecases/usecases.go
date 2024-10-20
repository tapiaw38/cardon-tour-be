package usecases

import (
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/business"
	businessimage "github.com/tapiaw38/cardon-tour-be/internal/usecases/business/bisiness_image"
	businesstype "github.com/tapiaw38/cardon-tour-be/internal/usecases/business/business_type"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/event"
	eventschedule "github.com/tapiaw38/cardon-tour-be/internal/usecases/event/event_schedule"
	eventtype "github.com/tapiaw38/cardon-tour-be/internal/usecases/event/event_type"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/location"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/profile"
	profilesite "github.com/tapiaw38/cardon-tour-be/internal/usecases/profile/profile_site"
	profiletype "github.com/tapiaw38/cardon-tour-be/internal/usecases/profile/profile_type"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/site"
	sitebusinesstype "github.com/tapiaw38/cardon-tour-be/internal/usecases/site/business_type"
)

type UseCases struct {
	Profile  Profile
	Location Location
	Site     Site
	Business Business
	Event    Event
}

type Profile struct {
	GetByUserIDUseCase profile.GetByUserIDUsecase
	CreateUseCase      profile.CreateUsecase
	UpdateUseCase      profile.UpdateUsecase
	Types              ProfileType
	Sites              ProfileSite
}

type ProfileType struct {
	CreateUseCase profiletype.CreateUsecase
	DeleteUseCase profiletype.DeleteUsecase
	ListUseCase   profiletype.ListUsecase
}

type ProfileSite struct {
	CreateUseCase profilesite.CreateUsecase
	DeleteUseCase profilesite.DeleteUsecase
}

type Location struct {
	GetProvinceBySlugUseCase location.GetProvinceBySlugUsecase
	ListProvinceUseCase      location.ListProvinceUsecase
}

type Site struct {
	CreateUseCase    site.CreateUsecase
	ListUseCase      site.ListUsecase
	GetUseCase       site.GetUsecase
	GetBySlugUseCase site.GetBySlugUsecase
	BusinessType     SiteBusinessType
}

type SiteBusinessType struct {
	CreateUseCase sitebusinesstype.CreateUsecase
	DeleteUseCase sitebusinesstype.DeleteUsecase
}

type Business struct {
	CreateUseCase business.CreateUsecase
	GetUseCase    business.GetUsecase
	ListUseCase   business.ListUsecase
	Types         BusinessType
	Images        BusinessImage
}

type BusinessType struct {
	CreateUseCase    businesstype.CreateUsecase
	ListUseCase      businesstype.ListUsecase
	GetUseCase       businesstype.GetUsecase
	GetBySlugUseCase businesstype.GetBySlugUsecase
}

type BusinessImage struct {
	CreateUseCase businessimage.CreateUsecase
	GetUseCase    businessimage.GetUsecase
}

type Event struct {
	GetUseCase  event.GetUsecase
	ListUseCase event.ListUsecase
	Types       EventType
	Schedule    EventSchedule
}

type EventType struct {
	CreateUseCase eventtype.CreateUsecase
	ListUseCase   eventtype.ListUsecase
	GetUseCase    eventtype.GetUsecase
}

type EventSchedule struct {
	CreateUseCase eventschedule.CreateUsecase
	GetUseCase    eventschedule.GetUsecase
	ListUseCase   eventschedule.ListUsecase
}

func CreateUsecases(contextFactory appcontext.Factory) *UseCases {
	return &UseCases{
		Profile: Profile{
			GetByUserIDUseCase: profile.NewGetByUserIDUseCase(contextFactory),
			CreateUseCase:      profile.NewCreateUseCase(contextFactory),
			UpdateUseCase:      profile.NewUpdateUseCase(contextFactory),
			Types: ProfileType{
				CreateUseCase: profiletype.NewCreateUseCase(contextFactory),
				DeleteUseCase: profiletype.NewDeleteUseCase(contextFactory),
				ListUseCase:   profiletype.NewListUseCase(contextFactory),
			},
			Sites: ProfileSite{
				CreateUseCase: profilesite.NewCreateUseCase(contextFactory),
				DeleteUseCase: profilesite.NewDeleteUseCase(contextFactory),
			},
		},
		Location: Location{
			GetProvinceBySlugUseCase: location.NewGetProvinceBySlugUseCase(contextFactory),
			ListProvinceUseCase:      location.NewListProvinceUseCase(contextFactory),
		},
		Site: Site{
			CreateUseCase:    site.NewCreateUsecase(contextFactory),
			ListUseCase:      site.NewListUseCase(contextFactory),
			GetUseCase:       site.NewGetUsecase(contextFactory),
			GetBySlugUseCase: site.NewGetBySlugUsecase(contextFactory),
			BusinessType: SiteBusinessType{
				CreateUseCase: sitebusinesstype.NewCreateUsecase(contextFactory),
				DeleteUseCase: sitebusinesstype.NewDeleteUsecase(contextFactory),
			},
		},
		Business: Business{
			CreateUseCase: business.NewCreateUseCase(contextFactory),
			GetUseCase:    business.NewGetUseCase(contextFactory),
			ListUseCase:   business.NewListUseCase(contextFactory),
			Types: BusinessType{
				CreateUseCase:    businesstype.NewCreateUsecase(contextFactory),
				ListUseCase:      businesstype.NewListUseCase(contextFactory),
				GetUseCase:       businesstype.NewGetUsecase(contextFactory),
				GetBySlugUseCase: businesstype.NewGetBySlugUsecase(contextFactory),
			},
			Images: BusinessImage{
				CreateUseCase: businessimage.NewCreateUseCase(contextFactory),
				GetUseCase:    businessimage.NewGetUsecase(contextFactory),
			},
		},
		Event: Event{
			GetUseCase:  event.NewGetUseCase(contextFactory),
			ListUseCase: event.NewListUseCase(contextFactory),
			Types: EventType{
				CreateUseCase: eventtype.NewCreateUsecase(contextFactory),
				ListUseCase:   eventtype.NewListUsecase(contextFactory),
				GetUseCase:    eventtype.NewGetUsecase(contextFactory),
			},
			Schedule: EventSchedule{
				CreateUseCase: eventschedule.NewCreateUsecase(contextFactory),
				GetUseCase:    eventschedule.NewGetUsecase(contextFactory),
				ListUseCase:   eventschedule.NewListUsecase(contextFactory),
			},
		},
	}
}
