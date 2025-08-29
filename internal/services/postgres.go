package services

import (
	"os"

	"github.com/gox7/gitbot/internal/repository"
	"github.com/gox7/gitbot/models"
	"github.com/jmoiron/sqlx"
)

type (
	Postgres struct {
		connect *sqlx.DB
	}
)

func NewPostgres(config *models.Config, postgres *Postgres) {
	connect, err := repository.NewConnectPostgres(
		config.POSTGRES_HOST,
		config.POSTGRES_HOST,
		config.POSTGRES_USER,
		config.POSTGRES_PASSWORD,
		config.POSTGRES_NAME,
	)
	if err != nil {
		os.Exit(1)
	}

	*postgres = Postgres{
		connect: connect,
	}
}
