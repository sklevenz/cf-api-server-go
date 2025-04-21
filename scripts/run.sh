#!/usr/bin/env bash


# Define the path to the main.go file
MAIN_FILE="./cmd/server/main.go"

# Run the Go server
echo "Starting the server..."
go run "$MAIN_FILE" $@

# Check the exit code and handle normal termination
EXIT_CODE=$?
if [ $EXIT_CODE -eq 0 ] || [ $EXIT_CODE -eq 130 ] || [ $EXIT_CODE -eq 1 ]; then
    # Exit code 130 indicates termination via Ctrl+C (SIGINT)
    echo "Server stopped gracefully."
else
    echo "Server encountered an error (exit code: $EXIT_CODE)."
    exit $EXIT_CODE
fi
