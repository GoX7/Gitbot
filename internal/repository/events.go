package repository

import (
	"database/sql"

	"github.com/gox7/gitbot/models"
	"github.com/jmoiron/sqlx"
)

func CreateUserTable(connect *sqlx.DB) (sql.Result, error) {
	return connect.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		user_id BIGINT UNIQUE NOT NULL,
		username VARCHAR(255) NOT NULL,
		github VARCHAR(255) NOT NULL
	)`)
}

func CreateUser(connect *sqlx.DB, id int64, username string, github string) error {
	_, err := connect.Exec("INSERT INTO users (user_id, username, github) VALUES ($1, $2, $3)",
		id, username,
		github,
	)
	return err
}

func SearchUser(connect *sqlx.DB, id int64) (*models.UserOrm, error) {
	var user models.UserOrm
	if err := connect.Get(&user, "SELECT * FROM users WHERE user_id=$1", id); err != nil {
		return nil, err
	}
	return &user, nil
}
