package location

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/location"
)

func (r *repository) ListProvince(ctx context.Context) ([]domain.Province, error) {
	rows, err := r.executeListProvinceQuery(ctx)
	if err != nil {
		return nil, err
	}

	var provinces []domain.Province
	for rows.Next() {
		var ID, name, slug, countryID string
		var imageURL, description sql.NullString
		var latitude, longitude sql.NullFloat64
		err = rows.Scan(
			&ID,
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

		provinces = append(provinces, domain.Province{
			ID:          ID,
			Name:        name,
			Slug:        slug,
			CountryID:   countryID,
			ImageURL:    imageURL.String,
			Latitude:    latitude.Float64,
			Longitude:   longitude.Float64,
			Description: description.String,
		})
	}

	return provinces, nil
}

func (r *repository) executeListProvinceQuery(ctx context.Context) (*sql.Rows, error) {
	query := `SELECT
			id,
			name,
			slug,
			country_id,
			image_url,
			latitude,
			longitude,
			description
		FROM provinces;`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
