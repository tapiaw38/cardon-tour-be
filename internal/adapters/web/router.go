package web

import (
	"github.com/gin-gonic/gin"
	businesstype "github.com/tapiaw38/cardon-tour-be/internal/adapters/web/handlers/business/business_type"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/web/handlers/profile"
	profiletype "github.com/tapiaw38/cardon-tour-be/internal/adapters/web/handlers/profile/profile_type"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/config"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases"
)

func RegisterApplicationRoutes(app *gin.Engine, usecases *usecases.UseCases, cfg *config.Config) {
	routeGroup := app.Group("/api")

	routeGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//routeGroup.Use(middlewares.CheckAuthMiddleware(*cfg))

	routeGroup.POST("/profiles", profile.NewCreateHandler(usecases.Profile.CreateUseCase))
	routeGroup.GET("/profiles/:id", profile.NewGetHandler(usecases.Profile.GetUseCase))
	routeGroup.PATCH("/profiles/:id", profile.NewUpdateHandler(usecases.Profile.UpdateUseCase))

	routeGroup.POST("/profiles/types", profiletype.NewCreateHandler(usecases.Profile.Types.CreateUseCase))
	routeGroup.DELETE("/profiles/types/:id", profiletype.NewDeleteHandler(usecases.Profile.Types.DeleteUseCase))
	routeGroup.GET("/profiles/types", profiletype.NewListHandler(usecases.Profile.Types.ListUseCase))

	routeGroup.GET("/businesses/types", businesstype.NewListHandler(usecases.Business.Types.ListUseCase))
}
