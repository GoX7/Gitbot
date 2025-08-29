package models

type (
	Config struct {
		TELEGRAM_TOKEN    string
		GITHUB_TOKEN      string
		POSTGRES_HOST     string
		POSTGRES_USER     string
		POSTGRES_PASSWORD string
		POSTGRES_NAME     string
		REDIS_HOST        string
		REDIS_PASSWORD    string
		REDIS_DB          string
	}
)
