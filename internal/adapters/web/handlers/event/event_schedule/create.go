package event_schedule

import (
	"github.com/gin-gonic/gin"
	eventschedule "github.com/tapiaw38/cardon-tour-be/internal/usecases/event/event_schedule"
	"net/http"
)

func NewCreateHandler(usecase eventschedule.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var eventScheduleInput EventScheduleInput
		if err := c.ShouldBindJSON(&eventScheduleInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		eventSchedule, err := usecase.Execute(
			c.Request.Context(),
			toEventScheduleInput(eventScheduleInput),
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, eventSchedule)
	}
}
