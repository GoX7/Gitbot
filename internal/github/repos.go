package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Repo struct {
	Name            string `json:"name"`
	StargazersCount int    `json:"stargazers_count"`
	ForksCount      int    `json:"forks_count"`
}

func GetStarsForks(name string) (stars int, forks int) {
	url := fmt.Sprintf("https://api.github.com/users/%s", name)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, 0
	}

	if token != "none" {
		request.Header.Set("Authorization", "Bearer "+token)
	}

	client := http.Client{}
	client.Timeout = time.Second * 10

	response, err := client.Do(request)
	if err != nil {
		return 0, 0
	}

	var repos []Repo
	if err := json.NewDecoder(response.Body).Decode(&repos); err != nil {
		return 0, 0
	}

	for _, r := range repos {
		stars += r.StargazersCount
		forks += r.ForksCount
	}

	return stars, forks
}
