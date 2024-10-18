package businesstype

import (
	"github.com/gin-gonic/gin"
	businesstype "github.com/tapiaw38/cardon-tour-be/internal/usecases/business/business_type"
)

func NewCreateHandler(usecase businesstype.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var businessTypeInput BusinessTypeInput
		if err := c.ShouldBindJSON(&businessTypeInput); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		businessTypeCreated, err := usecase.Execute(c.Request.Context(), toBusinessTypeInput(businessTypeInput))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, businessTypeCreated)
	}
}
