package event_type

import (
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/event/event_type"
)

func NewCreateHandler(usecase event_type.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventTypeInput := EventTypeInput{}
		if err := c.ShouldBindJSON(&eventTypeInput); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		eventType, err := usecase.Execute(c.Request.Context(), toEventTypeInput(eventTypeInput))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, eventType)
	}
}
