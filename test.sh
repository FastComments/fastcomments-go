#!/bin/bash

set -e

echo "Running FastComments Go SDK tests..."

# Check required environment variables
if [ -z "$FASTCOMMENTS_API_KEY" ] || [ -z "$FASTCOMMENTS_TENANT_ID" ]; then
  echo "Error: FASTCOMMENTS_API_KEY and FASTCOMMENTS_TENANT_ID environment variables must be set"
  exit 1
fi

# Run all tests
echo "Running unit tests..."
go test ./tests/sso_test.go -v

echo ""
echo "Running integration tests..."
go test ./tests/sso_integration_test.go -v

echo ""
echo "✓ All tests passed!"
