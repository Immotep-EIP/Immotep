#!/bin/bash

# Run tests and generate coverage report for all packages except prisma
go test `go list ./... | grep -v prisma` -coverprofile cover.out

if [ $? -ne 0 ]; then
    exit 1
fi

# Generate coverage htlm report
go tool cover -func=cover.out
go tool cover -html=cover.out $@
