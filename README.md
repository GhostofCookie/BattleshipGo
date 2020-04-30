[![Build Status](https://travis-ci.com/GhostofCookie/BattleshipGo.svg?branch=master)](https://travis-ci.com/GhostofCookie/BattleshipGo)
[![GhostofCookie](https://circleci.com/gh/GhostofCookie/BattleshipGo.svg?style=svg)](https://circleci.com/gh/GhostofCookie/BattleshipGo)
[![codecov](https://codecov.io/gh/GhostofCookie/BattleshipGo/branch/master/graph/badge.svg)](https://codecov.io/gh/GhostofCookie/BattleshipGo)

With the coverage posted above, it is worth noting that the remaining code that is not covered is only the code contained within the 'main' function in [Battleship.go](https://github.com/GhostofCookie/BattleshipGo/blob/master/Battleship.go#L12).

# Requirements

All this project requires is Golang(Go), and you can get it [here](https://golang.org/dl/).

# Building and Running the Game

## Build
Simply type
```bash
go build
```
into your teminal.

## Run
### Windows
```
./Battleship.exe
```
### MacOS, Linux
```
./Battleship
```

# Testing
To run the tests for the project, simply enter the following into the bash
```
go test
```
For tests with coverage, run the following
```
go test -coverprofile coverage.info
```
