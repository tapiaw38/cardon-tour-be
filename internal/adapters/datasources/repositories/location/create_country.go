package location

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/location"
)

func (r *repository) CreateCountry(ctx context.Context, country domain.Country) (string, error) {
	row, err := r.executeCreateCountryQuery(ctx, country)
	if err != nil {
		return "", err
	}

	var insertedID string
	err = row.Scan(&insertedID)

	return insertedID, err
}

func (r *repository) executeCreateCountryQuery(ctx context.Context, country domain.Country) (*sql.Row, error) {
	query := `INSERT INTO countries(
					id,
					name,
					code
				) VALUES ($1, $2, $3) RETURNING id`

	args := []any{
		country.ID,
		country.Name,
		country.Code,
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
