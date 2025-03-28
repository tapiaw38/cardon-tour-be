package site

import (
	"context"
	"database/sql"

	domain_city "github.com/tapiaw38/cardon-tour-be/internal/domain/location"
	domain_site "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

func (r *repository) GetBySlug(ctx context.Context, slug string) (*domain_site.Site, error) {
	rows, err := r.executeGetBySlugQuery(ctx, slug)
	if err != nil {
		return &domain_site.Site{}, err
	}

	var id, siteSlug, name, description, cityID string
	var imageURL sql.NullString
	var isPromoted bool

	var cityName, citySlug, cityProvinceID sql.NullString
	var cityLatitude, cityLongitude sql.NullFloat64

	businessTypeIDs := []string{}
	for rows.Next() {
		var businessTypeID sql.NullString
		if err := rows.Scan(
			&id,
			&siteSlug,
			&name,
			&description,
			&imageURL,
			&isPromoted,
			&cityID,
			&cityName,
			&citySlug,
			&cityProvinceID,
			&cityLatitude,
			&cityLongitude,
			&businessTypeID,
		); err != nil {
			return &domain_site.Site{}, err
		}

		if businessTypeID.Valid {
			businessTypeIDs = append(businessTypeIDs, businessTypeID.String)
		}
	}

	if err = rows.Err(); err != nil {
		return &domain_site.Site{}, err
	}

	return &domain_site.Site{
		ID:          id,
		Slug:        siteSlug,
		Name:        name,
		Description: description,
		ImageURL:    imageURL.String,
		IsPromoted:  isPromoted,
		CityID:      cityID,
		City: &domain_city.City{
			Name:       cityName.String,
			Slug:       citySlug.String,
			ProvinceID: cityProvinceID.String,
			Latitude:   cityLatitude.Float64,
			Longitude:  cityLongitude.Float64,
		},
		BusinessTypeID: businessTypeIDs,
	}, nil
}

func (r *repository) executeGetBySlugQuery(ctx context.Context, slug string) (*sql.Rows, error) {
	query := `SELECT
			s.id,
			s.slug,
			s.name,
			s.description,
			s.image_url,
			s.is_promoted,
			s.city_id,
			c.name AS city_name,
			c.slug AS city_slug,
			c.province_id AS city_province_id,
			c.latitude AS city_latitude,
			c.longitude AS city_longitude,
			bt.id AS business_type_id
		FROM sites s
		LEFT JOIN site_business_types sbt ON sbt.site_id = s.id
		LEFT JOIN business_types bt ON bt.id = sbt.business_type_id
		LEFT JOIN cities c ON c.id = s.city_id
		WHERE s.slug = $1`

	rows, err := r.db.QueryContext(ctx, query, slug)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
