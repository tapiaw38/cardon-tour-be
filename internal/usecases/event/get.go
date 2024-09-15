package event

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
		Data EventOutputData `json:"data"`
	}
)

func NewGetUseCase(contextFactory appcontext.Factory) GetUsecase {
	return &getUsecase{
		contextFactory: contextFactory,
	}
}

func (u *getUsecase) Execute(ctx context.Context, slug string) (GetOutput, error) {
	app := u.contextFactory()

	event, err := app.Repositories.Event.Get(ctx, slug)
	if err != nil {
		return GetOutput{}, err
	}

	return GetOutput{
		Data: toEventOutputData(*event),
	}, nil
}
