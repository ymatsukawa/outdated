package chain

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymatsukawa/jak/internal/http"
	se "github.com/ymatsukawa/jak/internal/sys_error"
)

func TestGJSONExtractor_Extract(t *testing.T) {
	tests := []struct {
		name         string
		jsonString   string
		path         string
		expect       string
		expectExists bool
	}{
		{
			name:         "simple path exists",
			jsonString:   `{"name": "test"}`,
			path:         "name",
			expect:       "test",
			expectExists: true,
		},
		{
			name:         "nested path exists",
			jsonString:   `{"user": {"name": "test", "age": 30}}`,
			path:         "user.name",
			expect:       "test",
			expectExists: true,
		},
		{
			name:         "array element exists",
			jsonString:   `{"users": ["test1", "test2"]}`,
			path:         "users.0",
			expect:       "test1",
			expectExists: true,
		},
		{
			name:         "path does not exist",
			jsonString:   `{"name": "test"}`,
			path:         "invalid.path",
			expect:       "",
			expectExists: false,
		},
		{
			name:         "empty json string",
			jsonString:   `{}`,
			path:         "name",
			expect:       "",
			expectExists: false,
		},
		{
			name:         "number value",
			jsonString:   `{"age": 25}`,
			path:         "age",
			expect:       "25",
			expectExists: true,
		},
		{
			name:         "boolean value",
			jsonString:   `{"active": true}`,
			path:         "active",
			expect:       "true",
			expectExists: true,
		},
		{
			name:         "null value",
			jsonString:   `{"data": null}`,
			path:         "data",
			expect:       "",
			expectExists: true,
		},
		{
			name:         "complex nested array",
			jsonString:   `{"users": [{"name": "test1"}, {"name": "test2"}]}`,
			path:         "users.1.name",
			expect:       "test2",
			expectExists: true,
		},
		{
			name:         "array index out of bounds",
			jsonString:   `{"users": ["test1"]}`,
			path:         "users.1",
			expect:       "",
			expectExists: false,
		},
		{
			name:         "invalid json string",
			jsonString:   `{invalid}`,
			path:         "name",
			expect:       "",
			expectExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &GJSONExtractor{}
			got, exists := e.Extract(tt.jsonString, tt.path)
			assert.Equal(t, tt.expect, got)
			assert.Equal(t, tt.expectExists, exists)
		})
	}
}

func TestExtractVariables(t *testing.T) {
	tests := []struct {
		name        string
		response    *http.Response
		extractions map[string]string
		expect      map[string]string
		expectErr   error
	}{
		{
			name:      "nil response returns error",
			response:  nil,
			expectErr: se.ErrNilResponse,
		},
		{
			name: "successfully extracts variables",
			response: &http.Response{
				Body: io.NopCloser(strings.NewReader(`{"name": "test", "age": 30}`)),
			},
			extractions: map[string]string{
				"username": "name",
				"userAge":  "age",
			},
			expect: map[string]string{
				"username": "test",
				"userAge":  "30",
			},
			expectErr: nil,
		},
		{
			name: "path not found returns error",
			response: &http.Response{
				Body: io.NopCloser(strings.NewReader(`{"name": "test"}`)),
			},
			extractions: map[string]string{
				"notFound": "invalid.path",
			},
			expectErr: se.ErrPathNotFound,
		},
		{
			name: "successfully extracts nested variables",
			response: &http.Response{
				Body: io.NopCloser(strings.NewReader(`{"user": {"name": "test", "details": {"age": 30, "active": true}}}`)),
			},
			extractions: map[string]string{
				"name":   "user.name",
				"age":    "user.details.age",
				"active": "user.details.active",
			},
			expect: map[string]string{
				"name":   "test",
				"age":    "30",
				"active": "true",
			},
			expectErr: nil,
		},
		{
			name: "successfully extracts array elements",
			response: &http.Response{
				Body: io.NopCloser(strings.NewReader(`{"users": [{"name": "test1"}, {"name": "test2"}]}`)),
			},
			extractions: map[string]string{
				"firstUser":  "users.0.name",
				"secondUser": "users.1.name",
			},
			expect: map[string]string{
				"firstUser":  "test1",
				"secondUser": "test2",
			},
			expectErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			extractor := newVariableExtractor(nil)
			got, err := extractor.ExtractVariables(context.Background(), tt.response, tt.extractions)

			if tt.expectErr != nil {
				assert.ErrorIs(t, err, tt.expectErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, got)
			}
		})
	}
}
