package location

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/location"
)

func NewGetProvinceBySlugHandler(usecase location.GetProvinceBySlugUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")
		province, err := usecase.Execute(c.Request.Context(), slug)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, province)
	}
}
