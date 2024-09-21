#!/bin/bash

readonly MOBILE_ANDROID_PATH=Code/Mobile_Android/
CURRENT_DIR=$(pwd)

# Check if the current directory ends with "Code/Mobile_Android"
if [[ "$CURRENT_DIR" == */Code/Mobile_Android ]]; then
    echo "You are in the correct directory: $CURRENT_DIR"
else
    echo "You are not in Code/Mobile_Android. Current directory is: $CURRENT_DIR"
	cd $MOBILE_ANDROID_PATH || (echo "The current directory is not the good for this script, exiting" && exit 1)
fi


echo "Running ktlint..."
./gradlew ktlintCheck || exit 1


echo "Running project build..."
./gradlew build || exit 1

echo "All checks passed!, ready to push !"
