package location

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/location"
)

func (r *repository) GetProvinceBySlug(ctx context.Context, provincesSlug string) (*domain.Province, error) {
	row, err := r.executeGetProvinceBySlugQuery(ctx, provincesSlug)
	if err != nil {
		return nil, err
	}

	var id, name, slug, countryID string
	var imageURL, description sql.NullString
	var latitude, longitude sql.NullFloat64

	err = row.Scan(
		&id,
		&name,
		&slug,
		&countryID,
		&imageURL,
		&latitude,
		&longitude,
		&description,
	)

	if err != nil {
		return nil, err
	}

	return &domain.Province{
		ID:          id,
		Name:        name,
		Slug:        slug,
		CountryID:   countryID,
		ImageURL:    imageURL.String,
		Latitude:    latitude.Float64,
		Longitude:   longitude.Float64,
		Description: description.String,
	}, nil
}

func (r *repository) executeGetProvinceBySlugQuery(ctx context.Context, slug string) (*sql.Row, error) {
	query := `SELECT
			id,
			name,
			slug,
			country_id,
			image_url,
			latitude,
			longitude,
			description
		FROM provinces
		WHERE slug = $1`

	row := r.db.QueryRowContext(ctx, query, slug)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
