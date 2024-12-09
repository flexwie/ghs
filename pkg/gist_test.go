package pkg

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	url := "https://api.github.com/users/flexwie/gists"
	gistName := "test.sh"

	file, gist, err := SearchForGistFile(url, gistName, context.TODO())

	assert.Nil(t, err)
	assert.NotNil(t, file)
	assert.NotNil(t, gist)

	url = "https://api.github.com/users/unknown/gists"
	_, _, err = SearchForGistFile(url, gistName, context.TODO())
	assert.NotNil(t, err)
}
