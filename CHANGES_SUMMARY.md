# Summary of Changes - Template Size Feature

## Overview

Successfully implemented template size presets and font/DPI configuration for the `mdsplit` tool, along with comprehensive sample outputs demonstrating all features.

## Changes Made

### 1. Core Library Updates (`mdsplit.go`)

#### Added Template Size Type and Constants
- New `TemplateSize` type for predefined templates
- Four template presets:
  - **Card**: 600×800px, ~25 lines (mobile/vertical)
  - **Horizontal Card**: 800×600px, ~18 lines (landscape/compact)
  - **Presentation**: 1920×1080px, ~40 lines (standard slides)
  - **A4**: 794×1123px, ~50 lines (printing/documents)

#### Enhanced SplitOptions Struct
- Added `TemplateSize` field for preset selection
- Added `FontSize` field (default: 12pt)
- Added `DPI` field (default: 96)

#### Updated Split Function
- Template size presets automatically set `MaxHeight` and `MaxWidth`
- Default values for `FontSize` (12) and `DPI` (96)
- Template sizes override manual height/width unless explicitly specified

### 2. CLI Updates (`cmd/mdsplit/main.go`)

#### New Command-Line Flags
- `-template-size`: Select preset (card, horizontal-card, presentation, a4)
- `-font-size`: Font size in points (default: 12)
- `-dpi`: DPI for rendering (default: 96)

#### Updated Existing Flags
- `-max-height`: Now defaults to 0, overridden by template size
- `-max-width`: Now defaults to 0, overridden by template size

### 3. Documentation Updates

#### README.md
- Updated flags table with new options
- Added "Template Size Presets" section with descriptions
- Added multiple examples demonstrating different use cases:
  - Custom height
  - Presentation template
  - Card template with larger font
  - A4 printing at high DPI

#### New Documentation Files
- **QUICKSTART.md**: Quick start guide for template sizes
- **samples/TEMPLATE_COMPARISON.md**: Detailed comparison of all templates
- **samples/output/README.md**: Documentation of sample outputs

### 4. Sample Generation

#### Updated `samples/generate_output.sh`
- Generates outputs for all template sizes
- Creates 36 output directories total:
  - 7 samples × 5 configurations (default + 4 templates)
  - Plus 1 custom font size example

#### Sample Output Statistics
- **Total files generated**: 102 markdown files
- **Template variations**: 5 (default + 4 presets)
- **Demonstrates**:
  - Table splitting with continuation notes
  - Code block splitting with fence preservation
  - Paragraph splitting
  - Different content densities per template

### 5. Test Updates (`mdsplit_test.go`)

#### Fixed README Split Test
- Updated expected file count from 3 to 4 (due to added documentation)
- Changed content check from `slide-3.md` to `slide-4.md`
- All tests passing ✓

## Sample Output Highlights

### Basic Sample Comparison
- **Default/Presentation/A4/Card**: 1 slide
- **Horizontal Card**: 2 slides (demonstrates splitting)

### Code Blocks Sample
- **Default/Presentation/A4**: 3 slides
- **Card**: 4 slides
- **Horizontal Card**: 6 slides (most aggressive)

### Long Table Sample
- **Default/Presentation/A4**: 3 slides with table continuation
- **Card**: 4 slides
- **Horizontal Card**: 5 slides

### Very Long File Sample
- **A4**: 5 slides (fewest - maximum capacity)
- **Default/Presentation**: 6 slides
- **Card**: 9 slides
- **Horizontal Card**: 13 slides (most - minimum capacity)

## Key Features Demonstrated

### 1. Intelligent Table Splitting
- Headers repeated on each slide
- Continuation notes added (e.g., "_Table continued (part 2)_")
- Proper row alignment maintained

### 2. Code Block Splitting
- Syntax highlighting fences preserved
- Language specifiers maintained
- Split across multiple slides when needed

### 3. Template Flexibility
- Presets for common use cases
- Override capability for custom needs
- Backward compatible with manual settings

## Use Case Recommendations

| Use Case | Template | Reason |
|----------|----------|--------|
| Mobile viewing | Card | Vertical orientation, 25 lines |
| Social media | Card | Compact, mobile-friendly |
| Presentations | Presentation | Standard 16:9, 40 lines |
| Desktop displays | Presentation | Large width, balanced height |
| Printing | A4 | Standard paper size, 50 lines |
| PDF generation | A4 | Maximum content per page |
| Dashboards | Horizontal Card | Landscape, very focused |
| Embedded content | Horizontal Card | Compact, 18 lines |

## Files Modified

1. `mdsplit.go` - Core library with template support
2. `cmd/mdsplit/main.go` - CLI with new flags
3. `mdsplit_test.go` - Updated tests
4. `README.md` - Updated documentation
5. `samples/generate_output.sh` - Enhanced sample generation

## Files Created

1. `QUICKSTART.md` - Quick start guide
2. `samples/TEMPLATE_COMPARISON.md` - Detailed comparison
3. `samples/output/README.md` - Sample documentation
4. 102 sample output files in `samples/output/`

## Testing

All tests passing:
```
=== RUN   TestSplit
=== RUN   TestSplit/simple_split
=== RUN   TestSplit/long_table
=== RUN   TestSplit/long_codeblock
=== RUN   TestSplit/comprehensive_split
=== RUN   TestSplit/size-based_split
=== RUN   TestSplit/readme_split
--- PASS: TestSplit (0.00s)
    --- PASS: TestSplit/simple_split (0.00s)
    --- PASS: TestSplit/long_table (0.00s)
    --- PASS: TestSplit/long_codeblock (0.00s)
    --- PASS: TestSplit/comprehensive_split (0.00s)
    --- PASS: TestSplit/size-based_split (0.00s)
    --- PASS: TestSplit/readme_split (0.00s)
PASS
```

## Backward Compatibility

✓ Existing usage with `-max-height` and `-max-width` still works
✓ Default behavior unchanged when no template specified
✓ All existing tests pass
✓ Library API extended, not changed

## Next Steps

Users can now:
1. Choose from 4 convenient template presets
2. Customize font size and DPI for rendering
3. Override template defaults when needed
4. See comprehensive examples in `samples/output/`
5. Reference documentation for use case guidance

## Integration Ready

The template sizes are designed to work seamlessly with `md2png`:
```bash
./mdsplit -in doc.md -out ./slides -template-size card
# Then render with md2png using matching dimensions
```
