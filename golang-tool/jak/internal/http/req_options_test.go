package http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithRequestOptions(t *testing.T) {
	t.Run("context option", testContextOption)
	t.Run("header options", testHeaderOptions)
	t.Run("body options", testBodyOptions)
}

func testContextOption(t *testing.T) {
	req := &Request{}
	ctx := context.Background()

	WithContext(ctx)(req)

	assert.Equal(t, context.Background(), req.Context)
}

func testHeaderOptions(t *testing.T) {
	t.Run("single header", func(t *testing.T) {
		req := &Request{}
		WithHeader("Content-Type: application/json")(req)

		assert.Equal(t, "application/json", req.Headers.Get("Content-Type"))
	})

	t.Run("multiple headers", func(t *testing.T) {
		req := &Request{}
		WithHeaders([]string{
			"Content-Type: application/json",
			"Accept: text/plain",
		})(req)

		assert.Equal(t, "application/json", req.Headers.Get("Content-Type"))
		assert.Equal(t, "text/plain", req.Headers.Get("Accept"))
	})
}

func testBodyOptions(t *testing.T) {
	t.Run("JSON body", func(t *testing.T) {
		req := &Request{}
		WithJsonBody(`{"key":"value"}`)(req)

		assert.Equal(t, "application/json", req.Body.ContentType())
		assert.Equal(t, `{"key":"value"}`, req.Body.Content())
	})

	t.Run("form body", func(t *testing.T) {
		req := &Request{}
		WithFormBody("key=value")(req)

		assert.Equal(t, "application/x-www-form-urlencoded", req.Body.ContentType())
		assert.Equal(t, "key=value", req.Body.Content())
	})

	t.Run("raw body", func(t *testing.T) {
		req := &Request{}
		WithRawBody("raw data", "text/plain")(req)

		assert.Equal(t, "text/plain", req.Body.ContentType())
		assert.Equal(t, "raw data", req.Body.Content())
	})
}
