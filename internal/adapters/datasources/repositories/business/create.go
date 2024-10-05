package business

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/utils"
	"time"
)

func (r *repository) Create(ctx context.Context, business domain.Business) (string, error) {
	row, err := r.executeCreateQuery(ctx, business)
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

func (r *repository) executeCreateQuery(ctx context.Context, business domain.Business) (*sql.Row, error) {
	query := `INSERT INTO businesses (
                        id,
                        profile_id, 
                        business_type_id, 
                        site_id, 
                        name, 
                        phone_number, 
                        email, 
                        description,
                        content,
                        address, 
                        active, 
                        latitude, 
                        longitude, 
                        created_at) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
            RETURNING id`

	var name, phoneNumber, email, description, content, address, createdAt *string
	var latitude, longitude *float64

	if business.Name != "" {
		name = utils.ToPointer(business.Name)
	}
	if business.PhoneNumber != "" {
		phoneNumber = utils.ToPointer(business.PhoneNumber)
	}
	if business.Email != "" {
		email = utils.ToPointer(business.Email)
	}
	if business.Description != "" {
		description = utils.ToPointer(business.Description)
	}
	if business.Content != "" {
		content = utils.ToPointer(business.Content)
	}
	if business.Address != "" {
		address = utils.ToPointer(business.Address)
	}
	if !business.CreatedAt.IsZero() {
		createdAt = utils.ToPointer(business.CreatedAt.Format("2006-01-02"))
	} else {
		createdAt = utils.ToPointer(time.Now().Format("2006-01-02"))
	}

	if business.Latitude != 0 {
		latitude = utils.ToPointer(business.Latitude)
	}
	if business.Longitude != 0 {
		longitude = utils.ToPointer(business.Longitude)
	}

	args := []any{
		business.ID,
		business.ProfileID,
		business.BusinessTypeID,
		business.SiteID,
		name,
		phoneNumber,
		email,
		description,
		content,
		address,
		business.Active,
		latitude,
		longitude,
		createdAt,
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
