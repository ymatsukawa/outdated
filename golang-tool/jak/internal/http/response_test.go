package http

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	t.Run("TestResponse", func(t *testing.T) {
		resp := &Response{
			Header:     make(http.Header),
			StatusCode: 200,
		}
		resp.Header.Set("Content-Type", "application/json")

		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	})
}
