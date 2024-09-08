package location

import (
	"context"

	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	ListProvinceUsecase interface {
		Execute(ctx context.Context) (ListProvinceOutputData, error)
	}

	listProvinceUsecase struct {
		contextFactory appcontext.Factory
	}

	ListProvinceOutputData struct {
		Data []ProvinceOutputData `json:"data"`
	}
)

func NewListProvinceUseCase(contextFactory appcontext.Factory) ListProvinceUsecase {
	return &listProvinceUsecase{
		contextFactory: contextFactory,
	}
}

func (u *listProvinceUsecase) Execute(ctx context.Context) (ListProvinceOutputData, error) {
	app := u.contextFactory()

	provinces, err := app.Repositories.Location.ListProvince(ctx)
	if err != nil {
		return ListProvinceOutputData{}, err
	}

	provinceOutputs := make([]ProvinceOutputData, 0, len(provinces))
	for _, province := range provinces {
		provinceOutputs = append(provinceOutputs, toProvinceOutputData(&province))
	}

	return ListProvinceOutputData{
		Data: provinceOutputs,
	}, nil
}
