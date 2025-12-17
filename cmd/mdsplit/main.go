package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/arran4/mdsplit"
)

func main() {
	// Define the command-line flags.
	in := flag.String("in", "", "Markdown input file, or stdin when empty")
	out := flag.String("out", ".", "Output directory for the split files")
	maxHeight := flag.Int("max-height", 0, "Maximum height of a slide in lines (overridden by -template-size)")
	maxWidth := flag.Int("max-width", 0, "Maximum width of a slide in pixels (overridden by -template-size)")
	theme := flag.String("theme", "light", "light or dark")
	templateSize := flag.String("template-size", "", "Predefined template size: card, horizontal-card, presentation, a4")
	fontSize := flag.Int("font-size", 12, "Font size in points")
	dpi := flag.Int("dpi", 96, "DPI for rendering")
	flag.Parse()

	// Read the input from the specified file or stdin.
	var data []byte
	var err error
	if *in == "" {
		data, err = io.ReadAll(os.Stdin)
	} else {
		data, err = os.ReadFile(*in)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Create the SplitOptions struct.
	opts := mdsplit.SplitOptions{
		OutDir:       *out,
		MaxHeight:    *maxHeight,
		MaxWidth:     *maxWidth,
		Theme:        *theme,
		TemplateSize: mdsplit.TemplateSize(*templateSize),
		FontSize:     *fontSize,
		DPI:          *dpi,
	}

	// Split the Markdown file.
	if err := mdsplit.Split(data, opts); err != nil {
		fmt.Fprintf(os.Stderr, "Error splitting Markdown: %v\n", err)
		os.Exit(1)
	}
}
