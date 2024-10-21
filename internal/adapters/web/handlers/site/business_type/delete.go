package site_business_type

import (
	"github.com/gin-gonic/gin"
	sitebusinesstype "github.com/tapiaw38/cardon-tour-be/internal/usecases/site/business_type"
	"net/http"
)

func NewDeleteHandler(usecase sitebusinesstype.DeleteUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		siteID := c.Param("site_id")
		businessTypeID := c.Param("business_type_id")

		if siteID == "" || businessTypeID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "site_id and business_type_id are required"})
			return
		}

		siteBusinessTypeDeleted, err := usecase.Execute(
			c.Request.Context(),
			siteID,
			businessTypeID,
		)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, siteBusinessTypeDeleted)
	}
}
