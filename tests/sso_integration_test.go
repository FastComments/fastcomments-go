package tests

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/fastcomments/fastcomments-go/client"
	"github.com/fastcomments/fastcomments-go/sso"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Environment variables
func getAPIKey() string {
	apiKey := os.Getenv("FASTCOMMENTS_API_KEY")
	if apiKey == "" {
		panic("FASTCOMMENTS_API_KEY environment variable is required")
	}
	return apiKey
}

func getTenantID() string {
	tenantID := os.Getenv("FASTCOMMENTS_TENANT_ID")
	if tenantID == "" {
		panic("FASTCOMMENTS_TENANT_ID environment variable is required")
	}
	return tenantID
}

// TestGetCommentsWithSecureSSO tests getting comments using secure SSO
func TestGetCommentsWithSecureSSO(t *testing.T) {
	apiKey := getAPIKey()
	tenantID := getTenantID()

	timestamp := time.Now().UnixMilli()
	user := &sso.SecureSSOUserData{
		UserID:   fmt.Sprintf("test-user-%d", timestamp),
		Email:    fmt.Sprintf("test-%d@example.com", timestamp),
		Username: fmt.Sprintf("testuser%d", timestamp),
		Avatar:   "https://example.com/avatar.jpg",
	}

	ssoInstance, err := sso.NewSecure(apiKey, user)
	require.NoError(t, err)

	token, err := ssoInstance.CreateToken()
	require.NoError(t, err)
	assert.NotEmpty(t, token)

	// Call PublicApi.GetCommentsPublic with the SSO token
	config := client.NewConfiguration()
	apiClient := client.NewAPIClient(config)

	response, httpResp, err := apiClient.PublicAPI.GetCommentsPublic(context.Background(), tenantID).
		UrlId("sdk-test-page-secure").
		Sso(token).
		Execute()

	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 200, httpResp.StatusCode)
}

// TestCreateCommentWithSecureSSO tests creating a comment using secure SSO
func TestCreateCommentWithSecureSSO(t *testing.T) {
	apiKey := getAPIKey()
	tenantID := getTenantID()

	timestamp := time.Now().UnixMilli()
	user := &sso.SecureSSOUserData{
		UserID:   fmt.Sprintf("test-user-%d", timestamp),
		Email:    fmt.Sprintf("test-%d@example.com", timestamp),
		Username: fmt.Sprintf("testuser%d", timestamp),
		Avatar:   "https://example.com/avatar.jpg",
	}

	ssoInstance, err := sso.NewSecure(apiKey, user)
	require.NoError(t, err)

	token, err := ssoInstance.CreateToken()
	require.NoError(t, err)

	// Call PublicApi.CreateCommentPublic with the SSO token
	config := client.NewConfiguration()
	apiClient := client.NewAPIClient(config)

	commentData := client.NewCommentData(
		user.Username,
		"Test comment with secure SSO from Go SDK",
		"https://example.com/test-page",
		"sdk-test-page-secure-comment",
	)
	commentData.Date = &timestamp

	response, httpResp, err := apiClient.PublicAPI.CreateCommentPublic(context.Background(), tenantID).
		UrlId("sdk-test-page-secure-comment").
		BroadcastId(fmt.Sprintf("test-%d", timestamp)).
		CommentData(*commentData).
		Sso(token).
		Execute()

	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 200, httpResp.StatusCode)
}

// TestGetCommentsWithDefaultAPI tests getting comments using authenticated DefaultApi
func TestGetCommentsWithDefaultAPI(t *testing.T) {
	apiKey := getAPIKey()
	tenantID := getTenantID()

	// Call DefaultApi.GetComments with API key authentication
	config := client.NewConfiguration()
	apiClient := client.NewAPIClient(config)

	// Properly authenticate using context with API key
	auth := context.WithValue(
		context.Background(),
		client.ContextAPIKeys,
		map[string]client.APIKey{
			"api_key": {Key: apiKey},
		},
	)

	response, httpResp, err := apiClient.DefaultAPI.GetComments(auth).
		TenantId(tenantID).
		UrlId("sdk-test-page-secure-admin").
		Execute()

	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 200, httpResp.StatusCode)
}

