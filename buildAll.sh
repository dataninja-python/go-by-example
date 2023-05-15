#!/bin/bash

# Recursively go through all directories from the current location
for file in $(find . -name "*.go" -type f); do
    # Check if the file is a regular file
    if [ -f "$file" ]; then
        # Build the file with the go command
        go build "$file"
    fi
done