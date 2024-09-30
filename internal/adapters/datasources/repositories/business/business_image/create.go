package business_image

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

func (r *repository) Create(ctx context.Context, businessImage domain.BusinessImage) (string, error) {
	row, err := r.executeCreateQuery(ctx, businessImage)
	if err != nil {
		return "", err
	}

	var id string
	err = row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *repository) executeCreateQuery(ctx context.Context, businessImage domain.BusinessImage) (*sql.Row, error) {
	query := `INSERT INTO business_images (
                        id,
                        business_id,
                        url) 
			VALUES ($1, $2, $3)
            RETURNING id`

	args := []any{
		businessImage.ID,
		businessImage.BusinessID,
		businessImage.URL,
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
