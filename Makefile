.PHONY: help build test clean run docker-build docker-run

# Default target
help:
	@echo "Available commands:"
	@echo "  build        - Build the Go application"
	@echo "  test         - Run tests with coverage"
	@echo "  clean        - Clean build artifacts"
	@echo "  run          - Run the application locally"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  lint         - Run linter (golangci-lint)"
	@echo "  fmt          - Format Go code"

# Build the application
build:
	@echo "Building Go application..."
	go build -o bin/devops-challenge main.go

# Run tests with coverage
test:
	@echo "Running tests..."
	go test -v -cover ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean -cache

# Run the application locally
run:
	@echo "Running application locally..."
	go run main.go

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t devops-challenge:latest .

# Run Docker container
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 -d PORT=8080 -e ENVIRONMENT=docker -e VERSION=1.0.0 devops-challenge:latest

# Run linter (requires golangci-lint to be installed)
lint:
	@echo "Running linter..."
	golangci-lint run

# Format Go code
fmt:
	@echo "Formatting Go code..."
	go fmt ./...
	go vet ./...

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

# Generate test coverage report
coverage:
	@echo "Generating coverage report..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
