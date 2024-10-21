package event_type

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
		Data EventTypeOutputData `json:"data"`
	}
)

func NewGetUsecase(contextFactory appcontext.Factory) GetUsecase {
	return &getUsecase{
		contextFactory: contextFactory,
	}
}

func (u *getUsecase) Execute(ctx context.Context, eventTypeID string) (*GetOutput, error) {
	app := u.contextFactory()

	eventType, err := app.Repositories.EventType.Get(ctx, eventTypeID)
	if err != nil {
		return nil, err
	}

	return &GetOutput{
		Data: toEventTypeOutputData(eventType),
	}, nil
}
