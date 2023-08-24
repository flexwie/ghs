package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/flexwie/ghs/pkg/executors"
	"github.com/flexwie/ghs/pkg/github"
)

func SearchForGist(ctx context.Context) (*github.Gist, *github.File, error) {
	url := ctx.Value("apiUrl").(string)
	gistName := ctx.Value("gist").(string)

	resp, err := http.Get(url)
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

func GetExecutor(ctx context.Context) (executors.Executor, error) {
	for _, e := range executors.ExecutorPipeline {
		file := ctx.Value("file").(*github.File)

		if e.Match(file) {
			return e, nil
		}
	}

	return nil, errors.New("no executor found")
}
