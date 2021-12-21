package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/golang-jwt/jwt"
	"io"
	"time"
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

func NewJSONWebToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		IssuedAt: time.Now().Unix(),
		Issuer:   "ApuliaCreativeHub",
		Subject:  "ECultureToolAuthorizationToken",
	})

	secret, err := GenerateRandomString(32)
	if err != nil {
		return "", err
	}
	hmacSecret := hmac.New(sha256.New, []byte(secret))
	return token.SignedString([]byte(hex.EncodeToString(hmacSecret.Sum(nil))))
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
