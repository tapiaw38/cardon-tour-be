package profile

import (
	"context"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	UpdateUsecase interface {
		Execute(context.Context, string, domain.Profile) (UpdateOutput, error)
	}

	updateUsecase struct {
		contextFactory appcontext.Factory
	}

	UpdateOutput struct {
		Data ProfileOutputData `json:"data"`
	}
)

func NewUpdateUseCase(contextFactory appcontext.Factory) UpdateUsecase {
	return &updateUsecase{
		contextFactory: contextFactory,
	}
}

func (u *updateUsecase) Execute(ctx context.Context, id string, profile domain.Profile) (UpdateOutput, error) {
	app := u.contextFactory()

	err := app.Repositories.Profile.Update(ctx, id, profile)
	if err != nil {
		return UpdateOutput{}, err
	}

	profileUpdated, err := app.Repositories.Profile.Get(ctx, id)
	if err != nil {
		return UpdateOutput{}, err
	}

	return UpdateOutput{
		Data: toProfileOutputData(profileUpdated),
	}, nil
}
