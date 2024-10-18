package site

import (
	"context"
	"github.com/google/uuid"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	CreateUsecase interface {
		Execute(context.Context, domain.Site) (*CreateOutput, error)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateOutput struct {
		Data SiteOutputData `json:"data"`
	}
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{
		contextFactory: contextFactory,
	}
}

func (u *createUsecase) Execute(ctx context.Context, siteInput domain.Site) (*CreateOutput, error) {
	app := u.contextFactory()

	siteID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	siteInput.ID = siteID.String()
	id, err := app.Repositories.Site.Create(ctx, siteInput)
	if err != nil {
		return nil, err
	}

	site, err := app.Repositories.Site.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{
		Data: toSiteOutputData(site),
	}, nil
}
