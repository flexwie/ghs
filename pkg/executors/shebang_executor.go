package executors

import (
	"os"
	"os/exec"
	"strings"
)

var _ Executor = ShebangExecutor{}

type ShebangExecutor struct {
}

func (n ShebangExecutor) Match(content string) bool {
	lines := strings.Split(content, "\n")

	if strings.HasPrefix(lines[0], "#!") {
		return true
	}

	return false
}

func (n ShebangExecutor) Execute(content string) error {
	cmd := exec.Command("sh", "-c", content)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
