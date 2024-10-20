package event_schedule

import (
	"context"
	"github.com/google/uuid"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	CreateUsecase interface {
		Execute(context.Context, domain.EventSchedule) (CreateOutput, error)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateOutput struct {
		Data EventScheduleOutputData `json:"data"`
	}
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{
		contextFactory: contextFactory,
	}
}

func (u *createUsecase) Execute(ctx context.Context, eventScheduleInput domain.EventSchedule) (CreateOutput, error) {
	app := u.contextFactory()

	id, err := uuid.NewUUID()
	if err != nil {
		return CreateOutput{}, err
	}
	eventScheduleInput.ID = id.String()
	eventScheduleID, err := app.Repositories.EventSchedule.Create(ctx, eventScheduleInput)
	if err != nil {
		return CreateOutput{}, err
	}

	eventSchedule, err := app.Repositories.EventSchedule.Get(ctx, eventScheduleID)
	if err != nil {
		return CreateOutput{}, err
	}

	return CreateOutput{
		Data: toEventScheduleOutputData(eventSchedule),
	}, nil
}
