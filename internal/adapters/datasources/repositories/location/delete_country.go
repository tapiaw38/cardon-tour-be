package location

import "context"

func (r *repository) DeleteCountry(ctx context.Context, id string) error {
	return r.executeDeleteCountryQuery(ctx, id)
}

func (r *repository) executeDeleteCountryQuery(ctx context.Context, id string) error {
	query := `DELETE FROM countries WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
