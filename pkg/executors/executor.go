package executors

import (
	"context"
	"os"
	"os/exec"

	"github.com/flexwie/ghs/pkg/github"
)

type Executor interface {
	Match(*github.File) bool
	Execute(*github.File, *github.Gist, context.Context) error
}

var ExecutorPipeline = []Executor{
	ShebangExecutor{},
	GolangExecutor{},
	NullExecutor{},
}

func BuildCommandExecutor(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd
}
