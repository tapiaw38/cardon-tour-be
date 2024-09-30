package business

import (
	"context"
	"database/sql"
	"time"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

func (r *repository) Get(ctx context.Context, businessID string) (domain.Business, error) {
	rows, err := r.executeGetQuery(ctx, businessID)
	if err != nil {
		return domain.Business{}, err
	}
	defer rows.Close()

	var business domain.Business

	for rows.Next() {
		var id, profileID, businessTypeID, siteID, name string
		var description, phoneNumber, email, address sql.NullString
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
			&address,
			&phoneNumber,
			&email,
			&active,
			&latitude,
			&longitude,
			&createdAt,
			&imageID,
			&imageURL,
		)
		if err != nil {
			return domain.Business{}, err
		}

		if business.ID == "" {
			business = domain.Business{
				ID:             id,
				ProfileID:      profileID,
				BusinessTypeID: businessTypeID,
				SiteID:         siteID,
				Name:           name,
				Description:    description.String,
				PhoneNumber:    phoneNumber.String,
				Email:          email.String,
				Address:        address.String,
				Active:         active,
				Latitude:       latitude.Float64,
				Longitude:      longitude.Float64,
				CreatedAt:      createdAt,
				Images:         []domain.BusinessImage{},
			}
		}

		if imageID != nil && imageURL != nil {
			business.Images = append(business.Images, domain.BusinessImage{
				ID:  *imageID,
				URL: *imageURL,
			})
		}
	}

	if err = rows.Err(); err != nil {
		return domain.Business{}, err
	}

	if business.ID == "" {
		return domain.Business{}, sql.ErrNoRows
	}

	return business, nil
}

func (r *repository) executeGetQuery(ctx context.Context, businessID string) (*sql.Rows, error) {
	query := `SELECT 
		bs.id,
		bs.profile_id,
		bs.business_type_id,
		bs.site_id,
		bs.name,
		bs.description,
		bs.address,
		bs.phone_number,
		bs.email,
		bs.active,
		bs.latitude,
		bs.longitude,
		bs.created_at,
		bis.id,
		bis.url
	FROM
		businesses bs
	LEFT JOIN business_images bis ON bis.business_id = bs.id
	WHERE
		bs.id = $1`

	rows, err := r.db.QueryContext(ctx, query, businessID)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
