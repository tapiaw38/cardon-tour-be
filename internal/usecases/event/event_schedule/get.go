package event_schedule

import (
	"context"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	GetUsecase interface {
		Execute(context.Context, string) (GetOutput, error)
	}

	getUsecase struct {
		contextFactory appcontext.Factory
	}

	GetOutput struct {
		Data EventScheduleOutputData `json:"data"`
	}
)

func NewGetUsecase(contextFactory appcontext.Factory) GetUsecase {
	return &getUsecase{
		contextFactory: contextFactory,
	}
}

func (u *getUsecase) Execute(ctx context.Context, slug string) (GetOutput, error) {
	app := u.contextFactory()

	eventSchedule, err := app.Repositories.EventSchedule.Get(ctx, slug)
	if err != nil {
		return GetOutput{}, err
	}

	return GetOutput{
		Data: toEventScheduleOutputData(eventSchedule),
	}, nil
}
