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

		profile, err := usecase.Execute(c.Request.Context(), claims.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, profile)
	}
}
