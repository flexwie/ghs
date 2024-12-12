package pkg

import (
	"context"
	"errors"
	"fmt"
	"regexp"
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

func ParseArgs(args string) (string, string, error) {
	// check if its the gist id
	re := regexp.MustCompile(`(?m)^[0-9a-z]{32}$`)

	var splitPath = strings.Split(args, "/")
	switch len(splitPath) {
	case 1:
		if re.Match([]byte(splitPath[0])) {
			// found id of gist
			return "", "", errors.ErrUnsupported
		} else {
			user := getGithubUsername()
			return fmt.Sprintf("https://api.github.com/users/%s/gists", user), splitPath[0], nil
		}
	case 2:
		return fmt.Sprintf("https://api.github.com/users/%s/gists", splitPath[0]), splitPath[1], nil
	default:
		return "", "", errors.New("malformed gist name")
	}
}
