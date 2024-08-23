package site

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
		Data []SiteOutputData `json:"data"`
	}
)

func NewListUseCase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{
		contextFactory: contextFactory,
	}
}

func (u *listUsecase) Execute(ctx context.Context) (ListOutput, error) {
	app := u.contextFactory()

	sites, err := app.Repositories.Site.List(ctx)
	if err != nil {
		return ListOutput{}, err
	}

	siteOutputs := make([]SiteOutputData, 0, len(sites))
	for _, site := range sites {
		siteOutputs = append(siteOutputs, toSiteOutputData(&site))
	}

	return ListOutput{
		Data: siteOutputs,
	}, nil
}
