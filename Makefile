.PHONY: all test run

# Default target
all: test

# Run tests with coverage
test:
	go test -cover ./...

# Start the application
run:
	go run main.go

e2e:
	@go run main.go < scenarioA.txt
	@go run main.go < scenarioB.txt
	@go run main.go < scenarioC.txt