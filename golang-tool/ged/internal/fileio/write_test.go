package fileio

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymatsukawa/ged/test_util"
)

func TestWriteFile(t *testing.T) {
	validPath := test_util.MockDataPath()
	invalidPath := "/tmp/invalid_path/test.txt"

	testCases := []struct {
		name    string
		file    File
		isError bool
	}{
		{
			name: "when valid and write as encrypt",
			file: File{
				TargetPath: validPath,
				Data:       []byte("test data"),
				CryptoAs:   Encrypt,
			},
			isError: false,
		},
		{
			name: "when valid and write as decrypt",
			file: File{
				TargetPath: validPath,
				Data:       []byte("test data"),
				CryptoAs:   Decrypt,
			},
			isError: false,
		},
		{
			name: "when invalid because file does not exist",
			file: File{
				TargetPath: invalidPath,
			},
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := WriteFile(tc.file)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
			if tc.file.CryptoAs == Encrypt {
				_, err := os.Stat(tc.file.TargetPath + ".enc")
				assert.Nil(t, err)
			} else if tc.file.CryptoAs == Decrypt {
				_, err := os.Stat(tc.file.TargetPath + ".dec")
				assert.Nil(t, err)
			}
		})
	}
}
