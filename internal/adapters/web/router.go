package web

import (
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/web/handlers/business"
	businessimages "github.com/tapiaw38/cardon-tour-be/internal/adapters/web/handlers/business/business_image"
	businesstype "github.com/tapiaw38/cardon-tour-be/internal/adapters/web/handlers/business/business_type"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/web/handlers/event"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/web/handlers/location"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/web/handlers/profile"
	profilesite "github.com/tapiaw38/cardon-tour-be/internal/adapters/web/handlers/profile/profile_site"
	profiletype "github.com/tapiaw38/cardon-tour-be/internal/adapters/web/handlers/profile/profile_type"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/web/handlers/site"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/web/middlewares"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases"
)

func RegisterApplicationRoutes(app *gin.Engine, usecases *usecases.UseCases) {
	routeGroup := app.Group("/api")

	routeGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routeGroup.GET("/locations/provinces/:slug", location.NewGetProvinceBySlugHandler(usecases.Location.GetProvinceBySlugUseCase))
	routeGroup.GET("/locations/provinces", location.NewListProvinceHandler(usecases.Location.ListProvinceUseCase))

	routeGroup.GET("/sites", site.NewListHandler(usecases.Site.ListUseCase))
	routeGroup.GET("/sites/sites-by-id/:id", site.NewGetHandler(usecases.Site.GetUseCase))
	routeGroup.GET("/sites/sites-by-slug/:slug", site.NewGetBySlugHandler(usecases.Site.GetBySlugUseCase))

	routeGroup.GET("/businesses/:id", business.NewGetHandler(usecases.Business.GetUseCase))
	routeGroup.GET("/businesses", business.NewListHandler(usecases.Business.ListUseCase))

	routeGroup.GET("/businesses/types", businesstype.NewListHandler(usecases.Business.Types.ListUseCase))
	routeGroup.GET("/businesses/types-by-id/:id", businesstype.NewGetHandler(usecases.Business.Types.GetUseCase))
	routeGroup.GET("/businesses/types-by-slug/:slug", businesstype.NewGetBySlugHandler(usecases.Business.Types.GetBySlugUseCase))

	routeGroup.GET("/events/:id", event.NewGetHandler(usecases.Event.GetUseCase))
	routeGroup.GET("/events", event.NewListHandler(usecases.Event.ListUseCase))

	routeGroup.Use(middlewares.CheckAuthMiddleware())
	routeGroup.POST("/profiles", profile.NewCreateHandler(usecases.Profile.CreateUseCase))
	routeGroup.GET("/profiles/me", profile.NewGetByUserIDHandler(usecases.Profile.GetByUserIDUseCase))
	routeGroup.PATCH("/profiles/:id", profile.NewUpdateHandler(usecases.Profile.UpdateUseCase))

	routeGroup.POST("/profiles/types", profiletype.NewCreateHandler(usecases.Profile.Types.CreateUseCase))
	routeGroup.DELETE("/profiles/types/:id", profiletype.NewDeleteHandler(usecases.Profile.Types.DeleteUseCase))
	routeGroup.GET("/profiles/types", profiletype.NewListHandler(usecases.Profile.Types.ListUseCase))

	routeGroup.POST("/profiles/sites", profilesite.NewCreateHandler(usecases.Profile.Sites.CreateUseCase))
	routeGroup.DELETE("/profiles/:profile_id/sites/:site_id", profilesite.NewDeleteHandler(usecases.Profile.Sites.DeleteUseCase))

	routeGroup.POST("/sites", site.NewCreateHandler(usecases.Site.CreateUseCase))

	routeGroup.POST("/businesses", business.NewCreateHandler(usecases.Business.CreateUseCase))

	routeGroup.POST("/businesses/types", businesstype.NewCreateHandler(usecases.Business.Types.CreateUseCase))

	routeGroup.POST("/businesses/images/:business_id", businessimages.NewCreateHandler(usecases.Business.Images.CreateUseCase))
}
