package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		isErr  bool
	}{
		{
			name: "valid config",
			config: Config{
				BaseUrl: "http://example.com",
				Request: []Request{
					{
						Name:   "test1",
						Method: "GET",
						Path:   "/test",
					},
				},
			},
			isErr: false,
		},
		{
			name: "valid config with all fields",
			config: Config{
				BaseUrl:     "http://example.com",
				Timeout:     60,
				Concurrency: true,
				IgnoreFail:  true,
				Request: []Request{
					{
						Name:      "test1",
						Method:    "POST",
						Path:      "/test",
						Headers:   []string{"Content-Type: application/json"},
						JsonBody:  strPtrTest(`{"key":"value"}`),
						Extract:   map[string]string{"token": "$.token"},
						DependsOn: "auth",
					},
				},
			},
			isErr: false,
		},
		{
			name: "valid config with raw body",
			config: Config{
				BaseUrl: "http://example.com",
				Request: []Request{
					{
						Name:    "test1",
						Method:  "POST",
						Path:    "/test",
						RawBody: strPtrTest("raw data"),
					},
				},
			},
			isErr: false,
		},
		{
			name: "valid config with form body",
			config: Config{
				BaseUrl: "http://example.com",
				Request: []Request{
					{
						Name:     "test1",
						Method:   "POST",
						Path:     "/test",
						FormBody: strPtrTest("key=value"),
					},
				},
			},
			isErr: false,
		},
		{
			name: "missing base URL",
			config: Config{
				Request: []Request{
					{
						Name:   "test1",
						Method: "GET",
						Path:   "/test",
					},
				},
			},
			isErr: true,
		},
		{
			name: "no requests",
			config: Config{
				BaseUrl: "http://example.com",
			},
			isErr: true,
		},
		{
			name: "request missing name",
			config: Config{
				BaseUrl: "http://example.com",
				Request: []Request{
					{
						Method: "GET",
						Path:   "/test",
					},
				},
			},
			isErr: true,
		},
		{
			name: "duplicate request names",
			config: Config{
				BaseUrl: "http://example.com",
				Request: []Request{
					{
						Name:   "test1",
						Method: "GET",
						Path:   "/test1",
					},
					{
						Name:   "test1",
						Method: "POST",
						Path:   "/test2",
					},
				},
			},
			isErr: true,
		},
		{
			name: "request missing method",
			config: Config{
				BaseUrl: "http://example.com",
				Request: []Request{
					{
						Name: "test1",
						Path: "/test",
					},
				},
			},
			isErr: true,
		},
		{
			name: "request missing path",
			config: Config{
				BaseUrl: "http://example.com",
				Request: []Request{
					{
						Name:   "test1",
						Method: "GET",
					},
				},
			},
			isErr: true,
		},
		{
			name: "multiple body types - raw and json",
			config: Config{
				BaseUrl: "http://example.com",
				Request: []Request{
					{
						Name:     "test1",
						Method:   "POST",
						Path:     "/test",
						RawBody:  strPtrTest("raw"),
						JsonBody: strPtrTest("json"),
					},
				},
			},
			isErr: true,
		},
		{
			name: "multiple body types - form and json",
			config: Config{
				BaseUrl: "http://example.com",
				Request: []Request{
					{
						Name:     "test1",
						Method:   "POST",
						Path:     "/test",
						FormBody: strPtrTest("form=data"),
						JsonBody: strPtrTest("json"),
					},
				},
			},
			isErr: true,
		},
		{
			name: "multiple body types - raw and form",
			config: Config{
				BaseUrl: "http://example.com",
				Request: []Request{
					{
						Name:     "test1",
						Method:   "POST",
						Path:     "/test",
						RawBody:  strPtrTest("raw"),
						FormBody: strPtrTest("form=data"),
					},
				},
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func strPtrTest(s string) *string {
	return &s
}
