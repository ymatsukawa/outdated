package engine

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymatsukawa/jak/internal/http"
	"github.com/ymatsukawa/jak/internal/rule"
	mock_engine "github.com/ymatsukawa/jak/internal/test/mock/engine"
	mock_http "github.com/ymatsukawa/jak/internal/test/mock/http"
	"go.uber.org/mock/gomock"
)

func TestExecuteBatchConcurrent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)
	mockRequest1 := &http.Request{}
	mockRequest2 := &http.Request{}
	mockResponse1 := &http.Response{StatusCode: 200}
	mockResponse2 := &http.Response{StatusCode: 201}

	config := &rule.Config{
		BaseUrl: "http://example.com",
		Request: []rule.Request{
			{
				Name:   "req1",
				Method: "GET",
				Path:   "/api1",
			},
			{
				Name:   "req2",
				Method: "POST",
				Path:   "/api2",
			},
		},
	}

	mockFactory.EXPECT().
		CreateFromConfig(config, &config.Request[0]).
		Return(mockRequest1, nil)

	mockClient.EXPECT().
		Do(mockRequest1).
		Return(mockResponse1, nil)

	mockFactory.EXPECT().
		CreateFromConfig(config, &config.Request[1]).
		Return(mockRequest2, nil)

	mockClient.EXPECT().
		Do(mockRequest2).
		Return(mockResponse2, nil)

	ctx := context.Background()
	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	err := executor.ExecuteBatchConcurrent(config)

	assert.NoError(t, err)
}

func TestExecuteBatchConcurrent_EmptyRequests(t *testing.T) {
	// Test parameters
	config := &rule.Config{
		BaseUrl: "http://example.com",
		Request: []rule.Request{},
	}

	ctx := context.Background()
	executor := NewExecutor(ctx)

	err := executor.ExecuteBatchConcurrent(config)

	assert.NoError(t, err)
}

func TestExecuteBatchConcurrent_MoreRequestsThanWorkers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)

	config := &rule.Config{
		BaseUrl: "http://example.com",
		Request: make([]rule.Request, DefaultMaxWorkers+5),
	}

	for i := 0; i < len(config.Request); i++ {
		config.Request[i] = rule.Request{
			Name:   fmt.Sprintf("req%d", i+1),
			Method: "GET",
			Path:   fmt.Sprintf("/api%d", i+1),
		}

		mockRequest := &http.Request{}
		mockResponse := &http.Response{StatusCode: 200}

		mockFactory.EXPECT().
			CreateFromConfig(config, &config.Request[i]).
			Return(mockRequest, nil)

		mockClient.EXPECT().
			Do(mockRequest).
			Return(mockResponse, nil)
	}

	ctx := context.Background()
	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	err := executor.ExecuteBatchConcurrent(config)

	assert.NoError(t, err)
}

func TestExecuteBatchConcurrent_LessRequestsThanWorkers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)

	requestCount := DefaultMaxWorkers - 2
	config := &rule.Config{
		BaseUrl: "http://example.com",
		Request: make([]rule.Request, requestCount),
	}

	for i := 0; i < len(config.Request); i++ {
		config.Request[i] = rule.Request{
			Name:   fmt.Sprintf("req%d", i+1),
			Method: "GET",
			Path:   fmt.Sprintf("/api%d", i+1),
		}

		mockRequest := &http.Request{}
		mockResponse := &http.Response{StatusCode: 200}

		mockFactory.EXPECT().
			CreateFromConfig(config, &config.Request[i]).
			Return(mockRequest, nil)

		mockClient.EXPECT().
			Do(mockRequest).
			Return(mockResponse, nil)
	}

	ctx := context.Background()
	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	err := executor.ExecuteBatchConcurrent(config)

	assert.NoError(t, err)
}

