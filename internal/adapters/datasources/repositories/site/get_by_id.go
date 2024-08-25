package site

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

func (r *repository) GetByID(ctx context.Context, id string) (*domain.Site, error) {
	row, err := r.executeGetByIDQuery(ctx, id)
	if err != nil {
		return &domain.Site{}, err
	}

	var siteID, siteSlug, name, description, CityID string
	var imageURL sql.NullString
	err = row.Scan(
		&siteID,
		&siteSlug,
		&name,
		&description,
		&imageURL,
		&CityID,
	)
	if err != nil {
		return &domain.Site{}, err
	}

	return &domain.Site{
		ID:          siteID,
		Slug:        siteSlug,
		Name:        name,
		Description: description,
		ImageURL:    imageURL.String,
		CityID:      CityID,
	}, nil
}

func (r *repository) executeGetByIDQuery(ctx context.Context, id string) (*sql.Row, error) {
	query := `SELECT
			id, slug, name, description, image_url, city_id
		FROM sites
		WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
