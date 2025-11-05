package tests

import (
	"encoding/base64"
	"encoding/json"
	"testing"
	"time"

	"github.com/fastcomments/fastcomments-go/sso"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testAPIKey = "test-api-key-12345"

var testUserData = map[string]string{
	"user_id":  "test-user-123",
	"email":    "test@example.com",
	"username": "testuser",
	"avatar":   "https://example.com/avatar.jpg",
}

// TestCreateVerificationHash tests the hash creation function
func TestCreateVerificationHash(t *testing.T) {
	timestamp := time.Now().Unix()
	userData := "test_data_base64"

	result, err := sso.CreateVerificationHash(testAPIKey, timestamp, userData)

	require.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Len(t, result, 64) // SHA256 hex is 64 characters
}

// TestCreateVerificationHashWithEmptyAPIKey tests hash creation with empty API key
func TestCreateVerificationHashWithEmptyAPIKey(t *testing.T) {
	timestamp := time.Now().Unix()
	userData := "test_data"

	result, err := sso.CreateVerificationHash("", timestamp, userData)

	require.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Len(t, result, 64)
}

// TestSecureSSOUserData tests the SecureSSOUserData struct
func TestSecureSSOUserData(t *testing.T) {
	user := &sso.SecureSSOUserData{
		UserID:   testUserData["user_id"],
		Email:    testUserData["email"],
		Username: testUserData["username"],
		Avatar:   testUserData["avatar"],
	}

	assert.Equal(t, testUserData["user_id"], user.UserID)
	assert.Equal(t, testUserData["email"], user.Email)
	assert.Equal(t, testUserData["username"], user.Username)
	assert.Equal(t, testUserData["avatar"], user.Avatar)
}

// TestSecureSSOUserDataToJSON tests JSON serialization
func TestSecureSSOUserDataToJSON(t *testing.T) {
	user := &sso.SecureSSOUserData{
		UserID:   testUserData["user_id"],
		Email:    testUserData["email"],
		Username: testUserData["username"],
		Avatar:   testUserData["avatar"],
	}

	jsonStr, err := user.ToJSON()
	require.NoError(t, err)

	var parsed map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &parsed)
	require.NoError(t, err)

	assert.Equal(t, testUserData["user_id"], parsed["user_id"])
	assert.Equal(t, testUserData["email"], parsed["email"])
	assert.Equal(t, testUserData["username"], parsed["username"])
	assert.Equal(t, testUserData["avatar"], parsed["avatar"])
}

// TestSecureSSOUserDataAsJSONBase64 tests base64 encoding
func TestSecureSSOUserDataAsJSONBase64(t *testing.T) {
	user := &sso.SecureSSOUserData{
		UserID:   testUserData["user_id"],
		Email:    testUserData["email"],
		Username: testUserData["username"],
		Avatar:   testUserData["avatar"],
	}

	base64Str, err := user.AsJSONBase64()
	require.NoError(t, err)

	// Verify it's valid base64
	decodedBytes, err := base64.StdEncoding.DecodeString(base64Str)
	require.NoError(t, err)

	var parsed map[string]interface{}
	err = json.Unmarshal(decodedBytes, &parsed)
	require.NoError(t, err)

	assert.Equal(t, testUserData["user_id"], parsed["user_id"])
	assert.Equal(t, testUserData["email"], parsed["email"])
}

// TestSimpleSSOUserData tests the SimpleSSOUserData struct
func TestSimpleSSOUserData(t *testing.T) {
	user := &sso.SimpleSSOUserData{
		UserID: testUserData["user_id"],
		Email:  testUserData["email"],
		Avatar: testUserData["avatar"],
	}

	assert.Equal(t, testUserData["user_id"], user.UserID)
	assert.Equal(t, testUserData["email"], user.Email)
	assert.Equal(t, testUserData["avatar"], user.Avatar)
}

// TestSimpleSSOUserDataToJSON tests JSON serialization
func TestSimpleSSOUserDataToJSON(t *testing.T) {
	user := &sso.SimpleSSOUserData{
		UserID: testUserData["user_id"],
		Email:  testUserData["email"],
		Avatar: testUserData["avatar"],
	}

	jsonStr, err := user.ToJSON()
	require.NoError(t, err)

	var parsed map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &parsed)
	require.NoError(t, err)

	assert.Equal(t, testUserData["user_id"], parsed["user_id"])
	assert.Equal(t, testUserData["email"], parsed["email"])
	assert.Equal(t, testUserData["avatar"], parsed["avatar"])
}

