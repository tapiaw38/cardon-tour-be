package event_schedule

import (
	"github.com/gin-gonic/gin"
	eventschedule "github.com/tapiaw38/cardon-tour-be/internal/usecases/event/event_schedule"
	"net/http"
)

func NewGetHandler(usecase eventschedule.GetUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		eventSchedule, err := usecase.Execute(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, eventSchedule)
	}
}
