package site_business_type

import (
	"context"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	DeleteUsecase interface {
		Execute(context.Context, string, string) (*DeleteOutput, error)
	}

	deleteUsecase struct {
		contextFactory appcontext.Factory
	}

	DeleteOutput struct {
		Data SiteBusinessTypeOutputData `json:"data"`
	}
)

func NewDeleteUsecase(contextFactory appcontext.Factory) DeleteUsecase {
	return &deleteUsecase{
		contextFactory: contextFactory,
	}
}

func (u *deleteUsecase) Execute(ctx context.Context, siteID, businessTypeID string) (*DeleteOutput, error) {
	app := u.contextFactory()

	siteBusinessType, err := app.Repositories.SiteBusinessType.Delete(ctx, siteID, businessTypeID)
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{
		Data: toSiteBusinessTypeOutputData(&siteBusinessType),
	}, nil
}
