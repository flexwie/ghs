package executors

import (
	"errors"
	"os"

	"github.com/flexwie/ghs/pkg/github"
)

var _ Executor = GolangExecutor{}

type GolangExecutor struct{}

func (e GolangExecutor) Match(file *github.File) bool {
	if file.Language == "Go" {
		return true
	}

	return false
}

func (e GolangExecutor) Execute(file *github.File, gist *github.Gist, args []string) error {
	// check if the gist includes a dependency file
	var hasDependencies bool = false
	for _, f := range gist.Files {
		if f.Filename == "go.mod" {
			hasDependencies = true
		}
	}

	if hasDependencies {
		// TODO: download dependencies

		return errors.New("go executor can't handle dependencies so far")
	} else {
		content, err := file.Content()
		if err != nil {
			return err
		}

		file, err := os.CreateTemp("", "ghs.*.go")
		if err != nil {
			return err
		}
		file.Write([]byte(content))

		cmd := BuildCommandExecutor("go", "run", file.Name())
		cmd.Dir = os.Getenv("PWD")

		err = cmd.Run()

		os.Remove(file.Name())

		return err
	}
}
