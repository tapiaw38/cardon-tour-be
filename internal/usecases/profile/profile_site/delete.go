package profilesite

import (
	"context"

	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	DeleteUsecase interface {
		Execute(context.Context, string, string) (DeleteOutput, error)
	}

	deleteUsecase struct {
		contextFactory appcontext.Factory
	}

	DeleteOutput struct {
		Data ProfileSiteOutputData `json:"data"`
	}
)

func NewDeleteUseCase(contextFactory appcontext.Factory) DeleteUsecase {
	return &deleteUsecase{
		contextFactory: contextFactory,
	}
}

func (u *deleteUsecase) Execute(ctx context.Context, profileID string, siteID string) (DeleteOutput, error) {
	app := u.contextFactory()

	profileSite, err := app.Repositories.ProfileSite.Delete(ctx, profileID, siteID)
	if err != nil {
		return DeleteOutput{}, err
	}

	return DeleteOutput{
		Data: toProfileSiteOutputData(profileSite),
	}, nil
}
