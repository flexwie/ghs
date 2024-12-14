package executors

import (
	"os"
	"os/exec"
	"path"

	"github.com/charmbracelet/log"
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
	dir, err := MakeTmpDir()
	defer os.RemoveAll(dir)

	if err != nil {
		return err
	}

	// check if the gist includes a dependency file
	for _, f := range gist.Files {
		if f.Filename == "go.mod" {
			// write go mod file and install dependencies
			content, err := f.Content()
			if err != nil {
				return err
			}

			modFile, err := os.Create(path.Join(dir, "go.mod"))
			if err != nil {
				return err
			}

			modFile.Write([]byte(content))
			cmd := exec.Command("go", "mod", "tidy")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin

			if err := cmd.Run(); err != nil {
				return err
			}

			log.Debug("installed dependencies")
		}
	}

	// get script content
	content, err := file.Content()
	if err != nil {
		return err
	}

	tmpFile, err := os.Create(path.Join(dir, "main.go"))
	if err != nil {
		return err
	}

	tmpFile.Write([]byte(content))

	// build command and run
	cmd := BuildCommandExecutor("go", "run", tmpFile.Name())
	cmd.Dir = os.Getenv("PWD")

	err = cmd.Run()

	return err

}
