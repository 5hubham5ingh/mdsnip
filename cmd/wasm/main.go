package main

import (
	"syscall/js"

	"github.com/5hubham5ingh/mdsnip/pkg/codec"
)

func compressMarkdown(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return "Error: No input provided"
	}

	var input []byte
	if args[0].Type() == js.TypeString {
		input = []byte(args[0].String())
	} else {
		// Handle Uint8Array input
		input = make([]byte, args[0].Get("length").Int())
		js.CopyBytesToGo(input, args[0])
	}

	encoded, err := codec.Compress(input)
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

	decodedBytes, err := codec.Decompress(encoded)
	if err != nil {
		return "Error: Decompression failed - " + err.Error()
	}

	// Return Uint8Array to JS to avoid UTF-8 string corruption
	uint8Array := js.Global().Get("Uint8Array").New(len(decodedBytes))
	js.CopyBytesToJS(uint8Array, decodedBytes)
	return uint8Array
}

func encryptMarkdown(this js.Value, args []js.Value) any {
	if len(args) < 2 {
		return "Error: Input and Password required"
	}
	input := args[0].String()
	password := args[1].String()

	encryptedBytes, err := codec.Encrypt([]byte(input), password)
	if err != nil {
		return "Error: Encryption failed - " + err.Error()
	}

	// Return Uint8Array to JS
	uint8Array := js.Global().Get("Uint8Array").New(len(encryptedBytes))
	js.CopyBytesToJS(uint8Array, encryptedBytes)
	return uint8Array
}

func decryptMarkdown(this js.Value, args []js.Value) any {
	if len(args) < 2 {
		return "Error: Data and Password required"
	}

	// Data here is expected to be a Uint8Array
	dataJS := args[0]
	if dataJS.Type() == js.TypeString {
		return "Error: Expected Uint8Array for encrypted data"
	}

	data := make([]byte, dataJS.Get("length").Int())
	js.CopyBytesToGo(data, dataJS)
	password := args[1].String()

	decryptedBytes, err := codec.Decrypt(data, password)
	if err != nil {
		return "Error: Decryption failed - check your password"
	}

	return string(decryptedBytes)
}

func main() {
	js.Global().Set("compressMarkdown", js.FuncOf(compressMarkdown))
	js.Global().Set("decompressMarkdown", js.FuncOf(decompressMarkdown))
	js.Global().Set("encryptMarkdown", js.FuncOf(encryptMarkdown))
	js.Global().Set("decryptMarkdown", js.FuncOf(decryptMarkdown))

	select {}
}
