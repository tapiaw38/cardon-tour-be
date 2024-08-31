package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/claim"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/profile"
)

func NewCreateHandler(usecase profile.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var profileInput inputProfile
		if err := c.ShouldBindJSON(&profileInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		claimUserID := c.MustGet("claims").(domain.Claims).UserId
		if claimUserID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		profileInput.UserID = claimUserID

		profile := toProfileInput(profileInput)

		profileCreated, err := usecase.Execute(c.Request.Context(), profile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, profileCreated)
	}
}
