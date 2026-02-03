# mdsnip

A high-compression Markdown sharing tool.

## Features
- **High Compression**: Uses DEFLATE (Level 9) + Base64URL encoding.
- **CLI Tool**: Generate shareable URLs from the command line.
- **WASM Powered**: Decompression happens in the browser via WebAssembly.
- **Minimalist**: Vanilla JS frontend with GitHub Markdown styling.
- **Offline First**: No server needed to decode (entirely client-side).

## Usage

### CLI Tool
Compress a markdown file and get a shareable URL:
```bash
go run cmd/cli/main.go your-file.md
```

### Build WASM
If you want to rebuild the WASM engine:
```bash
GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/wasm/main.go
```

## Development
The project structure:
- `pkg/codec`: Core compression logic.
- `cmd/cli`: CLI tool.
- `cmd/wasm`: WASM bridge for the frontend.
- `static/`: WASM and JS assets.
- `index.html`: Web viewer/editor.
