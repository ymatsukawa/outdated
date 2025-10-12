package fileio

import (
	"fmt"
	"os"
)

func ReadFile(f File) ([]byte, error) {
	absPath, err := AbsPath(f.TargetPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get file: %w", err)
	}

	bins, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return bins, nil
}
