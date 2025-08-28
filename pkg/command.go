package pkg

import (
	"errors"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/cli/go-gh/v2"
	"github.com/flexwie/ghs/pkg/github"
	"github.com/urfave/cli/v2"
)

type GistRequest struct {
	ApiEndpoint string
	User        string
	GistName    string

	Gist          *github.Gist
	ExecutionFile *github.File
}

func Run(args []string) error {
	_, err := gh.Path()
	if err != nil {
		return errors.New("gh cli is not installed")
	}

	gistname := args[0]
	if len(gistname) == 0 {
		return cli.Exit("no gist provided", 2)
	}

	ghclient := github.NewGithubClient()

	args = args[1:]

	username, gistname, err := ParseArgs(gistname)
	if err != nil {
		return err
	}
	log.Debug("parsing args", "user", username, "gist", gistname, "args", args)

	file, gist, err := github.SearchForGistFile(username, gistname)
	if err != nil {
		return err
	}

	exec, err := GetExecutor(file)
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
		user := github.GetUsername()
		return user, splitPath[0], nil
	case 2:
		return splitPath[0], splitPath[1], nil
	default:
		return "", "", errors.New("malformed gist name")
	}
}
