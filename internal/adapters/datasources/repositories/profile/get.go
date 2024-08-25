package profile

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
)

func (r *repository) Get(ctx context.Context, id string) (*domain.Profile, error) {
	rows, err := r.ExecuteGetQuery(ctx, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profileID, userID, profileName string
	var profileSitesID []string

	for rows.Next() {
		var siteID sql.NullString
		if err := rows.Scan(
			&profileID,
			&userID,
			&profileName,
			&siteID,
		); err != nil {
			return nil, err
		}

		if siteID.Valid {
			profileSitesID = append(profileSitesID, siteID.String)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	user := unmarshalProfile(profileID, userID, profileName, profileSitesID)

	return user, nil
}

func (r *repository) ExecuteGetQuery(ctx context.Context, id string) (*sql.Rows, error) {
	query := `SELECT
			p.id AS profile_id,
			p.user_id,
			pt.name AS profile_type_name,
			ps.site_id
		FROM
			profiles p
		LEFT JOIN
			profile_types pt ON pt.id = p.profile_type
		LEFT JOIN
			profile_sites ps ON ps.profile_id = p.id
		WHERE
			p.id = $1`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
