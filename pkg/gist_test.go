package pkg

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	username := "flexwie"
	gistName := "test.sh"

	file, gist, err := SearchForGistFile(username, gistName, context.TODO())

	assert.Nil(t, err)
	assert.NotNil(t, file)
	assert.NotNil(t, gist)

	username = "unknown"
	_, _, err = SearchForGistFile(username, gistName, context.TODO())
	assert.NotNil(t, err)
}
