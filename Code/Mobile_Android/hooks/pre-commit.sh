#!/bin/bash

echo "Running ktlint..."
./gradlew ktlintCheck || exit 1


echo "Running project build..."
./gradlew build || exit 1

echo "All checks passed!"
