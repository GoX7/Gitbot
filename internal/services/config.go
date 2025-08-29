package services

import (
	"os"

	"github.com/gox7/gitbot/models"
)

func NewConfig(config *models.Config) {
	*config = models.Config{
		TELEGRAM_TOKEN:    get("TELEGRAM_TOKEN", "none"),
		GITHUB_TOKEN:      get("GITHUB_TOKEN", "none"),
		POSTGRES_HOST:     get("POSTGRES_HOST", "127.0.0.1"),
		POSTGRES_USER:     get("POSTGRES_USER", "postgres"),
		POSTGRES_PASSWORD: get("POSTGRES_PASSWORD", "postgres"),
		POSTGRES_NAME:     get("POSTGRES_NAME", "postgres"),
		REDIS_HOST:        get("REDIS_HOST", "127.0.0.1"),
		REDIS_PASSWORD:    get("REDIS_PASSWORD", "redis"),
		REDIS_DB:          get("REDIS_DB", "0"),
	}
}

func get(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}
