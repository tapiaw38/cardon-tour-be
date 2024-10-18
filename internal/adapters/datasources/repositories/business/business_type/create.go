package business_type

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/utils"
)

func (r *repository) Create(ctx context.Context, businessType domain.BusinessType) (string, error) {
	row, err := r.executeCreateQuery(ctx, businessType)
	if err != nil {
		return "", err
	}

	var insertedID string
	err = row.Scan(&insertedID)

	return insertedID, err
}

func (r *repository) executeCreateQuery(ctx context.Context, businessType domain.BusinessType) (*sql.Row, error) {
	query := `INSERT INTO business_types(
					id, slug, name, color, description, image_url
				) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var name, color, description, imageURL *string

	if businessType.Name != "" {
		name = utils.ToPointer(businessType.Name)
	}
	if businessType.Color != "" {
		color = utils.ToPointer(businessType.Color)
	}
	if businessType.Description != "" {
		description = utils.ToPointer(businessType.Description)
	}
	if businessType.ImageURL != "" {
		imageURL = utils.ToPointer(businessType.ImageURL)
	}

	args := []any{
		businessType.ID,
		businessType.Slug,
		name,
		color,
		description,
		imageURL,
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
