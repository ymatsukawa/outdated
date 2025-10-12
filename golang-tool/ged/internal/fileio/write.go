package fileio

import (
	"fmt"
	"os"
	"strings"
)

var (
	encryptExt = ".enc"
	decryptExt = ".dec"
)

func WriteFile(f File) error {
	var err error

	path, err := AbsPath(f.TargetPath)
	if err != nil {
		return fmt.Errorf("failed to get file: %w", err)
	}

	if f.CryptoAs == Encrypt {
		path = path + encryptExt
	}
	if f.CryptoAs == Decrypt {
		path = strings.TrimSuffix(path, encryptExt) + decryptExt
	}

	err = os.WriteFile(path, f.Data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
