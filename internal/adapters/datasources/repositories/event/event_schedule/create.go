package event_schedule

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
)

func (r *repository) Create(ctx context.Context, schedule domain.EventSchedule) (string, error) {
	row, err := r.executeCreateQuery(ctx, schedule)
	if err != nil {
		return "", err
	}

	var id string
	err = row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *repository) executeCreateQuery(ctx context.Context, schedule domain.EventSchedule) (*sql.Row, error) {
	query := `
		INSERT INTO event_schedules (
			id,
			event_id, 
			active, 
		    start_at, 
		    end_at, 
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	args := []any{
		schedule.EventID,
		schedule.Active,
		schedule.StartAt,
		schedule.EndAt,
		schedule.CreatedAt,
	}

	rows := r.db.QueryRowContext(ctx, query, args...)

	return rows, nil
}
