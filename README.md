# Toy Robot Simulator

A Go (Golang) implementation as per [PROBLEM](doc/PROBLEM.md) requirements.

The coding challenge was created by [Jon Eaves](https://www.linkedin.com/in/joneaves/) at [ANZ](https://joneaves.wordpress.com/2014/07/21/toy-robot-coding-test/).

Over the past decade, it has been a common interview take home test for many Australian companies.

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

Golang is a beautiful language (subjective, I agree). Let me explain why I like it:
  * code formatting is universal, part of SDK. No more religious discussions about styles.
  * testing is out of the box. No need for third party testing frameworks. Makes it ideal for TDD
  * is less verbose than say, Java. No need for private, public, etc.
  * is statically typed language, so many checks done at compile time.
  * all variables must be used, cannot compile if a variable is declared and not used.
