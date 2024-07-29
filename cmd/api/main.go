package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/joho/godotenv"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/web"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/config"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/database"
	"github.com/tapiaw38/cardon-tour-be/internal/usecases"
)

func main() {
	cfg, err := initConfig()
	if err != nil {
		panic(err)
	}

	if err := run(cfg); err != nil {
		panic(err)
	}
}

func initConfig() (*config.Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error load env file")
	}

	cfg := config.NewConfig()
	if cfg.DatabaseURL == "" {
		return nil, errors.New("databaseURL is required")
	}
	if cfg.Port == "" {
		return nil, errors.New("port is required")
	}

	return &cfg, nil
}

func run(config *config.Config) error {
	sqlClient := database.NewSQLConfig(*config)

	db, err := sqlClient.GetSQLClientInstance()
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	err = sqlClient.Makemigration()
	if err != nil {
		return err
	}

	if config.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	app := gin.Default()

	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{"*"}
	cfg.AllowCredentials = true
	cfg.AllowMethods = []string{"*"}
	cfg.AllowHeaders = []string{"*"}
	cfg.ExposeHeaders = []string{"*"}

	app.Use(cors.New(cfg))

	bootstrap(app, db, config)

	return app.Run(":" + config.Port)
}

func bootstrap(app *gin.Engine, db *sql.DB, cfg *config.Config) {
	contextFactory := appcontext.NewFactory(db)
	uc := usecases.CreateUsecases(contextFactory)

	web.RegisterApplicationRoutes(app, uc, cfg)
}
