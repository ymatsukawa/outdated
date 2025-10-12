package fileio

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileCrypto uint8

const (
	Encrypt FileCrypto = iota
	Decrypt
)

type File struct {
	Data       []byte
	TargetPath string
	OutPath    *string
	CryptoAs   FileCrypto
}

func AbsPath(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("path is not specified")
	}

	absPath := path
	var err error
	if !filepath.IsAbs(path) {
		absPath, err = filepath.Abs(path)
		if err != nil {
			return "", fmt.Errorf("cannot convert relative path to abs path: %w", err)
		}
	}

	if err := isExistDirectory(absPath); err != nil {
		return "", fmt.Errorf("path search error: %w", err)
	}

	return absPath, nil
}

func isExistDirectory(path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %s", dir)
	}

	return nil
}
