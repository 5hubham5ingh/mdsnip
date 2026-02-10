package renderer

import (
	"fmt"
	"os"
)

// Render prints the markdown content to the terminal.
// Currently implements a basic pass-through, ready for ANSI styling later.
func Render(content []byte) {
	if len(content) == 0 {
		return
	}

	// Output the raw content
	fmt.Print(string(content))

	// Ensure the terminal prompt starts on a new line if the content doesn't end with one
	if content[len(content)-1] != '\n' {
		fmt.Println()
	}
}

// Error is a helper to standardize error reporting to stderr within the package
func Error(err error) {
	fmt.Fprintf(os.Stderr, "Render Error: %v\n", err)
}
