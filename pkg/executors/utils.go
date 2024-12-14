package executors

import (
	"os"

	"github.com/charmbracelet/log"
)

func MakeTmpFile() (*os.File, error) {
	file, err := os.CreateTemp("", "ghs")
	if err != nil {
		return nil, err
	}

	log.Debug("created temp file", "path", file.Name())

	return file, nil
}

func MakeTmpDir() (string, error) {
	dir, err := os.MkdirTemp("", "ghs")
	if err != nil {
		return "", err
	}

	log.Debug("created temp dir", "path", dir)

	return dir, nil
}
