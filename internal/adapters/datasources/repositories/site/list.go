package site

import (
	"context"
	"database/sql"

	domain_city "github.com/tapiaw38/cardon-tour-be/internal/domain/location"
	domain_site "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

type (
	ListFilterOptions struct {
		ProvinceID string
	}
)

func (r *repository) List(ctx context.Context, filters ListFilterOptions) ([]domain_site.Site, error) {
	rows, err := r.executeListQuery(ctx, filters)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	siteMap := make(map[string]*domain_site.Site)

	for rows.Next() {
		var id, siteSlug, name, description, cityID string
		var imageURL sql.NullString

		var cityName, citySlug, cityProvinceID sql.NullString
		var cityLatitude, cityLongitude sql.NullFloat64

		var businessTypeID sql.NullString

		if err := rows.Scan(
			&id,
			&siteSlug,
			&name,
			&description,
			&imageURL,
			&cityID,
			&cityName,
			&citySlug,
			&cityProvinceID,
			&cityLatitude,
			&cityLongitude,
			&businessTypeID,
		); err != nil {
			return nil, err
		}

		if site, exists := siteMap[id]; exists {
			if businessTypeID.Valid {
				site.BusinessTypeID = append(site.BusinessTypeID, businessTypeID.String)
			}
		} else {
			siteMap[id] = &domain_site.Site{
				ID:          id,
				Slug:        siteSlug,
				Name:        name,
				Description: description,
				ImageURL:    imageURL.String,
				CityID:      cityID,
				City: &domain_city.City{
					Name:       cityName.String,
					Slug:       citySlug.String,
					ProvinceID: cityProvinceID.String,
					Latitude:   cityLatitude.Float64,
					Longitude:  cityLongitude.Float64,
				},
				BusinessTypeID: []string{},
			}

			if businessTypeID.Valid {
				siteMap[id].BusinessTypeID = append(siteMap[id].BusinessTypeID, businessTypeID.String)
			}
		}
	}

	sites := make([]domain_site.Site, 0, len(siteMap))
	for _, site := range siteMap {
		sites = append(sites, *site)
	}

	return sites, nil
}

func (r *repository) executeListQuery(ctx context.Context, filters ListFilterOptions) (*sql.Rows, error) {
	query := `SELECT
			s.id,
			s.slug,
			s.name,
			s.description,
			s.image_url,
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
		LEFT JOIN cities c ON c.id = s.city_id`

	query += ` WHERE s.id = s.id`

	var args []any

	if filters.ProvinceID != "" {
		query += ` AND c.province_id = $1`
		args = append(args, filters.ProvinceID)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
