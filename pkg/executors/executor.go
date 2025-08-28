package executors

import (
	"os"
	"os/exec"

	"github.com/flexwie/ghs/pkg/github"
)

type Executor interface {
	Match(*github.File) bool
	Execute(*github.File, *github.Gist, []string) error
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

type Exec interface {
	Match() bool
	InstallDependencies()
	Run()
}
