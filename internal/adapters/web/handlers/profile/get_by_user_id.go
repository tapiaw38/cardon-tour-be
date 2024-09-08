package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/profile"

	web_utils "github.com/tapiaw38/cardon-tour-be/internal/platform/web"
)

func NewGetByUserIDHandler(usecase profile.GetByUserIDUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := web_utils.GetClaimsFromContext(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		userID := claims.UserId
		if userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		profile, err := usecase.Execute(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, profile)
	}
}
