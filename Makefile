# Todo App - TDD Compliant Makefile

APP_NAME=todo-backend
BINARY_NAME=todo-backend

.PHONY: help build run test test-all test-unit test-integration test-contract test-coverage clean deps

# Default target
all: help

help: ## Show available commands
	@echo "Todo App Commands (TDD Workflow):"
	@echo ""
	@echo "Building & Running:"
	@echo "  build              - Build the application"
	@echo "  run                - Run the application"
	@echo "  clean              - Clean build artifacts"
	@echo "  deps               - Install dependencies"
	@echo ""
	@echo "Testing (TDD Pyramid):"
	@echo "  test               - Run all tests (contract + integration + unit)"
	@echo "  test-contract      - Run contract tests (API contract validation)" 
	@echo "  test-integration   - Run integration tests (API + DB integration)"
	@echo "  test-unit          - Run unit tests (domain + application layer)"
	@echo "  test-coverage      - Generate test coverage report"
	@echo ""
	@echo "TDD Workflow:"
	@echo "  1. Write failing unit test (Red)"
	@echo "  2. Write failing integration test (Red)"  
	@echo "  3. Write failing contract test (Red)"
	@echo "  4. Write minimal code (Green)"
	@echo "  5. Refactor (Blue)"
	@echo "  6. Repeat until all tests pass"

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@go build -o $(BINARY_NAME) cmd/main.go
	@echo "Build completed!"

run: ## Run the application
	@echo "Starting $(APP_NAME)..."
	@go run cmd/main.go

# Test Commands - Following TDD Pyramid
test: test-unit test-integration test-contract ## Run all tests (TDD pyramid: unit → integration → contract)
	@echo "✅ All tests completed successfully!"

test-contract: ## Run contract tests (Consumer-Driven Contract validation)
	@echo "🤝 Running contract tests..."
	@go test -v ./test/contract/... -tags=contract
	@echo "✅ Contract tests completed!"

test-integration: ## Run integration tests (API + Database integration)
	@echo "🔗 Running integration tests..."
	@go test -v ./test/integration/... -tags=integration

test-unit: ## Run unit tests (Domain + Application layers)
	@echo "🧪 Running unit tests..."
	@go test -v ./test/unit/...
	@echo "✅ Unit tests completed!"

test-coverage: ## Generate test coverage report
	@echo "📊 Generating test coverage report..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "📋 Coverage report generated: coverage.html"

# TDD Workflow helpers
test-watch: ## Watch for changes and run tests (TDD Red-Green-Refactor)
	@echo "👀 Watching for changes (TDD workflow)..."
	@which air > /dev/null || (echo "Installing air..." && go install github.com/cosmtrek/air@latest)
	@air -c .air.toml

tdd-init: ## Initialize TDD workflow (run this before starting TDD)
	@echo "🚀 Initializing TDD workflow..."
	@echo "Step 1: Write your acceptance test first"
	@echo "Step 2: Run 'make test-acceptance' (should fail - Red)"
	@echo "Step 3: Write integration test" 
	@echo "Step 4: Run 'make test-integration' (should fail - Red)"
	@echo "Step 5: Write unit test"
	@echo "Step 6: Run 'make test-unit' (should fail - Red)"
	@echo "Step 7: Write minimal code to make tests pass"
	@echo "Step 8: Run 'make test' (should pass - Green)"
	@echo "Step 9: Refactor code while keeping tests green"
	@echo "Ready for TDD! 🎯"

clean: ## Clean build artifacts
	@echo "🧹 Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -f coverage.out coverage.html
	@go clean
	@echo "✅ Cleaned!"

deps: ## Install dependencies
	@echo "📦 Installing dependencies..."
	@go mod download
	@go mod tidy
	@echo "✅ Dependencies installed!"

# Development helpers
fmt: ## Format code
	@echo "🎨 Formatting code..."
	@go fmt ./...
	@echo "✅ Code formatted!"

lint: ## Run linter
	@echo "🔍 Running linter..."
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	@golangci-lint run
	@echo "✅ Linting completed!"

# Database helpers (for integration tests)
db-test-setup: ## Setup test database
	@echo "🗄️ Setting up test database..."
	@echo "Using in-memory SQLite for tests"
	@echo "✅ Test database ready!" 