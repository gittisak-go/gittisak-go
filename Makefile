# Build variables
BINARY_NAME=mcp-server
BINARY_PATH=bin/$(BINARY_NAME)
CMD_PATH=./cmd/mcp-server

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

.PHONY: all build clean test run help

all: clean build

## build: Build the MCP server binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p bin
	$(GOBUILD) -o $(BINARY_PATH) $(CMD_PATH)
	@echo "Build complete: $(BINARY_PATH)"

## clean: Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf bin/
	@echo "Clean complete"

## test: Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

## run: Build and run the server
run: build
	@echo "Running $(BINARY_NAME)..."
	@$(BINARY_PATH)

## tidy: Tidy go modules
tidy:
	@echo "Tidying modules..."
	$(GOMOD) tidy

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' Makefile
