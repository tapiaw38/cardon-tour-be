package site_business_type

import (
	"context"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	CreateUsecase interface {
		Execute(context.Context, string, string) (*CreateOutput, error)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateOutput struct {
		Data SiteBusinessTypeOutputData `json:"data"`
	}
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{
		contextFactory: contextFactory,
	}
}

func (u *createUsecase) Execute(ctx context.Context, siteID, businessTypeID string) (*CreateOutput, error) {
	app := u.contextFactory()

	siteBusinessType, err := app.Repositories.SiteBusinessType.Create(ctx, siteID, businessTypeID)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{
		Data: toSiteBusinessTypeOutputData(&siteBusinessType),
	}, nil
}
