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

# Generate outputs with default settings
for f in samples/*.md; do
  if [ "$(basename "$f")" = "README.md" ]; then
      continue
  fi
  FILENAME=$(basename "$f")
  DIRNAME="${FILENAME%.*}"
  echo "Splitting $f into $OUTPUT_DIR/$DIRNAME (default)..."
  mkdir -p "$OUTPUT_DIR/$DIRNAME"
  ./mdsplit -in "$f" -out "$OUTPUT_DIR/$DIRNAME"
done

# Generate outputs with different template sizes
TEMPLATE_SIZES=("card" "horizontal-card" "presentation" "a4")

for template in "${TEMPLATE_SIZES[@]}"; do
  echo ""
  echo "Generating samples with template size: $template"
  
  for f in samples/*.md; do
    if [ "$(basename "$f")" = "README.md" ]; then
        continue
    fi
    FILENAME=$(basename "$f")
    DIRNAME="${FILENAME%.*}_${template}"
    echo "  Splitting $f into $OUTPUT_DIR/$DIRNAME..."
    mkdir -p "$OUTPUT_DIR/$DIRNAME"
    ./mdsplit -in "$f" -out "$OUTPUT_DIR/$DIRNAME" -template-size "$template"
  done
done

# Generate a sample with custom font size
echo ""
echo "Generating sample with custom font size (16pt)..."
mkdir -p "$OUTPUT_DIR/basic_font16"
./mdsplit -in "samples/basic.md" -out "$OUTPUT_DIR/basic_font16" -font-size 16 -template-size "presentation"

echo ""
echo "Sample output generated in $OUTPUT_DIR"
