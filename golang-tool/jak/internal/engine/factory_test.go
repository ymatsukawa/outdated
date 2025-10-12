package engine

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymatsukawa/jak/internal/rule"
)

func TestCreateSimple(t *testing.T) {
	factory := NewFactory()
	var url = "http://example.com"

	tests := []struct {
		name      string
		url       string
		method    string
		header    string
		body      string
		expectErr bool
	}{
		{
			name:      "valid GET request",
			url:       url,
			method:    "get",
			header:    "Content-Type: application/json",
			body:      "test body",
			expectErr: false,
		},
		{
			name:      "valid POST request",
			url:       url + "/post",
			method:    "POST",
			header:    "Content-Type: application/json",
			body:      "{\"key\":\"value\"}",
			expectErr: false,
		},
		{
			name:      "valid PUT request with form data",
			url:       url + "/put",
			method:    "PUT",
			header:    "Content-Type: application/x-www-form-urlencoded",
			body:      "key1=value1&key2=value2",
			expectErr: false,
		},
		{
			name:      "valid DELETE request",
			url:       url + "/delete",
			method:    "DELETE",
			header:    "",
			body:      "",
			expectErr: false,
		},
		{
			name:      "valid request with multiple headers",
			url:       url,
			method:    "GET",
			header:    "Content-Type: application/json\nAuthorization: Bearer token",
			body:      "",
			expectErr: false,
		},
		{
			name:      "request with query parameters",
			url:       url + "/search?q=test&page=1",
			method:    "GET",
			header:    "",
			body:      "",
			expectErr: false,
		},
		{
			name:      "method with lowercase",
			url:       url,
			method:    "post",
			header:    "",
			body:      "",
			expectErr: false,
		},
		{
			name:      "empty URL",
			url:       "",
			method:    "GET",
			header:    "",
			body:      "",
			expectErr: false,
		},
		{
			name:      "invalid method",
			url:       url,
			method:    "invalid",
			header:    "",
			body:      "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := factory.CreateSimple(tt.url, tt.method, tt.header, tt.body)
			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.url, req.URL)
			assert.Equal(t, strings.ToUpper(tt.method), req.Method)
		})
	}
}

func TestCreateFromConfig(t *testing.T) {
	factory := NewFactory()
	var url = "http://example.com"

	type expect struct {
		url    string
		method string
	}

	baseConfig := &rule.Config{
		BaseUrl: url,
	}

	tests := []struct {
		name      string
		config    *rule.Config
		request   *rule.Request
		expect    expect
		expectErr bool
	}{
		{
			name:   "valid JSON request",
			config: baseConfig,
			request: &rule.Request{
				Name:     "post /test",
				Path:     "/test",
				Method:   "POST",
				Headers:  []string{"Content-Type: application/json"},
				JsonBody: strPtrTest(`{"key":"value"}`),
			},
			expect: expect{
				url:    url + "/test",
				method: "POST",
			},
			expectErr: false,
		},
		{
			name:   "form body request",
			config: baseConfig,
			request: &rule.Request{
				Name:     "post /form",
				Path:     "/form",
				Method:   "POST",
				FormBody: strPtrTest("key=value"),
			},
			expect: expect{
				url:    url + "/form",
				method: "POST",
			},
			expectErr: false,
		},
		{
			name:   "raw body request",
			config: baseConfig,
			request: &rule.Request{
				Name:    "put /raw",
				Path:    "/raw",
				Method:  "PUT",
				RawBody: strPtrTest("raw content"),
			},
			expect: expect{
				url:    url + "/raw",
				method: "PUT",
			},
			expectErr: false,
		},
		{
			name:   "GET request without body",
			config: baseConfig,
			request: &rule.Request{
				Name:   "get /test",
				Path:   "/test",
				Method: "GET",
			},
			expect: expect{
				url:    url + "/test",
				method: "GET",
			},
			expectErr: false,
		},
		{
			name:   "request with query parameters",
			config: baseConfig,
			request: &rule.Request{
				Name:   "get /test",
				Path:   "/test?param=value",
				Method: "GET",
			},
			expect: expect{
				url:    url + "/test?param=value",
				method: "GET",
			},
			expectErr: false,
		},
		{
			name:      "invalid config",
			config:    &rule.Config{},
			request:   &rule.Request{},
			expectErr: true,
		},
		{
			name:      "nil config",
			config:    nil,
			request:   &rule.Request{},
			expectErr: true,
		},
		{
			name:      "invalid method",
			config:    baseConfig,
			request:   &rule.Request{Method: "INVALID"},
			expectErr: true,
		},
		{
			name:      "nil request",
			config:    baseConfig,
			request:   nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.config != nil && tt.request != nil {
				tt.config.Request = []rule.Request{*tt.request}
			}

			req, err := factory.CreateFromConfig(tt.config, tt.request)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expect.url, req.URL)
			assert.Equal(t, tt.expect.method, req.Method)
		})
	}
}

func strPtrTest(s string) *string {
	return &s
}
