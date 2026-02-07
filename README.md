# mdsnip

A high-compression Markdown sharing tool that lives entirely in the URL with a complimentary CLI for terminal. No databases, no accounts, no server.

## The "CLI → Web" Workflow

Live in CLI share beautiful, rendered documents with others over web.

1. **Generate:** Run `mdsnip docs.md` in your terminal.
2. **Share:** Get an instant, highly-compressed URL.
3. **View:** Your recipient clicks the link and sees a perfectly rendered web page, **zero server-side fetching required.**

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
# generate link
# Use it in scripts to generate instant documentation links

```
[generated link](https://5hubham5ingh.github.io/mdsnip/#XJRRcttGDIbf9xSY6MVhJWrSNn3wm9vYHk_s1rXiym6nYy-5EInRcpddYKXIB-gBesSepANSctQ8SgQWwI_vxwQ6x4F6Y86gpaad1bHrEzJTDHBj09rFbQBubaLQgMToQVor4GmDDBiEEvodUABpEe7vrmFL0oIFfcdTh0Fs2sFP11ewigkEU0fB-hJ-juCs2Moy8hRCBFvXMQcZfzCmDabSmMkEPrUIb_SBf__-B5ZYvYFlTOuVj1tjrmmDWlw_a5MIFdostMp-CgmDw4QOXKyzdsJjc1FaTAxxgwm2WJXGvCuhKC4xYLKCp0UBdznA86iMZnPZuWets4v5aAjzreYttK4mXaKADUCBxQaZDnr63aui6FSf0nynSb8RbjXnUR9MWFNPGARqT_WaBy09hTXY4IARGSz0mFZYi999mWuLFfS2wSkUxQumg2wzJoewQqlbXVrCvzIldGVRGDObzQZRVbBPMXpjDmOP-tnKD2tkcJTGcqsUu68Gh9uxmWGnBzh8bHj6qrUVJUgF5Ckw1gmF4aSOnYZyjOHtsN0JXKlc3g_xxnyI2-CjdYMEFQWFR4sM9XtvZRVTN7b0xx16VHz-PGlFej6dzxuSNldlHbv5-zZXre3eU2ja-bjJedonaOnn5-fKcmsm8Gumej1uzXs4uaaQP887W_-yeGvEJph9flntXfK0eSgfyocnrzFPtnM_fF-KTWXzYji7CN1mHwjzzGnuY239vKIwH-qN897GLSa4Z9vgcRcDRWBhRR7N_pG787MPN-dl54yZwD0jkCiEXCfqhUEiNIfl7aH7Sn5liPe1D5tfYgVn_WD4jgJ11hPLVMnDICM66EhiGuDbs6ZWLIqPuIMLtJIT8qnSVEBR_I4pzhYSk23wleg6BlGeiSGg-oztBh3EAPbV25q8PFvcjIqg02TFUstqk7bvITOObmDbIRzfJo1yePwPhoYCqiqdpSCWwpBghSryJLu9-QOCDcMJGFu4yN7D-TCxdrB8XFwtHy-hO9y-vRj7s6Z3D_qEG8ItSGwaj-MzH3FXRZscXMQ68zjNbYobcsjANiCsDxHcxiR1Fi6_bOXHTN6pNQa0FzGn-n94DN9BAT2YgOEbUPlMOa_0Y8mteQ2MwR_5ps4p6TZ4x4LdUQLM3o10_BcAAP__)
---

## Web App

A minimalist, client-side editor and renderer.

**Key Features:**

* **Zero-Storage:** Your content is never saved on a server.
* **WASM Powered:** CLI and Web app uses the same compression and decompression engine to maintain compatibility with one another.
* **Full Editor:** WYSIWYG markdown editor with a live preview toggle.
* **Keyboard Focused:** Provides sane keyboard shortcuts.

---

## Building from Source

```bash
# Build all platforms + WASM
./build.sh

# Build only for your current system
./build.sh -1

```
