package pkg

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, "apiUrl", "https://api.github.com/users/flexwie/gists")
	ctx = context.WithValue(ctx, "gist", "test.go")

	gist, file, err := SearchForGist(ctx)

	assert.Nil(t, err)
	assert.NotNil(t, file)
	assert.NotNil(t, gist)

	// _, err = pkg.SearchGist("https://api.github.com/users/unknown/gists", "test.sh")
	// assert.NotNil(t, err)
}
