package site

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

func (r *repository) GetByID(ctx context.Context, id string) (*domain.Site, error) {
	rows, err := r.executeGetByIDQuery(ctx, id)
	if err != nil {
		return &domain.Site{}, err
	}

	var siteID, siteSlug, name, description, cityID string
	var imageURL sql.NullString
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
			&businessTypeSlug,
		); err != nil {
			return &domain.Site{}, err
		}

		if businessTypeSlug.Valid {
			businessTypeSlugs = append(businessTypeSlugs, businessTypeSlug.String)
		}
	}

	if err = rows.Err(); err != nil {
		return &domain.Site{}, err
	}

	return &domain.Site{
		ID:                id,
		Slug:              siteSlug,
		Name:              name,
		Description:       description,
		ImageURL:          imageURL.String,
		CityID:            cityID,
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
			bt.slug AS business_type_slug
		FROM sites s
		LEFT JOIN site_business_types sbt ON sbt.site_id = s.id
		LEFT JOIN business_types bt ON bt.id = sbt.business_type_id
		WHERE id = $1`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
