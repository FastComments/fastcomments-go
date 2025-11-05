package sso

import (
	"encoding/base64"
	"encoding/json"
)

// SecureSSOUserData represents user data for secure SSO
type SecureSSOUserData struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

// ToJSON converts the user data to JSON string
func (u *SecureSSOUserData) ToJSON() (string, error) {
	jsonBytes, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// AsJSONBase64 converts the user data to base64-encoded JSON
func (u *SecureSSOUserData) AsJSONBase64() (string, error) {
	jsonStr, err := u.ToJSON()
	if err != nil {
		return "", err
	}
	jsonBytes := []byte(jsonStr)
	result := base64.StdEncoding.EncodeToString(jsonBytes)
	return result, nil
}
