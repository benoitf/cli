#!/bin/sh
env GOOS=windows GOARCH=amd64 go build -o bin/che.exe src/github.com/benoitf/cli/che.go

