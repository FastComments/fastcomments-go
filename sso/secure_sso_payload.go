package sso

import (
	"encoding/json"
)

// SecureSSOPayload represents the secure SSO payload
type SecureSSOPayload struct {
	UserDataJSONBase64 string `json:"user_data_json_base64"`
	VerificationHash   string `json:"verification_hash"`
	Timestamp          int64  `json:"timestamp"`
}

// ToJSON converts the payload to JSON string
func (p *SecureSSOPayload) ToJSON() (string, error) {
	jsonBytes, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
