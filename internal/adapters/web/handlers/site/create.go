package site

import (
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/site"
	"net/http"
)

func NewCreateHandler(usecase site.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var siteInput SiteInput
		if err := c.ShouldBindJSON(&siteInput); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		siteCreated, err := usecase.Execute(c.Request.Context(), toSiteInput(siteInput))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, siteCreated)
	}
}
