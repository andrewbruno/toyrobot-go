# Toy Robot Simulator

A Go (Golang) implementation as per [PROBLEM](doc/PROBLEM.md) requirements.

## Prerequisites

The code has been tested with Go 1.8.3 and Go 1.11.1.  Ensure your Go version is 1.8+ by executing `go version`

If you don't have Go installed, you can download it from https://golang.org/dl/

Clone the project, and setup the GOPATH variable to the root of the project.

## Steps to setup and run

```
git clone https://bitbucket.org/andrewbruno/go-toyrobot.git
cd go-toyrobot
export GOPATH=`pwd`
cd $GOPATH/src/toyrobot
go run main.go
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
