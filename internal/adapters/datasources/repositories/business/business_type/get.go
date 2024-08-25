package business_type

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

func (r *repository) Get(ctx context.Context, id string) (domain.BusinessType, error) {
	row, err := r.executeGetQuery(ctx, id)
	if err != nil {
		return domain.BusinessType{}, err
	}

	var businessTypeID, slug, name, color, description string
	var imageURL sql.NullString
	err = row.Scan(
		&businessTypeID,
		&slug,
		&name,
		&color,
		&description,
		&imageURL,
	)
	if err != nil {
		return domain.BusinessType{}, err
	}

	return domain.BusinessType{
		ID:          businessTypeID,
		Slug:        slug,
		Name:        name,
		Color:       color,
		Description: description,
		ImageURL:    imageURL.String,
	}, nil
}

func (r *repository) executeGetQuery(ctx context.Context, id string) (*sql.Row, error) {
	query := `SELECT id, slug, name, color, description, image_url FROM business_types WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
