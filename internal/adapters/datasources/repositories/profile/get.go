package profile

import (
	"context"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
)

func (r *repository) Get(ctx context.Context, id string) (*domain.Profile, error) {
	user, err := r.ExecuteGetQuery(ctx, id)

	return user, err
}

func (r *repository) ExecuteGetQuery(ctx context.Context, id string) (*domain.Profile, error) {
	query := `SELECT
			p.id, p.user_id, pt.name
		FROM
			profiles p
		LEFT JOIN
			profile_types pt ON pt.id = p.profile_type
		WHERE
			p.id = $1`

	row := r.db.QueryRowContext(ctx, query, id)

	var profileID, userID, profileName string

	err := row.Scan(&profileID, &userID, &profileName)

	return unmarshalProfile(profileID, userID, profileName), err
}
