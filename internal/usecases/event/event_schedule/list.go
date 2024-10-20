package event_schedule

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
		Data []EventScheduleOutputData `json:"data"`
	}
)

func NewListUsecase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{
		contextFactory: contextFactory,
	}
}

func (u *listUsecase) Execute(ctx context.Context) (ListOutput, error) {
	app := u.contextFactory()

	eventSchedules, err := app.Repositories.EventSchedule.List(ctx)
	if err != nil {
		return ListOutput{}, err
	}

	eventScheduleOutputDatas := make([]EventScheduleOutputData, len(eventSchedules))
	for idx, eventSchedule := range eventSchedules {
		eventScheduleOutputDatas[idx] = toEventScheduleOutputData(&eventSchedule)
	}

	return ListOutput{
		Data: eventScheduleOutputDatas,
	}, nil
}
