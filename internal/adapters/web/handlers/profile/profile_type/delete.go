package profiletype

import (
	"net/http"

	"github.com/gin-gonic/gin"
	profiletype "github.com/tapiaw38/cardon-tour-be/internal/usecases/profile/profile_type"
)

func NewDeleteHandler(usecase profiletype.DeleteUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := usecase.Execute(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusOK)
	}
}
