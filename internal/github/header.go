package github

import "github.com/gox7/gitbot/models"

var (
	token string
)

func SetToken(config *models.Config) {
	token = config.GITHUB_TOKEN
}
