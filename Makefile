# Makefile

# Variables
APP_NAME := go-toyrobot
SRC_DIR := ./src
TEST_DIRS := $(shell find . -type d -name '*_test.go' -exec dirname {} \; | sort -u)

.PHONY: all test coverage run

# Default target
all: test

# Run tests with coverage
test:
	@echo "Running tests with coverage..."
	@go test -cover ./...

# Generate coverage report
coverage:
	@echo "Generating coverage report..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

# Start the application
run:
	@echo "Starting the application..."
	@go run main.go