package business

import (
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/business"
	"net/http"
)

func NewCreateHandler(usecase business.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var businessInput BusinessInput
		if err := c.ShouldBindJSON(&businessInput); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		business, err := usecase.Execute(c, toBusinessInput(businessInput))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, business)
	}
}
