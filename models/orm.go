package models

type (
	UserOrm struct {
		Id       int64  `db:"id"`
		UserId   int64  `db:"user_id"`
		Username string `db:"username"`
		Github   string `db:"github"`
	}
)
