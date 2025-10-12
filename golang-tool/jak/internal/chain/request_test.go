package chain

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymatsukawa/jak/internal/http"
	"github.com/ymatsukawa/jak/internal/rule"
	mock_engine "github.com/ymatsukawa/jak/internal/test/mock/engine"
	mock_http "github.com/ymatsukawa/jak/internal/test/mock/http"
	"go.uber.org/mock/gomock"
)

// MockVariableResolver is a mock implementation of the VariableResolver interface
type MockVariableResolver struct {
	ResolveFunc        func(string) string
	ResolveHeadersFunc func([]string) []string
	ResolveBodyFunc    func(*string) *string
	SetFunc            func(string, string) error
	GetFunc            func(string) (string, bool)
}

func (m *MockVariableResolver) Resolve(input string) string {
	if m.ResolveFunc != nil {
		return m.ResolveFunc(input)
	}
	return input
}

func (m *MockVariableResolver) ResolveHeaders(headers []string) []string {
	if m.ResolveHeadersFunc != nil {
		return m.ResolveHeadersFunc(headers)
	}
	return headers
}

func (m *MockVariableResolver) ResolveBody(body *string) *string {
	if m.ResolveBodyFunc != nil {
		return m.ResolveBodyFunc(body)
	}
	return body
}

func (m *MockVariableResolver) Set(name, value string) error {
	if m.SetFunc != nil {
		return m.SetFunc(name, value)
	}
	return nil
}

func (m *MockVariableResolver) Get(name string) (string, bool) {
	if m.GetFunc != nil {
		return m.GetFunc(name)
	}
	return "", false
}

func TestNewRequestProcessor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)
	mockResolver := &MockVariableResolver{}

	processor := NewRequestProcessor(mockFactory, mockClient, mockResolver)

	assert.NotNil(t, processor)
	// Check if the return type is DefaultRequestProcessor
	_, ok := processor.(*DefaultRequestProcessor)
	assert.True(t, ok, "Expected DefaultRequestProcessor type")
}

