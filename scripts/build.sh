#!/usr/bin/env bash

# Set variables
OUTPUT_DIR="gen"
OUTPUT_FILE="$OUTPUT_DIR/cf-api-server"
MAIN_FILE="./cmd/server/main.go" # Pfad zur main.go-Datei

# Create the output directory if it doesn't exist
mkdir -p $OUTPUT_DIR


# Check for build errors in the entire project
echo "Checking for build errors in the entire project..."
go build ./...
if [ $? -ne 0 ]; then
    echo "Build check failed. Please fix the errors and try again."
    exit 1
fi

# Build the Go server
echo "Building the server..."
go build -ldflags "-X 'main.SemanticVersion=1.0.0' -X 'main.CommitID=abcd1234'" -o $OUTPUT_FILE $MAIN_FILE

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo "Build successful. Executable created at $OUTPUT_FILE"
else
    echo "Build failed."
    exit 1
fi