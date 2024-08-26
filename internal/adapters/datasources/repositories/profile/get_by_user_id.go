package profile

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
)

func (r *repository) GetByUserID(ctx context.Context, userID string) (*domain.Profile, error) {
	rows, err := r.executeGetByUserIDQuery(ctx, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profileID, profileUserID, profileName string
	var profileSitesID []string

	for rows.Next() {
		var siteID sql.NullString
		if err := rows.Scan(
			&profileID,
			&profileUserID,
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

	user := unmarshalProfile(profileID, profileUserID, profileName, profileSitesID)

	return user, nil
}

func (r *repository) executeGetByUserIDQuery(ctx context.Context, userID string) (*sql.Rows, error) {
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
			p.user_id = $1`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
