package location

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/location"
)

func (r *repository) CreateProvince(ctx context.Context, province domain.Province, countryID string) (string, error) {
	row, err := r.executeCreateProvinceQuery(ctx, province, countryID)
	if err != nil {
		return "", err
	}

	var insertedID string
	err = row.Scan(&insertedID)

	return insertedID, err
}

func (r *repository) executeCreateProvinceQuery(ctx context.Context, province domain.Province, countryID string) (*sql.Row, error) {
	query := `INSERT INTO provinces(
					id,
					name,
					slug,
					country_id
				) VALUES ($1, $2, $3, $4) RETURNING id`

	args := []any{
		province.ID,
		province.Name,
		province.Slug,
		countryID,
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
