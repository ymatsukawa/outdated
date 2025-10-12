package engine

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ymatsukawa/jak/internal/http"
	"github.com/ymatsukawa/jak/internal/rule"
	mock_engine "github.com/ymatsukawa/jak/internal/test/mock/engine"
	mock_http "github.com/ymatsukawa/jak/internal/test/mock/http"
	"go.uber.org/mock/gomock"
)

func TestNewExecutor(t *testing.T) {
	ctx := context.Background()
	executor := NewExecutor(ctx)

	assert.NotNil(t, executor)
	assert.NotNil(t, executor.client)
	assert.NotNil(t, executor.factory)
	assert.Equal(t, ctx, executor.ctx)
}

func TestWithClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_http.NewMockClient(ctrl)

	executor := NewExecutor(context.Background())
	result := executor.WithClient(mockClient)

	assert.Equal(t, mockClient, executor.client)
	assert.Equal(t, executor, result)
}

func TestWithFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)

	executor := NewExecutor(context.Background())
	result := executor.WithFactory(mockFactory)

	assert.Equal(t, mockFactory, executor.factory)
	assert.Equal(t, executor, result)
}

func TestMethodChaining(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_http.NewMockClient(ctrl)
	mockFactory := mock_engine.NewMockFactory(ctrl)

	executor := NewExecutor(context.Background())

	result := executor.WithClient(mockClient).WithFactory(mockFactory)

	assert.Equal(t, mockClient, executor.client)
	assert.Equal(t, mockFactory, executor.factory)
	assert.Equal(t, executor, result)
}

func TestExecuteSimple(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)
	mockRequest := &http.Request{}
	mockResponse := &http.Response{StatusCode: 200}

	// Test parameters
	url := "http://example.com"
	method := "GET"
	header := "Content-Type: application/json"
	body := `{"key":"value"}`

	mockFactory.EXPECT().
		CreateSimple(url, method, header, body).
		Return(mockRequest, nil)

	mockClient.EXPECT().
		Do(mockRequest).
		Return(mockResponse, nil)

	ctx := context.Background()
	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	resp, err := executor.ExecuteSimple(url, method, header, body)

	assert.NoError(t, err)
	assert.Equal(t, mockResponse, resp)
}

func TestExecuteSimple_FactoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)

	url := "http://example.com"
	method := "INVALID"
	header := ""
	body := ""

	expectedErr := errors.New("invalid method")

	mockFactory.EXPECT().
		CreateSimple(url, method, header, body).
		Return(nil, expectedErr)

	ctx := context.Background()
	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	resp, err := executor.ExecuteSimple(url, method, header, body)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), expectedErr.Error())
}

func TestExecuteSimple_ClientError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)
	mockRequest := &http.Request{}

	url := "http://example.com"
	method := "GET"
	header := ""
	body := ""

	expectedErr := errors.New("connection error")

	mockFactory.EXPECT().
		CreateSimple(url, method, header, body).
		Return(mockRequest, nil)

	mockClient.EXPECT().
		Do(mockRequest).
		Return(nil, expectedErr)

	ctx := context.Background()
	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	resp, err := executor.ExecuteSimple(url, method, header, body)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), expectedErr.Error())
}

func TestExecuteConfigRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)
	mockRequest := &http.Request{}
	mockResponse := &http.Response{StatusCode: 200}

	config := &rule.Config{
		BaseUrl: "http://example.com",
	}
	reqConfig := &rule.Request{
		Name:   "test",
		Method: "GET",
		Path:   "/api",
	}

	mockFactory.EXPECT().
		CreateFromConfig(config, reqConfig).
		Return(mockRequest, nil)

	mockClient.EXPECT().
		Do(mockRequest).
		Return(mockResponse, nil)

	ctx := context.Background()
	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	resp, err := executor.executeConfigRequest(config, reqConfig)

	assert.NoError(t, err)
	assert.Equal(t, mockResponse, resp)
}

func TestExecuteConfigRequest_FactoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)

	config := &rule.Config{
		BaseUrl: "http://example.com",
	}
	reqConfig := &rule.Request{
		Name:   "test",
		Method: "INVALID",
		Path:   "/api",
	}

	expectedErr := errors.New("invalid method")

	mockFactory.EXPECT().
		CreateFromConfig(config, reqConfig).
		Return(nil, expectedErr)

	ctx := context.Background()
	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	resp, err := executor.executeConfigRequest(config, reqConfig)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), expectedErr.Error())
}

func TestExecuteBatchSequential(t *testing.T) {
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

	gomock.InOrder(
		mockFactory.EXPECT().
			CreateFromConfig(config, &config.Request[0]).
			Return(mockRequest1, nil),

		mockClient.EXPECT().
			Do(mockRequest1).
			Return(mockResponse1, nil),

		mockFactory.EXPECT().
			CreateFromConfig(config, &config.Request[1]).
			Return(mockRequest2, nil),

		mockClient.EXPECT().
			Do(mockRequest2).
			Return(mockResponse2, nil),
	)

	ctx := context.Background()
	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	err := executor.ExecuteBatchSequential(config)

	assert.NoError(t, err)
}

func TestExecuteBatchSequential_EmptyRequests(t *testing.T) {
	config := &rule.Config{
		BaseUrl: "http://example.com",
		Request: []rule.Request{},
	}

	ctx := context.Background()
	executor := NewExecutor(ctx)

	err := executor.ExecuteBatchSequential(config)

	assert.NoError(t, err)
}

func TestExecuteBatchSequential_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)
	mockRequest := &http.Request{}

	config := &rule.Config{
		BaseUrl: "http://example.com",
		Request: []rule.Request{
			{
				Name:   "req1",
				Method: "GET",
				Path:   "/path1",
			},
			{
				Name:   "req2",
				Method: "POST",
				Path:   "/path2",
			},
		},
	}

	mockFactory.EXPECT().
		CreateFromConfig(config, &config.Request[0]).
		Return(mockRequest, nil)

	expectedErr := errors.New("request failed")

	mockClient.EXPECT().
		Do(mockRequest).
		Return(nil, expectedErr)

	ctx := context.Background()
	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	err := executor.ExecuteBatchSequential(config)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErr.Error())
	assert.Contains(t, err.Error(), "batch request failed")
}

func TestExecuteBatchSequential_IgnoreFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)
	mockRequest1 := &http.Request{}
	mockRequest2 := &http.Request{}
	mockResponse2 := &http.Response{StatusCode: 201}

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

	gomock.InOrder(
		mockFactory.EXPECT().
			CreateFromConfig(config, &config.Request[0]).
			Return(mockRequest1, nil),

		mockClient.EXPECT().
			Do(mockRequest1).
			Return(nil, errors.New("first request failed")),

		mockFactory.EXPECT().
			CreateFromConfig(config, &config.Request[1]).
			Return(mockRequest2, nil),

		mockClient.EXPECT().
			Do(mockRequest2).
			Return(mockResponse2, nil),
	)

	ctx := context.Background()
	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	err := executor.ExecuteBatchSequential(config)

	assert.NoError(t, err)
}

func TestContextTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Setup mocks
	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)

	// Test parameters
	config := &rule.Config{
		BaseUrl: "http://example.com",
		Request: []rule.Request{
			{
				Name:   "req1",
				Method: "GET",
				Path:   "/api1",
			},
		},
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	time.Sleep(5 * time.Millisecond)

	executor := NewExecutor(ctx)
	executor.factory = mockFactory
	executor.client = mockClient

	err := executor.ExecuteBatchSequential(config)
	assert.Error(t, err)
	assert.Equal(t, context.DeadlineExceeded, err)
}
