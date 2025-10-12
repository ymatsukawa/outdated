package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidMethod(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		expected bool
	}{
		{"valid GET", "GET", true},
		{"valid POST", "POST", true},
		{"valid lowercase", "post", true},
		{"invalid method", "INVALID", false},
		{"empty method", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidMethod(tt.method)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsBodyRequired(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		expected bool
	}{
		{"POST requires body", "POST", true},
		{"PUT requires body", "PUT", true},
		{"PATCH requires body", "PATCH", true},
		{"GET no body required", "GET", false},
		{"lowercase post", "post", true},
		{"invalid method", "INVALID", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsBodyRequired(tt.method)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNormalizeMethod(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		expected    string
		expectError bool
	}{
		{"valid GET", "GET", "GET", false},
		{"lowercase to upper", "post", "POST", false},
		{"mixed case", "PaTcH", "PATCH", false},
		{"invalid method", "INVALID", "", true},
		{"empty method", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NormalizeMethod(tt.method)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
