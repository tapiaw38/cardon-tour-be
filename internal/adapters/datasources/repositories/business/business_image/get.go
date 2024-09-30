package business_image

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

func (r *repository) Get(ctx context.Context, id string) (domain.BusinessImage, error) {
	row, err := r.executeGetQuery(ctx, id)
	if err != nil {
		return domain.BusinessImage{}, err
	}

	var businessImageID, businessID, url string
	err = row.Scan(&businessImageID, &businessID, &url)
	if err != nil {
		return domain.BusinessImage{}, err
	}

	return domain.BusinessImage{
		ID:         businessImageID,
		BusinessID: businessID,
		URL:        url,
	}, nil
}

func (r *repository) executeGetQuery(ctx context.Context, id string) (*sql.Row, error) {
	query := `SELECT 
		id,
		business_id,
		url
	FROM
		business_images
	WHERE
		id = $1`

	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
