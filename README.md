# Toy Robot Simulator

A Go (Golang) implementation as per [PROBLEM](doc/PROBLEM.md) requirements.

## Prerequisites

The code has been tested with the latest Go version. Ensure your Go version is up-to-date by executing `go version`.

If you don't have Go installed, you can download it from https://golang.org/dl/

Clone the project and initialize the Go module dependencies.

## Steps to setup and run

```
git clone https://github.com/andrewbruno/go-toyrobot.git
cd go-toyrobot
go mod tidy
cd src/toyrobot
go run main.go
```

## Extra Features

CLI accepts "q" to quit program

## Testing via CLI

```
cd src/toyrobot
go run main.go 
```

## Testing via file input

```
cd src/toyrobot
go run main.go < scenarioA.txt
go run main.go < scenarioB.txt
go run main.go < scenarioC.txt
```

## Testing code coverage via GO

```
go test -cover ./...
```

Expected out:

```
?       toyrobot        [no test files]
ok      toyrobot/cmd    0.006s  coverage: 100.0% of statements
ok      toyrobot/compass        0.007s  coverage: 100.0% of statements
ok      toyrobot/table  0.007s  coverage: 100.0% of statements
ok      toyrobot/toy    0.011s  coverage: 100.0% of statements
ok      toyrobot/ui     0.013s  coverage: 100.0% of statements
ok      toyrobot/unit   0.010s  coverage: 100.0% of statements
```

## Why Go?

Golang is a beautiful language (subjective, I agree).  Let me explain why I like it:
  * code formatting is universal, part of SDK.  No more religious discussions about styles.
  * testing is out of the box.  No need for third party testing frameworks. Makes it ideal for TDD
  * is less verbose than say, Java.  No need for private, public, etc.
  * is statically typed language, so many checks done at compile time.
  * all variables must be used, cannot compile if a variable is declared and not used.
