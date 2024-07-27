package usecases

import (
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/profile"
	profiletype "github.com/tapiaw38/cardon-tour-be/internal/usecases/profile/profile_type"
)

type UseCases struct {
	Profile Profile
}

type Profile struct {
	GetUseCase    profile.GetUsecase
	CreateUseCase profile.CreateUsecase
	ProfileType   ProfileType
}

type ProfileType struct {
	CreateUseCase profiletype.CreateUsecase
}

func CreateUsecases(contextFactory appcontext.Factory) *UseCases {
	return &UseCases{
		Profile: Profile{
			GetUseCase:    profile.NewGetUseCase(contextFactory),
			CreateUseCase: profile.NewCreateUseCase(contextFactory),
			ProfileType: ProfileType{
				CreateUseCase: profiletype.NewCreateUseCase(contextFactory),
			},
		},
	}
}
