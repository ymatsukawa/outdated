package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildFromConfig(t *testing.T) {
	type buildFromConfigTest struct {
		name          string
		method        string
		headers       []string
		json          *string
		form          *string
		raw           *string
		expectOptsLen int
	}

	tests := []buildFromConfigTest{
		{
			name:          "No options",
			method:        "GET",
			expectOptsLen: 0,
		},
		{
			name:          "With headers only",
			method:        "GET",
			headers:       []string{"Accept: application/json"},
			expectOptsLen: 1,
		},
		{
			name:          "Multiple headers",
			method:        "GET",
			headers:       []string{"Accept: application/json", "Authorization: Bearer token"},
			expectOptsLen: 1,
		},
		{
			name:          "With JSON body",
			method:        "POST",
			json:          strPtrTest(`{"foo":"bar"}`),
			expectOptsLen: 1,
		},
		{
			name:          "With form body",
			method:        "POST",
			form:          strPtrTest("foo=bar"),
			expectOptsLen: 1,
		},
		{
			name:          "With raw body",
			method:        "POST",
			raw:           strPtrTest("raw content"),
			expectOptsLen: 1,
		},
		{
			name:          "With headers and JSON body",
			method:        "POST",
			headers:       []string{"Content-Type: application/json"},
			json:          strPtrTest(`{"foo":"bar"}`),
			expectOptsLen: 2,
		},
		{
			name:          "With headers and form body",
			method:        "PUT",
			headers:       []string{"Content-Type: application/x-www-form-urlencoded"},
			form:          strPtrTest("key=value"),
			expectOptsLen: 2,
		},
		{
			name:          "With headers and raw body",
			method:        "PATCH",
			headers:       []string{"Content-Type: text/plain"},
			raw:           strPtrTest("raw data"),
			expectOptsLen: 2,
		},
		{
			name:          "DELETE with headers",
			method:        "DELETE",
			headers:       []string{"Authorization: Basic xyz"},
			expectOptsLen: 1,
		},
	}

	builder := NewRequestBuilder()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := builder.BuildFromConfig(tt.method, tt.headers, tt.json, tt.form, tt.raw)
			assert.Equal(t, tt.expectOptsLen, len(opts))
		})
	}
}

func TestBuildFromConfigEmptyInput(t *testing.T) {
	type buildFromConfigTest struct {
		name          string
		method        string
		headers       []string
		json          *string
		form          *string
		raw           *string
		expectOptsLen int
	}

	tests := []buildFromConfigTest{
		{
			name:          "Empty method",
			expectOptsLen: 0,
		},
		{
			name:          "All empty",
			headers:       []string{},
			expectOptsLen: 0,
		},
		{
			name:          "Empty headers with JSON body",
			method:        "POST",
			json:          strPtrTest("{}"),
			expectOptsLen: 1,
		},
		{
			name:          "Empty headers with form body",
			method:        "POST",
			form:          strPtrTest("a=b"),
			expectOptsLen: 1,
		},
		{
			name:          "Empty headers with raw body",
			method:        "POST",
			raw:           strPtrTest("test"),
			expectOptsLen: 1,
		},
		{
			name:          "Headers with empty bodies",
			method:        "POST",
			headers:       []string{"Content-Type: application/json"},
			json:          strPtrTest(""),
			form:          strPtrTest(""),
			raw:           strPtrTest(""),
			expectOptsLen: 2,
		},
		{
			name:          "Multiple bodies provided",
			method:        "POST",
			json:          strPtrTest("{}"),
			form:          strPtrTest("a=b"),
			raw:           strPtrTest("test"),
			expectOptsLen: 1,
		},
	}

	builder := NewRequestBuilder()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := builder.BuildFromConfig(tt.method, tt.headers, tt.json, tt.form, tt.raw)
			assert.Equal(t, tt.expectOptsLen, len(opts))
		})
	}
}

func TestBuildFromConfig_IllegalCase_ApiClientShouldValidate(t *testing.T) {
	type buildFromConfigTest struct {
		name    string
		method  string
		headers []string
		json    *string
		form    *string
		raw     *string
	}

	tests := []buildFromConfigTest{
		{
			name:   "Empty JSON body",
			method: "POST",
			json:   strPtrTest(""),
		},
		{
			name:   "Empty form body",
			method: "POST",
			form:   strPtrTest(""),
		},
		{
			name:   "Empty raw body",
			method: "POST",
			raw:    strPtrTest(""),
		},
		{
			name:    "Invalid header format",
			method:  "GET",
			headers: []string{"Invalid"},
		},
	}

	builder := NewRequestBuilder()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := builder.BuildFromConfig(tt.method, tt.headers, tt.json, tt.form, tt.raw)
			assert.Equal(t, 1, len(opts))
		})
	}
}

func TestDefaultRequestBuilder_BuildFromSimple(t *testing.T) {
	tests := []struct {
		name          string
		method        string
		header        string
		body          string
		expectOptsLen int
	}{
		{
			name:          "No options",
			method:        "GET",
			expectOptsLen: 0,
		},
		{
			name:          "With header only",
			method:        "GET",
			header:        "Accept: application/json",
			expectOptsLen: 1,
		},
		{
			name:          "With body only",
			method:        "POST",
			body:          `{"foo":"bar"}`,
			expectOptsLen: 1,
		},
	}

	builder := NewRequestBuilder()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := builder.BuildFromSimple(tt.method, tt.header, tt.body)
			assert.Equal(t, tt.expectOptsLen, len(opts))
		})
	}
}

func TestBuildFromSimpleEmptyInput(t *testing.T) {
	tests := []struct {
		name   string
		method string
		header string
		body   string
	}{
		{
			name:   "Empty body",
			method: "POST",
		},
		{
			name:   "Empty header",
			method: "POST",
		},
		{
			name: "Empty method",
		},
		{
			name: "All empty",
		},
	}

	builder := NewRequestBuilder()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := builder.BuildFromSimple(tt.method, tt.header, tt.body)
			assert.Equal(t, 0, len(opts))
		})
	}
}

func strPtrTest(s string) *string {
	return &s
}
