package profilesite

import (
	"context"

	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	CreateUsecase interface {
		Execute(context.Context, string, string) (CreateOutput, error)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateOutput struct {
		Data ProfileSiteOutputData `json:"data"`
	}
)

func NewCreateUseCase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{
		contextFactory: contextFactory,
	}
}

func (u *createUsecase) Execute(ctx context.Context, profileID string, siteID string) (CreateOutput, error) {
	app := u.contextFactory()

	profileSite, err := app.Repositories.ProfileSite.Create(ctx, profileID, siteID)
	if err != nil {
		return CreateOutput{}, err
	}

	return CreateOutput{
		Data: toProfileSiteOutputData(profileSite),
	}, nil
}
