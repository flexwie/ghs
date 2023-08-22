package pkg

import (
	"errors"
	"fmt"
	"strings"
)

func ExecGist(gist string) error {
	var splitPath = strings.Split(gist, "/")
	if len(splitPath) != 2 {
		return errors.New("malformed gist name")
	}

	apiUrl := fmt.Sprintf("https://api.github.com/users/%s/gists", splitPath[0])

	file, err := SearchGist(apiUrl, splitPath[1])
	if err != nil {
		return err
	}

	content, err := FetchGistContent(file)
	if err != nil {
		return err
	}

	return ExecuteGist(content)
}
