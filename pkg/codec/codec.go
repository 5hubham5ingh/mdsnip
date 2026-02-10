package codec

import (
	"bytes"
	"compress/flate"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
)

// Encrypt secures the data using AES-GCM with a password-derived key.
// The output format is: [12-byte Nonce][Encrypted Data + 16-byte Auth Tag]
func Encrypt(input []byte, password string) ([]byte, error) {
	// 1. Derive a 32-byte key from the password
	key := sha256.Sum256([]byte(password))

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 2. Create a random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// 3. Encrypt and seal
	// Seal appends the result to the prefix (nonce), so we get [nonce][ciphertext]
	return gcm.Seal(nonce, nonce, input, nil), nil
}

func Decrypt(input []byte, password string) ([]byte, error) {
	key := sha256.Sum256([]byte(password))

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(input) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := input[:nonceSize], input[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func Compress(input []byte) (string, error) {
	var b bytes.Buffer
	w, err := flate.NewWriter(&b, flate.BestCompression)
	if err != nil {
		return "", err
	}
	if _, err := w.Write(input); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b.Bytes()), nil
}

func Decompress(encoded string) ([]byte, error) {
	data, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	r := flate.NewReader(bytes.NewReader(data))
	defer r.Close()

	var out bytes.Buffer
	if _, err := io.Copy(&out, r); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
