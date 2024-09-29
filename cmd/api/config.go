package main

import (
	"github.com/joho/godotenv"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/config"
	"log"
	"os"
)

func initConfig() error {
	configService, err := readConfig()
	if err != nil {
		return err
	}
	config.InitConfigService(configService)
	return nil
}

func readConfig() (*config.ConfigurationService, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error load env file")
	}

	configService := &config.ConfigurationService{
		ServerConfig: config.ServerConfig{
			GinMode:   config.GinModeServer(getEnv("GIN_MODE", "release")),
			Port:      getEnv("PORT", "8080"),
			Host:      getEnv("HOST", "localhost"),
			JWTSecret: getEnv("JWT_SECRET", "secret"),
		},
		DBConfig: config.DBConfig{
			DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/cardontour?sslmode=disable"),
		},
		S3Config: config.S3Config{
			AWSRegion:          getEnv("AWS_REGION", ""),
			AWSBucket:          getEnv("AWS_BUCKET", ""),
			AWSAccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", ""),
			AWSSecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", ""),
		},
	}

	return configService, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
