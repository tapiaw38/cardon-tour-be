package site

import (
	"context"
	"database/sql"

	domain_city "github.com/tapiaw38/cardon-tour-be/internal/domain/location"
	domain_site "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

func (r *repository) GetByID(ctx context.Context, id string) (*domain_site.Site, error) {
	rows, err := r.executeGetByIDQuery(ctx, id)
	if err != nil {
		return &domain_site.Site{}, err
	}

	var siteID, siteSlug, name, description, cityID string
	var imageURL sql.NullString

	var cityName, cityCode, cityProvinceID sql.NullString
	var cityLatitude, cityLongitude sql.NullFloat64

	var businessTypeSlugs []string

	for rows.Next() {
		var businessTypeSlug sql.NullString
		if err := rows.Scan(
			&siteID,
			&siteSlug,
			&name,
			&description,
			&imageURL,
			&cityID,
			&cityName,
			&cityCode,
			&cityProvinceID,
			&cityLatitude,
			&cityLongitude,
			&businessTypeSlug,
		); err != nil {
			return &domain_site.Site{}, err
		}

		if businessTypeSlug.Valid {
			businessTypeSlugs = append(businessTypeSlugs, businessTypeSlug.String)
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
		CityID:      cityID,
		City: &domain_city.City{
			Name:       cityName.String,
			Code:       cityCode.String,
			ProvinceID: cityProvinceID.String,
			Latitude:   cityLatitude.Float64,
			Longitude:  cityLongitude.Float64,
		},
		BusinessTypeSlugs: businessTypeSlugs,
	}, nil
}
func (r *repository) executeGetByIDQuery(ctx context.Context, id string) (*sql.Rows, error) {
	query := `SELECT
			s.id, 
			s.slug, 
			s.name, 
			s.description, 
			s.image_url, 
			s.city_id,
			c.name AS city_name,
			c.code AS city_code,
			c.province_id AS city_province_id,
			c.latitude AS city_latitude,
			c.longitude AS city_longitude,
			bt.slug AS business_type_slug
		FROM sites s
		LEFT JOIN site_business_types sbt ON sbt.site_id = s.id
		LEFT JOIN business_types bt ON bt.id = sbt.business_type_id
		LEFT JOIN cities c ON c.id = s.city_id
		WHERE s.id = $1`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
