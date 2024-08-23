package site

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

func (r *repository) List(ctx context.Context) ([]domain.Site, error) {
	rows, err := r.executeListQuery(ctx)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sites []domain.Site
	for rows.Next() {
		var id, slug, name, description string
		var imageURL sql.NullString
		err = rows.Scan(
			&id,
			&slug,
			&name,
			&description,
			&imageURL,
		)
		if err != nil {
			return nil, err
		}

		sites = append(sites, domain.Site{
			ID:          id,
			Slug:        slug,
			Name:        name,
			Description: description,
			ImageURL:    imageURL.String,
		})
	}

	return sites, nil
}

func (r *repository) executeListQuery(ctx context.Context) (*sql.Rows, error) {
	query := `SELECT id, slug, name, description, image_url FROM sites`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
