#!/bin/bash

# -e Exit immediately when a command returns a non-zero status.
set -e

git switch backend
git pull
git switch -
git merge backend

cd Server/
go run github.com/steebchen/prisma-client-go migrate deploy
go run github.com/steebchen/prisma-client-go generate
cd ..
