# mdsnip

**Terminal speed meets web beauty.** A high-compression Markdown sharing tool that lives entirely in the URL. No databases, no accounts, no server storage—just pure data.

## The "CLI → Web" Workflow

The core power of **mdsnip** is the seamless bridge between your terminal and the browser. It’s designed for developers who live in the CLI but want to share beautiful, rendered documents with others.

1. **Generate:** Run `mdsnip docs.md` in your terminal.
2. **Share:** Get an instant, highly-compressed URL.
3. **View:** Your recipient clicks the link and sees a perfectly rendered web page—**zero server-side fetching required.**

---

## CLI Tool

Generate shareable URLs directly from your terminal or pipes. Perfect for sharing logs, documentation snips, or README previews.

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

A minimalist, client-side editor and renderer. When you open an **mdsnip** link, the web app uses WebAssembly to decompress the data directly from the URL hash.

**Key Features:**

* **Zero-Storage:** Your content is never saved on a server. If you have the link, you have the data.
* **WASM Powered:** High-performance decompression happens locally in your browser.
* **Full Editor:** WYSIWYG markdown editor with a live preview toggle.
* **Keyboard Focused:**
* `Ctrl+Shift+S`: Copy compressed link
* `Ctrl+Shift+P`: Preview mode
* `Ctrl+Shift+E`: Edit mode

---

## How It Works

**mdsnip** treats the URL as a portable database:

* **Compression:** Uses **DEFLATE (Level 9)** to shrink your Markdown to the absolute minimum size.
* **Encoding:** Converts the binary blob into a URL-safe Base64 string.
* **Portability:** The CLI and Web App share the exact same logic. You can create a link in the terminal and edit it in the browser, or vice versa.

---

## Building from Source

```bash
# Build all platforms + WASM
./build.sh

# Build only for your current system
./build.sh -1

```
