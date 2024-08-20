package business_type

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

func (r *repository) List(ctx context.Context) ([]domain.BusinessType, error) {
	rows, err := r.ExecuteListQuery(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var businessTypes []domain.BusinessType
	for rows.Next() {
		var id, slug, name, color, description string
		var imageURL sql.NullString
		err = rows.Scan(
			&id,
			&slug,
			&name,
			&color,
			&description,
			&imageURL,
		)
		if err != nil {
			return nil, err
		}

		businessTypes = append(businessTypes, domain.BusinessType{
			ID:          id,
			Slug:        slug,
			Name:        name,
			Color:       color,
			Description: description,
			ImageURL:    imageURL.String,
		})
	}

	return businessTypes, nil
}

func (r *repository) ExecuteListQuery(ctx context.Context) (*sql.Rows, error) {
	query := `SELECT id, slug, name, color, description, image_url FROM business_types`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
