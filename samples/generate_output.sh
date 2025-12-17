#!/bin/bash

# This script generates the output for the sample Markdown files.

set -e

# Navigate to the root of the repository
cd "$(dirname "$0")"/..

echo "Building mdsplit..."
go build ./cmd/mdsplit

OUTPUT_DIR="samples/output"

echo "Creating output directory: $OUTPUT_DIR"
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

for f in samples/*.md; do
  if [ "$(basename "$f")" = "README.md" ]; then
      continue
  fi
  FILENAME=$(basename "$f")
  DIRNAME="${FILENAME%.*}"
  echo "Splitting $f into $OUTPUT_DIR/$DIRNAME..."
  mkdir -p "$OUTPUT_DIR/$DIRNAME"
  ./mdsplit -in "$f" -out "$OUTPUT_DIR/$DIRNAME"
done

echo "Sample output generated in $OUTPUT_DIR"
