package profile

import (
	"context"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	GetByUserIDUsecase interface {
		Execute(context.Context, string) (GetByUserIDOutput, error)
	}

	getByUserIDUsecase struct {
		contextFactory appcontext.Factory
	}

	GetByUserIDOutput struct {
		Data ProfileOutputData `json:"data"`
	}
)

func NewGetByUserIDUseCase(contextFactory appcontext.Factory) GetByUserIDUsecase {
	return &getByUserIDUsecase{
		contextFactory: contextFactory,
	}
}

func (u *getByUserIDUsecase) Execute(ctx context.Context, userID string) (GetByUserIDOutput, error) {
	app := u.contextFactory()

	profile, err := app.Repositories.Profile.GetByUserID(ctx, userID)
	if err != nil {
		return GetByUserIDOutput{}, err
	}

	var profileSites []domain.Site
	for _, siteID := range profile.ProfileSitesID {
		site, err := app.Repositories.Site.Get(ctx, siteID)
		if err != nil {
			return GetByUserIDOutput{}, err
		}
		profileSites = append(profileSites, *site)
	}

	profile.ProfileSites = profileSites

	return GetByUserIDOutput{
		Data: toProfileOutputData(profile),
	}, nil
}
