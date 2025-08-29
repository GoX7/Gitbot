package services

import (
	"fmt"

	"github.com/gox7/gitbot/internal/repository"
	"github.com/gox7/gitbot/models"
	"github.com/jmoiron/sqlx"
)

type (
	PostgresUsers struct {
		connect *sqlx.DB
	}
)

func NewPostgresUsers(postgres *Postgres, users *PostgresUsers) {
	fmt.Println("[+] postgres.users: ok")
	repository.CreateUserTable(postgres.connect)
	*users = PostgresUsers{
		connect: postgres.connect,
	}
}

func (users *PostgresUsers) Register(user_id int64, username string, github string) error {
	err := repository.CreateUser(users.connect, user_id, username, github)
	return err
}

func (users *PostgresUsers) Search(user_id int64) (*models.UserOrm, error) {
	user, err := repository.SearchUser(users.connect, user_id)
	return user, err
}
