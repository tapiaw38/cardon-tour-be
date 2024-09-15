package businesstype

import (
	"net/http"

	"github.com/gin-gonic/gin"

	businesstype "github.com/tapiaw38/cardon-tour-be/internal/usecases/business/business_type"
)

func NewGetHandler(usecase businesstype.GetUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		businessType, err := usecase.Execute(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, businessType)
	}
}
