package site

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tapiaw38/cardon-tour-be/internal/usecases/site"
)

func NewGetHandler(usecase site.GetUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		site, err := usecase.Execute(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, site)
	}
}
