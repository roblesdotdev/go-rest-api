# Variables
BIN_NAME := app
PKG := ./cmd/server/main.go
TEST_TARGET := ./...

# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOLINT := golangci-lint
GOBIN := $(shell $(GOCMD) env GOBIN)
GOFILES := $(shell find . -name "*.go" -type f)

# Build the project
build:
	$(GOBUILD) -o $(BIN_NAME) $(PKG)

# Run unit tests
test:
	$(GOTEST) -v $(TEST_TARGET)

# Lint the project
lint:
	$(GOLINT) run

# Run the project
run:
	docker compose up --build

# Install dependencies
deps:
	$(GOCMD) mod tidy

# Clean the project
clean:
	$(GOCMD) clean
	rm -f $(BIN_NAME)

.PHONY: build test lint run deps clean

