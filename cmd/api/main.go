package main

import (
	"database/sql"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/storeservice"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/web"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/config"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/database"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/objectstore"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases"
	"log"
)

func main() {
	scope := config.GetScope()

	log.Printf("scope identifier: %s", scope)

	if err := initConfig(); err != nil {
		panic(err)
	}

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	configService := config.GetConfigService()

	db, err := database.GetSQLClientInstance()
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	err = database.Makemigration()
	if err != nil {
		return err
	}

	s3, err := objectstore.GetS3SessionInstance()
	if err != nil {
		return err
	}

	if configService.ServerConfig.GinMode == config.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowOrigins = []string{"*"}
	ginConfig.AllowCredentials = true
	ginConfig.AllowMethods = []string{"*"}
	ginConfig.AllowHeaders = []string{"*"}
	ginConfig.ExposeHeaders = []string{"*"}
	app.Use(cors.New(ginConfig))

	bootstrap(app, db, s3, &configService)

	return app.Run(":" + configService.ServerConfig.Port)
}

func bootstrap(app *gin.Engine, db *sql.DB, s3 *session.Session, configService *config.ConfigurationService) {
	storeService := storeservice.NewStoreService(s3, configService)
	contextFactory := appcontext.NewFactory(db, storeService)
	uc := usecases.CreateUsecases(contextFactory)
	web.RegisterApplicationRoutes(app, uc)
}
