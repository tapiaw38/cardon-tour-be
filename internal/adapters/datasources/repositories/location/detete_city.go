package location

import "context"

func (r *repository) DeleteCity(ctx context.Context, id string) error {
	return r.executeDeleteCityQuery(ctx, id)
}

func (r *repository) executeDeleteCityQuery(ctx context.Context, id string) error {
	query := `DELETE FROM cities WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
