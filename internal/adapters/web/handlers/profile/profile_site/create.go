package profilesite

import (
	"net/http"

	"github.com/gin-gonic/gin"
	profilesite "github.com/tapiaw38/cardon-tour-be/internal/usecases/profile/profile_site"
)

func NewCreateHandler(usecase profilesite.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var profileSiteInput inputProfileSite
		if err := c.ShouldBindJSON(&profileSiteInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := usecase.Execute(c.Request.Context(), profileSiteInput.ProfileID, profileSiteInput.SiteID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusCreated)
	}
}