func TestExecuteBatchConcurrent_FailWithoutIgnore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)
	mockRequest1 := &http.Request{}
	mockRequest2 := &http.Request{}
	mockRequest3 := &http.Request{}
	mockResponse1 := &http.Response{StatusCode: 200}
	mockResponse3 := &http.Response{StatusCode: 200}

	config := &rule.Config{
		BaseUrl:    "http://example.com",
		IgnoreFail: false,
		Request: []rule.Request{
			{
				Name:   "req1",
				Method: "GET",
				Path:   "/api1",
			},
			{
				Name:   "req2",
				Method: "POST",
				Path:   "/api2",
			},
			{
				Name:   "req3",
				Method: "GET",
				Path:   "/api3",
			},
		},
	}

	expectedErr := errors.New("second request failed")

	mockFactory.EXPECT().
		CreateFromConfig(config, &config.Request[0]).
		Return(mockRequest1, nil)

	mockClient.EXPECT().
		Do(mockRequest1).
		Return(mockResponse1, nil)

	// fail case
	mockFactory.EXPECT().
		CreateFromConfig(config, &config.Request[1]).
		Return(mockRequest2, nil)

	mockClient.EXPECT().
		Do(mockRequest2).
		Return(nil, expectedErr)

	mockFactory.EXPECT().
		CreateFromConfig(config, &config.Request[2]).
		Return(mockRequest3, nil).
		AnyTimes()

	mockClient.EXPECT().
		Do(mockRequest3).
		Return(mockResponse3, nil).
		AnyTimes()

	ctx := context.Background()
	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	err := executor.ExecuteBatchConcurrent(config)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErr.Error())
	assert.Contains(t, err.Error(), "batch request failed")
}

func TestExecuteBatchConcurrent_AllFailWithIgnore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)
	mockRequest1 := &http.Request{}
	mockRequest2 := &http.Request{}
	mockRequest3 := &http.Request{}

	config := &rule.Config{
		BaseUrl:    "http://example.com",
		IgnoreFail: true,
		Request: []rule.Request{
			{
				Name:   "req1",
				Method: "GET",
				Path:   "/api1",
			},
			{
				Name:   "req2",
				Method: "POST",
				Path:   "/api2",
			},
			{
				Name:   "req3",
				Method: "GET",
				Path:   "/api3",
			},
		},
	}

	mockFactory.EXPECT().
		CreateFromConfig(config, &config.Request[0]).
		Return(mockRequest1, nil)
	mockClient.EXPECT().
		Do(mockRequest1).
		Return(nil, errors.New("first request failed"))

	mockFactory.EXPECT().
		CreateFromConfig(config, &config.Request[1]).
		Return(mockRequest2, nil)
	mockClient.EXPECT().
		Do(mockRequest2).
		Return(nil, errors.New("second request failed"))

	mockFactory.EXPECT().
		CreateFromConfig(config, &config.Request[2]).
		Return(mockRequest3, nil)
	mockClient.EXPECT().
		Do(mockRequest3).
		Return(nil, errors.New("third request failed"))

	ctx := context.Background()
	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	err := executor.ExecuteBatchConcurrent(config)

	assert.NoError(t, err)
}

// Tests concurrent batch execution with ignore fail flag
func TestExecuteBatchConcurrent_IgnoreFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Setup mocks
	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)
	mockRequest1 := &http.Request{}
	mockRequest2 := &http.Request{}
	mockResponse2 := &http.Response{StatusCode: 201}

	// Test parameters
	config := &rule.Config{
		BaseUrl:    "http://example.com",
		IgnoreFail: true,
		Request: []rule.Request{
			{
				Name:   "req1",
				Method: "GET",
				Path:   "/api1",
			},
			{
				Name:   "req2",
				Method: "POST",
				Path:   "/api2",
			},
		},
	}

	// Set expectations - order not guaranteed
	mockFactory.EXPECT().
		CreateFromConfig(config, &config.Request[0]).
		Return(mockRequest1, nil)

	mockClient.EXPECT().
		Do(mockRequest1).
		Return(nil, errors.New("first request failed"))

	mockFactory.EXPECT().
		CreateFromConfig(config, &config.Request[1]).
		Return(mockRequest2, nil)

	mockClient.EXPECT().
		Do(mockRequest2).
		Return(mockResponse2, nil)

	// Run test
	ctx := context.Background()
	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	err := executor.ExecuteBatchConcurrent(config)

	// Assert results
	assert.NoError(t, err)
}
