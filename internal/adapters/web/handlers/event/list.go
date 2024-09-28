package event

import (
	"github.com/gin-gonic/gin"
	web_utils "github.com/tapiaw38/cardon-tour-be/internal/platform/web"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/event"
	"net/http"
	"net/url"
	"time"
)

func NewListHandler(usecase event.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		filters := parseListFilterOptions(c.Request.URL.Query())
		events, err := usecase.Execute(c.Request.Context(), filters)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, events)
	}
}

func parseListFilterOptions(queries url.Values) event.ListFilterOptions {
	return event.ListFilterOptions{
		Name:        queries.Get("name"),
		SiteID:      queries.Get("site_id"),
		EventTypeID: queries.Get("event_type_id"),
		Active:      web_utils.ParseBoolPointerQueryValue(queries.Get("active")),
		StartAt:     web_utils.ParseTimeQueryValue(queries.Get("start_at"), time.RFC3339),
		EndAt:       web_utils.ParseTimeQueryValue(queries.Get("end_at"), time.RFC3339),
	}
}
