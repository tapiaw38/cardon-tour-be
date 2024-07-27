package profile

import (
	"context"

	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	GetUsecase interface {
		Execute(context.Context, string) (GetOutput, error)
	}

	getUsecase struct {
		contextFactory appcontext.Factory
	}

	GetOutput struct {
		Data ProfileOutputData `json:"data"`
	}
)

func NewGetUseCase(contextFactory appcontext.Factory) GetUsecase {
	return &getUsecase{
		contextFactory: contextFactory,
	}
}

func (u *getUsecase) Execute(ctx context.Context, id string) (GetOutput, error) {
	app := u.contextFactory()

	user, err := app.Repositories.Profile.Get(ctx, id)
	if err != nil {
		return GetOutput{}, err
	}

	return GetOutput{
		Data: toProfileOutputData(user),
	}, nil
}
