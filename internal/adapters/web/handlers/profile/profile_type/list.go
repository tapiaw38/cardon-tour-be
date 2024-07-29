package profiletype

import (
	"net/http"

	"github.com/gin-gonic/gin"
	profiletype "github.com/tapiaw38/cardon-tour-be/internal/usecases/profile/profile_type"
)

func NewListHandler(usecase profiletype.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		profileTypes, err := usecase.Execute(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, profileTypes)
	}
}
