package profiletype

import (
	"context"

	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	ListUsecase interface {
		Execute(context.Context) (ListOutput, error)
	}

	listUsecase struct {
		contextFactory appcontext.Factory
	}

	ListOutput struct {
		Data []ProfileTypeOutputData `json:"data"`
	}
)

func NewListUseCase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{
		contextFactory: contextFactory,
	}
}

func (u *listUsecase) Execute(ctx context.Context) (ListOutput, error) {
	app := u.contextFactory()

	profileTypes, err := app.Repositories.ProfileType.List(ctx)
	if err != nil {
		return ListOutput{}, err
	}

	profileTypeOutputs := make([]ProfileTypeOutputData, 0, len(profileTypes))
	for _, profileType := range profileTypes {
		profileTypeOutputs = append(profileTypeOutputs, toProfileTypeOutputData(&profileType))
	}

	return ListOutput{
		Data: profileTypeOutputs,
	}, nil
}
