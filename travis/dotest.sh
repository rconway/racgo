#!/usr/bin/env bash

# shell options to ensure that the script bails out with a non-zero exit code on failure
# Travis relies upon this to know the status of the step
#set -euov pipefail

echo "Running tests..."
go test
echo "...tests done"
