# mdsnip

**Markdown Snippets. Compressed, encrypted, truly serverless, & private pastebin*.**

A high-compression Markdown sharing tool that lives entirely in the URL. No databases, no accounts, no server persistence.
Your data is encrypted and stored only in the url so server has no knowledge and access to it.

## The Workflow

Share beautiful, rendered documents with others safely and instantly.

1. **Generate:** Run `mdsnip docs.md` or use the Web Editor.
2. **Compress & Encrypt:** Content is optimized using DEFLATE and can be password-protected with AES-GCM.
3. **Share:** Get an instant URL where the content is embedded in the hash fragment.
4. **View:** Your recipient clicks the link and decompress/decrypts the content **locally in their browser** using WebAssembly.

## Web App

A minimalist, premium editor and renderer built with performance and privacy in mind.

### Key Features
- **Zero-Storage:** Your data never touches a server disk.
- **WASM Engine:** Blazing fast decompression and decryption via WebAssembly.
- **Smart Outline:** Automatic Table of Contents sidebar for easy navigation.
- **Zen Mode:** A distraction-free, full-screen writing experience.
- **Secure Modals:** Custom themed dialogs for password entry that respect browser security policies.
- **Live Toggle:** Effortlessly switch between WYSIWYG editing and clean preview.

### Keyboard Shortcuts
| Shortcut | Action |
|----------|--------|
| `Ctrl + Space` | Toggle between Edit & View mode |
| `Ctrl + Enter` | Share / Generate Link |
| `Ctrl + Shift + L` | Switch Theme (Light/Dark) |
| `Ctrl + Shift + Z` | Toggle Zen Mode |
| `Ctrl + Shift + S` | Toggle Sidebar/Outline |

## CLI Tool

The command-line companion for power users.

### Installation
Download the latest binary for your platform from [Releases](https://github.com/5hubham5ingh/mdsnip/releases).

```bash
# Example for Linux/macOS
tar -xzf mdsnip_vX.X.X_linux_amd64.tar.gz
sudo mv mdsnip /usr/local/bin/
```

### Usage
```bash
# Standard sharing
mdsnip document.md

# Encrypted sharing (prompts for password)
mdsnip -e secret.md

# Create a link from piped content
cat Doc.md | mdsnip

# Decrypt and view an mdsnip URL in terminal
mdsnip "https://5hubham5ingh.github.io/mdsnip/#.compressed_data..."
```

## Development

### Building from Source
The project uses a unified build system for the CLI (Multi-platform Go) and the Web App (Go/WASM + Minification).

```bash
# Build all platforms + WASM
./build.sh

# Build only for your current system
./build.sh -1
```

### Local Development
To run the web app locally with automatic minification:
```bash
./serve.sh
```

## Technical Details
- **Compression**: flate (DEFLATE Level 9)
- **Encryption**: AES-GCM (256-bit)
- **Encoding**: Base64 Raw URL Encoding
- **Frontend**: Vanilla JS + [OverType](https://github.com/5hubham5ingh/overtype) + WebAssembly
- **Security**: Local-only decryption. Passwords are never sent to any server.

