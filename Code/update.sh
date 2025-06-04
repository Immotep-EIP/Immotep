#!/bin/bash

# -e Exit immediately when a command returns a non-zero status.
set -e

git switch backend
git pull
git switch -
git merge backend

cd Server/
make db_update
cd ..
