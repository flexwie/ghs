package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/cli/go-gh/v2"
	"github.com/flexwie/ghs/pkg/executors"
	"github.com/flexwie/ghs/pkg/github"
)

func SearchForGistFile(url string, gistName string, ctx context.Context) (*github.File, *github.Gist, error) {
	token := getGithubToken()

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

	return nil, nil, errors.New("unable to find requested gist")
}

func getGithubToken() string {
	x, _, err := gh.Exec("auth", "token")
	if err != nil {
		log.Warn("unauthorized")
	}

	return strings.TrimSuffix(x.String(), "\n")

	cmd := exec.Command("which", "gh")
	if err := cmd.Run(); err != nil {
		log.Warn("GitHub cli is not installed. Your private gists will not be found.")
		return ""
	}

	cmd = exec.Command("gh", "auth", "token")
	writer := new(bytes.Buffer)
	cmd.Stdout = writer

	errOut := new(bytes.Buffer)
	cmd.Stderr = errOut

	if err := cmd.Run(); err != nil {
		log.Warn("GitHub cli authentication failed. Your private gists will not be found.", "err", errOut.String())
		return ""
	}

	return writer.String()
}

func GetExecutor(file *github.File, ctx context.Context) (executors.Executor, error) {
	for _, e := range executors.ExecutorPipeline {
		if e.Match(file) {
			return e, nil
		}
	}

	return nil, errors.New("no executor found")
}
