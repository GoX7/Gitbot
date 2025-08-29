package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func NewConnectRedis(addr string, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	response, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("[-] redis.ping: " + err.Error())
		return nil, err
	}

	fmt.Println("[+] redis.ping: " + response)
	return client, nil
}

func NewConnectPostgres(host, port, user, password, dbname string) (*sqlx.DB, error) {
	psc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port,
		user, password,
		dbname,
	)

	var connect *sqlx.DB
	var err error
	for range 10 {
		connect, err = sqlx.Connect("postgres", psc)
		if err != nil {
			time.Sleep(250 * time.Millisecond)
			continue
		}

		fmt.Println("[+] postgres.connect: ok")
		break
	}
	if err != nil {
		fmt.Println("[-] postgres.connect: " + err.Error())
		return nil, err
	}

	for range 10 {
		err = connect.Ping()
		if err != nil {
			fmt.Println("[-] postgres.ping: " + err.Error())
			time.Sleep(250 * time.Millisecond)
			continue
		}

		fmt.Println("[+] postgres.ping: ok")
		break
	}
	if err != nil {
		return nil, err
	}

	return connect, nil
}
