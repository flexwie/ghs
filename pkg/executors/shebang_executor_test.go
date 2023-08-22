package executors_test

import (
	"testing"

	"github.com/flexwie/ghs/pkg/executors"
	"github.com/stretchr/testify/assert"
)

func TestMatchShebang(t *testing.T) {
	exec := executors.ShebangExecutor{}

	works := `#!/bin/bash
echo hi`

	fails := `///bin/bash
echo hi`

	assert.True(t, exec.Match(works))
	assert.False(t, exec.Match(fails))
}

func TestExecuteShebang(t *testing.T) {
	exec := executors.ShebangExecutor{}

	err := exec.Execute("")
	assert.Nil(t, err)
}
