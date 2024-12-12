package pkg

import (
	"context"
	"errors"
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

func ExecGist(ctx context.Context, name string, args []string) error {
	apiUrl, gistName, err := ParseArgs(name)
	if err != nil {
		return err
	}

	file, gist, err := SearchForGistFile(apiUrl, gistName, ctx)
	if err != nil {
		return err
	}

	exec, err := GetExecutor(file, ctx)
	if err != nil {
		return err
	}

	return exec.Execute(file, gist, args)
}

// return the username and the gist name or id
func ParseArgs(args string) (string, string, error) {
	var splitPath = strings.Split(args, "/")
	switch len(splitPath) {
	case 1:
		user := getGithubUsername()
		return user, splitPath[0], nil
	case 2:
		return splitPath[0], splitPath[1], nil
	default:
		return "", "", errors.New("malformed gist name")
	}
}
