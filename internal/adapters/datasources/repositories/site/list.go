package site

import (
	"context"
	"database/sql"

	domain_city "github.com/tapiaw38/cardon-tour-be/internal/domain/location"
	domain_site "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

func (r *repository) List(ctx context.Context) ([]domain_site.Site, error) {
	rows, err := r.executeListQuery(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	siteMap := make(map[string]*domain_site.Site)

	for rows.Next() {
		var id, siteSlug, name, description, cityID string
		var imageURL sql.NullString

		var cityName, cityCode, cityProvinceID sql.NullString
		var cityLatitude, cityLongitude sql.NullFloat64

		var businessTypeSlug sql.NullString

		if err := rows.Scan(
			&id,
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
			return nil, err
		}

		if site, exists := siteMap[id]; exists {
			if businessTypeSlug.Valid {
				site.BusinessTypeSlugs = append(site.BusinessTypeSlugs, businessTypeSlug.String)
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
					Code:       cityCode.String,
					ProvinceID: cityProvinceID.String,
					Latitude:   cityLatitude.Float64,
					Longitude:  cityLongitude.Float64,
				},
				BusinessTypeSlugs: []string{},
			}

			if businessTypeSlug.Valid {
				siteMap[id].BusinessTypeSlugs = append(siteMap[id].BusinessTypeSlugs, businessTypeSlug.String)
			}
		}
	}

	sites := make([]domain_site.Site, 0, len(siteMap))
	for _, site := range siteMap {
		sites = append(sites, *site)
	}

	return sites, nil
}

func (r *repository) executeListQuery(ctx context.Context) (*sql.Rows, error) {
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
		LEFT JOIN cities c ON c.id = s.city_id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
