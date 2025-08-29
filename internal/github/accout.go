package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gox7/gitbot/models"
)

func CheckAccout(name string) (bool, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", name)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if token != "none" {
		request.Header.Set("Authorization", "Bearer "+token)
	}

	client := http.Client{}
	client.Timeout = time.Second * 10

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if response.StatusCode == 404 {
		return false, nil
	}

	return true, nil
}

func GetAccout(name string) (*models.GitHubUser, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", name)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if token != "none" {
		request.Header.Set("Authorization", "Bearer "+token)
	}

	client := http.Client{}
	client.Timeout = time.Second * 10

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	var model models.GitHubUser
	if err := json.NewDecoder(response.Body).Decode(&model); err != nil {
		return nil, fmt.Errorf("%s", "ANF")
	}

	return &model, nil
}
