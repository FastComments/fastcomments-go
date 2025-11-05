#!/bin/bash

# Remove previously generated code
rm -rf ./client

# Try to download the OpenAPI spec from the running FastComments server
if wget -q http://localhost:3001/js/swagger.json -O ./openapi.json; then
    echo "Downloaded OpenAPI spec from running server"
    SPEC_FILE="./openapi.json"
else
    echo "Server not running, using local OpenAPI spec"
    SPEC_FILE="./openapi.yaml"
fi

# Generate the Go client using openapi-generator
npx @openapitools/openapi-generator-cli generate \
    -i "$SPEC_FILE" \
    -g go \
    -o ./client \
    -c config.json

echo "Generated Go client in ./client"

# Initialize Go modules
go mod tidy

# Fix naming conflict in generated code (field HasChildren vs method HasChildren)
# Rename the HasChildren() method that checks Children array to ChildrenIsSet()
if [ -f "./client/model_public_comment.go" ]; then
    sed -i 's/func (o \*PublicComment) HasChildren() bool {/func (o *PublicComment) ChildrenIsSet() bool {/' ./client/model_public_comment.go
    sed -i 's/\/\/ HasChildren returns a boolean if a field has been set\.$/\/\/ ChildrenIsSet returns a boolean if a field has been set./' ./client/model_public_comment.go
    echo "Applied naming conflict fix to model_public_comment.go"
fi

echo "✓ Go SDK updated successfully!"
