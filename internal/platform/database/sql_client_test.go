package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/config"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/database"
)

func TestNewSQLConfig(t *testing.T) {
	cfg := config.Config{DatabaseURL: "postgres://localhost/testdb"}
	sqlConfig := database.NewSQLConfig(cfg)

	assert.Equal(t, cfg.DatabaseURL, sqlConfig.DatabaseURL)
}
