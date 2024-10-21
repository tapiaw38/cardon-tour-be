package event_type

import (
	"github.com/gin-gonic/gin"
	eventtype "github.com/tapiaw38/cardon-tour-be/internal/usecases/event/event_type"
)

func NewGetHandler(usecase eventtype.GetUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		eventType, err := usecase.Execute(c.Request.Context(), id)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, eventType)
	}
}
