#!/bin/bash

# -e Exit immediately when a command returns a non-zero status.
# -x Print commands before they are executed
set -ex

go run github.com/steebchen/prisma-client-go migrate reset
go run prisma/seed.go
