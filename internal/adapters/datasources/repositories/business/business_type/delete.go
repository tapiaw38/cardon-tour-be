package business_type

import (
	"context"
	"database/sql"
)

func (r *repository) Delete(ctx context.Context, id string) (string, error) {
	row, err := r.executeDeleteQuery(ctx, id)
	if err != nil {
		return "", err
	}

	var deletedID string
	err = row.Scan(&deletedID)

	return deletedID, err
}

func (r *repository) executeDeleteQuery(ctx context.Context, id string) (*sql.Row, error) {
	query := `DELETE FROM business_types WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
