package businesstype

import (
	"net/http"

	"github.com/gin-gonic/gin"
	businesstype "github.com/tapiaw38/cardon-tour-be/internal/usecases/business/business_type"
)

func NewListHandler(usecase businesstype.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		businessTypes, err := usecase.Execute(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, businessTypes)
	}
}
