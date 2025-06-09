# Toy Robot Simulator

A Go (Golang) implementation as per [PROBLEM](doc/PROBLEM.md) requirements.

The coding challenge was created by [Jon Eaves](https://www.linkedin.com/in/joneaves/) at [ANZ](https://joneaves.wordpress.com/2014/07/21/toy-robot-coding-test/).

Over the past decade, it has been a common interview take home test for many companies.

I have decided to open source it because today many companies are using [HackerRank](https://www.hackerrank.com/), [LeetCode](https://leetcode.com/), [CodeSignal](https://codesignal.com/) or their own custom tests and there are close to [1000 open source projects](https://github.com/search?q=toyrobot&type=repositories).

## Architecture Overview

The project follows a clean, modular architecture, adhering to separation of concerns.

Below is a structural outline:

```
go-toyrobot/
├── main.go              # Entry point
├── ui/                  # User interface layer
├── cmd/                 # Command parsing
├── table/               # Game board logic
├── toy/                 # Robot entity
├── compass/             # Direction handling
├── unit/                # Coordinate types
└── doc/                 # Documentation
```

> ✅ Reviewed with [Amazon Q for Dev CLI](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-installing.html) on 9 June 2025 for best practices and structural integrity.

### Strengths

1. Excellent Test Coverage: 100% coverage across all packages
2. Clean Architecture: Well-separated concerns with clear interfaces
3. Type Safety: Custom types for coordinates and directions prevent errors
4. Robust Input Validation: Regex-based command parsing
5. Interface Design: Good use of interfaces for testability


## Prerequisites

The code was originally written for Go 1.1 in Nov 2017, but in Mar 2025 I trialed Copilot Agent mode as a learning exercise to see if it could fix the code and upgrade it to Go version 1.23.1, as per this [commit](https://github.com/andrewbruno/go-toyrobot/commit/37a724f532ead157ade0eea3bd698cd4a058111d). Copilot Agent mode successfully upgraded the code to Go 1.23.1 and passed all tests within 30 minutes.

Ensure your Go version is up-to-date by executing `go version`.

If you don't have Go installed, you can download it from https://golang.org/dl/

Clone the project and initialize the Go module dependencies.

## Steps to setup and run

```
git clone https://github.com/andrewbruno/go-toyrobot.git
cd go-toyrobot
make run
```

## Extra Features

CLI accepts "q" to quit program

## Testing via CLI

```
make test
```

## Testing via file input

```
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
        github.com/andrewbruno/go-toyrobot          coverage: 0.0% of statements
ok      github.com/andrewbruno/go-toyrobot/cmd      coverage: 100.0% of statements
ok      github.com/andrewbruno/go-toyrobot/compass  coverage: 100.0% of statements
ok      github.com/andrewbruno/go-toyrobot/table    coverage: 100.0% of statements
ok      github.com/andrewbruno/go-toyrobot/toy      coverage: 100.0% of statements
ok      github.com/andrewbruno/go-toyrobot/ui       coverage: 100.0% of statements
ok      github.com/andrewbruno/go-toyrobot/unit     coverage: 100.0% of statements
```

## Why Go?

I started to learn Go in July 2016 at Raven Fly, while working remotely on [ManyMe](https://manyme.com/), a greenfield online privacy platform.

Why I enjoyed using Go:

- **Built-in formatting** eliminates style debates.
- **Testing support out of the box** encourages TDD.
- **Less verbose** than Java — no need for access modifiers.
- **Statically typed** for compile-time safety.
- **Enforced use of variables** prevents bloat.
