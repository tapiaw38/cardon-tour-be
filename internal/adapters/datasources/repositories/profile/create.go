package profile

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
)

func (r *repository) Create(ctx context.Context, profile domain.Profile) (string, error) {
	row, err := r.executeCreateQuery(ctx, profile)
	if err != nil {
		return "", err
	}

	var insertedID string
	err = row.Scan(&insertedID)

	return insertedID, err
}

func (r *repository) executeCreateQuery(ctx context.Context, profile domain.Profile) (*sql.Row, error) {
	query := `INSERT INTO profiles(
					id, user_id, profile_type
				) VALUES ($1, $2, $3) RETURNING id`

	args := []any{
		profile.ID,
		profile.UserID,
		profile.ProfileTypeID,
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
