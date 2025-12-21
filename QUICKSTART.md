# Quick Start Guide - Template Sizes

## TL;DR - Choose Your Template

```bash
# Mobile/vertical cards
./mdsplit -in document.md -out ./slides -template-size card

# Landscape/compact
./mdsplit -in document.md -out ./slides -template-size horizontal-card

# Presentations (default)
./mdsplit -in document.md -out ./slides -template-size presentation

# Documents/printing
./mdsplit -in document.md -out ./slides -template-size a4
```

## What Changed?

### New Flags

- **`-template-size`**: Choose from `card`, `horizontal-card`, `presentation`, or `a4`
- **`-font-size`**: Font size in points (default: 12)
- **`-dpi`**: DPI for rendering (default: 96)

### Template Presets

| Template | Lines | Width | Use Case |
|----------|-------|-------|----------|
| `card` | 25 | 600px | Mobile, social media |
| `horizontal-card` | 18 | 800px | Landscape displays |
| `presentation` | 40 | 1920px | Standard slides |
| `a4` | 50 | 794px | Printing, PDFs |

## Examples

### Basic Usage
```bash
# Use presentation template (same as default)
./mdsplit -in README.md -out ./slides -template-size presentation
```

### Mobile-Friendly Cards
```bash
# Create vertical cards for mobile viewing
./mdsplit -in article.md -out ./cards -template-size card
```

### High-Quality Print
```bash
# A4 format with high DPI for printing
./mdsplit -in document.md -out ./pages -template-size a4 -dpi 300
```

### Large Font for Accessibility
```bash
# Presentation with larger font
./mdsplit -in slides.md -out ./output -template-size presentation -font-size 16
```

### Custom Override
```bash
# Start with card template but adjust height
./mdsplit -in content.md -out ./output -template-size card -max-height 30
```

## Sample Outputs

The `samples/output/` directory now contains examples for all template sizes:

- `basic_card/` - Card template example
- `basic_horizontal-card/` - Horizontal card example  
- `basic_presentation/` - Presentation template example
- `basic_a4/` - A4 template example
- `basic_font16/` - Custom font size example

### Comparing Outputs

Check how the same content splits differently:

```bash
# Compare slide counts
ls samples/output/code_blocks/        # 3 slides (default)
ls samples/output/code_blocks_card/   # 4 slides (more splits)
ls samples/output/code_blocks_horizontal-card/  # 6 slides (most splits)
ls samples/output/code_blocks_a4/     # 3 slides (same as default)
```

## Key Features

### Intelligent Table Splitting
Long tables are split across multiple slides with:
- Headers repeated on each slide
- Continuation notes (e.g., "_Table continued (part 2)_")
- Proper row alignment maintained

### Code Block Splitting
Long code blocks are split while:
- Preserving syntax highlighting fences
- Maintaining language specifiers
- Keeping code readable

### Paragraph Handling
Long paragraphs are split intelligently to fit within height limits.

## Migration from Old Flags

### Before
```bash
./mdsplit -in file.md -out ./slides -max-height 40 -max-width 1024
```

### After (equivalent)
```bash
./mdsplit -in file.md -out ./slides -template-size presentation
```

### Custom Settings Still Work
```bash
# You can still use manual settings if you prefer
./mdsplit -in file.md -out ./slides -max-height 35 -max-width 1200
```

## Regenerating Samples

To see all template sizes in action:

```bash
cd samples
./generate_output.sh
```

This generates:
- Default outputs for all samples
- Template-specific outputs (card, horizontal-card, presentation, a4)
- Custom font size example

## Next Steps

1. **Try different templates** with your content to see which works best
2. **Check the samples** in `samples/output/` for real examples
3. **Read the comparison** in `samples/TEMPLATE_COMPARISON.md` for detailed analysis
4. **Integrate with md2png** to render slides as images

## Tips

- **Start with a template**: Use `-template-size` for quick setup
- **Horizontal card creates most slides**: Use for very focused content
- **A4 creates fewest slides**: Use when you want maximum content per page
- **Font size affects readability**: Increase for accessibility or large displays
- **DPI matters for rendering**: Use 300 for print, 96 for web

## Getting Help

```bash
./mdsplit -help
```

Shows all available flags and their defaults.
