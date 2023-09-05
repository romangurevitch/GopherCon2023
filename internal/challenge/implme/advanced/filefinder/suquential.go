package filefinder

import (
	"context"
	"errors"
	"os"
	"path/filepath"
)

type sequential struct {
	FileFinder
}

func NewSequential() FileFinder {
	return &sequential{}
}

// FindFile searches for a file named filename starting at startPath.
func (s *sequential) FindFile(ctx context.Context, rootPath, filename string) (string, error) {
	// Try to find the file in the current directory.
	found, err := findInDir(rootPath, filename)
	if err == nil {
		return found, nil // File found
	}
	if !errors.Is(err, ErrNotFound) {
		return "", err // Some other error occurred
	}

	// If the file isn't found, search in subdirectories.
	entries, err := os.ReadDir(rootPath)
	if err != nil {
		return "", err // Return the error if we can't read the directory
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue // We are only interested in directories
		}

		// Recursively search in the subdirectory.
		subPath := filepath.Join(rootPath, entry.Name())
		found, err = s.FindFile(ctx, subPath, filename)
		if err == nil {
			return found, nil // File found in subdirectory
		}
		if !errors.Is(err, ErrNotFound) {
			return "", err // Some other error occurred in a subdirectory
		}
	}

	// If we reach this point, the file has not been found in any subdirectories.
	return "", ErrNotFound
}
