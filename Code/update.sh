#!/bin/bash

# -e Exit immediately when a command returns a non-zero status.
set -e

git pull

cd Server/
go run github.com/steebchen/prisma-client-go migrate deploy
go run github.com/steebchen/prisma-client-go generate
cd ..
