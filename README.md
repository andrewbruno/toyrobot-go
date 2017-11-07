# Toy Robot Simulation

A Go (Golang) implementation as per [PROBLEM](doc/PROBLEM.md) requirements.

## Setup

1. Install Go 1.8+

2. Ensure GOPATH is set to root of the project

e.g. the following command should work

```
cd $GOPATH/src/toyrobot
```

## Extra Features

CLI accepts "q" to quit program

## Testing via CLI

```
cd $GOPATH/src/toyrobot
go run main.go 
```

## Testing via file input

```
cd $GOPATH/src/toyrobot
go run main.go < cases.txt -- TODO ---
```

## Testing code coverage via GO

```
go test -cover ./...
```
