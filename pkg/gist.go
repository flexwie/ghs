package pkg

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/flexwie/ghs/pkg/executors"
)

type Gist struct {
	Url    string
	Id     string
	Files  map[string]File
	Public bool
}

type File struct {
	Filename string
	Type     string
	Language string
	RawUrl   string `json:"raw_url"`
}

func SearchGist(gistEndpoint, filename string) (*File, error) {
	resp, err := http.Get(gistEndpoint)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []Gist
	json.Unmarshal(body, &data)

	for _, gists := range data {
		for _, file := range gists.Files {
			if file.Filename == filename {
				return &file, nil
			}
		}
	}

	return nil, errors.New("unable to find requested gist")
}

func FetchGistContent(file *File) (string, error) {
	resp, err := http.Get(file.RawUrl)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ExecuteGist(content string) error {
	for _, e := range executors.ExecutorPipeline {
		if e.Match(content) {
			return e.Execute(content)
		}
	}

	return errors.New("no executor found")
}
