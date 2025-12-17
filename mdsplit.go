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
		// Handle long tables by splitting them at the AST level.
		if table, ok := node.(*extast.Table); ok && (table.ChildCount()-1) > opts.MaxHeight {
			if currentSlide.Len() > 0 {
				if err := writeSlide(opts.OutDir, slideCount, &currentSlide); err != nil {
					return err
				}
				slideCount++
				currentSlide.Reset()
				currentLineCount = 0
			}

			var header ast.Node
			var rows []ast.Node
			for child := table.FirstChild(); child != nil; child = child.NextSibling() {
				if child.Kind() == extast.KindTableHeader {
					header = child
				} else if child.Kind() == extast.KindTableRow {
					rows = append(rows, child)
				}
			}

			if header != nil {
				tablePart := 1
				for len(rows) > 0 {
					// Account for header, separator, and continuation note.
					chunkSize := opts.MaxHeight - 2
					if chunkSize <= 0 {
						chunkSize = 1
					}
					if len(rows) < chunkSize {
						chunkSize = len(rows)
					}
					chunk := rows[:chunkSize]
					rows = rows[chunkSize:]

					newTable := extast.NewTable()
					newTable.Alignments = table.Alignments
					// Manually reconstruct the header for each new table chunk.
					newTable.AppendChild(newTable, manuallyCloneHeader(header, data))
					for _, row := range chunk {
						// We can move the row nodes directly to the new table,
						// but it's safer to clone them as well.
						// For now, we'll move them.
						row.SetParent(nil)
						newTable.AppendChild(newTable, row)
					}

					var slideContent bytes.Buffer
					if err := renderer.Render(&slideContent, data, newTable); err != nil {
						return err
					}
					// Add continuation note if there are more rows.
					if len(rows) > 0 || tablePart > 1 {
						note := fmt.Sprintf("\n_Table continued (part %d)_", tablePart)
						if len(rows) == 0 {
							note = fmt.Sprintf("\n_Table continued (part %d - final part)_", tablePart)
						}
						slideContent.WriteString(note)
					}

					if err := writeSlide(opts.OutDir, slideCount, &slideContent); err != nil {
						return err
					}
					slideCount++
					tablePart++
				}
				continue
			}
		}

		// Process regular nodes.
		var nodeContent bytes.Buffer
		if err := renderer.Render(&nodeContent, data, node); err != nil {
			return err
		}
		nodeLineCount := bytes.Count(nodeContent.Bytes(), []byte{'\n'})

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

// manuallyCloneHeader creates a deep copy of a table header node.
func manuallyCloneHeader(header ast.Node, source []byte) ast.Node {
	headerRow := header.FirstChild()
	newHeaderRow := extast.NewTableRow(nil)
	for cell := headerRow.FirstChild(); cell != nil; cell = cell.NextSibling() {
		newCell := extast.NewTableCell()
		for textNode := cell.FirstChild(); textNode != nil; textNode = textNode.NextSibling() {
			if text, ok := textNode.(*ast.Text); ok {
				newText := ast.NewText()
				newText.Segment = text.Segment
				newCell.AppendChild(newCell, newText)
			}
		}
		newHeaderRow.AppendChild(newHeaderRow, newCell)
	}
	return extast.NewTableHeader(newHeaderRow)
}
