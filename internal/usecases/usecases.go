package usecases

import (
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/business"
	businesstype "github.com/tapiaw38/cardon-tour-be/internal/usecases/business/business_type"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/location"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/profile"
	profilesite "github.com/tapiaw38/cardon-tour-be/internal/usecases/profile/profile_site"
	profiletype "github.com/tapiaw38/cardon-tour-be/internal/usecases/profile/profile_type"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/site"
)

type UseCases struct {
	Profile  Profile
	Location Location
	Site     Site
	Business Business
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
	ListProvinceUseCase location.ListProvinceUsecase
}

type Site struct {
	ListUseCase      site.ListUsecase
	GetUseCase       site.GetUsecase
	GetBySlugUseCase site.GetBySlugUsecase
}

type Business struct {
	GetUseCase  business.GetUsecase
	ListUseCase business.ListUsecase
	Types       BusinessType
}

type BusinessType struct {
	ListUseCase      businesstype.ListUsecase
	GetUseCase       businesstype.GetUsecase
	GetBySlugUseCase businesstype.GetBySlugUsecase
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
			ListProvinceUseCase: location.NewListProvinceUseCase(contextFactory),
		},
		Site: Site{
			ListUseCase:      site.NewListUseCase(contextFactory),
			GetUseCase:       site.NewGetUsecase(contextFactory),
			GetBySlugUseCase: site.NewGetBySlugUsecase(contextFactory),
		},
		Business: Business{
			GetUseCase:  business.NewGetUseCase(contextFactory),
			ListUseCase: business.NewListUseCase(contextFactory),
			Types: BusinessType{
				ListUseCase:      businesstype.NewListUseCase(contextFactory),
				GetUseCase:       businesstype.NewGetUsecase(contextFactory),
				GetBySlugUseCase: businesstype.NewGetBySlugUsecase(contextFactory),
			},
		},
	}
}
