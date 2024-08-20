package businesstype

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
		Data []BusinessTypeOutputData `json:"data"`
	}
)

func NewListUseCase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{
		contextFactory: contextFactory,
	}
}

func (u *listUsecase) Execute(ctx context.Context) (ListOutput, error) {
	app := u.contextFactory()

	businessTypes, err := app.Repositories.BusinessType.List(ctx)
	if err != nil {
		return ListOutput{}, err
	}

	businessTypeOutputs := make([]BusinessTypeOutputData, 0, len(businessTypes))
	for _, businessType := range businessTypes {
		businessTypeOutputs = append(businessTypeOutputs, toBusinessTypeOutputData(&businessType))
	}

	return ListOutput{
		Data: businessTypeOutputs,
	}, nil
}
