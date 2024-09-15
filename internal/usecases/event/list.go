package event

import (
	"context"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/event"
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
		Data []EventOutputData `json:"data"`
	}

	ListFilterOptions event.ListFilterOptions
)

func NewListUseCase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{
		contextFactory: contextFactory,
	}
}

func (u *listUsecase) Execute(ctx context.Context, filter ListFilterOptions) (ListOutput, error) {
	app := u.contextFactory()

	events, err := app.Repositories.Event.List(ctx, event.ListFilterOptions(filter))
	if err != nil {
		return ListOutput{}, err
	}

	outputData := make([]EventOutputData, len(events))
	for i, event := range events {
		outputData[i] = toEventOutputData(*event)
	}

	return ListOutput{
		Data: outputData,
	}, nil
}
