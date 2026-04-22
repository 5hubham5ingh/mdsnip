#!/bin/bash

# run.sh - unified build and release script for mdsnip

set -e

show_menu() {
    echo "------------------------------------------"
    echo " mdsnip - Management Script"
    echo "------------------------------------------"
    echo "1) Start Dev Server (Minify + Serve)"
    echo "2) Release New Version"
    echo "3) Re-release Last Version"
    echo "4) Exit"
    echo "------------------------------------------"
    read -p "Choose an option [1-4]: " choice
}

start_dev_server() {
    echo "Minifying main.html -> index.html..."
    if command -v npx &>/dev/null; then
        npx html-minifier-next main.html \
            --collapse-whitespace \
            --remove-comments \
            --minify-js true \
            --minify-css true \
            -o index.html
    else
        echo "Warning: npx not found. Copying main.html to index.html without minification."
        cp main.html index.html
    fi

    echo "Starting server at http://localhost:8080"
    if command -v python3 &>/dev/null; then
        python3 -m http.server 8080
    elif command -v npx &>/dev/null; then
        npx serve .
    else
        echo "Error: No simple HTTP server found (python3 or npx serve)."
        exit 1
    fi
}

release_version() {
    echo "Last version:"
    git describe --tags --abbrev=0 || echo "No tags found."
    
    read -p "Enter new version (e.g., v1.0.1): " version
    if [ -z "$version" ]; then
        echo "Version cannot be empty."
        return
    fi

    read -p "Begin building version $version? [y/N]: " confirm
    [[ "$confirm" =~ ^[Yy]$ ]] || return

    read -p "Compile for release (run build.sh)? [y/N]: " compile
    if [[ "$compile" =~ ^[Yy]$ ]]; then
        bash build.sh
    fi

    echo "Preparing git commit..."
    git add .
    
    read -p "Confirm commit and push? [y/N]: " push_confirm
    [[ "$push_confirm" =~ ^[Yy]$ ]] || return

    git commit -m "Release $version"
    git push origin main
    git tag "$version"
    git push origin "$version"

    echo "Version $version released successfully!"
}

re_release_last() {
    last_version=$(git describe --tags --abbrev=0)
    read -p "Re-release version \"$last_version\"? [y/N]: " confirm
    [[ "$confirm" =~ ^[Yy]$ ]] || return

    echo "Deleting tag $last_version locally and remotely..."
    git tag -d "$last_version" || true
    git push origin --delete "$last_version" || true
    
    echo "Pushing tag $last_version again..."
    git tag "$last_version"
    git push origin "$last_version"
}

# Main loop
if [ -n "$1" ]; then
    # Direct command execution if argument provided
    case "$1" in
        dev) start_dev_server ;;
        release) release_version ;;
        *) echo "Unknown command: $1" ;;
    esac
else
    while true; do
        show_menu
        case $choice in
            1) start_dev_server ;;
            2) release_version ;;
            3) re_release_last ;;
            4) exit 0 ;;
            *) echo "Invalid option." ;;
        esac
        echo ""
    done
fi
