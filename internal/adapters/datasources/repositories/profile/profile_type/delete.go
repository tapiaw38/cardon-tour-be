package profiletype

import "context"

func (r *repository) Delete(ctx context.Context, id string) error {
	return r.executeDeleteQuery(ctx, id)
}

func (r *repository) executeDeleteQuery(ctx context.Context, id string) error {
	query := `DELETE FROM profile_types WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
