package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type File struct {
	Filename string
	Type     string
	Language string
	RawUrl   string `json:"raw_url"`

	httpClient *http.Client
}

func NewFile(client *http.Client) *File {
	file := &File{
		httpClient: client,
	}

	return file
}

func (f *File) Get() error {
	// TODO: fix
	url := fmt.Sprintf("https://api.github.com/users/%s/gists", "")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(body, f)

	return nil
}

func (f *File) Content() (string, error) {
	resp, err := f.httpClient.Get(f.RawUrl)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