// TestGetCommentsWithSimpleSSO tests getting comments using simple SSO
func TestGetCommentsWithSimpleSSO(t *testing.T) {
	tenantID := getTenantID()

	timestamp := time.Now().UnixMilli()
	user := &sso.SimpleSSOUserData{
		UserID: fmt.Sprintf("simple-user-%d", timestamp),
		Email:  fmt.Sprintf("simple-%d@example.com", timestamp),
		Avatar: "https://example.com/simple-avatar.jpg",
	}

	ssoInstance := sso.NewSimple(user)
	token, err := ssoInstance.CreateToken()
	require.NoError(t, err)
	assert.NotEmpty(t, token)

	// Call PublicApi.GetCommentsPublic with simple SSO token
	config := client.NewConfiguration()
	apiClient := client.NewAPIClient(config)

	response, httpResp, err := apiClient.PublicAPI.GetCommentsPublic(context.Background(), tenantID).
		UrlId("sdk-test-page-simple").
		Sso(token).
		Execute()

	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 200, httpResp.StatusCode)
}

// TestCreateCommentWithSimpleSSO tests creating a comment using simple SSO
func TestCreateCommentWithSimpleSSO(t *testing.T) {
	tenantID := getTenantID()

	timestamp := time.Now().UnixMilli()
	user := &sso.SimpleSSOUserData{
		UserID: fmt.Sprintf("simple-user-%d", timestamp),
		Email:  fmt.Sprintf("simple-%d@example.com", timestamp),
		Avatar: "https://example.com/simple-avatar.jpg",
	}

	ssoInstance := sso.NewSimple(user)
	token, err := ssoInstance.CreateToken()
	require.NoError(t, err)

	// Call PublicApi.CreateCommentPublic with simple SSO token
	config := client.NewConfiguration()
	apiClient := client.NewAPIClient(config)

	commentData := client.NewCommentData(
		user.UserID,
		"Test comment with simple SSO from Go SDK",
		"https://example.com/test-page",
		"sdk-test-page-simple-comment",
	)
	commentData.Date = &timestamp

	response, httpResp, err := apiClient.PublicAPI.CreateCommentPublic(context.Background(), tenantID).
		UrlId("sdk-test-page-simple-comment").
		BroadcastId(fmt.Sprintf("test-%d", timestamp)).
		CommentData(*commentData).
		Sso(token).
		Execute()

	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 200, httpResp.StatusCode)
}

// TestErrorHandlingInvalidTenant tests error handling with invalid tenant ID
func TestErrorHandlingInvalidTenant(t *testing.T) {
	apiKey := getAPIKey()

	timestamp := time.Now().UnixMilli()
	user := &sso.SecureSSOUserData{
		UserID:   fmt.Sprintf("test-user-%d", timestamp),
		Email:    fmt.Sprintf("test-%d@example.com", timestamp),
		Username: fmt.Sprintf("testuser%d", timestamp),
		Avatar:   "https://example.com/avatar.jpg",
	}

	ssoInstance, err := sso.NewSecure(apiKey, user)
	require.NoError(t, err)

	token, err := ssoInstance.CreateToken()
	require.NoError(t, err)

	// Call API with invalid tenant ID "invalid-tenant-123"
	config := client.NewConfiguration()
	apiClient := client.NewAPIClient(config)

	_, httpResp, err := apiClient.PublicAPI.GetCommentsPublic(context.Background(), "invalid-tenant-123").
		UrlId("sdk-test-page-secure").
		Sso(token).
		Execute()

	// Should return 400+ error
	require.Error(t, err, "Expected error for invalid tenant")
	if httpResp != nil {
		assert.GreaterOrEqual(t, httpResp.StatusCode, 400, "Expected 400+ error code")
	}
}
