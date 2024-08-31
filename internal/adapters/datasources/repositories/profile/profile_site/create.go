package profilesite

import "context"

func (r *repository) Create(ctx context.Context, profileID string, siteID string) error {
	err := r.executeCreateQuery(ctx, profileID, siteID)

	return err
}

func (r *repository) executeCreateQuery(ctx context.Context, profileID string, siteID string) error {
	query := `INSERT INTO profile_sites(
					profile_id,
					site_id
				) VALUES ($1, $2)`

	args := []any{
		profileID,
		siteID,
	}

	_, err := r.db.ExecContext(ctx, query, args...)

	return err
}