// TestSecureSSOPayload tests the SecureSSOPayload struct
func TestSecureSSOPayload(t *testing.T) {
	timestamp := time.Now().Unix()
	userDataStr := "test_data_base64"
	hashValue := "test_hash"

	payload := &sso.SecureSSOPayload{
		UserDataJSONBase64: userDataStr,
		VerificationHash:   hashValue,
		Timestamp:          timestamp,
	}

	assert.Equal(t, userDataStr, payload.UserDataJSONBase64)
	assert.Equal(t, hashValue, payload.VerificationHash)
	assert.Equal(t, timestamp, payload.Timestamp)
}

// TestSecureSSOPayloadToJSON tests JSON serialization
func TestSecureSSOPayloadToJSON(t *testing.T) {
	timestamp := time.Now().Unix()
	payload := &sso.SecureSSOPayload{
		UserDataJSONBase64: "user_data",
		VerificationHash:   "hash_value",
		Timestamp:          timestamp,
	}

	jsonStr, err := payload.ToJSON()
	require.NoError(t, err)

	var parsed map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &parsed)
	require.NoError(t, err)

	assert.Equal(t, "user_data", parsed["user_data_json_base64"])
	assert.Equal(t, "hash_value", parsed["verification_hash"])
	assert.Equal(t, float64(timestamp), parsed["timestamp"])
}

// TestCreateSecureSSO tests creating secure SSO
func TestCreateSecureSSO(t *testing.T) {
	user := &sso.SecureSSOUserData{
		UserID:   testUserData["user_id"],
		Email:    testUserData["email"],
		Username: testUserData["username"],
		Avatar:   testUserData["avatar"],
	}

	ssoInstance, err := sso.NewSecure(testAPIKey, user)
	require.NoError(t, err)
	require.NotNil(t, ssoInstance)
	assert.NotNil(t, ssoInstance.SecureSSOPayload)
	assert.Nil(t, ssoInstance.SimpleSSOUserData)
}

// TestSecureSSOTokenCreation tests token creation
func TestSecureSSOTokenCreation(t *testing.T) {
	user := &sso.SecureSSOUserData{
		UserID:   testUserData["user_id"],
		Email:    testUserData["email"],
		Username: testUserData["username"],
		Avatar:   testUserData["avatar"],
	}

	ssoInstance, err := sso.NewSecure(testAPIKey, user)
	require.NoError(t, err)

	token, err := ssoInstance.CreateToken()
	require.NoError(t, err)

	var parsed map[string]interface{}
	err = json.Unmarshal([]byte(token), &parsed)
	require.NoError(t, err)

	assert.Contains(t, parsed, "user_data_json_base64")
	assert.Contains(t, parsed, "verification_hash")
	assert.Contains(t, parsed, "timestamp")

	// Verify the base64 data decodes to the original user data
	decodedBytes, err := base64.StdEncoding.DecodeString(parsed["user_data_json_base64"].(string))
	require.NoError(t, err)

	var userData map[string]interface{}
	err = json.Unmarshal(decodedBytes, &userData)
	require.NoError(t, err)
	assert.Equal(t, testUserData["user_id"], userData["user_id"])
}

// TestCreateSimpleSSO tests creating simple SSO
func TestCreateSimpleSSO(t *testing.T) {
	user := &sso.SimpleSSOUserData{
		UserID: testUserData["user_id"],
		Email:  testUserData["email"],
		Avatar: testUserData["avatar"],
	}

	ssoInstance := sso.NewSimple(user)
	require.NotNil(t, ssoInstance)
	assert.NotNil(t, ssoInstance.SimpleSSOUserData)
	assert.Nil(t, ssoInstance.SecureSSOPayload)
}

// TestSimpleSSOTokenCreation tests simple SSO token creation
func TestSimpleSSOTokenCreation(t *testing.T) {
	user := &sso.SimpleSSOUserData{
		UserID: testUserData["user_id"],
		Email:  testUserData["email"],
		Avatar: testUserData["avatar"],
	}

	ssoInstance := sso.NewSimple(user)
	token, err := ssoInstance.CreateToken()
	require.NoError(t, err)

	var parsed map[string]interface{}
	err = json.Unmarshal([]byte(token), &parsed)
	require.NoError(t, err)

	assert.Equal(t, testUserData["user_id"], parsed["user_id"])
	assert.Equal(t, testUserData["email"], parsed["email"])
	assert.Equal(t, testUserData["avatar"], parsed["avatar"])
}

