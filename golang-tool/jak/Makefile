GO = go
BINARY_NAME = jak
GO_VERSION = 1.23.8
GOFLAGS = -v
PACKAGE = github.com/ymatsukawa/jak
BUILD_DIR = build
MAIN_FILE = main.go

# Test related variables
TEST_PACKAGES = ./internal/... ./cmd/...
COVERAGE_FILE = coverage.out
COVERAGE_HTML = coverage.html
MOCKGEN = go run go.uber.org/mock/mockgen

# Command related variables
RM = rm -f
MKDIR = mkdir -p

# Default target
.PHONY: all
all: clean build test

# Create build directory
$(BUILD_DIR):
	$(MKDIR) $(BUILD_DIR)

# Build the application
.PHONY: build
build: $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

# Development build (for local execution)
.PHONY: dev
dev:
	$(GO) build $(GOFLAGS) -o $(BINARY_NAME) $(MAIN_FILE)

# Run unit tests
.PHONY: test
test:
	$(GO) test $(GOFLAGS) $(TEST_PACKAGES)

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	$(GO) test $(GOFLAGS) -coverprofile=$(COVERAGE_FILE) $(TEST_PACKAGES)
	$(GO) tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "Coverage report generated at $(COVERAGE_HTML)"

# Install dependencies
.PHONY: deps
deps:
	$(GO) mod download
	$(GO) mod tidy

# Clean build files
.PHONY: clean
clean:
	$(RM) $(BINARY_NAME)
	$(RM) $(COVERAGE_FILE)
	$(RM) $(COVERAGE_HTML)
	$(RM) -r $(BUILD_DIR)

# Install (place binary in $GOPATH/bin)
.PHONY: install
install: build
	$(GO) install $(PACKAGE)
