package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases/profile"
)

func NewCreateHandler(usecase profile.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var profileInput inputProfile
		if err := c.ShouldBindJSON(&profileInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		profile := toUserInput(profileInput)

		profileCreated, err := usecase.Execute(c.Request.Context(), profile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, profileCreated)
	}
}
