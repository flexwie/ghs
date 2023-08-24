package github

import (
	"io"
	"net/http"
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

func (f *File) Content() (string, error) {
	resp, err := http.Get(f.RawUrl)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
