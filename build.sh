#!/bin/bash
set -e

APP_NAME="mdsnip"
VERSION="v0.1.0"
DIST_DIR="dist"

# Handle single build flag
if [ "$1" == "-1" ]; then
    echo "Building CLI for current system..."
    OUTPUT="mdsnip"
    if [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
        OUTPUT="mdsnip.exe"
    fi
    go build -o "$OUTPUT" cmd/cli/main.go
    echo "Done! Binary created at ./${OUTPUT}"
    exit 0
fi

# Clean previous builds
echo "Cleaning up..."
rm -rf $DIST_DIR
mkdir -p $DIST_DIR

# 1. Build WASM for Frontend
echo "Building WASM..."
mkdir -p static

if command -v tinygo &> /dev/null; then
    echo "  - Using TinyGo for optimized build..."
    TINYGO_ROOT=$(tinygo env TINYGOROOT)
    cp "${TINYGO_ROOT}/targets/wasm_exec.js" static/
    # -no-debug: strips debug symbols
    # -panic=trap: replaces panic messages with a simple processor trap
    tinygo build -o static/main.wasm -target wasm -no-debug -panic=trap cmd/wasm/main.go
    
    if command -v wasm-opt &> /dev/null; then
        echo "  - Running wasm-opt for maximum optimization..."
        wasm-opt -Oz static/main.wasm -o static/main.wasm
    fi
else
    echo "  - TinyGo not found, falling back to standard Go (binary will be larger)..."
    cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" static/
    GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/wasm/main.go
fi

# 2. Build CLI for multiple platforms
PLATFORMS=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
)

echo "Building CLI for multiple platforms..."

for PLATFORM in "${PLATFORMS[@]}"; do
    IFS="/" read -r OS ARCH <<< "$PLATFORM"
    
    OUTPUT_NAME="${APP_NAME}"
    if [ "$OS" == "windows" ]; then
        OUTPUT_NAME="${OUTPUT_NAME}.exe"
    fi

    # Create platform-specific build dir
    BUILD_DIR="${DIST_DIR}/${APP_NAME}_${OS}_${ARCH}"
    mkdir -p "${BUILD_DIR}"

    echo "  - Building for ${OS}/${ARCH}..."
    GOOS=$OS GOARCH=$ARCH go build -o "${BUILD_DIR}/${OUTPUT_NAME}" cmd/cli/main.go
    
    # Copy README for the release package
    cp README.md "${BUILD_DIR}/"

    # Create Archive
    cd "${DIST_DIR}"
    if [ "$OS" == "windows" ]; then
        ZIP_NAME="${APP_NAME}_${VERSION}_${OS}_${ARCH}.zip"
        zip -r "${ZIP_NAME}" "${APP_NAME}_${OS}_${ARCH}" > /dev/null
        echo "    Created ${ZIP_NAME}"
    else
        TAR_NAME="${APP_NAME}_${VERSION}_${OS}_${ARCH}.tar.gz"
        tar -czf "${TAR_NAME}" "${APP_NAME}_${OS}_${ARCH}"
        echo "    Created ${TAR_NAME}"
    fi
    # Cleanup the temp build directory after archiving
    rm -rf "${APP_NAME}_${OS}_${ARCH}"
    cd ..
done

echo "Done! All builds are in the '${DIST_DIR}' directory."
ls -lh $DIST_DIR
