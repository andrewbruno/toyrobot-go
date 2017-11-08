# Toy Robot Simulator

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

