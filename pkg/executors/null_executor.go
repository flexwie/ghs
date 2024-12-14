package executors

import (
	"github.com/charmbracelet/log"
	"github.com/flexwie/ghs/pkg/github"
)

type NullExecutor struct {
}

func (n NullExecutor) Match(_ *github.File) bool {
	return true
}

func (n NullExecutor) Execute(_ *github.File, _ *github.Gist, _ []string) error {
	log.Warn("no executor found")
	return nil
}
