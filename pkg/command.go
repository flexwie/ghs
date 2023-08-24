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

	ctx = context.WithValue(ctx, "apiUrl", apiUrl)
	ctx = context.WithValue(ctx, "username", splitPath[0])
	ctx = context.WithValue(ctx, "gist", splitPath[1])

	gist, file, err := SearchForGist(ctx)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, "gist", gist)
	ctx = context.WithValue(ctx, "file", file)

	exec, err := GetExecutor(ctx)

	return exec.Execute(ctx)
}
