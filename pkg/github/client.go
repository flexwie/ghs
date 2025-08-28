package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/cli/go-gh"
)

type GithubClient struct {
	httpClient *http.Client
	username   string
}

func NewGithubClient(ops ...GithubClientOption) *GithubClient {
	client := &GithubClient{
		httpClient: http.DefaultClient,
		username:   username(),
	}

	for _, o := range ops {
		o(client)
	}

	return client
}

func (g *GithubClient) GetGistWithFile(username string, gistName string) (*Gist, error) {
	token := getGithubToken()

	url := fmt.Sprintf("https://api.github.com/users/%s/gists", username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []Gist
	json.Unmarshal(body, &data)

	for _, gist := range data {
		for _, file := range gist.Files {
			if file.Filename == gistName {
				return &gist, nil
			}
		}
	}

	return nil, errors.New(fmt.Sprintf("unable to find requested gist: %s", url))
}

func getGithubToken() string {
	x, _, err := gh.Exec("auth", "token")
	if err != nil {
		log.Warn("unauthorized")
	}

	return strings.TrimSuffix(x.String(), "\n")
}

func username() string {
	x, _, err := gh.Exec("api", "user", "-q", ".login")
	if err != nil {
		// do smth
	}

	return strings.TrimSpace(x.String())
}
