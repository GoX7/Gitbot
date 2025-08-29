package services

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/gox7/gitbot/internal/repository"
	"github.com/gox7/gitbot/models"
	"github.com/redis/go-redis/v9"
)

type (
	Redis struct {
		client *redis.Client
	}
)

func NewRedis(config *models.Config, redis *Redis) {
	db, err := strconv.Atoi(config.REDIS_DB)
	if err != nil {
		fmt.Println("[-] config.redis: atoi error - " + err.Error())
		os.Exit(1)
	}

	client, err := repository.NewConnectRedis(
		config.REDIS_HOST+":"+"6379",
		config.REDIS_PASSWORD,
		db,
	)
	if err != nil {
		os.Exit(1)
	}

	*redis = Redis{
		client: client,
	}
}

func (redis *Redis) Get(key string) (string, error) {
	return redis.client.Get(context.Background(), key).Result()
}

func (redis *Redis) Set(key string, value string) (string, error) {
	return redis.client.Set(context.Background(), key, value, 0).Result()
}
