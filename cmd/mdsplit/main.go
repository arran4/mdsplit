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
	maxHeight := flag.Int("max-height", 40, "Maximum height of a slide in lines")
	maxWidth := flag.Int("max-width", 1024, "Maximum width of a slide in pixels")
	theme := flag.String("theme", "light", "light or dark")
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
		OutDir:    *out,
		MaxHeight: *maxHeight,
		MaxWidth:  *maxWidth,
		Theme:     *theme,
	}

	// Split the Markdown file.
	if err := mdsplit.Split(data, opts); err != nil {
		fmt.Fprintf(os.Stderr, "Error splitting Markdown: %v\n", err)
		os.Exit(1)
	}
}
