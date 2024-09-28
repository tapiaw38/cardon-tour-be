package event_schedule

import (
	"context"
	"database/sql"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
)

func (r *repository) List(ctx context.Context, eventID string) ([]domain.EventSchedule, error) {
	rows, err := r.executeListQuery(ctx, eventID)
	if err != nil {
		return nil, err
	}

	var schedules []domain.EventSchedule
	for rows.Next() {
		var schedule domain.EventSchedule
		err = rows.Scan(
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
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

func (r *repository) executeListQuery(ctx context.Context, eventID string) (*sql.Rows, error) {
	query := `
		SELECT
			id,
			event_id, 
			active, 
		    start_at, 
		    end_at, 
			created_at
		FROM event_schedules
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
