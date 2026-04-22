#!/bin/bash
set -e

# Mock wl-copy and wl-paste if not available
if ! command -v wl-copy &> /dev/null; then
    CLIPBOARD_FILE=$(mktemp)
    wl-copy() { cat > "$CLIPBOARD_FILE"; }
    wl-paste() { cat "$CLIPBOARD_FILE"; }
    export -f wl-copy wl-paste
fi

# 1
echo "Test 1: mdsnip ./README.md"
LINK1=$(go run cmd/cli/main.go ./README.md)
if [[ $LINK1 == https://5hubham5ingh.github.io/mdsnip/#* ]]; then
    echo "  Pass"
else
    echo "  Fail: $LINK1"
    exit 1
fi

# 2
echo "Test 2: mdsnip <link>"
RENDER2=$(go run cmd/cli/main.go "$LINK1")
if [[ $RENDER2 == *"mdsnip"* ]]; then
    echo "  Pass"
else
    echo "  Fail: Rendered output doesn't contain expected content"
    exit 1
fi

# 3
echo "Test 3: cat ./README.md | mdsnip"
LINK3=$(cat ./README.md | go run cmd/cli/main.go)
if [[ $LINK3 == https://5hubham5ingh.github.io/mdsnip/#* ]]; then
    echo "  Pass"
else
    echo "  Fail: $LINK3"
    exit 1
fi

# 4
echo "Test 4: pipe and wl-copy/paste"
go run cmd/cli/main.go ./README.md | wl-copy
RENDER4=$(wl-paste | go run cmd/cli/main.go)
if [[ $RENDER4 == *"mdsnip"* ]]; then
    echo "  Pass"
else
    echo "  Fail: Rendered output doesn't contain expected content"
    exit 1
fi

# 5
echo "Test 5: content in clipboard"
cat ./README.md | wl-copy
LINK5=$(wl-paste | go run cmd/cli/main.go)
if [[ $LINK5 == https://5hubham5ingh.github.io/mdsnip/#* ]]; then
    echo "  Pass"
else
    echo "  Fail: $LINK5"
    exit 1
fi

# Additional tests for bug fixes
echo "Test 6: Kitty fallback"
export TERM=xterm-kitty
export KITTY_WINDOW_ID=123
RENDER6=$(echo "# Header" | go run cmd/cli/main.go | go run cmd/cli/main.go)
if [[ $RENDER6 == *$'\x1b]66;'* ]]; then
    echo "  Pass (Kitty sequence detected)"
else
    echo "  Fail (Kitty sequence NOT detected)"
    exit 1
fi

unset TERM
unset KITTY_WINDOW_ID
RENDER7=$(echo "# Header" | go run cmd/cli/main.go | go run cmd/cli/main.go)
if [[ $RENDER7 == *"# Header"* && $RENDER7 != *$'\x1b]66;'* ]]; then
    echo "  Pass (Fallback detected)"
else
    echo "  Fail (Fallback NOT detected)"
    exit 1
fi

echo "Test 7: Clickable links (OSC 8)"
RENDER8=$(echo "[Google](https://google.com)" | go run cmd/cli/main.go | go run cmd/cli/main.go)
if [[ $RENDER8 == *$'\x1b]8;;https://google.com'* ]]; then
    echo "  Pass (OSC 8 detected)"
else
    echo "  Fail (OSC 8 NOT detected)"
    exit 1
fi

echo "All tests passed successfully!"
