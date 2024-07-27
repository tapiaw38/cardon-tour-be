package usecases

import (
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/profile"
)

type UseCases struct {
	Profile Profile
}

type Profile struct {
	GetUseCase    profile.GetUsecase
	CreateUseCase profile.CreateUsecase
}

func CreateUsecases(contextFactory appcontext.Factory) *UseCases {
	return &UseCases{
		Profile: Profile{
			GetUseCase: profile.NewGetUseCase(contextFactory),
		},
	}
}
