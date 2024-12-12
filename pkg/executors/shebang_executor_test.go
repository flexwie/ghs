package executors_test

import (
	"testing"

	"github.com/flexwie/ghs/pkg/executors"
	"github.com/flexwie/ghs/pkg/github"
	"github.com/stretchr/testify/assert"
)

var testUrl = "https://gist.githubusercontent.com/flexwie/b9a1e66ac9dfcd2ffe4ce1e0a5f8ab46/raw/4c51d0592a37d35b1b57263f637ed49bee241fd7/test.sh"

func TestMatchShebang(t *testing.T) {
	exec := executors.ShebangExecutor{}

	works := &github.File{
		RawUrl: testUrl,
	}

	assert.True(t, exec.Match(works))
	//assert.False(t, exec.Match(fails))
}

func TestExecuteShebang(t *testing.T) {
	exec := executors.ShebangExecutor{}
	file := &github.File{RawUrl: testUrl}
	gist := &github.Gist{}

	err := exec.Execute(file, gist, []string{})
	assert.Nil(t, err)
}
