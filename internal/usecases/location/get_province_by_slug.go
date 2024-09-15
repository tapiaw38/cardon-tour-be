package location

import (
	"context"

	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	GetProvinceBySlugUsecase interface {
		Execute(context.Context, string) (GetProvinceBySlugOutput, error)
	}

	getProvinceBySlugUsecase struct {
		contextFactory appcontext.Factory
	}

	GetProvinceBySlugOutput struct {
		Data ProvinceOutputData `json:"data"`
	}
)

func NewGetProvinceBySlugUseCase(contextFactory appcontext.Factory) GetProvinceBySlugUsecase {
	return &getProvinceBySlugUsecase{
		contextFactory: contextFactory,
	}
}

func (u *getProvinceBySlugUsecase) Execute(ctx context.Context, provincesSlug string) (GetProvinceBySlugOutput, error) {
	app := u.contextFactory()

	province, err := app.Repositories.Location.GetProvinceBySlug(ctx, provincesSlug)
	if err != nil {
		return GetProvinceBySlugOutput{}, err
	}

	return GetProvinceBySlugOutput{
		Data: toProvinceOutputData(province),
	}, nil
}
