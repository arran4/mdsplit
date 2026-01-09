package mdsplit

import (
	"fmt"
	"io"
	"os"
)

// Run is a subcommand `mdsplit`
//
// Flags:
//   in:           --in            (default: "")           Markdown input file, or stdin when empty
//   out:          --out           (default: ".")          Output directory for the split files
//   maxHeight:    --max-height    (default: 0)            Maximum height of a slide in lines. Overridden by template selection.
//   maxWidth:     --max-width     (default: 0)            Maximum width of a slide in pixels. Overridden by template selection.
//   theme:        --theme         (default: "light")      light or dark
//   templateSize: --template-size (default: "")           Predefined template size.
//   fontSize:     --font-size     (default: 12)           Font size in points
//   dpi:          --dpi           (default: 96)           DPI for rendering
//
// Valid template sizes are: card, horizontal-card, presentation, a4.
func Run(in string, out string, maxHeight int, maxWidth int, theme string, templateSize string, fontSize int, dpi int) error {
	// Read the input from the specified file or stdin.
	var data []byte
	var err error
	if in == "" {
		data, err = io.ReadAll(os.Stdin)
	} else {
		data, err = os.ReadFile(in)
	}
	if err != nil {
		return fmt.Errorf("error reading input: %v", err)
	}

	// Create the SplitOptions struct.
	opts := SplitOptions{
		OutDir:       out,
		MaxHeight:    maxHeight,
		MaxWidth:     maxWidth,
		Theme:        theme,
		TemplateSize: TemplateSize(templateSize),
		FontSize:     fontSize,
		DPI:          dpi,
	}

	// Split the Markdown file.
	if err := Split(data, opts); err != nil {
		return fmt.Errorf("error splitting Markdown: %v", err)
	}
	return nil
}
