#!/bin/bash

# -e Exit immediately when a command returns a non-zero status.
set -e

./run_linter.sh

list=$(go list ./... | grep -ve prisma -ve docs -ve database)

# Run tests and generate coverage report for all packages except prisma
if [[ $1 == "debug" ]]; then
    go test -v $list -coverprofile cover.out
elif [[ $1 == "no-interactive" ]]; then
    go test $list -coverprofile cover.out
    go tool cover -func=cover.out
else
    go test $list -coverprofile cover.out
    go tool cover -func=cover.out
    go tool cover -html=cover.out
fi
