#!/usr/bin/env bash

echo "Running unit and integration tests..."
go test ./... $@
if [ $? -ne 0 ]; then
    echo "Unit tests failed."
    exit 1
fi
echo "Unit tests passed."

