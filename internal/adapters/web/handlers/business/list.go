package business

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/business"
)

func NewListHandler(usecase business.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		filters := parseListFilterOptions(c.Request.URL.Query())

		businessList, err := usecase.Execute(c.Request.Context(), filters)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, businessList)
	}
}

func parseListFilterOptions(query url.Values) business.ListFilterOptions {
	return business.ListFilterOptions{
		SiteSlug:         query.Get("site_slug"),
		BusinessTypeSlug: query.Get("business_type_slug"),
		Search:           query.Get("search"),
	}
}
