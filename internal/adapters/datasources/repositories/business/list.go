package business

import (
	"context"
	"database/sql"
	"time"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

func (r *repository) List(ctx context.Context, filter ListFilterOptions) ([]domain.Business, error) {
	rows, err := r.executeListQuery(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	businessMap := make(map[string]*domain.Business)

	for rows.Next() {
		var id, profileID, businessTypeID, siteID, name string
		var description, address, phoneNumber, email sql.NullString
		var active bool
		var latitude, longitude sql.NullFloat64
		var createdAt time.Time
		var imageID, imageURL *string

		err = rows.Scan(
			&id,
			&profileID,
			&businessTypeID,
			&siteID,
			&name,
			&description,
			&phoneNumber,
			&email,
			&address,
			&active,
			&latitude,
			&longitude,
			&createdAt,
			&imageID,
			&imageURL,
		)
		if err != nil {
			return nil, err
		}

		if business, exists := businessMap[id]; exists {
			if imageID != nil && imageURL != nil {
				business.Images = append(business.Images, domain.BusinessImage{
					ID:  *imageID,
					URL: *imageURL,
				})
			}
		} else {
			businessMap[id] = &domain.Business{
				ID:             id,
				ProfileID:      profileID,
				BusinessTypeID: businessTypeID,
				SiteID:         siteID,
				Name:           name,
				Description:    description.String,
				PhoneNumber:    phoneNumber.String,
				Email:          email.String,
				Active:         active,
				Latitude:       latitude.Float64,
				Longitude:      longitude.Float64,
				CreatedAt:      createdAt,
				Images:         []domain.BusinessImage{},
			}

			if imageID != nil && imageURL != nil {
				businessMap[id].Images = append(businessMap[id].Images, domain.BusinessImage{
					ID:  *imageID,
					URL: *imageURL,
				})
			}
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	businessList := make([]domain.Business, 0, len(businessMap))
	for _, business := range businessMap {
		businessList = append(businessList, *business)
	}

	return businessList, nil
}

func (r *repository) executeListQuery(ctx context.Context, filter ListFilterOptions) (*sql.Rows, error) {
	query := `SELECT 
		bs.id,
		bs.profile_id,
		bs.business_type_id,
		bs.site_id,
		bs.name,
		bs.description,
		bs.phone_number,
		bs.email,
		bs.address,
		bs.active,
		bs.latitude,
		bs.longitude,
		bs.created_at,
		bis.id,
		bis.url
	FROM
		businesses bs
	LEFT JOIN business_images bis ON bis.business_id = bs.id`

	query += " WHERE bs.id = bs.id"

	var args []any

	if filter.SiteSlug != "" {
		query += " AND bs.site_id = (SELECT id FROM sites WHERE slug = $1)"
		args = append(args, filter.SiteSlug)
	}
	if filter.BusinessTypeSlug != "" {
		query += " AND bs.business_type_id = (SELECT id FROM business_types WHERE slug = $2)"
		args = append(args, filter.BusinessTypeSlug)
	}

	query += " ORDER BY bs.id DESC;"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
