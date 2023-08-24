package executors_test

import (
	"context"
	"testing"

	"github.com/flexwie/ghs/pkg/executors"
	"github.com/flexwie/ghs/pkg/github"
	"github.com/stretchr/testify/assert"
)

func TestMatchNull(t *testing.T) {
	exec := executors.NullExecutor{}

	assert.True(t, exec.Match(&github.File{}))
}

func TestExecuteNull(t *testing.T) {
	exec := executors.NullExecutor{}

	err := exec.Execute(context.TODO())
	assert.NotNil(t, err)
}
