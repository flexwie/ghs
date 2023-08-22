package pkg_test

import (
	"testing"

	"github.com/flexwie/ghs/pkg"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	file, err := pkg.SearchGist("https://api.github.com/users/flexwie/gists", "test.sh")

	assert.Nil(t, err)
	assert.NotNil(t, file)

	_, err = pkg.SearchGist("https://api.github.com/users/unknown/gists", "test.sh")
	assert.NotNil(t, err)
}

func TestFetch(t *testing.T) {
	file, _ := pkg.SearchGist("https://api.github.com/users/flexwie/gists", "test.sh")

	content, err := pkg.FetchGistContent(file)

	assert.Nil(t, err)
	assert.Greater(t, len(content), 0)
}