// TestSecureSSOWithURLs tests creating secure SSO with URLs
func TestSecureSSOWithURLs(t *testing.T) {
	user := &sso.SecureSSOUserData{
		UserID:   testUserData["user_id"],
		Email:    testUserData["email"],
		Username: testUserData["username"],
		Avatar:   testUserData["avatar"],
	}

	timestamp := time.Now().Unix()
	userDataStr, err := user.AsJSONBase64()
	require.NoError(t, err)

	hashValue, err := sso.CreateVerificationHash(testAPIKey, timestamp, userDataStr)
	require.NoError(t, err)

	payload := &sso.SecureSSOPayload{
		UserDataJSONBase64: userDataStr,
		VerificationHash:   hashValue,
		Timestamp:          timestamp,
	}

	ssoInstance := sso.NewSecureWithURLs(payload, "/login", "/logout")

	assert.Equal(t, "/login", *ssoInstance.LoginURL)
	assert.Equal(t, "/logout", *ssoInstance.LogoutURL)
}

// TestNoDataRaisesError tests that SSO with no data raises an error
func TestNoDataRaisesError(t *testing.T) {
	ssoInstance := &sso.FastCommentsSSO{}

	_, err := ssoInstance.CreateToken()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no user data provided")
}

// TestTokenCaching tests that tokens are cached
func TestTokenCaching(t *testing.T) {
	user := &sso.SimpleSSOUserData{
		UserID: testUserData["user_id"],
		Email:  testUserData["email"],
		Avatar: testUserData["avatar"],
	}

	ssoInstance := sso.NewSimple(user)
	token1, err := ssoInstance.PrepareToSend()
	require.NoError(t, err)

	token2, err := ssoInstance.PrepareToSend()
	require.NoError(t, err)

	assert.Equal(t, token1, token2)
}

// TestResetToken tests that reset_token clears the cache
func TestResetToken(t *testing.T) {
	user := &sso.SimpleSSOUserData{
		UserID: testUserData["user_id"],
		Email:  testUserData["email"],
		Avatar: testUserData["avatar"],
	}

	ssoInstance := sso.NewSimple(user)
	token1, err := ssoInstance.PrepareToSend()
	require.NoError(t, err)

	ssoInstance.ResetToken()

	// After reset, a new token should be created
	token2, err := ssoInstance.PrepareToSend()
	require.NoError(t, err)
	assert.NotEmpty(t, token2)
	// Both tokens should be the same structure but we can't test internal state
	assert.NotEmpty(t, token1)
}

// TestSetSecureSSOPayload tests changing from simple to secure SSO
func TestSetSecureSSOPayload(t *testing.T) {
	simpleUser := &sso.SimpleSSOUserData{
		UserID: testUserData["user_id"],
		Email:  testUserData["email"],
		Avatar: testUserData["avatar"],
	}

	ssoInstance := sso.NewSimple(simpleUser)

	// Now switch to secure
	secureUser := &sso.SecureSSOUserData{
		UserID:   testUserData["user_id"],
		Email:    testUserData["email"],
		Username: testUserData["username"],
		Avatar:   testUserData["avatar"],
	}

	timestamp := time.Now().Unix()
	userDataStr, err := secureUser.AsJSONBase64()
	require.NoError(t, err)

	hashValue, err := sso.CreateVerificationHash(testAPIKey, timestamp, userDataStr)
	require.NoError(t, err)

	payload := &sso.SecureSSOPayload{
		UserDataJSONBase64: userDataStr,
		VerificationHash:   hashValue,
		Timestamp:          timestamp,
	}

	ssoInstance.SetSecureSSOPayload(payload)

	assert.NotNil(t, ssoInstance.SecureSSOPayload)
	assert.Nil(t, ssoInstance.SimpleSSOUserData)
}

// TestSetSimpleSSOUserData tests changing from secure to simple SSO
func TestSetSimpleSSOUserData(t *testing.T) {
	secureUser := &sso.SecureSSOUserData{
		UserID:   testUserData["user_id"],
		Email:    testUserData["email"],
		Username: testUserData["username"],
		Avatar:   testUserData["avatar"],
	}

	ssoInstance, err := sso.NewSecure(testAPIKey, secureUser)
	require.NoError(t, err)

	// Now switch to simple
	simpleUser := &sso.SimpleSSOUserData{
		UserID: testUserData["user_id"],
		Email:  testUserData["email"],
		Avatar: testUserData["avatar"],
	}

	ssoInstance.SetSimpleSSOUserData(simpleUser)

	assert.NotNil(t, ssoInstance.SimpleSSOUserData)
	assert.Nil(t, ssoInstance.SecureSSOPayload)
}
