package profilesite

import (
	"context"
	"database/sql"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
)

func (r *repository) Delete(ctx context.Context, profileID string, siteID string) (domain.ProfileSite, error) {
	row, err := r.executeDeleteQuery(ctx, profileID, siteID)
	if err != nil {
		return domain.ProfileSite{}, err
	}

	var profile, site string
	err = row.Scan(&profile, &site)

	return domain.ProfileSite{
		ProfileID: profile,
		SiteID:    site,
	}, err
}

func (r *repository) executeDeleteQuery(ctx context.Context, profileID string, siteID string) (*sql.Row, error) {
	query := `DELETE FROM profile_sites 
		WHERE profile_id = $1 AND site_id = $2
		RETURNING profile_id, site_id`

	args := []any{
		profileID,
		siteID,
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
