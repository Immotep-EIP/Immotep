#!/bin/bash

# -e Exit immediately when a command returns a non-zero status.
set -e

./run_linter.sh

echo "Running tests..." >&2
# Run tests and generate coverage report for all packages except prisma
if [[ $1 == "debug" ]]; then
    go test -v `go list ./... | grep -v prisma` -coverprofile cover.out
elif [[ $1 == "no-interactive" ]]; then
    go test `go list ./... | grep -v prisma` -coverprofile cover.out
    go tool cover -func=cover.out
else
    go test `go list ./... | grep -v prisma` -coverprofile cover.out
    go tool cover -func=cover.out
    go tool cover -html=cover.out
fi
