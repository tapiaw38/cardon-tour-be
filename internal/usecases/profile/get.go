package profile

import (
	"context"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	GetUsecase interface {
		Execute(context.Context, string) (GetOutput, error)
	}

	getUsecase struct {
		contextFactory appcontext.Factory
	}

	GetOutput struct {
		Data ProfileOutputData `json:"data"`
	}
)

func NewGetUseCase(contextFactory appcontext.Factory) GetUsecase {
	return &getUsecase{
		contextFactory: contextFactory,
	}
}

func (u *getUsecase) Execute(ctx context.Context, id string) (GetOutput, error) {
	app := u.contextFactory()

	profile, err := app.Repositories.Profile.Get(ctx, id)
	if err != nil {
		return GetOutput{}, err
	}

	var profileSites []domain.Site
	for _, siteID := range profile.ProfileSitesID {
		site, err := app.Repositories.Site.Get(ctx, siteID)
		if err != nil {
			return GetOutput{}, err
		}
		profileSites = append(profileSites, *site)
	}

	profile.ProfileSites = profileSites

	return GetOutput{
		Data: toProfileOutputData(profile),
	}, nil
}
