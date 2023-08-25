package executors

import (
	"context"
	"testing"

	"github.com/flexwie/ghs/pkg/github"
	"github.com/stretchr/testify/assert"
)

var testUrl string = "https://gist.githubusercontent.com/flexwie/3251028d488d877baadb2a7fc33a6a84/raw/1e0cee31ef9f5f75a4e146fc1a46dbba5172c2f2/test.go"

func TestMatch(t *testing.T) {
	file := github.File{
		Language: "Go",
	}
	exec := GolangExecutor{}

	assert.True(t, exec.Match(&file))
}

func TestCanExecuteSingleFile(t *testing.T) {
	exec := GolangExecutor{}
	ctx := context.WithValue(context.TODO(), "file", &github.File{RawUrl: testUrl})
	ctx = context.WithValue(ctx, "gist", &github.Gist{})

	err := exec.Execute(ctx)
	assert.Nil(t, err)
}
