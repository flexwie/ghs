package executors

import "errors"

type NullExecutor struct {
}

func (n NullExecutor) Match(_ string) bool {
	return true
}

func (n NullExecutor) Execute(content string) error {
	return errors.New("null executor can't actually execute")
}
