package eventtype

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
)

func (r *repository) Create(ctx context.Context, eventType domain.EventType) (string, error) {
	row, err := r.executeCreateQuery(ctx, eventType)
	if err != nil {
		return "", err
	}

	var eventTypeID string
	err = row.Scan(&eventTypeID)

	return eventTypeID, err
}

func (r *repository) executeCreateQuery(ctx context.Context, eventType domain.EventType) (*sql.Row, error) {
	query := `INSERT INTO event_types(
                    id,
					name
				) VALUES ($1, $2)
			RETURNING id`

	args := []any{
		eventType.ID,
		eventType.Name,
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
