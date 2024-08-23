package usecases

import (
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
	businesstype "github.com/tapiaw38/cardon-tour-be/internal/usecases/business/business_type"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/profile"
	profiletype "github.com/tapiaw38/cardon-tour-be/internal/usecases/profile/profile_type"
)

type UseCases struct {
	Profile  Profile
	Business Business
}

type Profile struct {
	GetUseCase    profile.GetUsecase
	CreateUseCase profile.CreateUsecase
	UpdateUseCase profile.UpdateUsecase
	Types         ProfileType
}

type ProfileType struct {
	CreateUseCase profiletype.CreateUsecase
	DeleteUseCase profiletype.DeleteUsecase
	ListUseCase   profiletype.ListUsecase
}

type Business struct {
	Types BusinessType
}

type BusinessType struct {
	ListUseCase      businesstype.ListUsecase
	GetBySlugUseCase businesstype.GetBySlugUsecase
}

func CreateUsecases(contextFactory appcontext.Factory) *UseCases {
	return &UseCases{
		Profile: Profile{
			GetUseCase:    profile.NewGetUseCase(contextFactory),
			CreateUseCase: profile.NewCreateUseCase(contextFactory),
			UpdateUseCase: profile.NewUpdateUseCase(contextFactory),
			Types: ProfileType{
				CreateUseCase: profiletype.NewCreateUseCase(contextFactory),
				DeleteUseCase: profiletype.NewDeleteUseCase(contextFactory),
				ListUseCase:   profiletype.NewListUseCase(contextFactory),
			},
		},
		Business: Business{
			Types: BusinessType{
				ListUseCase:      businesstype.NewListUseCase(contextFactory),
				GetBySlugUseCase: businesstype.NewGetBySlugUsecase(contextFactory),
			},
		},
	}
}
