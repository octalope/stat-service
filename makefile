# Define variables
GO_BINARY_NAME := stat-service
GO_SOURCE_FILES := $(shell find . -type f -name "*.go" -not -path "./bin/*")

# Default target
all: clean build test docker

prebuild:
	@echo "Preparing build environment..."
	go mod download
	go fmt ./...

# Build target
build: prebuild
	@echo "Building $(GO_BINARY_NAME)..."
	go build -o bin/$(GO_BINARY_NAME) ./modules

rebuild: clean build

docker: 
	@echo "Building Docker image for $(GO_BINARY_NAME)..."
	docker build -t $(GO_BINARY_NAME) -f Dockerfile .

run: build
	@echo "Running $(GO_BINARY_NAME)..."
	./bin/$(GO_BINARY_NAME)

run-docker: stop-docker docker
	@echo "Running Docker container for $(GO_BINARY_NAME)..."
	docker run --name $(GO_BINARY_NAME) --rm -d -p 8080:8080 $(GO_BINARY_NAME)

stop-docker:
	@echo "Stopping Docker container for $(GO_BINARY_NAME)..."
	docker stop $(GO_BINARY_NAME) >& /dev/null || true

test:
	@echo "Running tests..."
	go test ./...

clean:
	@echo "Cleaning up..."
	rm -f bin/$(GO_BINARY_NAME)
	go clean
	docker kill $(GO_BINARY_NAME) >& /dev/null || true
	docker rm -f $(GO_BINARY_NAME) >& /dev/null || true
	docker rmi -f $(GO_BINARY_NAME) >& /dev/null || true

# Help target
help:
	@echo "Available targets:"
	@echo "  all:           Build and run the application (default)"
	@echo "  build:         Build the Go binary"
	@echo "  clean:         Remove compiled binary, Go build cache, and Docker images/containers"
	@echo "  docker:        Build Docker image"
	@echo "  help:          Display this help message"
	@echo "  prebuild:      Prepare the build environment (download modules, format code)"
	@echo "  rebuild:       Clean and build the Go binary"
	@echo "  run-docker:    Run the application in a Docker container"
	@echo "  run:           Build and run the Go binary"
	@echo "  stop-docker:   Stop the running Docker container"
	@echo "  test:          Run all Go tests"