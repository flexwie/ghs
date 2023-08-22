package executors

type Executor interface {
	Match(string) bool
	Execute(string) error
}

var ExecutorPipeline = []Executor{
	ShebangExecutor{},
	NullExecutor{},
}
