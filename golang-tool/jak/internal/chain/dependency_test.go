package chain

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymatsukawa/jak/internal/rule"
	se "github.com/ymatsukawa/jak/internal/sys_error"
)

func TestNewDependencyResolver(t *testing.T) {
	dr := NewDependencyResolver()

	assert.NotNil(t, dr)
	assert.Empty(t, dr.requests)
	assert.Empty(t, dr.dependencies)
	assert.Empty(t, dr.dependsOn)
}

func TestBuildRequestGraph(t *testing.T) {
	tests := []struct {
		name      string
		config    *rule.Config
		expectErr error
	}{
		{
			name: "valid graph without dependencies",
			config: &rule.Config{
				Request: []rule.Request{
					{Name: "req1"},
					{Name: "req2"},
				},
			},
			expectErr: nil,
		},
		{
			name: "valid graph with dependencies",
			config: &rule.Config{
				Request: []rule.Request{
					{Name: "req1"},
					{Name: "req2", DependsOn: "req1"},
				},
			},
			expectErr: nil,
		},
		{
			name: "unknown dependency",
			config: &rule.Config{
				Request: []rule.Request{
					{Name: "req1", DependsOn: "nop"},
				},
			},
			expectErr: se.ErrUnknownDependency,
		},
		{
			name: "cyclic dependency",
			config: &rule.Config{
				Request: []rule.Request{
					{Name: "req1", DependsOn: "req2"},
					{Name: "req2", DependsOn: "req1"},
				},
			},
			expectErr: se.ErrCyclicDependency,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dr := NewDependencyResolver()
			ctx := context.Background()

			err := dr.BuildRequestGraph(ctx, tt.config)

			if tt.expectErr != nil {
				assert.ErrorIs(t, err, tt.expectErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// unstable. should be recreate testcase
func TestUnstable_CalculateExecutionOrder(t *testing.T) {
	tests := []struct {
		name        string
		config      *rule.Config
		expectOrder []string
		expectErr   error
	}{
		{
			name: "single request",
			config: &rule.Config{
				Request: []rule.Request{
					{Name: "req1"},
				},
			},
			expectOrder: []string{"req1"},
			expectErr:   nil,
		},
		{
			name: "multiple branches",
			config: &rule.Config{
				Request: []rule.Request{
					{Name: "root"},
					{Name: "A1", DependsOn: "root"},
					{Name: "A2", DependsOn: "A1"},
					{Name: "B1", DependsOn: "root"},
					{Name: "B2", DependsOn: "B1"},
					{Name: "edge", DependsOn: "A2"},
				},
			},
			// edge -> B2
			// B2 -> B1
			// B1 -> root
			// A2 -> A1
			// A1 -> root
			expectOrder: []string{"edge", "A2", "A1", "B2", "B1", "root"},
			expectErr:   nil,
		},
		{
			name: "squad pattern",
			config: &rule.Config{
				Request: []rule.Request{
					{Name: "root"},
					{Name: "A1", DependsOn: "root"},
					{Name: "A2", DependsOn: "A1"},
					{Name: "B1", DependsOn: "root"},
					{Name: "B2", DependsOn: "B1"},
					{Name: "C1", DependsOn: "root"},
					{Name: "C2", DependsOn: "C1"},
					{Name: "edgeA", DependsOn: "A2"},
					{Name: "edgeB", DependsOn: "B2"},
					{Name: "edgeC", DependsOn: "C2"},
				},
			},
			expectOrder: []string{"edgeC", "C2", "C1", "edegB", "B2", "B1", "edgeA", "A2", "A1", "root"},
			expectErr:   nil,
		},
		{
			name: "mesh pattern",
			config: &rule.Config{
				Request: []rule.Request{
					{Name: "A"},
					{Name: "B", DependsOn: "A"},
					{Name: "C", DependsOn: "A"},
					{Name: "D", DependsOn: "A"},
					{Name: "E", DependsOn: "B"},
					{Name: "E", DependsOn: "C"},
					{Name: "E", DependsOn: "D"},
					{Name: "F", DependsOn: "B"},
					{Name: "F", DependsOn: "C"},
					{Name: "F", DependsOn: "D"},
				},
			},
			// E,F -> (B,C,D) -> A
			expectOrder: []string{"F", "E", "D", "C", "B", "A"},
			expectErr:   nil,
		},
		{
			name: "cyclic dependency",
			config: &rule.Config{
				Request: []rule.Request{
					{Name: "req1", DependsOn: "req2"},
					{Name: "req2", DependsOn: "req3"},
					{Name: "req3", DependsOn: "req1"},
				},
			},
			expectOrder: nil,
			expectErr:   se.ErrCyclicDependency,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dr := NewDependencyResolver()
			ctx := context.Background()

			err := dr.BuildRequestGraph(ctx, tt.config)
			if err != nil {
				assert.ErrorIs(t, err, tt.expectErr)
				return
			}

			order, err := dr.CalculateExecutionOrder(ctx)
			if tt.expectErr != nil {
				assert.ErrorIs(t, err, tt.expectErr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectOrder, order)
		})
	}
}
