package http

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBaseHeaders(t *testing.T) {
	headers := NewBaseHeaders()
	assert.NotNil(t, headers)
	assert.Empty(t, headers.Headers)
}

func TestNewFromStringHeaders(t *testing.T) {
	tests := []struct {
		name     string
		headers  []string
		expected map[string]string
		isErr    bool
	}{
		{
			name:     "valid headers",
			headers:  []string{"Key1:Value1", "Key2:Value2"},
			expected: map[string]string{"Key1": "Value1", "Key2": "Value2"},
			isErr:    false,
		},
		{
			name:     "invalid format",
			headers:  []string{"InvalidHeader"},
			expected: nil,
			isErr:    true,
		},
		{
			name:     "empty key",
			headers:  []string{":Value"},
			expected: nil,
			isErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers, err := NewFromStringHeaders(tt.headers)
			if tt.isErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, headers.Headers)
		})
	}
}

func TestBaseHeaders_AddSetGet(t *testing.T) {
	headers := NewBaseHeaders()

	tests := []struct {
		name        string
		operation   string
		key         string
		value       string
		isErr       bool
		expectValue string
	}{
		{
			name:        "add valid header",
			operation:   "add",
			key:         "Key1",
			value:       "Value1",
			isErr:       false,
			expectValue: "Value1",
		},
		{
			name:        "add empty key",
			operation:   "add",
			key:         "",
			value:       "Value",
			isErr:       true,
			expectValue: "",
		},
		{
			name:        "set new value",
			operation:   "set",
			key:         "Key1",
			value:       "NewValue",
			isErr:       false,
			expectValue: "NewValue",
		},
		{
			name:        "set empty key",
			operation:   "set",
			key:         "",
			value:       "Value",
			isErr:       true,
			expectValue: "",
		},
		{
			name:        "get non-existent key",
			operation:   "get",
			key:         "NonExistentKey",
			value:       "",
			isErr:       false,
			expectValue: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			switch tt.operation {
			case "add":
				err = headers.Add(tt.key, tt.value)
			case "set":
				err = headers.Set(tt.key, tt.value)
			}

			if tt.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectValue, headers.Get(tt.key))
			}
		})
	}
}

func TestBaseHeaders_Apply(t *testing.T) {
	tests := []struct {
		name    string
		headers map[string]string
	}{
		{
			name:    "empty headers",
			headers: map[string]string{},
		},
		{
			name: "single header",
			headers: map[string]string{
				"Key1": "Value1",
			},
		},
		{
			name: "multiple headers",
			headers: map[string]string{
				"Key1": "Value1",
				"Key2": "Value2",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			baseHeaders := NewBaseHeaders()
			request := &http.Request{Header: make(http.Header)}

			for key, value := range test.headers {
				baseHeaders.Add(key, value)
			}

			baseHeaders.Apply(request)

			for key, expectedValue := range test.headers {
				actualValue := request.Header.Get(key)
				assert.Equal(t, expectedValue, actualValue)
			}
		})
	}
}

func TestBaseHeaders_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*BaseHeaders)
		expected bool
	}{
		{
			name:     "empty headers",
			setup:    func(h *BaseHeaders) {},
			expected: true,
		},
		{
			name: "non-empty headers",
			setup: func(h *BaseHeaders) {
				h.Add("Key", "Value")
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := NewBaseHeaders()
			tt.setup(headers)
			assert.Equal(t, tt.expected, headers.IsEmpty())
		})
	}
}

func TestBaseHeaders_GetAllHeaders(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected map[string]string
	}{
		{
			name:     "empty headers",
			headers:  map[string]string{},
			expected: map[string]string{},
		},
		{
			name: "multiple headers",
			headers: map[string]string{
				"Key1": "Value1",
				"Key2": "Value2",
			},
			expected: map[string]string{
				"Key1": "Value1",
				"Key2": "Value2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := NewBaseHeaders()
			for k, v := range tt.headers {
				headers.Add(k, v)
			}

			result := headers.GetAllHeaders()
			assert.Equal(t, tt.expected, result)

			result["NewKey"] = "NewValue"
			assert.Equal(t, tt.expected, headers.Headers)
		})
	}
}
