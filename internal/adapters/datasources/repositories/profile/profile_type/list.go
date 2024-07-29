package profiletype

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
)

func (r *repository) List(ctx context.Context) ([]domain.ProfileType, error) {
	rows, err := r.executeListQuery(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profileTypes []domain.ProfileType
	for rows.Next() {
		var profileType domain.ProfileType
		err = rows.Scan(
			&profileType.ID,
			&profileType.Name,
		)
		if err != nil {
			return nil, err
		}

		profileTypes = append(profileTypes, profileType)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return profileTypes, nil
}

func (r *repository) executeListQuery(ctx context.Context) (*sql.Rows, error) {
	query := `SELECT id, name FROM profile_types`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
