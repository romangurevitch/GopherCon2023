package filefinder

import (
	"context"
	"errors"
	"os"
	"path/filepath"
)

var ErrNotFound = errors.New("file not found")

type FileFinder interface {
	FindFile(ctx context.Context, rootPath, filename string) (string, error)
}

func findInDir(dir, filename string) (string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	for _, entry := range entries {
		if entry.Name() == filename {
			return filepath.Join(dir, filename), nil
		}
	}
	return "", ErrNotFound
}
