package chain

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymatsukawa/jak/internal/engine"
	"github.com/ymatsukawa/jak/internal/http"
	"github.com/ymatsukawa/jak/internal/rule"
	mock_engine "github.com/ymatsukawa/jak/internal/test/mock/engine"
	mock_http "github.com/ymatsukawa/jak/internal/test/mock/http"
	"go.uber.org/mock/gomock"
)

// MockRequestProcessor is a mock implementation of the RequestProcessor interface
type MockRequestProcessor struct {
	ProcessRequestFunc func(ctx context.Context, request *rule.Request, config *rule.Config) (*ExecutionResult, error)
}

func (m *MockRequestProcessor) ProcessRequest(ctx context.Context, request *rule.Request, config *rule.Config) (*ExecutionResult, error) {
	if m.ProcessRequestFunc != nil {
		return m.ProcessRequestFunc(ctx, request, config)
	}
	return &ExecutionResult{StatusCode: 200}, nil
}

func TestNewChainExecutor(t *testing.T) {
	executor := NewChainExecutor()

	assert.NotNil(t, executor)
	assert.NotNil(t, executor.factory)
	assert.NotNil(t, executor.client)
	assert.NotNil(t, executor.variableResolver)
	assert.NotNil(t, executor.requestProcessor)
}

func TestChainExecutor_WithFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	executor := NewChainExecutor()
	mockFactory := mock_engine.NewMockFactory(ctrl)

	result := executor.WithFactory(mockFactory)

	assert.Equal(t, mockFactory, executor.factory)
	assert.Equal(t, executor, result, "Method should return the executor for chaining")
}

func TestChainExecutor_WithClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	executor := NewChainExecutor()
	mockClient := mock_http.NewMockClient(ctrl)

	result := executor.WithClient(mockClient)

	assert.Equal(t, mockClient, executor.client)
	assert.Equal(t, executor, result, "Method should return the executor for chaining")
}

func TestChainExecutor_WithVariables(t *testing.T) {
	executor := NewChainExecutor()
	mockResolver := &MockVariableResolver{}

	result := executor.WithVariables(mockResolver)

	assert.Equal(t, mockResolver, executor.variableResolver)
	assert.Equal(t, executor, result, "Method should return the executor for chaining")
}

func TestChainExecutorUpdateProcessor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	executor := NewChainExecutor()
	initialProcessor := executor.requestProcessor

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)
	mockResolver := &MockVariableResolver{}

	// Chain calls to update all dependencies
	executor.WithFactory(mockFactory).WithClient(mockClient).WithVariables(mockResolver)

	// Verify processor has been updated
	assert.NotEqual(t, initialProcessor, executor.requestProcessor,
		"Request processor should be updated when dependencies change")
}

func TestExecute_Success(t *testing.T) {
	// Create a config with two independent requests
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
				Method: "GET",
				Path:   "/api2",
			},
		},
	}

	// Create a custom executor with mock processor
	executor := NewChainExecutor()

	// Replace the processor with our mock
	processedRequests := make([]string, 0)
	mockProcessor := &MockRequestProcessor{
		ProcessRequestFunc: func(ctx context.Context, request *rule.Request, cfg *rule.Config) (*ExecutionResult, error) {
			processedRequests = append(processedRequests, request.Name)
			return &ExecutionResult{StatusCode: 200}, nil
		},
	}
	executor.requestProcessor = mockProcessor

	// Execute the chain
	ctx := context.Background()
	err := executor.Execute(ctx, config)

	// Verify execution
	assert.NoError(t, err)
	assert.Len(t, processedRequests, 2)
	assert.Contains(t, processedRequests, "req1")
	assert.Contains(t, processedRequests, "req2")
}

func TestExecute_RequestProcessingError(t *testing.T) {
	// Create a config with two requests
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
				Method: "GET",
				Path:   "/api2",
			},
		},
	}

	// Create a custom executor with mock processor that fails on the first request
	executor := NewChainExecutor()

	expectedErr := errors.New("request processing error")
	mockProcessor := &MockRequestProcessor{
		ProcessRequestFunc: func(ctx context.Context, request *rule.Request, cfg *rule.Config) (*ExecutionResult, error) {
			if request.Name == "req1" {
				return nil, expectedErr
			}
			return &ExecutionResult{StatusCode: 200}, nil
		},
	}
	executor.requestProcessor = mockProcessor

	// Execute the chain
	ctx := context.Background()
	err := executor.Execute(ctx, config)

	// Verify error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErr.Error())
}

func TestExecute_IgnoreFailFlag(t *testing.T) {
	// Create a config with two requests and ignore_fail flag
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
				Method: "GET",
				Path:   "/api2",
			},
		},
	}

	// Create a custom executor with mock processor that fails on the first request
	executor := NewChainExecutor()

	processedRequests := make([]string, 0)
	mockProcessor := &MockRequestProcessor{
		ProcessRequestFunc: func(ctx context.Context, request *rule.Request, cfg *rule.Config) (*ExecutionResult, error) {
			if request.Name == "req1" {
				return nil, errors.New("request failed but should be ignored")
			}
			processedRequests = append(processedRequests, request.Name)
			return &ExecutionResult{StatusCode: 200}, nil
		},
	}
	executor.requestProcessor = mockProcessor

	// Execute the chain
	ctx := context.Background()
	err := executor.Execute(ctx, config)

	// Should not return error, and should process req2
	assert.NoError(t, err)
	assert.Len(t, processedRequests, 1)
	assert.Contains(t, processedRequests, "req2")
}

func TestChainMethodFluency(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)
	mockResolver := &MockVariableResolver{}

	// Test method chaining fluency
	executor := NewChainExecutor().
		WithFactory(mockFactory).
		WithClient(mockClient).
		WithVariables(mockResolver)

	// Verify all components have been set
	assert.Equal(t, mockFactory, executor.factory)
	assert.Equal(t, mockClient, executor.client)
	assert.Equal(t, mockResolver, executor.variableResolver)
}

// TestCustomDependenciesExecution verifies that custom dependencies are properly used when executing
func TestCustomDependenciesExecution(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock dependencies
	mockFactory := mock_engine.NewMockFactory(ctrl)
	mockClient := mock_http.NewMockClient(ctrl)
	mockResolver := &MockVariableResolver{}

	// Set up executor with custom dependencies
	executor := NewChainExecutor().
		WithFactory(mockFactory).
		WithClient(mockClient).
		WithVariables(mockResolver)

	originalProcessor := executor.requestProcessor

	// Create a new processor to compare
	standardProcessor := NewRequestProcessor(engine.NewFactory(), http.NewClient(), NewVariableResolver())

	// The processors should be different types or have different field values
	assert.NotEqual(t, fmt.Sprintf("%v", standardProcessor), fmt.Sprintf("%v", originalProcessor),
		"Custom dependencies should result in a different processor configuration")
}
