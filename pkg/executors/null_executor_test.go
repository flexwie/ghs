package executors_test

import (
	"testing"

	"github.com/flexwie/ghs/pkg/executors"
	"github.com/stretchr/testify/assert"
)

func TestMatchNull(t *testing.T) {
	exec := executors.NullExecutor{}

	assert.True(t, exec.Match(""))
}

func TestExecuteNull(t *testing.T) {
	exec := executors.NullExecutor{}

	err := exec.Execute("")
	assert.NotNil(t, err)
}
