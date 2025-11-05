package sso

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// CreateHashError represents an error that occurred during hash creation
type CreateHashError struct {
	Message string
}

func (e *CreateHashError) Error() string {
	return fmt.Sprintf("failed to create verification hash: %s", e.Message)
}

// CreateVerificationHash creates an HMAC-SHA256 hash for SSO verification
func CreateVerificationHash(apiKey string, timestamp int64, userDataJSONBase64 string) (string, error) {
	// Create message string by concatenating timestamp and base64 data
	messageStr := fmt.Sprintf("%d%s", timestamp, userDataJSONBase64)

	// Create HMAC using SHA256 hash function
	h := hmac.New(sha256.New, []byte(apiKey))
	_, err := h.Write([]byte(messageStr))
	if err != nil {
		return "", &CreateHashError{Message: err.Error()}
	}

	// Get digest as bytes then convert to hex
	bytesResult := h.Sum(nil)
	return GetBytesAsHex(bytesResult), nil
}

// GetBytesAsHex converts bytes to hex string
func GetBytesAsHex(bytesData []byte) string {
	return hex.EncodeToString(bytesData)
}
