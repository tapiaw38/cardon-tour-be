package profilesite

import (
	"context"

	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	CreateUsecase interface {
		Execute(context.Context, string, string) error
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}
)

func NewCreateUseCase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{
		contextFactory: contextFactory,
	}
}

func (u *createUsecase) Execute(ctx context.Context, profileID string, siteID string) error {
	app := u.contextFactory()

	return app.Repositories.ProfileSite.Create(ctx, profileID, siteID)
}
