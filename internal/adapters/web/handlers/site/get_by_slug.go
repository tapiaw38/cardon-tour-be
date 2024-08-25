package site

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tapiaw38/cardon-tour-be/internal/usecases/site"
)

func NewGetBySlugHandler(usecase site.GetBySlugUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")
		site, err := usecase.Execute(c.Request.Context(), slug)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, site)
	}
}
