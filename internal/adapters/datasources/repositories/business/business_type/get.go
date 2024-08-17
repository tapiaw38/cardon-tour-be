package business_type

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

func (r *repository) Get(ctx context.Context, id string) (domain.BusinessType, error) {
	row, err := r.executeGetQuery(ctx, id)
	if err != nil {
		return domain.BusinessType{}, err
	}

	var businessType domain.BusinessType
	err = row.Scan(
		&businessType.ID,
		&businessType.Slug,
		&businessType.Name,
		&businessType.Color,
		&businessType.Description,
		&businessType.ImageURL,
	)

	return businessType, err
}

func (r *repository) executeGetQuery(ctx context.Context, id string) (*sql.Row, error) {
	query := `SELECT id, slug, name, color, description, image_url FROM business_types WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
