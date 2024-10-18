package businesstype

import (
	"context"
	"github.com/google/uuid"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	CreateUsecase interface {
		Execute(context.Context, domain.BusinessType) (*CreateOutput, error)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateOutput struct {
		Data BusinessTypeOutputData `json:"data"`
	}
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{
		contextFactory: contextFactory,
	}
}

func (u *createUsecase) Execute(ctx context.Context, businessTypeInput domain.BusinessType) (*CreateOutput, error) {
	app := u.contextFactory()

	businessTypeID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	businessTypeInput.ID = businessTypeID.String()
	id, err := app.Repositories.BusinessType.Create(ctx, businessTypeInput)
	if err != nil {
		return nil, err
	}

	businessType, err := app.Repositories.BusinessType.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{
		Data: toBusinessTypeOutputData(&businessType),
	}, nil
}
