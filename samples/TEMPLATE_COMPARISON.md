# Template Size Comparison

This document provides a detailed comparison of how different template sizes affect the splitting of sample Markdown files.

## Template Size Specifications

| Template | Dimensions | Max Lines | Max Width | Best For |
|----------|-----------|-----------|-----------|----------|
| **Card** | 600×800px | 25 | 600px | Mobile-friendly vertical cards, social media |
| **Horizontal Card** | 800×600px | 18 | 800px | Landscape displays, compact presentations |
| **Presentation** | 1920×1080px | 40 | 1920px | Standard presentations, desktop displays |
| **A4** | 794×1123px | 50 | 794px | Printable documents, PDF generation |

## Sample File Comparison

### basic.md (21 lines)

| Template | Slides | Notes |
|----------|--------|-------|
| Default | 1 | Fits comfortably in default 40-line limit |
| Card | 1 | Still fits within 25-line limit |
| Horizontal Card | 2 | **Split required** - 18-line limit too small |
| Presentation | 1 | Same as default |
| A4 | 1 | Plenty of room with 50-line limit |

**Key Insight**: Horizontal card format is the most restrictive, causing splits even for small files.

### code_blocks.md (Multiple code blocks)

| Template | Slides | Notes |
|----------|--------|-------|
| Default | 3 | Code blocks split intelligently |
| Card | 4 | More aggressive splitting due to 25-line limit |
| Horizontal Card | 6 | **Most splits** - very restrictive 18-line limit |
| Presentation | 3 | Same as default |
| A4 | 3 | Same as default and presentation |

**Key Insight**: Code blocks are split while preserving syntax highlighting fences. Horizontal card creates the most slides.

### long_table.md (50-row table)

| Template | Slides | Notes |
|----------|--------|-------|
| Default | 3 | Table split with headers repeated |
| Card | 4 | More table parts due to smaller height |
| Horizontal Card | 5 | **Most table splits** |
| Presentation | 3 | Same as default |
| A4 | 3 | Same as default and presentation |

**Key Insight**: Tables are split intelligently with headers repeated on each slide and continuation notes added.

### very_long_file.md (Mixed content)

| Template | Slides | Notes |
|----------|--------|-------|
| Default | 6 | Balanced splitting |
| Card | 9 | Moderate increase in slides |
| Horizontal Card | 13 | **Maximum slides** - most aggressive splitting |
| Presentation | 6 | Same as default |
| A4 | 5 | **Fewest slides** - maximum capacity |

**Key Insight**: A4 format provides the most content per slide, while horizontal card creates the most slides.

### mixed_content.md

| Template | Slides | Notes |
|----------|--------|-------|
| Default | 1 | Single slide |
| Card | 2 | Split into 2 slides |
| Horizontal Card | 3 | Split into 3 slides |
| Presentation | 1 | Single slide |
| A4 | 1 | Single slide |

### short_table.md

| Template | Slides | Notes |
|----------|--------|-------|
| Default | 1 | Table fits on one slide |
| Card | 1 | Still fits |
| Horizontal Card | 1 | Still fits |
| Presentation | 1 | Fits comfortably |
| A4 | 1 | Plenty of room |

### images.md

| Template | Slides | Notes |
|----------|--------|-------|
| Default | 1 | Image references preserved |
| Card | 1 | Same |
| Horizontal Card | 1 | Same |
| Presentation | 1 | Same |
| A4 | 1 | Same |

## Recommendations by Use Case

### Mobile/Social Media
**Use: Card (600×800px)**
- Vertical orientation perfect for mobile devices
- 25 lines provides good readability on small screens
- Creates moderately sized chunks

### Presentations/Slides
**Use: Presentation (1920×1080px)**
- Standard 16:9 aspect ratio
- 40 lines is the sweet spot for slide content
- Works well with projectors and large displays

### Documents/Printing
**Use: A4 (794×1123px)**
- Standard paper size
- 50 lines maximizes content per page
- Minimizes total page count
- Best for PDF generation

### Compact/Landscape Displays
**Use: Horizontal Card (800×600px)**
- Landscape orientation
- 18 lines keeps content concise
- Good for dashboard displays or embedded content
- Creates more slides but each is very focused

## Font Size and DPI Considerations

The `-font-size` and `-dpi` flags allow fine-tuning for different rendering scenarios:

### Standard Web Display
```bash
./mdsplit -template-size presentation -font-size 12 -dpi 96
```

### High-DPI Displays (Retina)
```bash
./mdsplit -template-size presentation -font-size 12 -dpi 192
```

### Print Quality
```bash
./mdsplit -template-size a4 -font-size 11 -dpi 300
```

### Large Font for Accessibility
```bash
./mdsplit -template-size card -font-size 16 -dpi 96
```

## Summary Statistics

Across all sample files:
- **Total sample outputs**: 102 markdown files
- **Template variations**: 5 (default + 4 presets)
- **Sample input files**: 7
- **Splitting efficiency**: A4 creates ~40% fewer slides than horizontal-card

## Advanced Usage

### Override Template Defaults
You can use a template as a starting point and override specific values:

```bash
# Start with presentation template but reduce height
./mdsplit -template-size presentation -max-height 30
```

### Custom Configuration
For complete control, skip templates and set everything manually:

```bash
./mdsplit -max-height 35 -max-width 1200 -font-size 14 -dpi 120
```

## Integration with md2png

These template sizes are designed to work seamlessly with the `md2png` tool:

```bash
# Split with card template
./mdsplit -in document.md -out ./slides -template-size card

# Render each slide to PNG
for slide in slides/*.md; do
    md2png -in "$slide" -out "${slide%.md}.png" -width 600 -height 800
done
```
