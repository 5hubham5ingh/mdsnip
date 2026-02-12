# mdsnip

**Markdown Snippets: Compressed, encrypted, serverless, & private pastebin* with a complimentary CLI.**

A high-compression Markdown sharing tool that lives entirely in the URL. No databases, no accounts, no server persistence.
Your data is encrypted and stored only in the url so server has no knowledge and access to it.

## The Workflow

Share beautiful, rendered documents with others safely and instantly.

1. **Generate link:** Run `mdsnip docs.md` or use the Web Editor.
2. **Compress & Encrypt:** Content is optimized using DEFLATE and can be password-protected with AES-GCM.
3. **Share:** Get an instant URL where the content is embedded in the hash fragment.
4. **View:** Your recipient clicks the link and decompress/decrypts the content **locally in their browser** using WebAssembly.

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
# Copy a shareable link to clipboard
mdsnip document.md | wl-copy

# Encrypted sharing (prompts for password)
mdsnip -e secret.md

# Create a link from piped content
cat Doc.md | mdsnip

# Decrypt and view an mdsnip URL in terminal
mdsnip "https://5hubham5ingh.github.io/mdsnip/#.compressed_data..."

# View shared link for the Markdown in terminal
wl-paste | mdsnip
```

## Web App

A minimalist, premium editor and renderer built with performance and privacy in mind.

### Key Features
- **Zero-Storage:** Your data never touches a server disk.
- **WASM Engine:** Blazing fast decompression and decryption via WebAssembly.
- **Smart Outline:** Automatic Table of Contents sidebar for easy navigation.
- **Zen Mode:** A distraction-free, full-screen writing experience.
- **Secure Modals:** Custom themed dialogs for password entry that respect browser security policies.
- **Live Toggle:** Effortlessly switch between WYSIWYG editing and clean preview.
- **Keyboard driven**: Access all features with keyboard shortcuts.

### Keyboard Shortcuts
| Shortcut | Action |
|----------|--------|
| `Ctrl + Space` | Toggle between Edit & View mode |
| `Ctrl + Enter` | Share / Generate Link |
| `Ctrl + Shift + L` | Switch Theme (Light/Dark) |
| `Ctrl + Shift + Z` | Toggle Zen Mode |
| `Ctrl + Shift + S` | Toggle Sidebar/Outline |

## Why

Before this tool, I use to write documents in Markdown using Nvim. Sharing this would require following steps:
1. Copy the Markdown content.
2. Open browser.
3. Open google docs.
4. Create new document.
5. Right click and paste as Markdown.
6. Change document access permission.
7. Copy the shareable URL.
8. Paste it in messaging app like Teams.

Using this tool, now I follow these steps.
1. Run `mdsnip document_file_name.md | wl-copy`
2. Paste it in Teams.

**Q:** Why not just paste the Markdown content directly in messaging app?
- It makes the chat history dirty and external document is preferable over it.
- The shareable link is only visible in single line as most chat app show only a portion of it from the start.
- It is much easier and quicker to search from links saved as bookmarks in the browser than open google docs and search them.
- The links are auto saved as browser history for easy recoverability.

**Q:** What about the URL length limit?

* **Server-side:** Where limits are most restrictive, but irrelevant here as **URL hashes (`#`)** stay on the client and never reach the server.
* **Chromium based browsers (Chrome/Edge/Brave):** Very lenient (~2,000,000 characters).
* **Firefox:** Effectively unlimited.

## Development

### Building from Source
The project uses a unified build system for the CLI (Multi-platform Go) and the Web App (Go/WASM + Minification).

```bash
# Build all platforms + WASM
./build.sh

# Build only for your current system
./build.sh -1
```

## Technical Details
- **Compression**: flate (DEFLATE Level 9)
- **Encryption**: AES-GCM (256-bit)
- **Encoding**: Base64 Raw URL Encoding
- **Frontend**: Vanilla JS + [OverType](https://github.com/5hubham5ingh/overtype) + WebAssembly
- **Security**: Local-only decryption. Passwords are never sent to any server.

