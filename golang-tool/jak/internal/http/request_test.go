package http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	var testURL = "http://example.com"

	t.Run("NewRequest creates request with basic options", func(t *testing.T) {
		req := NewRequest(testURL, "GET")
		assert.Equal(t, testURL, req.GetURL())
		assert.Equal(t, "GET", req.GetMethod())
	})

	t.Run("GetBody returns nil for empty body", func(t *testing.T) {
		req := NewRequest(testURL, "GET")
		body, err := req.GetBody()
		assert.Nil(t, err)
		assert.Nil(t, body)
	})

	t.Run("GetContentType returns empty for no body", func(t *testing.T) {
		req := NewRequest(testURL, "GET")
		assert.Equal(t, "", req.GetContentType())
	})

	t.Run("GetContentLength returns 0 for no body", func(t *testing.T) {
		req := NewRequest(testURL, "GET")
		assert.Equal(t, 0, req.GetContentLength())
	})

	t.Run("GetContentLengthString returns '0' for no body", func(t *testing.T) {
		req := NewRequest(testURL, "GET")
		assert.Equal(t, "0", req.GetContentLengthString())
	})

	t.Run("GetHeaders returns request headers", func(t *testing.T) {
		headers := NewBaseHeaders()
		headers.Add("Content-Type", "application/json")
		headers.Add("Authorization", "Bearer token")

		req := NewRequest(testURL, "GET")
		req.Headers = headers
		assert.Equal(t, headers, req.GetHeaders())
	})

	t.Run("WithContext sets context", func(t *testing.T) {
		req := NewRequest(testURL, "GET")
		ctx := context.Background()
		req = req.WithContext(ctx)
		assert.Equal(t, ctx, req.GetContext())
	})
}
