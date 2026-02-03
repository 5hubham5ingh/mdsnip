package codec

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
)

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

func Decompress(encoded string) (string, error) {
	data, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	r := flate.NewReader(bytes.NewReader(data))
	defer r.Close()

	var out bytes.Buffer
	if _, err := io.Copy(&out, r); err != nil {
		return "", err
	}

	return out.String(), nil
}
