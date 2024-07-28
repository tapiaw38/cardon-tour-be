package profiletype

import (
	"context"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
)

func (r *repository) Create(ctx context.Context, profileType domain.ProfileType) error {
	return r.executeCreateQuery(ctx, profileType)
}

func (r *repository) executeCreateQuery(ctx context.Context, profileType domain.ProfileType) error {
	query := `INSERT INTO profile_types(
					id,
					name
				) VALUES ($1, $2)`

	args := []any{
		profileType.ID,
		profileType.Name,
	}

	_, err := r.db.ExecContext(ctx, query, args...)

	return err
}
