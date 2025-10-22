package utils

import (
	"crypto/rand"
	"crypto/sha256"
)

func GenerateSecureRandomString() (string, error) {
	const alphabet = "abcdefghijkmnpqrstuvwxyz23456789"
	const length = 24

	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	id := make([]byte, length)
	for i, b := range bytes {
		id[i] = alphabet[b>>3]
	}

	return string(id), nil
}

func HashSecret(secret string) []byte {
	hash := sha256.Sum256([]byte(secret))
	return hash[:]
}
