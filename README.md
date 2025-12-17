# mdsplit – Markdown Splitting (Go CLI & Library)

`mdsplit` splits large Markdown files into smaller "slides" for easier viewing on mobile devices. It ships as a CLI and as a library so you can call it from your own code. No Node, headless browsers, or helper scripts.

---

## What it does

- Parses Markdown with `goldmark` and splits the result into multiple Markdown files.
- Intelligently splits content based on a maximum line count.
- Handles long tables by splitting them and adding a header to each part with a continuation note.
- Customizable slide size (vertical and horizontal).

---

## Install

Clone and build:

```bash
git clone https://github.com/arran4/mdsplit.git
cd mdsplit
go build ./cmd/mdsplit
```

Dependencies are managed in `go.mod` and will be automatically downloaded by the `go build` command.

Requires Go 1.22 or newer.

---

## CLI usage

```bash
./mdsplit -in README.md -out ./slides
```

### Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-in` | Markdown input file, or stdin when empty | — |
| `-out` | Output directory for the split files | `.` |
| `-max-height` | Maximum height of a slide in lines | 40 |
| `-max-width` | Maximum width of a slide in pixels (not yet implemented) | 1024 |
| `-theme` | `light` or `dark` (not yet implemented) | `light` |

### Examples

Split a Markdown file into slides:

```bash
./mdsplit -in example.md -out ./slides -max-height 50
```

---

## Library usage

```go
package main

import (
        "os"

        "github.com/arran4/mdsplit"
)

func main() {
        err := mdsplit.Split([]byte("# Hello\nThis is a large Markdown file!"), mdsplit.SplitOptions{OutDir: "./slides", MaxHeight: 50})
        if err != nil {
                panic(err)
        }
}
```

`SplitOptions` exposes the same knobs as the CLI. Set custom dimensions.

---

## How it works

1. Parse Markdown with [`yuin/goldmark`](https://github.com/yuin/goldmark) and the [`goldmark-gfm`](https://github.com/yuin/goldmark-gfm) extension.
2. Walk the AST and split the content into multiple smaller Markdown files based on a maximum line count.
3. If a table is too long, it is split into multiple slides, with the header repeated on each slide.
4. Write the split Markdown files to the output directory.

Everything happens in memory; there is no HTML renderer or external process.

---

## Roadmap

- [ ] Use `md2png`'s rendering engine to accurately measure slide height.
- [ ] Implement `-max-width` to control the width of the slides.
- [ ] Intelligent splitting of lists and code blocks.
- [ ] Support for different output formats (e.g., a single HTML file with multiple sections).
- [ ] Configurable themes via YAML/JSON.

---

## License

`mdsplit` is available under the [MIT License](LICENSE).
