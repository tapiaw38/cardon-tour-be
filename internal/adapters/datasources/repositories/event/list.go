package event

import (
	"context"
	"database/sql"
	"fmt"
	event_domain "github.com/tapiaw38/cardon-tour-be/internal/domain/event"
	site_domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
	"time"
)

type (
	ListFilterOptions struct {
		Name        string
		SiteID      string
		EventTypeID string
		Active      *bool
		StartAt     time.Time
		EndAt       time.Time
	}
)

func (r *repository) List(ctx context.Context, filter ListFilterOptions) ([]*event_domain.Event, error) {
	rows, err := r.executeListQuery(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make(map[string]*event_domain.Event)

	for rows.Next() {
		var id, name, createdBy string
		var description sql.NullString
		var createdAt time.Time
		var eventTypeID string
		var eventTypeName sql.NullString
		var siteID, siteSlug, siteName sql.NullString
		var scheduleID string
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
			&scheduleID,
			&scheduleStartAt,
			&scheduleEndAt,
			&scheduleActive,
		)
		if err != nil {
			return nil, err
		}

		if _, exists := events[id]; !exists {
			events[id] = &event_domain.Event{
				ID:          id,
				Name:        name,
				Description: description.String,
				CreatedAt:   createdAt,
				CreatedBy:   createdBy,
				EventType: &event_domain.EventType{
					ID:   eventTypeID,
					Name: eventTypeName.String,
				},
				SiteID: siteID.String,
				Site: &site_domain.Site{
					ID:   siteID.String,
					Slug: siteSlug.String,
					Name: siteName.String,
				},
				Schedule: []event_domain.EventSchedule{},
			}
		}

		schedule := event_domain.EventSchedule{
			ID:      scheduleID,
			StartAt: scheduleStartAt,
			EndAt:   scheduleEndAt,
			Active:  scheduleActive,
		}
		events[id].Schedule = append(events[id].Schedule, schedule)
	}

	eventList := make([]*event_domain.Event, 0, len(events))
	for _, event := range events {
		eventList = append(eventList, event)
	}

	return eventList, nil
}

func (r *repository) executeListQuery(ctx context.Context, filter ListFilterOptions) (*sql.Rows, error) {
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
			sev.id AS schedule_id,
			sev.start_at AS schedule_start_at,
			sev.end_at AS schedule_end_at,
			sev.active AS schedule_active
		FROM events ev
		INNER JOIN event_types evt ON evt.id = ev.event_type_id
		INNER JOIN sites st ON st.id = ev.site_id
		LEFT JOIN event_schedules sev ON sev.event_id = ev.id
		WHERE 1=1
	`

	args := []any{}
	argIndex := 1

	if filter.Name != "" {
		query += " AND ev.name = $" + fmt.Sprintf("%d", argIndex)
		args = append(args, filter.Name)
		argIndex++
	}

	if filter.SiteID != "" {
		query += " AND ev.site_id = $" + fmt.Sprintf("%d", argIndex)
		args = append(args, filter.SiteID)
		argIndex++
	}

	if filter.EventTypeID != "" {
		query += " AND ev.event_type_id = $" + fmt.Sprintf("%d", argIndex)
		args = append(args, filter.EventTypeID)
		argIndex++
	}

	if filter.Active != nil {
		query += " AND sev.active = $" + fmt.Sprintf("%d", argIndex)
		args = append(args, *filter.Active)
		argIndex++
	}

	if !filter.StartAt.IsZero() {
		query += " AND sev.start_at >= $" + fmt.Sprintf("%d", argIndex)
		args = append(args, filter.StartAt)
		argIndex++

		if !filter.EndAt.IsZero() {
			query += " AND sev.start_at <= $" + fmt.Sprintf("%d", argIndex)
			args = append(args, filter.EndAt)
			argIndex++
		}
	}

	query += " ORDER BY ev.created_at DESC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
