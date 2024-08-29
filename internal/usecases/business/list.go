package business

import (
	"context"

	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/business"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	ListUsecase interface {
		Execute(context.Context, ListFilterOptions) (ListOutput, error)
	}

	listUsecase struct {
		contextFactory appcontext.Factory
	}

	ListOutput struct {
		Data []BusinessOutputData `json:"data"`
	}

	ListFilterOptions business.ListFilterOptions
)

func NewListUseCase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{
		contextFactory: contextFactory,
	}
}

func (u *listUsecase) Execute(ctx context.Context, filters ListFilterOptions) (ListOutput, error) {
	app := u.contextFactory()

	businesses, err := app.Repositories.Business.List(ctx, business.ListFilterOptions(filters))
	if err != nil {
		return ListOutput{}, err
	}

	businessOutputs := make([]BusinessOutputData, 0, len(businesses))
	for _, business := range businesses {
		businessOutputs = append(businessOutputs, toBusinessOutputData(&business))
	}

	return ListOutput{
		Data: businessOutputs,
	}, nil
}
