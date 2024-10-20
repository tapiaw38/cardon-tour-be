package site_business_type

import (
	"context"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

func (r *repository) Delete(ctx context.Context, siteID, businessTypeID string) (domain.SiteBusinessType, error) {
	query := `DELETE 
			FROM site_business_types 
			WHERE site_id = $1 AND business_type_id = $2
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
