package sso

import (
	"errors"
	"time"
)

// FastCommentsSSO represents the main SSO handler
type FastCommentsSSO struct {
	SecureSSOPayload   *SecureSSOPayload
	SimpleSSOUserData  *SimpleSSOUserData
	LoginURL           *string
	LogoutURL          *string
	cachedToken        *string
}

// NewSecure creates a new FastCommentsSSO instance with secure SSO
func NewSecure(apiKey string, secureUserData *SecureSSOUserData) (*FastCommentsSSO, error) {
	timestamp := time.Now().Unix()

	userDataStr, err := secureUserData.AsJSONBase64()
	if err != nil {
		return nil, err
	}

	hash, err := CreateVerificationHash(apiKey, timestamp, userDataStr)
	if err != nil {
		return nil, err
	}

	payload := &SecureSSOPayload{
		UserDataJSONBase64: userDataStr,
		VerificationHash:   hash,
		Timestamp:          timestamp,
	}

	return &FastCommentsSSO{
		SecureSSOPayload:  payload,
		SimpleSSOUserData: nil,
	}, nil
}

// NewSimple creates a new FastCommentsSSO instance with simple SSO
func NewSimple(simpleUserData *SimpleSSOUserData) *FastCommentsSSO {
	return &FastCommentsSSO{
		SecureSSOPayload:  nil,
		SimpleSSOUserData: simpleUserData,
	}
}

// NewSecureWithURLs creates a new FastCommentsSSO instance with secure SSO and URLs
func NewSecureWithURLs(securePayload *SecureSSOPayload, loginURL, logoutURL string) *FastCommentsSSO {
	return &FastCommentsSSO{
		SecureSSOPayload:  securePayload,
		SimpleSSOUserData: nil,
		LoginURL:          &loginURL,
		LogoutURL:         &logoutURL,
	}
}

// CreateToken creates a token from the SSO data
func (f *FastCommentsSSO) CreateToken() (string, error) {
	if f.SecureSSOPayload != nil {
		return f.SecureSSOPayload.ToJSON()
	} else if f.SimpleSSOUserData != nil {
		return f.SimpleSSOUserData.ToJSON()
	}
	return "", errors.New("no user data provided")
}

// ResetToken clears the cached token
func (f *FastCommentsSSO) ResetToken() {
	f.cachedToken = nil
}

// PrepareToSend prepares the token for sending, using cache if available
func (f *FastCommentsSSO) PrepareToSend() (string, error) {
	if f.cachedToken != nil {
		return *f.cachedToken, nil
	}

	token, err := f.CreateToken()
	if err != nil {
		return "", err
	}

	f.cachedToken = &token
	return token, nil
}

// SetSecureSSOPayload sets the secure SSO payload
func (f *FastCommentsSSO) SetSecureSSOPayload(payload *SecureSSOPayload) {
	f.SecureSSOPayload = payload
	f.SimpleSSOUserData = nil
	f.ResetToken()
}

// SetSimpleSSOUserData sets the simple SSO user data
func (f *FastCommentsSSO) SetSimpleSSOUserData(userData *SimpleSSOUserData) {
	f.SimpleSSOUserData = userData
	f.SecureSSOPayload = nil
	f.ResetToken()
}
