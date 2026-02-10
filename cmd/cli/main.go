package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"

	"github.com/5hubham5ingh/mdsnip/pkg/codec"
	"github.com/5hubham5ingh/mdsnip/pkg/renderer"
	"golang.org/x/term"
)

const baseURL = "https://5hubham5ingh.github.io/mdsnip/"

func main() {
	// 1. Encryption flag
	encryptFlag := flag.Bool("e", false, "Encrypt the content or decrypt the URL")
	flag.Parse()

	var input []byte
	var err error
	isURL := false

	// 2. Determine Input Source (Argument vs. Stdin)
	args := flag.Args()
	if len(args) > 0 {
		// CLI Argument Case: Could be a URL or a File Path
		arg := args[0]
		if strings.HasPrefix(arg, baseURL) {
			input = []byte(arg)
			isURL = true
		} else {
			// Must be a file path
			input, err = os.ReadFile(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
				os.Exit(1)
			}
		}
	} else {
		// Stdin Case: Could be a URL or Raw Markdown
		if isPiped() {
			input, err = io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
				os.Exit(1)
			}
			input = bytes.TrimSpace(input)
			// Check if single-line stdin is a URL
			if !bytes.Contains(input, []byte("\n")) && strings.HasPrefix(string(input), baseURL) {
				isURL = true
			}
		} else {
			fmt.Fprintf(os.Stderr, "Usage: mdsnip <file_path|url> or pipe content to mdsnip\n")
			os.Exit(1)
		}
	}

	if isURL {
		handleDecode(string(input), *encryptFlag)
	} else {
		handleEncode(input, *encryptFlag)
	}
}

func handleEncode(content []byte, encrypt bool) {
	var dataToEncode = content
	var prefix = ""

	if encrypt {
		password := getPassword("Enter password to encrypt: ")
		encrypted, err := codec.Encrypt(content, password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Encryption error: %v\n", err)
			os.Exit(1)
		}
		dataToEncode = encrypted
		prefix = "."
	}

	encoded, err := codec.Compress(dataToEncode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Compression error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s#%s%s\n", baseURL, prefix, encoded)
}

func handleDecode(urlStr string, forceDecrypt bool) {
	parts := strings.Split(urlStr, "#")
	if len(parts) < 2 {
		fmt.Fprintf(os.Stderr, "Invalid mdsnip URL\n")
		os.Exit(1)
	}

	fragment := parts[1]
	isEncrypted := strings.HasPrefix(fragment, ".")

	if isEncrypted {
		fragment = strings.TrimPrefix(fragment, ".")
	}

	decodedBytes, err := codec.Decompress(fragment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Decompression error: %v\n", err)
		os.Exit(1)
	}

	if isEncrypted || forceDecrypt {
		password := getPassword("Content is encrypted. Enter password: ")
		decrypted, err := codec.Decrypt(decodedBytes, password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Decryption failed: %v\n", err)
			os.Exit(1)
		}
		decodedBytes = decrypted
	}

	renderer.Render(decodedBytes)
}

// ---- Utility ----
func isPiped() bool {
	info, _ := os.Stdin.Stat()
	return (info.Mode() & os.ModeCharDevice) == 0
}

func getPassword(prompt string) string {
	fmt.Fprint(os.Stderr, prompt)

	fd := int(syscall.Stdin)

	// If Stdin is not a terminal (e.g. redirected or piped),
	// we try to open the controlling terminal /dev/tty
	if !term.IsTerminal(fd) {
		tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
		if err == nil {
			defer tty.Close()
			fd = int(tty.Fd())
		}
	}

	bytePassword, err := term.ReadPassword(fd)
	fmt.Fprintln(os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading password: %v\n", err)
		os.Exit(1)
	}
	return string(bytePassword)
}
