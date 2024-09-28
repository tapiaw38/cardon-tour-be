package event_schedule

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
)

func (r *repository) Update(ctx context.Context, schedule domain.EventSchedule) (string, error) {
	row, err := r.executeUpdateQuery(ctx, schedule)
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

func (r *repository) executeUpdateQuery(ctx context.Context, schedule domain.EventSchedule) (*sql.Row, error) {
	query := `
		UPDATE event_schedules
		SET
			active = COALESCE(active, $2),
			start_at = COALESCE(start_at, $3),
			end_at = COALESCE(end_at, $4),
			created_at = COALESCE(created_at, $5)
		WHERE id = $1
		RETURNING id
	`
	args := []any{
		schedule.ID,
		schedule.Active,
		schedule.StartAt,
		schedule.EndAt,
		schedule.CreatedAt,
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
