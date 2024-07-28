package profile

import (
	"context"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
)

func (r *repository) Create(ctx context.Context, profile domain.Profile) (string, error) {
	insertedID, err := r.executeCreateQuery(ctx, profile)
	if err != nil {
		return "", err
	}

	return insertedID, nil
}

func (r *repository) executeCreateQuery(ctx context.Context, profile domain.Profile) (string, error) {
	query := `INSERT INTO profiles(
					id, user_id, profile_type
				) VALUES ($1, $2, $3) RETURNING id`

	args := []any{
		profile.ID,
		profile.UserID,
		profile.ProfileTypeID,
	}

	row := r.db.QueryRowContext(ctx, query, args...)

	var lastID string

	err := row.Scan(&lastID)

	return lastID, err
}
