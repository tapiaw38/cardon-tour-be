package profile

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	CreateUsecase interface {
		Execute(context.Context, domain.Profile) (CreateOutput, error)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateOutput struct {
		Data ProfileOutputData `json:"data"`
	}
)

func NewCreateUseCase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{
		contextFactory: contextFactory,
	}
}

func (u *createUsecase) Execute(ctx context.Context, profile domain.Profile) (CreateOutput, error) {
	app := u.contextFactory()

	profileID, err := uuid.NewUUID()
	if err != nil {
		return CreateOutput{}, err
	}

	profile.ID = profileID.String()

	id, err := app.Repositories.Profile.Create(ctx, profile)
	if err != nil {
		return CreateOutput{}, err
	}

	userCreated, err := app.Repositories.Profile.Get(ctx, id)
	if err != nil {
		return CreateOutput{}, err
	}

	return CreateOutput{
		Data: toProfileOutputData(userCreated),
	}, nil
}
