package mdsplit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	readmeContent := readReadme(t)
	testCases := []struct {
		name                 string
		input                string
		opts                 SplitOptions
		expectedFileCount    int
		expectedContentCheck map[string]string // Check content for a subset of files
	}{
		{
			name:              "simple split",
			input:             "# Page 1\n\nSome content.\n\n# Page 2\n\nMore content.",
			opts:              SplitOptions{MaxHeight: 5},
			expectedFileCount: 2,
			expectedContentCheck: map[string]string{
				"slide-1.md": "# Page 1\n\nSome content.\n",
				"slide-2.md": "# Page 2\n\nMore content.",
			},
		},
		{
			name: "long table",
			input: `| Header 1 | Header 2 |
|---|---|
` + strings.Repeat("| a | b |\n", 50),
			opts:              SplitOptions{MaxHeight: 40},
			expectedFileCount: 2,
			expectedContentCheck: map[string]string{
				"slide-1.md": "| Header 1 | Header 2 |\n|---|---|\n" + strings.Repeat("| a | b |\n", 37) + "\n_Table continued (part 1)_",
				"slide-2.md": "| Header 1 | Header 2 |\n|---|---|\n" + strings.Repeat("| a | b |\n", 13) + "\n_Table continued (part 2)_",
			},
		},
		{
			name:              "size-based split",
			input:             strings.Repeat("This is a line of text.\n", 100),
			opts:              SplitOptions{MaxHeight: 40},
			expectedFileCount: 3,
			expectedContentCheck: map[string]string{
				"slide-1.md": strings.Repeat("This is a line of text.\n", 40),
				"slide-2.md": strings.Repeat("This is a line of text.\n", 40),
				"slide-3.md": strings.Repeat("This is a line of text.\n", 20),
			},
		},
		{
			name:              "readme split",
			input:             readmeContent,
			opts:              SplitOptions{MaxHeight: 40},
			expectedFileCount: 3,
			expectedContentCheck: map[string]string{
				"slide-1.md": `# mdsplit – Markdown Splitting (Go CLI & Library)

` + "`mdsplit`" + ` splits large Markdown files into smaller "slides" for easier viewing on mobile devices. It ships as a CLI and as a library so you can call it from your own code. No Node, headless browsers, or helper scripts.

---

## What it does

- Parses Markdown with ` + "`goldmark`" + ` and splits the result into multiple Markdown files.
- Intelligently splits content based on a maximum line count.
- Handles long tables by splitting them and adding a header to each part with a continuation note.
- Customizable slide size (vertical and horizontal).

---

## Install

Clone and build:

` + "```" + `bash
git clone https://github.com/arran4/mdsplit.git
cd mdsplit
go build ./cmd/mdsplit
` + "```" + `

Dependencies are managed in ` + "`go.mod`" + ` and will be automatically downloaded by the ` + "`go build`" + ` command.

Requires Go 1.22 or newer.`,
				"slide-3.md": `---

## Roadmap

- [ ] Use ` + "`md2png`" + `'s rendering engine to accurately measure slide height.
- [ ] Implement ` + "`-max-width`" + ` to control the width of the slides.
- [ ] Intelligent splitting of lists and code blocks.
- [ ] Support for different output formats (e.g., a single HTML file with multiple sections).
- [ ] Configurable themes via YAML/JSON.

---

## License

` + "`mdsplit`" + ` is available under the [MIT License](LICENSE).`,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tmpDir, err := os.MkdirTemp("", "mdsplit-test")
			if err != nil {
				t.Fatalf("Failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tmpDir)

			tc.opts.OutDir = tmpDir

			if err := Split([]byte(tc.input), tc.opts); err != nil {
				t.Fatalf("Split failed: %v", err)
			}

			if tc.expectedFileCount != countFiles(tmpDir) {
				files, _ := os.ReadDir(tmpDir)
				t.Fatalf("Expected %d files, but found %d: %v", tc.expectedFileCount, countFiles(tmpDir), files)
			}

			for expectedFile, expectedContent := range tc.expectedContentCheck {
				path := filepath.Join(tmpDir, expectedFile)
				if _, err := os.Stat(path); os.IsNotExist(err) {
					t.Errorf("Expected file %s was not created", expectedFile)
					continue
				}

				actualContent, err := os.ReadFile(path)
				if err != nil {
					t.Errorf("Failed to read file %s: %v", expectedFile, err)
					continue
				}

				if strings.TrimSpace(string(actualContent)) != strings.TrimSpace(expectedContent) {
					t.Errorf("File %s has incorrect content.\nExpected:\n%s\n\nActual:\n%s", expectedFile, expectedContent, actualContent)
				}
			}
		})
	}
}

func countFiles(dir string) int {
	files, err := os.ReadDir(dir)
	if err != nil {
		return 0
	}
	return len(files)
}

func readReadme(t *testing.T) string {
	t.Helper()
	content, err := os.ReadFile("README.md")
	if err != nil {
		t.Fatalf("Failed to read README.md: %v", err)
	}
	return string(content)
}
