package executors

import (
	"context"

	"github.com/flexwie/ghs/pkg/github"
)

var _ Executor = GolangExecutor{}

type GolangExecutor struct{}

func (e GolangExecutor) Match(*github.File) bool {
	return false
}

func (e GolangExecutor) Execute(context.Context) error {
	return nil
}
