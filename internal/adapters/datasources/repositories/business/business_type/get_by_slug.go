package business_type

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

func (r *repository) GetBySlug(ctx context.Context, slug string) (domain.BusinessType, error) {
	row, err := r.executeGetBySlugQuery(ctx, slug)
	if err != nil {
		return domain.BusinessType{}, err
	}

	var id, name, color, description string
	var imageURL sql.NullString
	err = row.Scan(
		&id,
		&slug,
		&name,
		&color,
		&description,
		&imageURL,
	)

	if err != nil {
		return domain.BusinessType{}, err
	}

	return domain.BusinessType{
		ID:          id,
		Slug:        slug,
		Name:        name,
		Color:       color,
		Description: description,
		ImageURL:    imageURL.String,
	}, nil
}

func (r *repository) executeGetBySlugQuery(ctx context.Context, slug string) (*sql.Row, error) {
	query := `SELECT id, slug, name, color, description, image_url FROM business_types WHERE slug = $1`

	row := r.db.QueryRowContext(ctx, query, slug)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
