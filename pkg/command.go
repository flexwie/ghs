package pkg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/flexwie/ghs/pkg/github"
)

type GistRequest struct {
	ApiEndpoint string
	User        string
	GistName    string

	Gist          *github.Gist
	ExecutionFile *github.File
}

func ExecGist(ctx context.Context, name string) error {
	var splitPath = strings.Split(name, "/")
	if len(splitPath) != 2 {
		return errors.New("malformed gist name")
	}

	apiUrl := fmt.Sprintf("https://api.github.com/users/%s/gists", splitPath[0])

	file, gist, err := SearchForGistFile(apiUrl, splitPath[1], ctx)
	if err != nil {
		return err
	}

	exec, err := GetExecutor(file, ctx)
	if err != nil {
		return err
	}

	return exec.Execute(file, gist, ctx)
}
