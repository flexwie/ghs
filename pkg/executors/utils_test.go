package executors_test

import (
	"os"
	"strings"
	"testing"

	"github.com/flexwie/ghs/pkg/executors"
	"github.com/stretchr/testify/assert"
)

func TestCreateFile(t *testing.T) {
	file, err := executors.MakeTmpFile()
	defer os.Remove(file.Name())

	assert.Nil(t, err)
	assert.FileExists(t, file.Name())
	assert.True(t, strings.Contains(file.Name(), "ghs"))
}

func TestCreateDir(t *testing.T) {
	dir, err := executors.MakeTmpDir()
	defer os.Remove(dir)

	assert.Nil(t, err)
	assert.DirExists(t, dir)
	assert.True(t, strings.Contains(dir, "ghs"))
}
