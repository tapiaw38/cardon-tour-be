package profile

import (
	"context"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/utils"
)

func (r *repository) Update(ctx context.Context, id string, profile domain.Profile) error {
	return r.executeUpdateQuery(ctx, id, profile)
}

func (r *repository) executeUpdateQuery(ctx context.Context, id string, profile domain.Profile) error {
	query := `UPDATE profile
				SET
					profile_type_id = COALESCE($1, profile_type_id)
				WHERE
					id = $2`

	var profileID *string

	if profile.ProfileTypeID != "" {
		profileID = utils.ToPointer(profile.ProfileTypeID)
	}

	args := []any{
		profileID,
		id,
	}

	_, err := r.db.ExecContext(ctx, query, args...)

	return err
}
