package location

import "context"

func (r *repository) DeleteProvince(ctx context.Context, id string) error {
	return r.executeDeleteProvinceQuery(ctx, id)
}

func (r *repository) executeDeleteProvinceQuery(ctx context.Context, id string) error {
	query := `DELETE FROM provinces WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