func TestProcessRequest(t *testing.T) {
	// Test case: Successful request processing without variable extraction
	t.Run("successful request processing without variable extraction", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Setup mocks
		mockFactory := mock_engine.NewMockFactory(ctrl)
		mockClient := mock_http.NewMockClient(ctrl)
		mockResolver := &MockVariableResolver{
			ResolveFunc: func(input string) string {
				return input
			},
		}

		// Prepare request and response
		config := &rule.Config{BaseUrl: "http://example.com"}
		request := &rule.Request{
			Name:   "test",
			Method: "GET",
			Path:   "/api",
		}
		mockHttpReq := &http.Request{}
		mockResp := &http.Response{StatusCode: 200}

		// Set mock expectations
		mockFactory.EXPECT().
			CreateFromConfig(config, gomock.Any()).
			Return(mockHttpReq, nil)

		mockClient.EXPECT().
			Do(mockHttpReq).
			Return(mockResp, nil)

		// Execute test
		ctx := context.Background()
		processor := NewRequestProcessor(mockFactory, mockClient, mockResolver)
		result, err := processor.ProcessRequest(ctx, request, config)

		// Verify results
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 200, result.StatusCode)
		assert.Empty(t, result.Variables)
	})

	// Test case: Successful request processing with variable extraction
	t.Run("successful request processing with variable extraction", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Setup mocks
		mockFactory := mock_engine.NewMockFactory(ctrl)
		mockClient := mock_http.NewMockClient(ctrl)

		// Record variable settings
		setVars := make(map[string]string)
		mockResolver := &MockVariableResolver{
			ResolveFunc: func(input string) string {
				return input
			},
			SetFunc: func(name, value string) error {
				setVars[name] = value
				return nil
			},
		}

		// Prepare request and response
		config := &rule.Config{BaseUrl: "http://example.com"}
		request := &rule.Request{
			Name:    "test",
			Method:  "GET",
			Path:    "/api",
			Extract: map[string]string{"username": "name"},
		}
		mockHttpReq := &http.Request{}

		// Response body with JSON string
		responseBody := `{"name":"test user","age":30}`
		mockResp := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(responseBody)),
		}

		// Set mock expectations
		mockFactory.EXPECT().
			CreateFromConfig(config, gomock.Any()).
			Return(mockHttpReq, nil)

		mockClient.EXPECT().
			Do(mockHttpReq).
			Return(mockResp, nil)

		// Execute test
		ctx := context.Background()
		processor := NewRequestProcessor(mockFactory, mockClient, mockResolver)
		result, err := processor.ProcessRequest(ctx, request, config)

		// Verify results
		// Note: Actual verification depends on variableExtractor
		// so we only check that there's no error
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 200, result.StatusCode)
	})

	// Test case: Context cancellation
	t.Run("context cancellation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Setup mocks
		mockFactory := mock_engine.NewMockFactory(ctrl)
		mockClient := mock_http.NewMockClient(ctrl)
		mockResolver := &MockVariableResolver{}

		// Execute test
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel the context

		processor := NewRequestProcessor(mockFactory, mockClient, mockResolver)
		result, err := processor.ProcessRequest(ctx, &rule.Request{}, &rule.Config{})

		// Verify results
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, ctx.Err(), err)
	})

	// Test case: Error during request preparation
	t.Run("prepare request error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Setup mocks
		mockFactory := mock_engine.NewMockFactory(ctrl)
		mockClient := mock_http.NewMockClient(ctrl)

		mockResolver := &MockVariableResolver{
			ResolveFunc: func(input string) string {
				return input
			},
		}

		// Execute test
		ctx := context.Background()
		processor := NewRequestProcessor(mockFactory, mockClient, mockResolver)

		// Cause an error with nil request
		result, err := processor.ProcessRequest(ctx, nil, &rule.Config{})

		// Verify results
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to prepare request")
	})

	// Test case: Request execution error - factory error
	t.Run("execute request error - factory error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Setup mocks
		mockFactory := mock_engine.NewMockFactory(ctrl)
		mockClient := mock_http.NewMockClient(ctrl)
		mockResolver := &MockVariableResolver{
			ResolveFunc: func(input string) string {
				return input
			},
		}

		// Prepare request and error
		config := &rule.Config{BaseUrl: "http://example.com"}
		request := &rule.Request{
			Name:   "test",
			Method: "GET",
			Path:   "/api",
		}
		expectedErr := errors.New("factory error")

		// Set mock expectations
		mockFactory.EXPECT().
			CreateFromConfig(config, gomock.Any()).
			Return(nil, expectedErr)

		// Execute test
		ctx := context.Background()
		processor := NewRequestProcessor(mockFactory, mockClient, mockResolver)
		result, err := processor.ProcessRequest(ctx, request, config)

		// Verify results
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to create request")
	})

	// Test case: Request execution error - client error
	t.Run("execute request error - client error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Setup mocks
		mockFactory := mock_engine.NewMockFactory(ctrl)
		mockClient := mock_http.NewMockClient(ctrl)
		mockResolver := &MockVariableResolver{
			ResolveFunc: func(input string) string {
				return input
			},
		}

		// Prepare request and error
		config := &rule.Config{BaseUrl: "http://example.com"}
		request := &rule.Request{
			Name:   "test",
			Method: "GET",
			Path:   "/api",
		}
		mockHttpReq := &http.Request{}
		expectedErr := errors.New("client error")

		// Set mock expectations
		mockFactory.EXPECT().
			CreateFromConfig(config, gomock.Any()).
			Return(mockHttpReq, nil)

		mockClient.EXPECT().
			Do(mockHttpReq).
			Return(nil, expectedErr)

		// Execute test
		ctx := context.Background()
		processor := NewRequestProcessor(mockFactory, mockClient, mockResolver)
		result, err := processor.ProcessRequest(ctx, request, config)

		// Verify results
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "request execution failed")
	})

	// Test case: Variable extraction error
	t.Run("variable extraction error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Setup mocks
		mockFactory := mock_engine.NewMockFactory(ctrl)
		mockClient := mock_http.NewMockClient(ctrl)
		mockResolver := &MockVariableResolver{
			ResolveFunc: func(input string) string {
				return input
			},
		}

		// Prepare request and response
		config := &rule.Config{BaseUrl: "http://example.com"}
		request := &rule.Request{
			Name:    "test",
			Method:  "GET",
			Path:    "/api",
			Extract: map[string]string{"invalid": "$.nonexistent.path"},
		}
		mockHttpReq := &http.Request{}

		// Cause a variable extraction error with an invalid path
		responseBody := `{"data":{"value":"test"}}`
		mockResp := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(responseBody)),
		}

		// Set mock expectations
		mockFactory.EXPECT().
			CreateFromConfig(config, gomock.Any()).
			Return(mockHttpReq, nil)

		mockClient.EXPECT().
			Do(mockHttpReq).
			Return(mockResp, nil)

		// Execute test
		ctx := context.Background()
		processor := NewRequestProcessor(mockFactory, mockClient, mockResolver)
		result, err := processor.ProcessRequest(ctx, request, config)

		// Verify results - should have a variable extraction error
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to extract variables")
	})
}
