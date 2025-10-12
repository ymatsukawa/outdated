package chain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	se "github.com/ymatsukawa/jak/internal/sys_error"
)

func TestVariableResolver(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*DefaultVariableResolver)
		input    string
		expected string
	}{
		{
			name:     "empty input returns empty",
			setup:    func(r *DefaultVariableResolver) {},
			input:    "",
			expected: "",
		},
		{
			name: "simple variable substitution",
			setup: func(r *DefaultVariableResolver) {
				r.Set("name", "value")
			},
			input:    "Hello ${name}",
			expected: "Hello value",
		},
		{
			name: "nested variable substitution",
			setup: func(r *DefaultVariableResolver) {
				r.Set("inner", "world")
				r.Set("outer", "Hello ${inner}")
			},
			input:    "${outer}!",
			expected: "Hello world!",
		},
		{
			name:     "undefined variable remains unchanged",
			setup:    func(r *DefaultVariableResolver) {},
			input:    "Hello ${undefined}",
			expected: "Hello ${undefined}",
		},
		{
			name: "multiple variables in string",
			setup: func(r *DefaultVariableResolver) {
				r.Set("first", "Hello")
				r.Set("second", "world")
			},
			input:    "${first} ${second}!",
			expected: "Hello world!",
		},
		{
			name: "resolve headers",
			setup: func(r *DefaultVariableResolver) {
				r.Set("key", "value")
			},
			input:    "Header: ${key}",
			expected: "Header: value",
		},
		{
			name: "truncate long values",
			setup: func(r *DefaultVariableResolver) {
				r.Set("long", "a"+strings.Repeat("b", maxVariableValueLength+10))
			},
			input:    "${long}",
			expected: "a" + strings.Repeat("b", maxVariableValueLength-1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resolver := NewVariableResolver()
			tt.setup(resolver)

			result := resolver.Resolve(tt.input)

			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSetEmptyVariableName(t *testing.T) {
	resolver := NewVariableResolver()
	err := resolver.Set("", "value")

	assert.Error(t, err)
	assert.Equal(t, se.ErrEmptyVariableName, err)
}

func TestGetVariable(t *testing.T) {
	resolver := NewVariableResolver()
	resolver.Set("key", "value")
	val, exists := resolver.Get("key")
	assert.True(t, exists)
	assert.Equal(t, "value", val)

	val, exists = resolver.Get("nop")
	assert.False(t, exists)
	assert.Empty(t, val)
}

func TestResolveHeadersSlice(t *testing.T) {
	resolver := NewVariableResolver()
	resolver.Set("key", "value")
	headers := []string{"Header1: ${key}", "Header2: static"}
	resolved := resolver.ResolveHeaders(headers)

	assert.Equal(t, []string{"Header1: value", "Header2: static"}, resolved)
}

func TestResolveBody(t *testing.T) {
	resolver := NewVariableResolver()
	resolver.Set("key", "value")
	body := "Body with ${key}"
	resolved := resolver.ResolveBody(&body)
	assert.Equal(t, "Body with value", *resolved)

	nilBody := (*string)(nil)
	assert.Nil(t, resolver.ResolveBody(nilBody))

	emptyBody := ""
	assert.Equal(t, &emptyBody, resolver.ResolveBody(&emptyBody))
}
