#!/bin/bash

# Run tests and generate coverage report for all packages except prisma
go test `go list ./... | grep -v prisma` -coverprofile cover.out

# Generate coverage htlm report
go tool cover -html=cover.out $@
