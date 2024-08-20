package business_type

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

func (r *repository) Update(ctx context.Context, id string, businessType domain.BusinessType) (string, error) {
	row, err := r.executeUpdateQuery(ctx, id, businessType)
	if err != nil {
		return "", err
	}

	var updatedID string
	err = row.Scan(&updatedID)

	return updatedID, err
}

func (r *repository) executeUpdateQuery(ctx context.Context, id string, businessType domain.BusinessType) (*sql.Row, error) {
	query := `UPDATE business_types
	SET
		name = COALESCE($1, name),
		color = COALESCE($2, color),
		description = COALESCE($3, description),
		image_url = COALESCE($4, image_url)
	WHERE
		id = $5`

	args := []any{
		businessType.Name,
		businessType.Color,
		businessType.Description,
		businessType.ImageURL,
		id,
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
