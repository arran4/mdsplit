# Sample Outputs

This directory contains sample outputs demonstrating the `mdsplit` tool with various template sizes and configurations.

## Directory Structure

The outputs are organized by input file and template configuration:

- **`<filename>/`** - Default settings (40 lines max height)
- **`<filename>_card/`** - Card template (25 lines, 600px width)
- **`<filename>_horizontal-card/`** - Horizontal card template (18 lines, 800px width)
- **`<filename>_presentation/`** - Presentation template (40 lines, 1920px width)
- **`<filename>_a4/`** - A4 template (50 lines, 794px width)
- **`basic_font16/`** - Custom example with 16pt font size

## Template Size Comparison

### Card (600×800px, ~25 lines)
Best for: Mobile-friendly vertical cards, social media posts, portrait displays

Example: `basic_card/` splits the basic sample into 1 slide (fits within 25 lines)

### Horizontal Card (800×600px, ~18 lines)
Best for: Landscape cards, horizontal mobile displays, compact presentations

Example: `basic_horizontal-card/` splits the basic sample into 2 slides (18 line limit requires split)

### Presentation (1920×1080px, ~40 lines)
Best for: Standard presentations, desktop displays, default use case

Example: `basic_presentation/` splits the basic sample into 1 slide (fits within 40 lines)

### A4 (794×1123px, ~50 lines)
Best for: Printable documents, PDF generation, maximum content per page

Example: `basic_a4/` splits the basic sample into 1 slide (fits within 50 lines)

## Sample Files

### basic.md
A simple Markdown file with headings, paragraphs, and lists.
- **Default**: 1 slide
- **Card**: 1 slide
- **Horizontal Card**: 2 slides (demonstrates splitting due to smaller height)
- **Presentation**: 1 slide
- **A4**: 1 slide

### code_blocks.md
Contains multiple code blocks with different languages.
- **Default**: 3 slides
- **Card**: 4 slides (code blocks split more aggressively)
- **Horizontal Card**: 6 slides (smallest height causes most splits)
- **Presentation**: 3 slides
- **A4**: 3 slides

### long_table.md
A long table that exceeds typical slide heights.
- **Default**: 3 slides (table split with continuation notes)
- **Card**: 4 slides
- **Horizontal Card**: 5 slides
- **Presentation**: 3 slides
- **A4**: 3 slides

### very_long_file.md
A comprehensive test with mixed content types.
- **Default**: 6 slides
- **Card**: 9 slides
- **Horizontal Card**: 13 slides (most aggressive splitting)
- **Presentation**: 6 slides
- **A4**: 5 slides (largest capacity)

## Font Size and DPI

The `basic_font16/` directory demonstrates using a custom font size (16pt) with the presentation template. While the current implementation uses font size and DPI for future rendering features, they're stored in the configuration for tools like `md2png` to use.

## Regenerating Samples

To regenerate all sample outputs:

```bash
./samples/generate_output.sh
```

This script will:
1. Build the `mdsplit` binary
2. Generate outputs for all sample files with default settings
3. Generate outputs for all template sizes (card, horizontal-card, presentation, a4)
4. Generate a custom example with 16pt font size

## Notes

- The line counts are approximate and based on typical rendering at 96 DPI
- Actual visual height may vary depending on the rendering engine
- Template sizes are designed to work well with the `md2png` tool
- Long tables and code blocks are intelligently split across multiple slides
- Each split table includes the header and a continuation note
