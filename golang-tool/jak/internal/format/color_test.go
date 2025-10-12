package format

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColorizeByStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   int
		text     string
		expected string
	}{
		{"success", 200, "OK", Green + "OK" + Reset},
		{"redirect", 301, "Moved", Yellow + "Moved" + Reset},
		{"client error", 404, "Not Found", Red + "Not Found" + Reset},
		{"server error", 500, "Server Error", Bold + Red + "Server Error" + Reset},
		{"unknown", 0, "Unknown", "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ColorizeByStatus(tt.status, tt.text)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestColorFormatting(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		fn       func(string) string
	}{
		{"error", "Error", Red + "Error" + Reset, ColorizeError},
		{"warning", "Warning", Yellow + "Warning" + Reset, ColorizeWarning},
		{"success", "Success", Green + "Success" + Reset, ColorizeSuccess},
		{"info", "Info", Cyan + "Info" + Reset, ColorizeInfo},
		{"header", "Header", Bold + Blue + "Header" + Reset, ColorizeHeader},
		{"name", "Name", Bold + Magenta + "Name" + Reset, ColorizeName},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestColorizeMethod(t *testing.T) {
	tests := []struct {
		method   string
		expected string
	}{
		{"GET", Bold + Green + "GET" + Reset},
		{"POST", Bold + Yellow + "POST" + Reset},
		{"PUT", Bold + Blue + "PUT" + Reset},
		{"DELETE", Bold + Red + "DELETE" + Reset},
		{"PATCH", Bold + Cyan + "PATCH" + Reset},
		{"HEAD", Bold + Magenta + "HEAD" + Reset},
		{"OPTIONS", Bold + Magenta + "OPTIONS" + Reset},
		{"CUSTOM", Bold + "CUSTOM" + Reset},
	}

	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			result := ColorizeMethod(tt.method)
			assert.Equal(t, tt.expected, result)
		})
	}
}
