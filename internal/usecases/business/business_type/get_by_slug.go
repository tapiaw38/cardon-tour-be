package businesstype

import (
	"context"

	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	GetBySlugUsecase interface {
		Execute(context.Context, string) (*GetBySlugOutput, error)
	}

	getBySlugUsecase struct {
		contextFactory appcontext.Factory
	}

	GetBySlugOutput struct {
		Data BusinessTypeOutputData `json:"data"`
	}
)

func NewGetBySlugUsecase(contextFactory appcontext.Factory) GetBySlugUsecase {
	return &getBySlugUsecase{
		contextFactory: contextFactory,
	}
}

func (u *getBySlugUsecase) Execute(ctx context.Context, slug string) (*GetBySlugOutput, error) {
	app := u.contextFactory()

	businessType, err := app.Repositories.BusinessType.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	return &GetBySlugOutput{
		Data: toBusinessTypeOutputData(&businessType),
	}, nil
}
