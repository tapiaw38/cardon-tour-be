package profilesite

import "context"

func (r *repository) Delete(ctx context.Context, profileID string, siteID string) error {
	err := r.executeDeleteQuery(ctx, profileID, siteID)

	return err
}

func (r *repository) executeDeleteQuery(ctx context.Context, profileID string, siteID string) error {
	query := `DELETE FROM profile_sites WHERE profile_id = $1 AND site_id = $2`

	args := []any{
		profileID,
		siteID,
	}

	_, err := r.db.ExecContext(ctx, query, args...)

	return err
}
