#!/bin/bash

## -e Exit immediately when a command returns a non-zero status.
set -e

## These were used when working with one branch for each app. Not needed anymore because we create separate branches for features from the dev branch.
# git switch backend
# git pull
# git switch -
# git merge backend

cd Server/
make db_update
cd ..
