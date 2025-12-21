# Template Size Examples - Command Reference

## Quick Reference

### Template Presets

```bash
# Card (600×800px, 25 lines) - Mobile/Vertical
./mdsplit -in document.md -out ./output -template-size card

# Horizontal Card (800×600px, 18 lines) - Landscape/Compact
./mdsplit -in document.md -out ./output -template-size horizontal-card

# Presentation (1920×1080px, 40 lines) - Standard Slides
./mdsplit -in document.md -out ./output -template-size presentation

# A4 (794×1123px, 50 lines) - Printing/Documents
./mdsplit -in document.md -out ./output -template-size a4
```

## Real-World Examples

### Blog Post to Mobile Cards
```bash
# Convert a blog post to mobile-friendly vertical cards
./mdsplit -in blog-post.md -out ./mobile-cards -template-size card

# Result: Each card fits nicely on a mobile screen
# Example: 21-line article → 1 card, 50-line article → 2 cards
```

### Documentation to Presentation
```bash
# Convert documentation to presentation slides
./mdsplit -in API-docs.md -out ./presentation -template-size presentation

# Result: Standard 16:9 slides ready for projection
# Example: 100-line doc → 3 slides
```

### Article to Printable Pages
```bash
# Convert article to A4 pages for printing
./mdsplit -in article.md -out ./print -template-size a4 -dpi 300

# Result: High-quality printable pages
# Example: 100-line article → 2 pages
```

### Social Media Snippets
```bash
# Create compact landscape snippets for social media
./mdsplit -in tips.md -out ./snippets -template-size horizontal-card

# Result: Focused, bite-sized content
# Example: 50-line tips → 3 snippets
```

## Advanced Usage

### Large Font for Accessibility
```bash
# Presentation with 16pt font for better readability
./mdsplit -in slides.md -out ./accessible -template-size presentation -font-size 16

# Good for: Visually impaired users, large venues
```

### High-DPI Printing
```bash
# A4 pages at 300 DPI for professional printing
./mdsplit -in report.md -out ./print-ready -template-size a4 -dpi 300

# Good for: Professional documents, high-quality prints
```

### Custom Template Override
```bash
# Start with card template but increase height
./mdsplit -in content.md -out ./custom -template-size card -max-height 30

# Good for: Fine-tuning based on specific needs
```

### Retina Display Optimization
```bash
# Presentation optimized for retina displays
./mdsplit -in slides.md -out ./retina -template-size presentation -dpi 192

# Good for: MacBook Pro, high-DPI monitors
```

## Batch Processing

### Process Multiple Files with Same Template
```bash
# Convert all markdown files to cards
for file in *.md; do
    ./mdsplit -in "$file" -out "./cards/$(basename "$file" .md)" -template-size card
done
```

### Generate Multiple Formats
```bash
# Create both presentation and print versions
./mdsplit -in document.md -out ./presentation -template-size presentation
./mdsplit -in document.md -out ./print -template-size a4
```

## Integration with md2png

### Complete Workflow
```bash
# Step 1: Split markdown into slides
./mdsplit -in presentation.md -out ./slides -template-size presentation

# Step 2: Render each slide to PNG (using md2png)
for slide in slides/*.md; do
    md2png -in "$slide" -out "${slide%.md}.png" -width 1920 -height 1080
done
```

### Card Format Workflow
```bash
# Step 1: Create mobile cards
./mdsplit -in article.md -out ./cards -template-size card

# Step 2: Render to images
for card in cards/*.md; do
    md2png -in "$card" -out "${card%.md}.png" -width 600 -height 800
done
```

## Comparison Examples

### Same Content, Different Templates
```bash
# Generate all variations for comparison
./mdsplit -in sample.md -out ./sample-default
./mdsplit -in sample.md -out ./sample-card -template-size card
./mdsplit -in sample.md -out ./sample-hcard -template-size horizontal-card
./mdsplit -in sample.md -out ./sample-pres -template-size presentation
./mdsplit -in sample.md -out ./sample-a4 -template-size a4

# Compare slide counts
ls sample-*/
```

### Actual Results from Samples
```
basic.md (21 lines):
  - default/presentation/a4/card: 1 slide
  - horizontal-card: 2 slides

very_long_file.md (200+ lines):
  - a4: 5 slides (most content per slide)
  - default/presentation: 6 slides
  - card: 9 slides
  - horizontal-card: 13 slides (most slides, least content per slide)
```

## Common Patterns

### Conference Presentation
```bash
# Standard presentation format
./mdsplit -in talk.md -out ./slides -template-size presentation -font-size 14
```

### Technical Documentation
```bash
# A4 format for PDF generation
./mdsplit -in manual.md -out ./pages -template-size a4
```

### Mobile App Content
```bash
# Vertical cards for in-app display
./mdsplit -in content.md -out ./app-cards -template-size card -font-size 14
```

### Dashboard Widgets
```bash
# Compact horizontal cards
./mdsplit -in metrics.md -out ./widgets -template-size horizontal-card
```

### Email Newsletter
```bash
# Mobile-friendly vertical format
./mdsplit -in newsletter.md -out ./sections -template-size card
```

## Tips

1. **Start with a template**: Choose the closest preset to your needs
2. **Override when needed**: Use `-max-height` or `-max-width` to fine-tune
3. **Check samples**: Look at `samples/output/` for real examples
4. **Font size matters**: Larger fonts need fewer lines per slide
5. **DPI for rendering**: 96 for screen, 300 for print
6. **Horizontal card is aggressive**: Creates the most slides, use for very focused content
7. **A4 is efficient**: Creates the fewest slides, use when you want maximum content

## Troubleshooting

### Too Many Slides
```bash
# Use a larger template
./mdsplit -in file.md -out ./output -template-size a4  # Instead of card
```

### Too Few Slides
```bash
# Use a smaller template
./mdsplit -in file.md -out ./output -template-size horizontal-card  # Instead of a4
```

### Content Doesn't Fit Well
```bash
# Fine-tune with custom height
./mdsplit -in file.md -out ./output -template-size presentation -max-height 35
```

## Help

```bash
# Show all available options
./mdsplit -help
```
