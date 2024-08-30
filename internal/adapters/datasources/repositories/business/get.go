package business

import (
	"context"
	"database/sql"
	"time"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

func (r *repository) Get(ctx context.Context, id string) (domain.Business, error) {
	row, err := r.executeGetQuery(ctx, id)
	if err != nil {
		return domain.Business{}, err
	}

	var businessID, profileID, businessTypeID, siteID, name string
	var description, phoneNumber, email sql.NullString
	var active bool
	var latitude, longitude sql.NullFloat64
	var createdAt time.Time
	err = row.Scan(
		&businessID,
		&profileID,
		&businessTypeID,
		&siteID,
		&name,
		&description,
		&phoneNumber,
		&email,
		&active,
		&latitude,
		&longitude,
		&createdAt,
	)
	if err != nil {
		return domain.Business{}, err
	}

	return domain.Business{
		ID:             businessID,
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
	}, nil
}

func (r *repository) executeGetQuery(ctx context.Context, id string) (*sql.Row, error) {
	query := `SELECT 
		id,
		profile_id,
		business_type_id,
		site_id,
		name,
		description,
		phone_number,
		email,
		active,
		latitude,
		longitude,
		created_at
	FROM
		businesses
	WHERE
		id = $1`

	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
