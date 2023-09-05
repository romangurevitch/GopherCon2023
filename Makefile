# Makefile for Golang Project

# Project Variables
PROJECT_NAME := myproject
GO := go
GOFMT := gofmt -s
GOFILES := $(shell find . -name "*.go" -type f)
GOTEST := $(GO) test
GOBENCH := $(GO) test -run=^$$ -bench .
GOLANGCI_LINT_VERSION := v1.55.1
GOLANGCI_LINT_FILE := bin/golangci-lint
GOLANGCI_LINT_VERSIONED := $(GOLANGCI_LINT_FILE)-$(GOLANGCI_LINT_VERSION)
GOLINT := $(GOLANGCI_LINT_VERSIONED) run

# Phony Targets
.PHONY: help build test bench lint fmt tidy setup clean

# Help
help:
	@echo "Choose a command run in "$(PROJECT_NAME)":"
	@echo "  test"
	@echo "    Run tests on a compiled project."
	@echo "  bench"
	@echo "    Run benchmarks on a compiled project."
	@echo "  lint"
	@echo "    Run all linters on the project."
	@echo "  fmt"
	@echo "    Run gofmt on all source files."
	@echo "  tidy"
	@echo "    Remove unused dependencies."
	@echo "  setup"
	@echo "    Setup project dependencies."
	@echo "  clean"
	@echo "    Remove temporary files and compiled binaries."

# Test Targets for different components
test:
	@echo "Testing concurrency components..."
	@$(GOTEST) ./internal/concurrency/...
	@echo "Testing pattern components..."
	@$(GOTEST) ./internal/pattern/...

# Test Targets for different components
test-basic:
	@echo "Running basic tests..."
	@$(GOTEST) -v -run=^TestBasic ./internal/challenge/fixme/...

# Test Targets for different components
test-intermediate:
	@echo "Running intermediate tests..."
	@$(GOTEST) -v -run=^TestIntermediate ./internal/challenge/fixme/...

# Test Targets for different components
test-advanced:
	@echo "Running advanced tests..."
	@$(GOTEST) -v -run=^TestAdvanced ./internal/challenge/fixme/...

# Benchmark Targets for different components
bench:
	@echo "Benchmarking concurrency components..."
	@$(GOBENCH) ./internal/concurrency/...
	@echo "Benchmarking pattern components..."
	@$(GOBENCH) ./internal/pattern/...

# Linting
lint: $(GOLANGCI_LINT_VERSIONED)
	@echo "Linting the code..."
	@$(GOLINT)

# Formatting
fmt:
	@echo "Formatting the code..."
	@$(GOFMT) -w $(shell grep -L -e '^// Code generated' $(GOFILES))

# Go mod tidy
tidy:
	@echo "Tidying up the go.mod and go.sum files..."
	@$(GO) mod tidy

# Setup
$(GOLANGCI_LINT_VERSIONED):
	@echo "Setting up..."
	@mkdir -p bin
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin $(GOLANGCI_LINT_VERSION)
	@mv $(GOLANGCI_LINT_FILE) $(GOLANGCI_LINT_VERSIONED)

# Cleaning up
clean:
	@echo "Cleaning up..."
	@rm -f $(PROJECT_NAME)
	@rm -f $(GOLANGCI_LINT_VERSIONED)
	@rm -f coverage.out
