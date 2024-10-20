package event_type

import (
	"context"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	ListUsecase interface {
		Execute(context.Context) ([]EventTypeOutputData, error)
	}

	listUsecase struct {
		contextFactory appcontext.Factory
	}

	ListOutput struct {
		Data []EventTypeOutputData `json:"data"`
	}
)

func NewListUsecase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{
		contextFactory: contextFactory,
	}
}

func (u *listUsecase) Execute(ctx context.Context) ([]EventTypeOutputData, error) {
	app := u.contextFactory()

	eventTypes, err := app.Repositories.EventType.List(ctx)
	if err != nil {
		return nil, err
	}

	output := make([]EventTypeOutputData, len(eventTypes))
	for i, eventType := range eventTypes {
		output[i] = toEventTypeOutputData(eventType)
	}

	return output, nil
}
