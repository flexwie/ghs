package executors

import (
	"context"
	"errors"

	"github.com/flexwie/ghs/pkg/github"
)

type NullExecutor struct {
}

func (n NullExecutor) Match(_ *github.File) bool {
	return true
}

func (n NullExecutor) Execute(_ *github.File, _ *github.Gist, _ context.Context) error {
	return errors.New("null executor can't actually execute")
}
