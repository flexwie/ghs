package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os/exec"

	"github.com/charmbracelet/log"
	"github.com/flexwie/ghs/pkg/executors"
	"github.com/flexwie/ghs/pkg/github"
)

func SearchForGist(ctx context.Context) (*github.Gist, *github.File, error) {
	url := ctx.Value("apiUrl").(string)
	gistName := ctx.Value("gist").(string)

	token := getGithubToken()

	req, err := http.NewRequest("GET", url, nil)
	if token != "" {
		req.Header.Add("Authorization", token)
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
				return &gist, &file, nil
			}
		}
	}

	return nil, nil, errors.New("unable to find requested gist")
}

func getGithubToken() string {
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

func GetExecutor(ctx context.Context) (executors.Executor, error) {
	for _, e := range executors.ExecutorPipeline {
		file := ctx.Value("file").(*github.File)

		if e.Match(file) {
			return e, nil
		}
	}

	return nil, errors.New("no executor found")
}
