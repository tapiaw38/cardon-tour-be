package profilesite

import (
	"context"

	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	DeleteUsecase interface {
		Execute(context.Context, string, string) error
	}

	deleteUsecase struct {
		contextFactory appcontext.Factory
	}
)

func NewDeleteUseCase(contextFactory appcontext.Factory) DeleteUsecase {
	return &deleteUsecase{
		contextFactory: contextFactory,
	}
}

func (u *deleteUsecase) Execute(ctx context.Context, profileID string, siteID string) error {
	app := u.contextFactory()

	return app.Repositories.ProfileSite.Delete(ctx, profileID, siteID)
}
