package executors

import (
	"os"
	"os/exec"
	"strings"

	"github.com/flexwie/ghs/pkg/github"
)

var _ Executor = ShebangExecutor{}

type ShebangExecutor struct {
}

func (n ShebangExecutor) Match(file *github.File) bool {
	content, err := file.Content()
	if err != nil {
		panic(err)
	}

	lines := strings.Split(content, "\n")

	if strings.HasPrefix(lines[0], "#!") {
		return true
	}

	return false
}

func (n ShebangExecutor) Execute(file *github.File, gist *github.Gist, args []string) error {
	content, err := file.Content()
	if err != nil {
		return err
	}

	dfile, err := os.CreateTemp("", "ghs")
	if err != nil {
		return err
	}
	defer os.Remove(dfile.Name())

	// write content to file and make executable
	_, err = dfile.Write([]byte(content))
	if err != nil {
		return err
	}

	err = dfile.Chmod(0777)
	if err != nil {
		return err
	}

	// run file as executable
	cmd := exec.Command(dfile.Name(), args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
