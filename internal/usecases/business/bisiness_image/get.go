package bisiness_image

import (
	"context"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	GetUsecase interface {
		Execute(context.Context, string) (*GetOutput, error)
	}

	getUsecase struct {
		contextFactory appcontext.Factory
	}

	GetOutput struct {
		Data BusinessImageOutputData `json:"data"`
	}
)

func NewGetUsecase(contextFactory appcontext.Factory) GetUsecase {
	return &getUsecase{
		contextFactory: contextFactory,
	}
}

func (u *getUsecase) Execute(ctx context.Context, id string) (*GetOutput, error) {
	app := u.contextFactory()

	businessImage, err := app.Repositories.BusinessImage.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetOutput{
		Data: toBusinessImageOutputData(businessImage),
	}, nil
}
