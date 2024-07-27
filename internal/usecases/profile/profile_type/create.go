package profiletype

import (
	"context"

	"github.com/google/uuid"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
)

type (
	CreateUsecase interface {
		Execute(context.Context, domain.ProfileType) error
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

func (u *createUsecase) Execute(ctx context.Context, profileType domain.ProfileType) error {
	app := u.contextFactory()

	profileTypeID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	profileType.ID = profileTypeID.String()

	return app.Repositories.ProfileType.Create(ctx, profileType)
}
