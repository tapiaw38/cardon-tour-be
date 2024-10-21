package event_type

import (
	"context"
	"github.com/google/uuid"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	CreateUsecase interface {
		Execute(context.Context, domain.EventType) (*CreateOutput, error)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateOutput struct {
		Data EventTypeOutputData `json:"data"`
	}
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{
		contextFactory: contextFactory,
	}
}

func (u *createUsecase) Execute(ctx context.Context, eventTypeInput domain.EventType) (*CreateOutput, error) {
	app := u.contextFactory()

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	eventTypeInput.ID = id.String()
	eventTypeID, err := app.Repositories.EventType.Create(ctx, eventTypeInput)
	if err != nil {
		return nil, err
	}

	eventType, err := app.Repositories.EventType.Get(ctx, eventTypeID)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{
		Data: toEventTypeOutputData(eventType),
	}, nil
}
