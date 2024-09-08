package profilesite

import (
	"net/http"

	"github.com/gin-gonic/gin"
	profilesite "github.com/tapiaw38/cardon-tour-be/internal/usecases/profile/profile_site"
)

func NewDeleteHandler(usecase profilesite.DeleteUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		profileID := c.Param("profile_id")
		siteID := c.Param("site_id")

		if profileID == "" || siteID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "profile_id and site_id are required"})
			return
		}

		profileSite, err := usecase.Execute(c.Request.Context(), profileID, siteID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, profileSite)
	}
}
