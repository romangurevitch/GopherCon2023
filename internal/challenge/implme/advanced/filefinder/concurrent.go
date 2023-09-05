package filefinder

import (
	"context"
)

type concurrent struct {
	FileFinder
}

func NewConcurrent() FileFinder {
	return &concurrent{}
}

func (s *concurrent) FindFile(ctx context.Context, rootPath, filename string) (string, error) {
	panic("implement me!")
}
