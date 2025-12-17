package mdsplit

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	markdown "github.com/teekennedy/goldmark-markdown"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
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
	if opts.MaxHeight == 0 {
		opts.MaxHeight = 40
	}

	if err := os.MkdirAll(opts.OutDir, 0755); err != nil {
		return err
	}

	parser := goldmark.New(goldmark.WithExtensions(gfm.GFM)).Parser()
	renderer := markdown.NewRenderer()
	root := parser.Parse(text.NewReader(data))

	var currentSlide bytes.Buffer
	slideCount := 1
	currentLineCount := 0

	for node := root.FirstChild(); node != nil; node = node.NextSibling() {
		var nodeContent bytes.Buffer

		if err := safeRender(renderer, &nodeContent, data, node); err != nil {
			return err
		}

		// Trim leading newlines to avoid double padding accumulated from previous nodes
		trimmedBytes := bytes.TrimLeft(nodeContent.Bytes(), "\n")
		nodeContent.Reset()
		nodeContent.Write(trimmedBytes)

		nodeLineCount := bytes.Count(nodeContent.Bytes(), []byte{'\n'})

		// Handle paragraphs that are too long.
		// fmt.Printf("DEBUG: Current Total: %d, Max: %d, Will Add: %v\n", currentLineCount, opts.MaxHeight, currentLineCount+nodeLineCount >= opts.MaxHeight)

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
			// Remove trailing empty lines resulting from Split on string ending with newlines
			for len(lines) > 0 && lines[len(lines)-1] == "" {
				lines = lines[:len(lines)-1]
			}

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
			}

				var slideContent bytes.Buffer
				slideContent.WriteString(header)
				slideContent.WriteString(strings.Join(rows[:chunkSize], "\n"))
				slideContent.WriteString("\n")
				slideContent.WriteString(continuationNote)

				if err := writeSlide(opts.OutDir, slideCount, &slideContent); err != nil {
					return err
				}
				continue
			}
		}

		// Handle paragraphs that are too long.
		if node.Kind() == ast.KindParagraph && nodeLineCount > opts.MaxHeight {
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
			// Remove trailing empty lines
			for len(lines) > 0 && lines[len(lines)-1] == "" {
				lines = lines[:len(lines)-1]
			}

			// Split paragraph into chunks
			for len(lines) > 0 {
				chunkSize := opts.MaxHeight
				if len(lines) < chunkSize {
					chunkSize = len(lines)
				}

				var slideContent bytes.Buffer
				slideContent.WriteString(strings.Join(lines[:chunkSize], "\n"))
				// Append newline if needed? strings.Join doesn't add trailing newline.
				// But original lines didn't have it (Split removed it).
				// We should add it back?
				// Paragraphs implies text.
				slideContent.WriteString("\n")

				if err := writeSlide(opts.OutDir, slideCount, &slideContent); err != nil {
					return err
				}

				slideCount++
				lines = lines[chunkSize:]
			}
			continue
		}

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

	if currentSlide.Len() > 0 {
		if err := writeSlide(opts.OutDir, slideCount, &currentSlide); err != nil {
			return err
		}
	}

	return nil
}

func writeSlide(outDir string, slideCount int, content *bytes.Buffer) error {
	filename := fmt.Sprintf("slide-%d.md", slideCount)
	filepath := filepath.Join(outDir, filename)
	return os.WriteFile(filepath, content.Bytes(), 0644)
}

func safeRender(renderer *markdown.Renderer, w *bytes.Buffer, source []byte, n ast.Node) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// Fallback: extract raw lines from source by finding the range covered by the node and its children
			start, stop := getNodeBounds(n)
			if start != -1 && stop != -1 {
				// Expand to full lines
				for start > 0 && source[start-1] != '\n' {
					start--
				}
				for stop < len(source) && source[stop] != '\n' {
					stop++
				}
				// Include the newline at the end if present
				if stop < len(source) && source[stop] == '\n' {
					stop++
				}

				if start < stop {
					content := source[start:stop]
					if n.Kind() == ast.KindParagraph {
						content = wrapText(content, 60)
					}
					w.Write(content)
					// Append newlines to mimic Block spacing usually added by renderer.
					w.Write([]byte("\n\n"))
				}
			}
			err = nil
		}
	}()
	err = renderer.Render(w, source, n)
	if err == nil {
		// Ensure block spacing even if renderer was tight
		if w.Len() > 0 && !bytes.HasSuffix(w.Bytes(), []byte("\n\n")) {
			if bytes.HasSuffix(w.Bytes(), []byte("\n")) {
				w.Write([]byte("\n"))
			} else {
				w.Write([]byte("\n\n"))
			}
		}
	}
	return err
}

func getNodeBounds(n ast.Node) (int, int) {
	start := -1
	stop := -1

	updateBounds := func(s, e int) {
		if start == -1 || s < start {
			start = s
		}
		if stop == -1 || e > stop {
			stop = e
		}
	}

	if n.Type() == ast.TypeBlock {
		lines := n.Lines()
		if lines != nil {
			for i := 0; i < lines.Len(); i++ {
				segment := lines.At(i)
				updateBounds(segment.Start, segment.Stop)
			}
		}
	}

	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		cStart, cStop := getNodeBounds(c)
		if cStart != -1 {
			updateBounds(cStart, cStop)
		}
	}

	return start, stop
}

func wrapText(text []byte, limit int) []byte {
	var result bytes.Buffer
	for _, line := range bytes.Split(text, []byte{'\n'}) {
		if len(line) == 0 {
			result.WriteByte('\n')
			continue
		}
		words := bytes.Fields(line)
		if len(words) == 0 {
			continue
		}
		currentLineLen := 0
		for i, word := range words {
			if currentLineLen+len(word)+1 > limit && currentLineLen > 0 {
				result.WriteByte('\n')
				currentLineLen = 0
			} else if i > 0 {
				result.WriteByte(' ')
				currentLineLen++
			}
			result.Write(word)
			currentLineLen += len(word)
		}
		result.WriteByte('\n')
	}
	return result.Bytes()
}
