package businesstype

import (
	"net/http"

	"github.com/gin-gonic/gin"

	businesstype "github.com/tapiaw38/cardon-tour-be/internal/usecases/business/business_type"
)

func NewGetBySlugHandler(usecase businesstype.GetBySlugUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")
		businessType, err := usecase.Execute(c.Request.Context(), slug)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, businessType)
	}
}
