package location

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/location"
)

func (r *repository) CreateCity(ctx context.Context, city domain.City, provinceID string) (string, error) {
	row, err := r.executeCreateCityQuery(ctx, city, provinceID)
	if err != nil {
		return "", err
	}

	var insertedID string
	err = row.Scan(&insertedID)

	return insertedID, err
}

func (r *repository) executeCreateCityQuery(ctx context.Context, city domain.City, provinceID string) (*sql.Row, error) {
	query := `INSERT INTO cities(
					id,
					name,
					code,
					province_id
				) VALUES ($1, $2, $3, $4) RETURNING id`

	args := []any{
		city.ID,
		city.Name,
		city.Code,
		provinceID,
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
