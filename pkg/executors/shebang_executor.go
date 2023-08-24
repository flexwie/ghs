package executors

import (
	"context"
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

func (n ShebangExecutor) Execute(ctx context.Context) error {
	file := ctx.Value("file").(*github.File)
	content, err := file.Content()
	if err != nil {
		return err
	}

	cmd := exec.Command("sh", "-c", content)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
