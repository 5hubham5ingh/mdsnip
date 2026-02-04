# mdsnip

A high-compression Markdown sharing tool that lives entirely in the URL with a complimentary CLI for terminal. No databases, no accounts, no server.

## The "CLI → Web" Workflow

Live in CLI share beautiful, rendered documents with others over web.

1. **Generate:** Run `mdsnip docs.md` in your terminal.
2. **Share:** Get an instant, highly-compressed URL.
3. **View:** Your recipient clicks the link and sees a perfectly rendered web page—**zero server-side fetching required.**

---

## CLI Tool

Generate shareable URLs directly from your terminal. Perfect for sharing logs, documentation snips, secrets (coming soon).

### Installation

Download the binary for your platform from [Releases](https://github.com/5hubham5ingh/mdsnip/releases).

```bash
# Quick install (Linux/macOS)
tar -xzf mdsnip_vX.X.X_linux_amd64.tar.gz
sudo mv mdsnip /usr/local/bin/

```

### Power Usage

```bash
# Share a file
mdsnip README.md

# Use it in scripts to generate instant documentation links

```

---

## Web App

A minimalist, client-side editor and renderer.

**Key Features:**

* **Zero-Storage:** Your content is never saved on a server.
* **WASM Powered:** CLI and Web app uses the same compression and decompression engine to maintain compatibility with one another.
* **Full Editor:** WYSIWYG markdown editor with a live preview toggle.
* **Keyboard Focused:**
* `Ctrl+Shift+S`: Copy compressed link
* `Ctrl+Shift+P`: Preview mode
* `Ctrl+Shift+E`: Edit mode


---

## Building from Source

```bash
# Build all platforms + WASM
./build.sh

# Build only for your current system
./build.sh -1

```
