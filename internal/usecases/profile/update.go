package profile

import (
	"context"

	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories/profile"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
)

type (
	UpdateUsecase interface {
		Execute(context.Context, string, domain.Profile) (UpdateOutput, error)
	}

	updateUsecase struct {
		repositories profile.Repository
	}

	UpdateOutput struct {
		Data ProfileOutputData `json:"data"`
	}
)

func NewUpdateUseCase(repository profile.Repository) UpdateUsecase {
	return &updateUsecase{
		repositories: repository,
	}
}

func (u *updateUsecase) Execute(ctx context.Context, id string, profile domain.Profile) (UpdateOutput, error) {
	err := u.repositories.Update(ctx, id, profile)
	if err != nil {
		return UpdateOutput{}, err
	}

	profileUpdated, err := u.repositories.Get(ctx, id)
	if err != nil {
		return UpdateOutput{}, err
	}

	return UpdateOutput{
		Data: toProfileOutputData(profileUpdated),
	}, nil
}
