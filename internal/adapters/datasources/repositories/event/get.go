package event

import (
	"context"
	"database/sql"
	event_domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
	site_domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
	"time"
)

func (r *repository) Get(ctx context.Context, eventID string) (*event_domain.Event, error) {
	rows, err := r.executeGetQuery(ctx, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var event *event_domain.Event
	schedules := []event_domain.EventSchedule{}

	for rows.Next() {
		var id, name, eventTypeID, eventTypeName, siteID, siteSlug, siteName, createdBy string
		var description sql.NullString
		var createdAt time.Time
		var scheduleStartAt, scheduleEndAt time.Time
		var scheduleActive bool

		err := rows.Scan(
			&id,
			&eventTypeID,
			&eventTypeName,
			&siteID,
			&siteSlug,
			&siteName,
			&name,
			&description,
			&createdAt,
			&createdBy,
			&scheduleStartAt,
			&scheduleEndAt,
			&scheduleActive,
		)
		if err != nil {
			return nil, err
		}

		if event == nil {
			event = &event_domain.Event{
				ID:          id,
				Name:        name,
				Description: description.String,
				CreatedAt:   createdAt,
				CreatedBy:   createdBy,
				EventType: &event_domain.EventType{
					ID:   eventTypeID,
					Name: eventTypeName,
				},
				SiteID: siteID,
				Site: &site_domain.Site{
					ID:   siteID,
					Slug: siteSlug,
					Name: siteName,
				},
				Schedule: []event_domain.EventSchedule{},
			}
		}

		schedule := event_domain.EventSchedule{
			StartAt: scheduleStartAt,
			EndAt:   scheduleEndAt,
			Active:  scheduleActive,
		}
		schedules = append(schedules, schedule)
	}

	if event != nil {
		event.Schedule = schedules
	}

	return event, nil
}

func (r *repository) executeGetQuery(ctx context.Context, eventID string) (*sql.Rows, error) {
	query := `
		SELECT
			ev.id,
			ev.event_type_id,
			evt.name AS event_type_name,
			ev.site_id,
			st.slug AS site_slug,
			st.name AS site_name,
			ev.name,
			ev.description,
			ev.created_at,
			ev.created_by,
			sev.start_at AS schedule_start_at,
			sev.end_at AS schedule_end_at,
			sev.active AS schedule_active
		FROM events ev
		INNER JOIN event_types evt ON evt.id = ev.event_type_id
		INNER JOIN sites st ON st.id = ev.site_id
		LEFT JOIN event_schedules sev ON sev.event_id = ev.id
		WHERE ev.id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, eventID)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
