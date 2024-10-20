package site_business_type

import (
	"context"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

func (r *repository) Create(ctx context.Context, siteID, businessTypeID string) (domain.SiteBusinessType, error) {
	query := `INSERT INTO 
    			site_business_types (site_id, business_type_id) 
			VALUES ($1, $2)
			RETURNING site_id, business_type_id
			`
	var businessType, site string

	args := []any{
		siteID,
		businessTypeID,
	}

	err := r.db.QueryRowContext(ctx, query, args...).Scan(&site, &businessType)
	if err != nil {
		return domain.SiteBusinessType{}, err
	}

	return domain.SiteBusinessType{
		SiteID:         site,
		BusinessTypeID: businessType,
	}, nil
}
