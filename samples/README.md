# Samples for mdsplit

This directory contains a variety of Markdown files to test the functionality of the `mdsplit` tool.

## Files

- `basic.md`: A simple file with headings, paragraphs, and lists.
- `long_table.md`: A file with a table longer than the default `max-height` to test table splitting.
- `short_table.md`: A file with a table shorter than `max-height`.
- `mixed_content.md`: A file with a mix of text, code blocks, lists, and tables.
- `code_blocks.md`: A file with large code blocks.
- `images.md`: A file with images.
- `very_long_file.md`: A file that will generate many slides.

## Usage

To use these samples, run the `mdsplit` command from the root of the repository. For example:

```bash
go build ./cmd/mdsplit
./mdsplit -in samples/basic.md -out ./slides
```

This will create a `slides` directory containing the split Markdown files.
