package main

import (
	"syscall/js"

	"github.com/5hubham5ingh/mdsnip/pkg/codec"
)

func compressMarkdown(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return "Error: No input provided"
	}
	input := args[0].String()
	encoded, err := codec.Compress([]byte(input))
	if err != nil {
		return "Error: " + err.Error()
	}
	return encoded
}

func decompressMarkdown(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return "Error: No input provided"
	}
	encoded := args[0].String()
	decoded, err := codec.Decompress(encoded)
	if err != nil {
		return "Error: Invalid Data"
	}
	return decoded
}

func main() {
	js.Global().Set("compressMarkdown", js.FuncOf(compressMarkdown))
	js.Global().Set("decompressMarkdown", js.FuncOf(decompressMarkdown))

	// Keep the Go process alive
	select {}
}
