package eventtype

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
)

func (r *repository) Get(ctx context.Context, eventTypeID string) (domain.EventType, error) {
	row, err := r.executeGetQuery(ctx, eventTypeID)
	if err != nil {
		return domain.EventType{}, err
	}

	var eventType domain.EventType
	err = row.Scan(&eventType.ID, &eventType.Name)

	return eventType, err
}

func (r *repository) executeGetQuery(ctx context.Context, eventTypeID string) (*sql.Row, error) {
	query := `SELECT id, name FROM event_types WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, eventTypeID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
