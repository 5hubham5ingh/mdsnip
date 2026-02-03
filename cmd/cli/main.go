package main

import (
	"fmt"
	"os"

	"github.com/5hubham5ingh/mdsnip/pkg/codec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mdsnip <file.md>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	encoded, err := codec.Compress(content)
	if err != nil {
		fmt.Printf("Error compressing content: %v\n", err)
		os.Exit(1)
	}

	baseURL := "https://5hubham5ingh.github.io/mdsnip/"
	fmt.Printf("%s#%s\n", baseURL, encoded)
}
