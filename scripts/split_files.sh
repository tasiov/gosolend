#!/bin/bash

# Check if input file is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <input_file>"
    exit 1
fi

input_file="$1"
current_file=""
mkdir -p temp_output

# Read the input file line by line
while IFS= read -r line; do
    # Check if line starts with "// File: "
    if [[ $line =~ ^//\ File:\ (.+)$ ]]; then
        # Extract the file path
        file_path="${BASH_REMATCH[1]}"
        # Create the directory structure
        mkdir -p "$(dirname "$file_path")"
        # Set the current file we're writing to
        current_file="$file_path"
        echo "Creating file: $current_file"
    elif [ ! -z "$current_file" ]; then
        # Write the line to the current file
        echo "$line" >> "$current_file"
    fi
done < "$input_file"

echo "Done! Files have been split successfully."