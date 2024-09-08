package location

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tapiaw38/cardon-tour-be/internal/usecases/location"
)

func NewListProvinceHandler(usecase location.ListProvinceUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		provinces, err := usecase.Execute(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, provinces)
	}
}
