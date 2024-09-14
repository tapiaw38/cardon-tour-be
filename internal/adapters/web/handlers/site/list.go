package site

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/site"
)

func NewListHandler(usecase site.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		filters := parseListFilterOptions(c.Request.URL.Query())
		sites, err := usecase.Execute(c.Request.Context(), filters)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, sites)
	}
}

func parseListFilterOptions(queries url.Values) site.ListFilterOptions {
	return site.ListFilterOptions{
		ProvinceID: queries.Get("province_id"),
	}
}
