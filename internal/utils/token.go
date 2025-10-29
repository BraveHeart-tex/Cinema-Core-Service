package utils

import (
	"errors"
	"strings"
)

// ParseSessionToken parses a session token and returns the session ID and secret.
// The token should be in the format "sessionID.secret".
// If the token is invalid, it returns an error.
func ParseSessionToken(token string) (sessionID string, secret string, err error) {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return "", "", errors.New("invalid token format")
	}
	return parts[0], parts[1], nil
}
