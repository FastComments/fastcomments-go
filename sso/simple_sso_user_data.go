package sso

import (
	"encoding/json"
)

// SimpleSSOUserData represents user data for simple SSO
type SimpleSSOUserData struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

// ToJSON converts the user data to JSON string
func (u *SimpleSSOUserData) ToJSON() (string, error) {
	jsonBytes, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
