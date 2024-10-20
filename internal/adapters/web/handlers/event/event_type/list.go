package event_type

import (
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/event/event_type"
)

func NewListHandler(usecase event_type.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventTypes, err := usecase.Execute(c.Request.Context())
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"data": eventTypes,
		})
	}
}
