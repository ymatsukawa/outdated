package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBaseBody(t *testing.T) {
	body := NewBaseBody("test content")
	assert.Equal(t, "test content", body.BodyContent)
}

func TestBaseBody_Content(t *testing.T) {
	body := BaseBody{BodyContent: "test"}
	assert.Equal(t, "test", body.Content())
}

func TestBaseBody_IsEmpty(t *testing.T) {
	body := BaseBody{BodyContent: ""}
	assert.True(t, body.IsEmpty())

	body.BodyContent = "test"
	assert.False(t, body.IsEmpty())
}

func TestRawBody(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		contentType string
		isEmpty     bool
		isErr       bool
	}{
		{
			name:        "valid raw body",
			content:     "content",
			contentType: "text/plain",
			isEmpty:     false,
			isErr:       false,
		},
		{
			name:        "empty raw body",
			content:     "",
			contentType: "",
			isEmpty:     true,
			isErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := NewRawBody(tt.content, tt.contentType)
			assert.Equal(t, tt.content, body.Content())
			assert.Equal(t, tt.contentType, body.ContentType())
			assert.Equal(t, tt.isEmpty, body.IsEmpty())
			if tt.isErr {
				assert.Error(t, body.Validate())
			} else {
				assert.NoError(t, body.Validate())
			}
		})
	}
}

func TestJsonBody(t *testing.T) {
	tests := []struct {
		name  string
		json  string
		isErr bool
	}{
		{
			name:  "valid JSON",
			json:  `{"key": "value"}`,
			isErr: false,
		},
		{
			name:  "invalid JSON",
			json:  "invalid json",
			isErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := NewJsonBody(tt.json)
			assert.Equal(t, ContentTypeJSON, body.ContentType())
			err := body.Validate()
			if tt.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFormBody(t *testing.T) {
	tests := []struct {
		name    string
		content string
		isEmpty bool
		isErr   bool
	}{
		{
			name:    "valid form data",
			content: "key=value",
			isEmpty: false,
			isErr:   false,
		},
		{
			name:    "empty form data",
			content: "",
			isEmpty: true,
			isErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := NewFormBody(tt.content)
			assert.Equal(t, ContentTypeFormURLEncoded, body.ContentType())
			assert.Equal(t, tt.content, body.Content())
			assert.Equal(t, tt.isEmpty, body.IsEmpty())
			if tt.isErr {
				assert.Error(t, body.Validate())
			} else {
				assert.NoError(t, body.Validate())
			}
		})
	}
}
