package business

import (
	"context"
	"github.com/google/uuid"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	CreateUsecase interface {
		Execute(context.Context, domain.Business) (*CreateOutput, error)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateOutput struct {
		Data BusinessOutputData `json:"data"`
	}
)

func NewCreateUseCase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{
		contextFactory: contextFactory,
	}
}

func (u *createUsecase) Execute(ctx context.Context, businessInput domain.Business) (*CreateOutput, error) {
	app := u.contextFactory()

	businessID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	businessInput.ID = businessID.String()
	id, err := app.Repositories.Business.Create(ctx, businessInput)
	if err != nil {
		return nil, err
	}

	business, err := app.Repositories.Business.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{
		Data: toBusinessOutputData(&business),
	}, nil
}
