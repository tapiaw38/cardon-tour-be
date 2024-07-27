package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/profile"
)

func NewGetHandler(usecase profile.GetUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		profile, err := usecase.Execute(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, profile)
	}
}
