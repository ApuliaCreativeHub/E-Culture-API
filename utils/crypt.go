package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

// CryptSHA1 Make secret's hash, using SHA1, and return its hex value
func CryptSHA1(secret string) (string, error) {
	h := sha1.New()
	_, err := io.WriteString(h, secret)
	if err != nil {
		return "", err
	}
	hash := h.Sum(nil)

	return hex.EncodeToString(hash), err
}
