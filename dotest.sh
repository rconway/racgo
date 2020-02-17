#!/usr/bin/env bash

set -euov pipefail

echo "Running tests..."
go test
echo "...tests done"
