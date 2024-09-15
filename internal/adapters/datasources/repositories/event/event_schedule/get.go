package event_schedule

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
)

func (r *repository) Get(ctx context.Context, id string) (*domain.EventSchedule, error) {
	row, err := r.executeGetQuery(ctx, id)
	if err != nil {
		return nil, err
	}

	var schedule domain.EventSchedule
	err = row.Scan(
		&schedule.ID,
		&schedule.EventID,
		&schedule.Active,
		&schedule.StartAt,
		&schedule.EndAt,
		&schedule.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (r *repository) executeGetQuery(ctx context.Context, id string) (*sql.Row, error) {
	query := `
		SELECT
			id,
			event_id, 
			active, 
		    start_at, 
		    end_at, 
			created_at
		FROM event_schedules
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}
