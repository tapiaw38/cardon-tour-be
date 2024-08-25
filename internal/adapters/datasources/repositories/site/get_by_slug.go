package site

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

func (r *repository) GetBySlug(ctx context.Context, slug string) (*domain.Site, error) {
	row, err := r.executeGetBySlugQuery(ctx, slug)
	if err != nil {
		return &domain.Site{}, err
	}

	var id, siteSlug, name, description, cityID string
	var imageURL sql.NullString
	err = row.Scan(
		&id,
		&siteSlug,
		&name,
		&description,
		&imageURL,
		&cityID,
	)
	if err != nil {
		return &domain.Site{}, err
	}

	return &domain.Site{
		ID:          id,
		Slug:        siteSlug,
		Name:        name,
		Description: description,
		ImageURL:    imageURL.String,
		CityID:      cityID,
	}, nil
}

func (r *repository) executeGetBySlugQuery(ctx context.Context, slug string) (*sql.Row, error) {
	query := `SELECT 
			id, slug, name, description, image_url, city_id
		FROM sites
		WHERE slug = $1`

	row := r.db.QueryRowContext(ctx, query, slug)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
