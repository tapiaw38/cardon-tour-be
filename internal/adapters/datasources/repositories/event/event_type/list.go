package eventtype

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
)

func (r *repository) List(ctx context.Context) ([]domain.EventType, error) {
	rows, err := r.executeListQuery(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventTypes []domain.EventType
	for rows.Next() {
		var eventType domain.EventType
		err := rows.Scan(&eventType.ID, &eventType.Name)
		if err != nil {
			return nil, err
		}
		eventTypes = append(eventTypes, eventType)
	}

	return eventTypes, nil
}

func (r *repository) executeListQuery(ctx context.Context) (*sql.Rows, error) {
	query := `SELECT 
		id, 
		name
		FROM event_types`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
