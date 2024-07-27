package profiletype

import (
	"net/http"

	"github.com/gin-gonic/gin"
	profiletype "github.com/tapiaw38/cardon-tour-be/internal/usecases/profile/profile_type"
)

func NewCreateHandler(usecase profiletype.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var profileTypeInput inputProfileType
		if err := c.ShouldBindJSON(&profileTypeInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		profileType := toProfileTypeInput(profileTypeInput)

		err := usecase.Execute(c.Request.Context(), profileType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusCreated)
	}
}
