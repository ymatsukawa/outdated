package fileio

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymatsukawa/ged/test_util"
)

func TestReadFile(t *testing.T) {
	validPath := test_util.MockDataPath()

	testCases := []struct {
		name    string
		file    File
		isError bool
	}{
		{
			name: "when valid",
			file: File{
				TargetPath: validPath,
			},
			isError: false,
		},
		{
			name: "when invalid because file does not found as absolute path",
			file: File{
				TargetPath: "/tmp/im_not_exist.txt",
			},
			isError: true,
		},
		{
			name: "when invalid because file does not found as relative path",
			file: File{
				TargetPath: "./also_im_not_exist.txt",
			},
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bins, err := ReadFile(tc.file)
			if tc.isError {
				assert.Error(t, err)
				assert.Nil(t, bins)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, bins)
			}
		})
	}
}
