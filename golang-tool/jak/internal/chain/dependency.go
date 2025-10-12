package chain

import (
	"context"
	"fmt"

	"github.com/ymatsukawa/jak/internal/rule"
	se "github.com/ymatsukawa/jak/internal/sys_error"
)

// DependencyResolver is responsible for analyzing and resolving dependencies between requests.
// It builds a dependency graph and calculates the optimal execution order.
type DependencyResolver struct {
	// requests maps request names to their corresponding request objects
	requests map[string]*rule.Request

	// dependencies maps a request name to the names of requests that depend on it
	// Example: If B depends on A, then dependencies["A"] contains "B"
	dependencies map[string][]string

	// dependsOn maps a request name to the name of the request it depends on
	// Example: If B depends on A, then dependsOn["B"] = "A"
	dependsOn map[string]string
}

// NewDependencyResolver creates a new dependency resolver with initialized maps.
// It returns a DependencyResolver ready to analyze request dependencies.
func NewDependencyResolver() *DependencyResolver {
	return &DependencyResolver{
		requests:     make(map[string]*rule.Request),
		dependencies: make(map[string][]string),
		dependsOn:    make(map[string]string),
	}
}

// BuildRequestGraph analyzes the provided configuration and builds a dependency graph.
// It validates that all dependencies exist and detects any cyclic dependencies.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - config: Configuration containing requests and their dependencies
//
// Returns:
//   - error: If there are unknown dependencies or cycles in the dependency graph
func (dr *DependencyResolver) BuildRequestGraph(ctx context.Context, config *rule.Config) error {
	// Check for context cancellation
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// Index requests by name for easy lookup
	for i := range config.Request {
		req := &config.Request[i]
		dr.requests[req.Name] = req
	}

	// Process dependencies and build the dependency maps
	return dr.buildDependencyMaps(ctx, config)
}

// buildDependencyMaps processes all requests and builds the dependency relationships.
// It validates that all referenced dependencies exist and then checks for cycles.
func (dr *DependencyResolver) buildDependencyMaps(ctx context.Context, config *rule.Config) error {
	// Check for context cancellation
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// Process each request's dependency
	for _, req := range config.Request {
		if req.DependsOn != "" {
			// Verify the dependency exists
			if _, exists := dr.requests[req.DependsOn]; !exists {
				return fmt.Errorf("request '%s' depends on unknown request '%s': %w",
					req.Name, req.DependsOn, se.ErrUnknownDependency)
			}

			// Record the dependency relationships in both maps
			dr.dependencies[req.DependsOn] = append(dr.dependencies[req.DependsOn], req.Name)
			dr.dependsOn[req.Name] = req.DependsOn
		}
	}

	// After building the dependency maps, check for cyclic dependencies
	return dr.detectCycles(ctx)
}

// detectCycles checks for cyclic dependencies in the dependency graph.
// A cyclic dependency occurs when a chain of dependencies loops back to itself.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//
// Returns:
//   - error: If a cyclic dependency is detected
func (dr *DependencyResolver) detectCycles(ctx context.Context) error {
	// Check for context cancellation
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// Check each request for cycles in its dependency chain
	for name := range dr.requests {
		visited := make(map[string]bool)
		if dr.hasCycle(name, visited) {
			return fmt.Errorf("cyclic dependency detected with request '%s': %w", name, se.ErrCyclicDependency)
		}
	}

	return nil
}

// hasCycle determines if there is a cycle in the dependency chain starting from nodeName.
// It uses depth-first traversal with a visited map to detect cycles.
//
// Parameters:
//   - nodeName: The name of the request to start checking from
//   - visited: A map tracking visited nodes in the current path
//
// Returns:
//   - bool: True if a cycle is detected, false otherwise
func (dr *DependencyResolver) hasCycle(nodeName string, visited map[string]bool) bool {
	// If we've already visited this node in the current path, we've found a cycle
	if visited[nodeName] {
		return true
	}

	// Get the dependency for this node
	dependsOn, exists := dr.dependsOn[nodeName]
	if !exists {
		// No dependency means no cycle possible from here
		return false
	}

	// Mark this node as visited in the current path
	visited[nodeName] = true

	// Recursively check the node we depend on
	if dr.hasCycle(dependsOn, visited) {
		return true
	}

	// Backtrack: remove from visited as we're leaving this path
	delete(visited, nodeName)
	return false
}

// CalculateExecutionOrder determines the order in which requests should be executed
// based on their dependencies. It returns a topologically sorted list of request names.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//
// Returns:
//   - []string: Ordered list of request names for execution
//   - error: If a cyclic dependency is detected or context is cancelled
func (dr *DependencyResolver) CalculateExecutionOrder(ctx context.Context) ([]string, error) {
	// Check for context cancellation
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// Prepare result and tracking maps
	result := make([]string, 0, len(dr.requests))
	visited := make(map[string]bool)  // Requests that have been fully processed
	visiting := make(map[string]bool) // Requests currently being processed (for cycle detection)

	// Define the depth-first traversal function
	var visit func(name string) error
	visit = func(name string) error {
		// Check if we're revisiting a node in the current traversal path (cycle)
		if visiting[name] {
			return fmt.Errorf("cyclic dependency detected with request '%s': %w", name, se.ErrCyclicDependency)
		}

		// Skip if already processed
		if visited[name] {
			return nil
		}

		// Mark as being visited in current traversal
		visiting[name] = true

		// First process any dependency this request has
		if dep, hasDep := dr.dependsOn[name]; hasDep {
			if err := visit(dep); err != nil {
				return err
			}
		}

		// Mark as done with this traversal path and fully processed
		delete(visiting, name)
		visited[name] = true

		// Add to result after all dependencies are processed
		result = append(result, name)
		return nil
	}

	// Process all requests
	for name := range dr.requests {
		if err := visit(name); err != nil {
			return nil, err
		}
	}

	// Reverse the result to get the correct execution order
	// (Dependencies should be executed before the requests that depend on them)
	dr.reverseStringSlice(result)

	return result, nil
}

// reverseStringSlice reverses the order of elements in a string slice in-place.
// This is used to convert the post-order traversal result to the correct execution order.
func (dr *DependencyResolver) reverseStringSlice(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
