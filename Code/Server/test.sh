#!/bin/bash

# -e Exit immediately when a command returns a non-zero status.
# -x Print commands before they are executed
set -e

OUTPUT=$(./run_tests.sh no-interactive)
echo "# Coverage report"
echo "$OUTPUT"
echo "COVERAGE=$(echo "$OUTPUT" | grep total | grep -o '[0-9]\+\(\.[0-9]\+\)\?' | cut -d. -f1)"
