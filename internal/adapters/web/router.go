package web

import (
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/web/handlers/profile"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases"
)

func RegisterApplicationRoutes(app *gin.Engine, usecases *usecases.UseCases) {
	routeGroup := app.Group("/api")

	routeGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routeGroup.POST("/profiles", profile.NewCreateHandler(usecases.Profile.CreateUseCase))
	routeGroup.GET("/profiles/:id", profile.NewGetHandler(usecases.Profile.GetUseCase))
}
