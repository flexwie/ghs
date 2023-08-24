package executors

import (
	"context"

	"github.com/flexwie/ghs/pkg/github"
)

type Executor interface {
	Match(*github.File) bool
	Execute(context.Context) error
}

var ExecutorPipeline = []Executor{
	ShebangExecutor{},
	NullExecutor{},
}
