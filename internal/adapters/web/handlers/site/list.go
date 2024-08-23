package site

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/site"
)

func NewListHandler(usecase site.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		sites, err := usecase.Execute(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, sites)
	}
}
