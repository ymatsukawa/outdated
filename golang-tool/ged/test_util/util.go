package test_util

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func MockDataPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get caller")
	}
	return filepath.Join(filepath.Dir(filename), "mock_data", "encrypt_me.txt")
}

func TempDirPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get caller")
	}
	return filepath.Join(filepath.Dir(filename), "..", "tmp", "test")
}

func TearDownTestDir() {
	tempPath := TempDirPath()
	testDir := filepath.Join(tempPath, "test")

	targetExts := []string{".txt", ".txt.enc", ".txt.dec"}

	err := filepath.Walk(testDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			for _, ext := range targetExts {
				if strings.HasSuffix(path, ext) {
					return os.Remove(path)
				}
			}
		}
		return nil
	})

	if err != nil {
		panic("failed to teardown: " + err.Error())
	}
}
