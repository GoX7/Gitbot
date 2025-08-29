package main

import (
	"log/slog"

	"github.com/gox7/gitbot/internal/github"
	"github.com/gox7/gitbot/internal/services"
	"github.com/gox7/gitbot/internal/transport"
	"github.com/gox7/gitbot/models"
)

func main() {
	config := new(models.Config)
	logger := new(slog.Logger)

	services.NewConfig(config)
	services.NewLogger("server", &logger)

	redis := new(services.Redis)
	postgres := new(services.Postgres)
	pqusers := new(services.PostgresUsers)
	services.NewRedis(config, redis)
	services.NewPostgres(config, postgres)
	services.NewPostgresUsers(postgres, pqusers)

	github.SetToken(config)

	transport.Resource(logger, pqusers, redis)
	transport.Register(config)
}
