#!/bin/bash

# Script to deduplicate identical consecutive lines in CHANGELOG.md
# Usage: ./deduplicate-changelog.sh [input_file] [output_file]

INPUT_FILE=${1:-CHANGELOG.md}
OUTPUT_FILE=${2:-CHANGELOG.md.tmp}

# Check if input file exists
if [ ! -f "$INPUT_FILE" ]; then
    echo "Error: Input file $INPUT_FILE does not exist"
    exit 1
fi

# Read the changelog file and process it
cat "$INPUT_FILE" | awk '
{
    # Remove trailing whitespace and newlines
    gsub(/[[:space:]]+$/, "", $0)
    
    if (NR > 1) {
        # Compare current line with previous line
        if (prev_line != $0) {
            print prev_line
        }
    }
    prev_line = $0
}
END {
    # Print the last line
    print prev_line
}' > "$OUTPUT_FILE"

# Replace the original file with the deduplicated version
mv "$OUTPUT_FILE" "$INPUT_FILE"

echo "Changelog deduplication completed successfully" 