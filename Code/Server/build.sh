#!/bin/bash

# -e Exit immediately when a command returns a non-zero status.
# -x Print commands before they are executed
set -ex

go run github.com/steebchen/prisma-client-go generate
swag fmt
swag init --parseDependency true
go build
