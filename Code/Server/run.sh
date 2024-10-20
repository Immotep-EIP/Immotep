#!/bin/bash

golangci-lint run

if [ $? -ne 0 ]; then
    echo "Linter failed"
    exit 1
fi

swag fmt
swag init --parseDependency true
go build && echo "Build successful" && ./backend
