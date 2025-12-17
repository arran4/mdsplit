package mdsplit

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	markdown "github.com/teekennedy/goldmark-markdown"
	"github.com/yuin/goldmark"
	gfm "github.com/yuin/goldmark/extension"
	extast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/text"
)

// SplitOptions holds the configuration for splitting the Markdown file.
type SplitOptions struct {
	OutDir    string
	MaxHeight int
	MaxWidth  int
	Theme     string
}

// Split takes a Markdown file as a byte slice and splits it into smaller files.
func Split(data []byte, opts SplitOptions) error {
	if opts.OutDir == "" {
		opts.OutDir = "."
	}
	// Use MaxHeight as a proxy for max lines.
	if opts.MaxHeight == 0 {
		opts.MaxHeight = 40 // Default to 40 lines if not set
	}

	// Create the output directory if it doesn't exist.
	if err := os.MkdirAll(opts.OutDir, 0755); err != nil {
		return err
	}

	// Create a new Goldmark parser and renderer.
	parser := goldmark.New(goldmark.WithExtensions(gfm.GFM)).Parser()
	renderer := markdown.NewRenderer()

	// Parse the Markdown into an AST.
	root := parser.Parse(text.NewReader(data))

	var currentSlide bytes.Buffer
	slideCount := 1
	currentLineCount := 0

	for node := root.FirstChild(); node != nil; node = node.NextSibling() {
		var nodeContent bytes.Buffer
		if err := renderer.Render(&nodeContent, data, node); err != nil {
			return err
		}

		nodeLineCount := bytes.Count(nodeContent.Bytes(), []byte{'\n'})

		// Handle tables that are too long.
		if node.Kind() == extast.KindTable && nodeLineCount > opts.MaxHeight {
			// Write the current slide if it has content.
			if currentSlide.Len() > 0 {
				if err := writeSlide(opts.OutDir, slideCount, &currentSlide); err != nil {
					return err
				}
				slideCount++
				currentSlide.Reset()
				currentLineCount = 0
			}

			lines := strings.Split(nodeContent.String(), "\n")
			header := lines[0] + "\n" + lines[1] + "\n"
			rows := lines[2:]

			tablePart := 1
			for len(rows) > 0 {
				continuationNote := fmt.Sprintf("\n_Table continued (part %d)_", tablePart)
				chunkSize := opts.MaxHeight - 3 // Account for header and continuation note.
				if chunkSize <= 0 {
					chunkSize = 1
				}
				if len(rows) <= chunkSize {
					chunkSize = len(rows)
				}

				var slideContent bytes.Buffer
				slideContent.WriteString(header)
				slideContent.WriteString(strings.Join(rows[:chunkSize], "\n"))
				slideContent.WriteString(continuationNote)

				if err := writeSlide(opts.OutDir, slideCount, &slideContent); err != nil {
					return err
				}

				slideCount++
				rows = rows[chunkSize:]
				tablePart++
			}
			continue
		}

		// If the current slide has content and adding the new node would exceed the max height,
		// write the current slide and start a new one.
		if currentSlide.Len() > 0 && currentLineCount+nodeLineCount > opts.MaxHeight {
			if err := writeSlide(opts.OutDir, slideCount, &currentSlide); err != nil {
				return err
			}
			slideCount++
			currentSlide.Reset()
			currentLineCount = 0
		}

		currentSlide.Write(nodeContent.Bytes())
		currentLineCount += nodeLineCount
	}

	// Write the last slide to a file.
	if currentSlide.Len() > 0 {
		if err := writeSlide(opts.OutDir, slideCount, &currentSlide); err != nil {
			return err
		}
	}

	return nil
}

// writeSlide writes the content of a slide to a file.
func writeSlide(outDir string, slideCount int, content *bytes.Buffer) error {
	filename := fmt.Sprintf("slide-%d.md", slideCount)
	filepath := filepath.Join(outDir, filename)
	return os.WriteFile(filepath, content.Bytes(), 0644)
}
