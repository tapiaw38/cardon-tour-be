package site

import (
	"context"
	"database/sql"
	domain_site "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/utils"
)

func (r *repository) Create(ctx context.Context, site domain_site.Site) (string, error) {
	row, err := r.executeCreateQuery(ctx, site)
	if err != nil {
		return "", err
	}

	var id string
	err = row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *repository) executeCreateQuery(ctx context.Context, site domain_site.Site) (*sql.Row, error) {
	query := `INSERT INTO sites (
                        id,
                        slug,
                        name,
                        description,
                        image_url,
                        city_id,
                        is_promoted)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
            RETURNING id`
	var name, description, imageURL, cityID *string
	var isPromoted *bool

	if site.Name != "" {
		name = utils.ToPointer(site.Name)
	}
	if site.Description != "" {
		description = utils.ToPointer(site.Description)
	}
	if site.ImageURL != "" {
		imageURL = utils.ToPointer(site.ImageURL)
	}
	if site.CityID != "" {
		cityID = utils.ToPointer(site.CityID)
	}
	if site.IsPromoted {
		isPromoted = &site.IsPromoted
	}

	args := []any{
		site.ID,
		site.Slug,
		name,
		description,
		imageURL,
		cityID,
		isPromoted,
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
