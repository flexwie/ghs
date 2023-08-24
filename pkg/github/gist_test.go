package github

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileContent(t *testing.T) {
	file := &File{
		RawUrl: "https://gist.githubusercontent.com/flexwie/b9a1e66ac9dfcd2ffe4ce1e0a5f8ab46/raw/4c51d0592a37d35b1b57263f637ed49bee241fd7/test.sh",
	}

	content, err := file.Content()

	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(content, "#!/bin/sh"))
}
