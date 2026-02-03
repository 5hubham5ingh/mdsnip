#!/bin/bash
echo "Minifying main.html -> index.html..."
npx html-minifier-terser main.html \
    --collapse-whitespace \
    --remove-comments \
    --minify-js true \
    --minify-css true \
    -o index.html

echo "Starting server at http://localhost:8080"
# Use Python if available, otherwise suggest alternatives
if command -v python3 &>/dev/null; then
    python3 -m http.server 8080
elif command -v npx &>/dev/null; then
    npx serve .
else
    echo "Please install a simple HTTP server (Python, Node/npx, or Caddy) to serve the current directory."
fi
