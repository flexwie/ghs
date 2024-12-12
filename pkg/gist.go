package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/cli/go-gh/v2"
	"github.com/flexwie/ghs/pkg/executors"
	"github.com/flexwie/ghs/pkg/github"
)

func SearchForGistFile(username string, gistName string, ctx context.Context) (*github.File, *github.Gist, error) {
	token := getGithubToken()

	url := fmt.Sprintf("https://api.github.com/users/%s/gists", username)

	req, err := http.NewRequest("GET", url, nil)
	if token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var data []github.Gist
	json.Unmarshal(body, &data)

	for _, gist := range data {
		for _, file := range gist.Files {
			if file.Filename == gistName {
				return &file, &gist, nil
			}
		}
	}

	return nil, nil, errors.New(fmt.Sprintf("unable to find requested gist: %s", url))

}

func getGithubToken() string {
	x, _, err := gh.Exec("auth", "token")
	if err != nil {
		log.Warn("unauthorized")
	}

	return strings.TrimSuffix(x.String(), "\n")
}

func getGithubUsername() string {
	x, _, err := gh.Exec("api", "user", "-q", ".login")
	if err != nil {
		// do smth
	}

	return strings.TrimSpace(x.String())
}

func GetExecutor(file *github.File, ctx context.Context) (executors.Executor, error) {
	for _, e := range executors.ExecutorPipeline {
		if e.Match(file) {
			return e, nil
		}
	}

	return nil, errors.New("no executor found")
}
