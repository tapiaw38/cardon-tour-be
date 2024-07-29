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
	UpdateUseCase profile.UpdateUsecase
	ProfileType   ProfileType
}

type ProfileType struct {
	CreateUseCase profiletype.CreateUsecase
	DeleteUseCase profiletype.DeleteUsecase
	ListUseCase   profiletype.ListUsecase
}

func CreateUsecases(contextFactory appcontext.Factory) *UseCases {
	return &UseCases{
		Profile: Profile{
			GetUseCase:    profile.NewGetUseCase(contextFactory),
			CreateUseCase: profile.NewCreateUseCase(contextFactory),
			UpdateUseCase: profile.NewUpdateUseCase(contextFactory),
			ProfileType: ProfileType{
				CreateUseCase: profiletype.NewCreateUseCase(contextFactory),
				DeleteUseCase: profiletype.NewDeleteUseCase(contextFactory),
				ListUseCase:   profiletype.NewListUseCase(contextFactory),
			},
		},
	}
}
