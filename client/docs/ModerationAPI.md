# \ModerationAPI

All URIs are relative to *https://fastcomments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteModerationVote**](ModerationAPI.md#DeleteModerationVote) | **Delete** /auth/my-account/moderate-comments/mod_api/vote/{commentId}/{voteId} | 
[**GetApiComments**](ModerationAPI.md#GetApiComments) | **Get** /auth/my-account/moderate-comments/mod_api/api/comments | 
[**GetApiExportStatus**](ModerationAPI.md#GetApiExportStatus) | **Get** /auth/my-account/moderate-comments/mod_api/api/export/status | 
[**GetApiIds**](ModerationAPI.md#GetApiIds) | **Get** /auth/my-account/moderate-comments/mod_api/api/ids | 
[**GetBanUsersFromComment**](ModerationAPI.md#GetBanUsersFromComment) | **Get** /auth/my-account/moderate-comments/mod_api/ban-users/from-comment/{commentId} | 
[**GetCommentBanStatus**](ModerationAPI.md#GetCommentBanStatus) | **Get** /auth/my-account/moderate-comments/mod_api/get-comment-ban-status/{commentId} | 
[**GetCommentChildren**](ModerationAPI.md#GetCommentChildren) | **Get** /auth/my-account/moderate-comments/mod_api/comment-children/{commentId} | 
[**GetCount**](ModerationAPI.md#GetCount) | **Get** /auth/my-account/moderate-comments/mod_api/count | 
[**GetCounts**](ModerationAPI.md#GetCounts) | **Get** /auth/my-account/moderate-comments/banned-users/mod_api/counts | 
[**GetLogs**](ModerationAPI.md#GetLogs) | **Get** /auth/my-account/moderate-comments/mod_api/logs/{commentId} | 
[**GetManualBadges**](ModerationAPI.md#GetManualBadges) | **Get** /auth/my-account/moderate-comments/mod_api/get-manual-badges | 
[**GetManualBadgesForUser**](ModerationAPI.md#GetManualBadgesForUser) | **Get** /auth/my-account/moderate-comments/mod_api/get-manual-badges-for-user | 
[**GetModerationComment**](ModerationAPI.md#GetModerationComment) | **Get** /auth/my-account/moderate-comments/mod_api/comment/{commentId} | 
[**GetModerationCommentText**](ModerationAPI.md#GetModerationCommentText) | **Get** /auth/my-account/moderate-comments/mod_api/get-comment-text/{commentId} | 
[**GetPreBanSummary**](ModerationAPI.md#GetPreBanSummary) | **Get** /auth/my-account/moderate-comments/mod_api/pre-ban-summary/{commentId} | 
[**GetSearchCommentsSummary**](ModerationAPI.md#GetSearchCommentsSummary) | **Get** /auth/my-account/moderate-comments/mod_api/search/comments/summary | 
[**GetSearchPages**](ModerationAPI.md#GetSearchPages) | **Get** /auth/my-account/moderate-comments/mod_api/search/pages | 
[**GetSearchSites**](ModerationAPI.md#GetSearchSites) | **Get** /auth/my-account/moderate-comments/mod_api/search/sites | 
[**GetSearchSuggest**](ModerationAPI.md#GetSearchSuggest) | **Get** /auth/my-account/moderate-comments/mod_api/search/suggest | 
[**GetSearchUsers**](ModerationAPI.md#GetSearchUsers) | **Get** /auth/my-account/moderate-comments/mod_api/search/users | 
[**GetTrustFactor**](ModerationAPI.md#GetTrustFactor) | **Get** /auth/my-account/moderate-comments/mod_api/get-trust-factor | 
[**GetUserBanPreference**](ModerationAPI.md#GetUserBanPreference) | **Get** /auth/my-account/moderate-comments/mod_api/user-ban-preference | 
[**GetUserInternalProfile**](ModerationAPI.md#GetUserInternalProfile) | **Get** /auth/my-account/moderate-comments/mod_api/get-user-internal-profile | 
[**PostAdjustCommentVotes**](ModerationAPI.md#PostAdjustCommentVotes) | **Post** /auth/my-account/moderate-comments/mod_api/adjust-comment-votes/{commentId} | 
[**PostApiExport**](ModerationAPI.md#PostApiExport) | **Post** /auth/my-account/moderate-comments/mod_api/api/export | 
[**PostBanUserFromComment**](ModerationAPI.md#PostBanUserFromComment) | **Post** /auth/my-account/moderate-comments/mod_api/ban-user/from-comment/{commentId} | 
[**PostBanUserUndo**](ModerationAPI.md#PostBanUserUndo) | **Post** /auth/my-account/moderate-comments/mod_api/ban-user/undo | 
[**PostBulkPreBanSummary**](ModerationAPI.md#PostBulkPreBanSummary) | **Post** /auth/my-account/moderate-comments/mod_api/bulk-pre-ban-summary | 
[**PostCommentsByIds**](ModerationAPI.md#PostCommentsByIds) | **Post** /auth/my-account/moderate-comments/mod_api/comments-by-ids | 
[**PostFlagComment**](ModerationAPI.md#PostFlagComment) | **Post** /auth/my-account/moderate-comments/mod_api/flag-comment/{commentId} | 
[**PostRemoveComment**](ModerationAPI.md#PostRemoveComment) | **Post** /auth/my-account/moderate-comments/mod_api/remove-comment/{commentId} | 
[**PostRestoreDeletedComment**](ModerationAPI.md#PostRestoreDeletedComment) | **Post** /auth/my-account/moderate-comments/mod_api/restore-deleted-comment/{commentId} | 
[**PostSetCommentApprovalStatus**](ModerationAPI.md#PostSetCommentApprovalStatus) | **Post** /auth/my-account/moderate-comments/mod_api/set-comment-approval-status/{commentId} | 
[**PostSetCommentReviewStatus**](ModerationAPI.md#PostSetCommentReviewStatus) | **Post** /auth/my-account/moderate-comments/mod_api/set-comment-review-status/{commentId} | 
[**PostSetCommentSpamStatus**](ModerationAPI.md#PostSetCommentSpamStatus) | **Post** /auth/my-account/moderate-comments/mod_api/set-comment-spam-status/{commentId} | 
[**PostSetCommentText**](ModerationAPI.md#PostSetCommentText) | **Post** /auth/my-account/moderate-comments/mod_api/set-comment-text/{commentId} | 
[**PostUnFlagComment**](ModerationAPI.md#PostUnFlagComment) | **Post** /auth/my-account/moderate-comments/mod_api/un-flag-comment/{commentId} | 
[**PostVote**](ModerationAPI.md#PostVote) | **Post** /auth/my-account/moderate-comments/mod_api/vote/{commentId} | 
[**PutAwardBadge**](ModerationAPI.md#PutAwardBadge) | **Put** /auth/my-account/moderate-comments/mod_api/award-badge | 
[**PutCloseThread**](ModerationAPI.md#PutCloseThread) | **Put** /auth/my-account/moderate-comments/mod_api/close-thread | 
[**PutRemoveBadge**](ModerationAPI.md#PutRemoveBadge) | **Put** /auth/my-account/moderate-comments/mod_api/remove-badge | 
[**PutReopenThread**](ModerationAPI.md#PutReopenThread) | **Put** /auth/my-account/moderate-comments/mod_api/reopen-thread | 
[**SetTrustFactor**](ModerationAPI.md#SetTrustFactor) | **Put** /auth/my-account/moderate-comments/mod_api/set-trust-factor | 



## DeleteModerationVote

> VoteDeleteResponse DeleteModerationVote(ctx, commentId, voteId).TenantId(tenantId).BroadcastId(broadcastId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	voteId := "voteId_example" // string | 
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.DeleteModerationVote(context.Background(), commentId, voteId).TenantId(tenantId).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.DeleteModerationVote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteModerationVote`: VoteDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.DeleteModerationVote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 
**voteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteModerationVoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**VoteDeleteResponse**](VoteDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiComments

> ModerationAPIGetCommentsResponse GetApiComments(ctx).TenantId(tenantId).Page(page).Count(count).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filters(filters).SearchFilters(searchFilters).Sorts(sorts).Demo(demo).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	page := float64(1.2) // float64 |  (optional)
	count := float64(1.2) // float64 |  (optional)
	textSearch := "textSearch_example" // string |  (optional)
	byIPFromComment := "byIPFromComment_example" // string |  (optional)
	filters := "filters_example" // string |  (optional)
	searchFilters := "searchFilters_example" // string |  (optional)
	sorts := "sorts_example" // string |  (optional)
	demo := true // bool |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetApiComments(context.Background()).TenantId(tenantId).Page(page).Count(count).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filters(filters).SearchFilters(searchFilters).Sorts(sorts).Demo(demo).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetApiComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiComments`: ModerationAPIGetCommentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetApiComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **page** | **float64** |  | 
 **count** | **float64** |  | 
 **textSearch** | **string** |  | 
 **byIPFromComment** | **string** |  | 
 **filters** | **string** |  | 
 **searchFilters** | **string** |  | 
 **sorts** | **string** |  | 
 **demo** | **bool** |  | 
 **sso** | **string** |  | 

### Return type

[**ModerationAPIGetCommentsResponse**](ModerationAPIGetCommentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiExportStatus

> ModerationExportStatusResponse GetApiExportStatus(ctx).TenantId(tenantId).BatchJobId(batchJobId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	batchJobId := "batchJobId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetApiExportStatus(context.Background()).TenantId(tenantId).BatchJobId(batchJobId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetApiExportStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiExportStatus`: ModerationExportStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetApiExportStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiExportStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **batchJobId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**ModerationExportStatusResponse**](ModerationExportStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiIds

> ModerationAPIGetCommentIdsResponse GetApiIds(ctx).TenantId(tenantId).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filters(filters).SearchFilters(searchFilters).AfterId(afterId).Demo(demo).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	textSearch := "textSearch_example" // string |  (optional)
	byIPFromComment := "byIPFromComment_example" // string |  (optional)
	filters := "filters_example" // string |  (optional)
	searchFilters := "searchFilters_example" // string |  (optional)
	afterId := "afterId_example" // string |  (optional)
	demo := true // bool |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetApiIds(context.Background()).TenantId(tenantId).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filters(filters).SearchFilters(searchFilters).AfterId(afterId).Demo(demo).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetApiIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiIds`: ModerationAPIGetCommentIdsResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetApiIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **textSearch** | **string** |  | 
 **byIPFromComment** | **string** |  | 
 **filters** | **string** |  | 
 **searchFilters** | **string** |  | 
 **afterId** | **string** |  | 
 **demo** | **bool** |  | 
 **sso** | **string** |  | 

### Return type

[**ModerationAPIGetCommentIdsResponse**](ModerationAPIGetCommentIdsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBanUsersFromComment

> GetBannedUsersFromCommentResponse GetBanUsersFromComment(ctx, commentId).TenantId(tenantId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetBanUsersFromComment(context.Background(), commentId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetBanUsersFromComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBanUsersFromComment`: GetBannedUsersFromCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetBanUsersFromComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBanUsersFromCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **sso** | **string** |  | 

### Return type

[**GetBannedUsersFromCommentResponse**](GetBannedUsersFromCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentBanStatus

> GetCommentBanStatusResponse GetCommentBanStatus(ctx, commentId).TenantId(tenantId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetCommentBanStatus(context.Background(), commentId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetCommentBanStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentBanStatus`: GetCommentBanStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetCommentBanStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentBanStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **sso** | **string** |  | 

### Return type

[**GetCommentBanStatusResponse**](GetCommentBanStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentChildren

> ModerationAPIChildCommentsResponse GetCommentChildren(ctx, commentId).TenantId(tenantId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetCommentChildren(context.Background(), commentId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetCommentChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentChildren`: ModerationAPIChildCommentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetCommentChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **sso** | **string** |  | 

### Return type

[**ModerationAPIChildCommentsResponse**](ModerationAPIChildCommentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCount

> ModerationAPICountCommentsResponse GetCount(ctx).TenantId(tenantId).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filter(filter).SearchFilters(searchFilters).Demo(demo).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	textSearch := "textSearch_example" // string |  (optional)
	byIPFromComment := "byIPFromComment_example" // string |  (optional)
	filter := "filter_example" // string |  (optional)
	searchFilters := "searchFilters_example" // string |  (optional)
	demo := true // bool |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetCount(context.Background()).TenantId(tenantId).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filter(filter).SearchFilters(searchFilters).Demo(demo).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCount`: ModerationAPICountCommentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **textSearch** | **string** |  | 
 **byIPFromComment** | **string** |  | 
 **filter** | **string** |  | 
 **searchFilters** | **string** |  | 
 **demo** | **bool** |  | 
 **sso** | **string** |  | 

### Return type

[**ModerationAPICountCommentsResponse**](ModerationAPICountCommentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCounts

> GetBannedUsersCountResponse GetCounts(ctx).TenantId(tenantId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetCounts(context.Background()).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCounts`: GetBannedUsersCountResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetCounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetBannedUsersCountResponse**](GetBannedUsersCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogs

> ModerationAPIGetLogsResponse GetLogs(ctx, commentId).TenantId(tenantId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetLogs(context.Background(), commentId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogs`: ModerationAPIGetLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **sso** | **string** |  | 

### Return type

[**ModerationAPIGetLogsResponse**](ModerationAPIGetLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualBadges

> GetTenantManualBadgesResponse GetManualBadges(ctx).TenantId(tenantId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetManualBadges(context.Background()).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetManualBadges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualBadges`: GetTenantManualBadgesResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetManualBadges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManualBadgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetTenantManualBadgesResponse**](GetTenantManualBadgesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualBadgesForUser

> GetUserManualBadgesResponse GetManualBadgesForUser(ctx).TenantId(tenantId).BadgesUserId(badgesUserId).CommentId(commentId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	badgesUserId := "badgesUserId_example" // string |  (optional)
	commentId := "commentId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetManualBadgesForUser(context.Background()).TenantId(tenantId).BadgesUserId(badgesUserId).CommentId(commentId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetManualBadgesForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualBadgesForUser`: GetUserManualBadgesResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetManualBadgesForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManualBadgesForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **badgesUserId** | **string** |  | 
 **commentId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetUserManualBadgesResponse**](GetUserManualBadgesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModerationComment

> ModerationAPICommentResponse GetModerationComment(ctx, commentId).TenantId(tenantId).IncludeEmail(includeEmail).IncludeIP(includeIP).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	includeEmail := true // bool |  (optional)
	includeIP := true // bool |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetModerationComment(context.Background(), commentId).TenantId(tenantId).IncludeEmail(includeEmail).IncludeIP(includeIP).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetModerationComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModerationComment`: ModerationAPICommentResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetModerationComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModerationCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **includeEmail** | **bool** |  | 
 **includeIP** | **bool** |  | 
 **sso** | **string** |  | 

### Return type

[**ModerationAPICommentResponse**](ModerationAPICommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModerationCommentText

> GetCommentTextResponse GetModerationCommentText(ctx, commentId).TenantId(tenantId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetModerationCommentText(context.Background(), commentId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetModerationCommentText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModerationCommentText`: GetCommentTextResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetModerationCommentText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModerationCommentTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **sso** | **string** |  | 

### Return type

[**GetCommentTextResponse**](GetCommentTextResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPreBanSummary

> PreBanSummary GetPreBanSummary(ctx, commentId).TenantId(tenantId).IncludeByUserIdAndEmail(includeByUserIdAndEmail).IncludeByIP(includeByIP).IncludeByEmailDomain(includeByEmailDomain).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	includeByUserIdAndEmail := true // bool |  (optional)
	includeByIP := true // bool |  (optional)
	includeByEmailDomain := true // bool |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetPreBanSummary(context.Background(), commentId).TenantId(tenantId).IncludeByUserIdAndEmail(includeByUserIdAndEmail).IncludeByIP(includeByIP).IncludeByEmailDomain(includeByEmailDomain).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetPreBanSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPreBanSummary`: PreBanSummary
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetPreBanSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPreBanSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **includeByUserIdAndEmail** | **bool** |  | 
 **includeByIP** | **bool** |  | 
 **includeByEmailDomain** | **bool** |  | 
 **sso** | **string** |  | 

### Return type

[**PreBanSummary**](PreBanSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchCommentsSummary

> ModerationCommentSearchResponse GetSearchCommentsSummary(ctx).TenantId(tenantId).Value(value).Filters(filters).SearchFilters(searchFilters).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	value := "value_example" // string |  (optional)
	filters := "filters_example" // string |  (optional)
	searchFilters := "searchFilters_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetSearchCommentsSummary(context.Background()).TenantId(tenantId).Value(value).Filters(filters).SearchFilters(searchFilters).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetSearchCommentsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchCommentsSummary`: ModerationCommentSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetSearchCommentsSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchCommentsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **value** | **string** |  | 
 **filters** | **string** |  | 
 **searchFilters** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**ModerationCommentSearchResponse**](ModerationCommentSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchPages

> ModerationPageSearchResponse GetSearchPages(ctx).TenantId(tenantId).Value(value).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	value := "value_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetSearchPages(context.Background()).TenantId(tenantId).Value(value).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetSearchPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchPages`: ModerationPageSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetSearchPages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **value** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**ModerationPageSearchResponse**](ModerationPageSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchSites

> ModerationSiteSearchResponse GetSearchSites(ctx).TenantId(tenantId).Value(value).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	value := "value_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetSearchSites(context.Background()).TenantId(tenantId).Value(value).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetSearchSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchSites`: ModerationSiteSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetSearchSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **value** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**ModerationSiteSearchResponse**](ModerationSiteSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchSuggest

> ModerationSuggestResponse GetSearchSuggest(ctx).TenantId(tenantId).TextSearch(textSearch).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	textSearch := "textSearch_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetSearchSuggest(context.Background()).TenantId(tenantId).TextSearch(textSearch).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetSearchSuggest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchSuggest`: ModerationSuggestResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetSearchSuggest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchSuggestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **textSearch** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**ModerationSuggestResponse**](ModerationSuggestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchUsers

> ModerationUserSearchResponse GetSearchUsers(ctx).TenantId(tenantId).Value(value).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	value := "value_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetSearchUsers(context.Background()).TenantId(tenantId).Value(value).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetSearchUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchUsers`: ModerationUserSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetSearchUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **value** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**ModerationUserSearchResponse**](ModerationUserSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustFactor

> GetUserTrustFactorResponse GetTrustFactor(ctx).TenantId(tenantId).UserId(userId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	userId := "userId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetTrustFactor(context.Background()).TenantId(tenantId).UserId(userId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetTrustFactor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustFactor`: GetUserTrustFactorResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetTrustFactor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **userId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetUserTrustFactorResponse**](GetUserTrustFactorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBanPreference

> APIModerateGetUserBanPreferencesResponse GetUserBanPreference(ctx).TenantId(tenantId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetUserBanPreference(context.Background()).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetUserBanPreference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBanPreference`: APIModerateGetUserBanPreferencesResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetUserBanPreference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBanPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**APIModerateGetUserBanPreferencesResponse**](APIModerateGetUserBanPreferencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserInternalProfile

> GetUserInternalProfileResponse GetUserInternalProfile(ctx).TenantId(tenantId).CommentId(commentId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetUserInternalProfile(context.Background()).TenantId(tenantId).CommentId(commentId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetUserInternalProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInternalProfile`: GetUserInternalProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetUserInternalProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInternalProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **commentId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetUserInternalProfileResponse**](GetUserInternalProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdjustCommentVotes

> AdjustVotesResponse PostAdjustCommentVotes(ctx, commentId).TenantId(tenantId).AdjustCommentVotesParams(adjustCommentVotesParams).BroadcastId(broadcastId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	adjustCommentVotesParams := *openapiclient.NewAdjustCommentVotesParams(float64(123)) // AdjustCommentVotesParams | 
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostAdjustCommentVotes(context.Background(), commentId).TenantId(tenantId).AdjustCommentVotesParams(adjustCommentVotesParams).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostAdjustCommentVotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdjustCommentVotes`: AdjustVotesResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostAdjustCommentVotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdjustCommentVotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **adjustCommentVotesParams** | [**AdjustCommentVotesParams**](AdjustCommentVotesParams.md) |  | 
 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**AdjustVotesResponse**](AdjustVotesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiExport

> ModerationExportResponse PostApiExport(ctx).TenantId(tenantId).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filters(filters).SearchFilters(searchFilters).Sorts(sorts).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	textSearch := "textSearch_example" // string |  (optional)
	byIPFromComment := "byIPFromComment_example" // string |  (optional)
	filters := "filters_example" // string |  (optional)
	searchFilters := "searchFilters_example" // string |  (optional)
	sorts := "sorts_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostApiExport(context.Background()).TenantId(tenantId).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filters(filters).SearchFilters(searchFilters).Sorts(sorts).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostApiExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiExport`: ModerationExportResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostApiExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **textSearch** | **string** |  | 
 **byIPFromComment** | **string** |  | 
 **filters** | **string** |  | 
 **searchFilters** | **string** |  | 
 **sorts** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**ModerationExportResponse**](ModerationExportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBanUserFromComment

> BanUserFromCommentResult PostBanUserFromComment(ctx, commentId).TenantId(tenantId).BanEmail(banEmail).BanEmailDomain(banEmailDomain).BanIP(banIP).DeleteAllUsersComments(deleteAllUsersComments).BannedUntil(bannedUntil).IsShadowBan(isShadowBan).UpdateId(updateId).BanReason(banReason).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	banEmail := true // bool |  (optional)
	banEmailDomain := true // bool |  (optional)
	banIP := true // bool |  (optional)
	deleteAllUsersComments := true // bool |  (optional)
	bannedUntil := "bannedUntil_example" // string |  (optional)
	isShadowBan := true // bool |  (optional)
	updateId := "updateId_example" // string |  (optional)
	banReason := "banReason_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostBanUserFromComment(context.Background(), commentId).TenantId(tenantId).BanEmail(banEmail).BanEmailDomain(banEmailDomain).BanIP(banIP).DeleteAllUsersComments(deleteAllUsersComments).BannedUntil(bannedUntil).IsShadowBan(isShadowBan).UpdateId(updateId).BanReason(banReason).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostBanUserFromComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBanUserFromComment`: BanUserFromCommentResult
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostBanUserFromComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBanUserFromCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **banEmail** | **bool** |  | 
 **banEmailDomain** | **bool** |  | 
 **banIP** | **bool** |  | 
 **deleteAllUsersComments** | **bool** |  | 
 **bannedUntil** | **string** |  | 
 **isShadowBan** | **bool** |  | 
 **updateId** | **string** |  | 
 **banReason** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**BanUserFromCommentResult**](BanUserFromCommentResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBanUserUndo

> APIEmptyResponse PostBanUserUndo(ctx).TenantId(tenantId).BanUserUndoParams(banUserUndoParams).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	banUserUndoParams := *openapiclient.NewBanUserUndoParams(*openapiclient.NewAPIBanUserChangeLog()) // BanUserUndoParams | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostBanUserUndo(context.Background()).TenantId(tenantId).BanUserUndoParams(banUserUndoParams).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostBanUserUndo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBanUserUndo`: APIEmptyResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostBanUserUndo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBanUserUndoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **banUserUndoParams** | [**BanUserUndoParams**](BanUserUndoParams.md) |  | 
 **sso** | **string** |  | 

### Return type

[**APIEmptyResponse**](APIEmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBulkPreBanSummary

> BulkPreBanSummary PostBulkPreBanSummary(ctx).TenantId(tenantId).BulkPreBanParams(bulkPreBanParams).IncludeByUserIdAndEmail(includeByUserIdAndEmail).IncludeByIP(includeByIP).IncludeByEmailDomain(includeByEmailDomain).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	bulkPreBanParams := *openapiclient.NewBulkPreBanParams([]string{"CommentIds_example"}) // BulkPreBanParams | 
	includeByUserIdAndEmail := true // bool |  (optional)
	includeByIP := true // bool |  (optional)
	includeByEmailDomain := true // bool |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostBulkPreBanSummary(context.Background()).TenantId(tenantId).BulkPreBanParams(bulkPreBanParams).IncludeByUserIdAndEmail(includeByUserIdAndEmail).IncludeByIP(includeByIP).IncludeByEmailDomain(includeByEmailDomain).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostBulkPreBanSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBulkPreBanSummary`: BulkPreBanSummary
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostBulkPreBanSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBulkPreBanSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **bulkPreBanParams** | [**BulkPreBanParams**](BulkPreBanParams.md) |  | 
 **includeByUserIdAndEmail** | **bool** |  | 
 **includeByIP** | **bool** |  | 
 **includeByEmailDomain** | **bool** |  | 
 **sso** | **string** |  | 

### Return type

[**BulkPreBanSummary**](BulkPreBanSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommentsByIds

> ModerationAPIChildCommentsResponse PostCommentsByIds(ctx).TenantId(tenantId).CommentsByIdsParams(commentsByIdsParams).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentsByIdsParams := *openapiclient.NewCommentsByIdsParams([]string{"Ids_example"}) // CommentsByIdsParams | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostCommentsByIds(context.Background()).TenantId(tenantId).CommentsByIdsParams(commentsByIdsParams).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostCommentsByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCommentsByIds`: ModerationAPIChildCommentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostCommentsByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCommentsByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **commentsByIdsParams** | [**CommentsByIdsParams**](CommentsByIdsParams.md) |  | 
 **sso** | **string** |  | 

### Return type

[**ModerationAPIChildCommentsResponse**](ModerationAPIChildCommentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFlagComment

> APIEmptyResponse PostFlagComment(ctx, commentId).TenantId(tenantId).BroadcastId(broadcastId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostFlagComment(context.Background(), commentId).TenantId(tenantId).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostFlagComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFlagComment`: APIEmptyResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostFlagComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFlagCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**APIEmptyResponse**](APIEmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRemoveComment

> PostRemoveCommentApiResponse PostRemoveComment(ctx, commentId).TenantId(tenantId).BroadcastId(broadcastId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostRemoveComment(context.Background(), commentId).TenantId(tenantId).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostRemoveComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRemoveComment`: PostRemoveCommentApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostRemoveComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRemoveCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostRemoveCommentApiResponse**](PostRemoveCommentApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRestoreDeletedComment

> APIEmptyResponse PostRestoreDeletedComment(ctx, commentId).TenantId(tenantId).BroadcastId(broadcastId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostRestoreDeletedComment(context.Background(), commentId).TenantId(tenantId).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostRestoreDeletedComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRestoreDeletedComment`: APIEmptyResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostRestoreDeletedComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRestoreDeletedCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**APIEmptyResponse**](APIEmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSetCommentApprovalStatus

> SetCommentApprovedResponse PostSetCommentApprovalStatus(ctx, commentId).TenantId(tenantId).Approved(approved).BroadcastId(broadcastId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	approved := true // bool |  (optional)
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostSetCommentApprovalStatus(context.Background(), commentId).TenantId(tenantId).Approved(approved).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostSetCommentApprovalStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSetCommentApprovalStatus`: SetCommentApprovedResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostSetCommentApprovalStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSetCommentApprovalStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **approved** | **bool** |  | 
 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**SetCommentApprovedResponse**](SetCommentApprovedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSetCommentReviewStatus

> APIEmptyResponse PostSetCommentReviewStatus(ctx, commentId).TenantId(tenantId).Reviewed(reviewed).BroadcastId(broadcastId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	reviewed := true // bool |  (optional)
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostSetCommentReviewStatus(context.Background(), commentId).TenantId(tenantId).Reviewed(reviewed).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostSetCommentReviewStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSetCommentReviewStatus`: APIEmptyResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostSetCommentReviewStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSetCommentReviewStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **reviewed** | **bool** |  | 
 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**APIEmptyResponse**](APIEmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSetCommentSpamStatus

> APIEmptyResponse PostSetCommentSpamStatus(ctx, commentId).TenantId(tenantId).Spam(spam).PermNotSpam(permNotSpam).BroadcastId(broadcastId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	spam := true // bool |  (optional)
	permNotSpam := true // bool |  (optional)
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostSetCommentSpamStatus(context.Background(), commentId).TenantId(tenantId).Spam(spam).PermNotSpam(permNotSpam).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostSetCommentSpamStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSetCommentSpamStatus`: APIEmptyResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostSetCommentSpamStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSetCommentSpamStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **spam** | **bool** |  | 
 **permNotSpam** | **bool** |  | 
 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**APIEmptyResponse**](APIEmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSetCommentText

> SetCommentTextResponse PostSetCommentText(ctx, commentId).TenantId(tenantId).SetCommentTextParams(setCommentTextParams).BroadcastId(broadcastId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	setCommentTextParams := *openapiclient.NewSetCommentTextParams("Comment_example") // SetCommentTextParams | 
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostSetCommentText(context.Background(), commentId).TenantId(tenantId).SetCommentTextParams(setCommentTextParams).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostSetCommentText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSetCommentText`: SetCommentTextResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostSetCommentText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSetCommentTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **setCommentTextParams** | [**SetCommentTextParams**](SetCommentTextParams.md) |  | 
 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**SetCommentTextResponse**](SetCommentTextResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUnFlagComment

> APIEmptyResponse PostUnFlagComment(ctx, commentId).TenantId(tenantId).BroadcastId(broadcastId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostUnFlagComment(context.Background(), commentId).TenantId(tenantId).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostUnFlagComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUnFlagComment`: APIEmptyResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostUnFlagComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUnFlagCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**APIEmptyResponse**](APIEmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVote

> VoteResponse PostVote(ctx, commentId).TenantId(tenantId).Direction(direction).BroadcastId(broadcastId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	commentId := "commentId_example" // string | 
	direction := "direction_example" // string |  (optional)
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostVote(context.Background(), commentId).TenantId(tenantId).Direction(direction).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostVote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVote`: VoteResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostVote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **direction** | **string** |  | 
 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**VoteResponse**](VoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAwardBadge

> AwardUserBadgeResponse PutAwardBadge(ctx).TenantId(tenantId).BadgeId(badgeId).UserId(userId).CommentId(commentId).BroadcastId(broadcastId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	badgeId := "badgeId_example" // string | 
	userId := "userId_example" // string |  (optional)
	commentId := "commentId_example" // string |  (optional)
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PutAwardBadge(context.Background()).TenantId(tenantId).BadgeId(badgeId).UserId(userId).CommentId(commentId).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PutAwardBadge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAwardBadge`: AwardUserBadgeResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PutAwardBadge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAwardBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **badgeId** | **string** |  | 
 **userId** | **string** |  | 
 **commentId** | **string** |  | 
 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**AwardUserBadgeResponse**](AwardUserBadgeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCloseThread

> APIEmptyResponse PutCloseThread(ctx).TenantId(tenantId).UrlId(urlId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	urlId := "urlId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PutCloseThread(context.Background()).TenantId(tenantId).UrlId(urlId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PutCloseThread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCloseThread`: APIEmptyResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PutCloseThread`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCloseThreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **urlId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**APIEmptyResponse**](APIEmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRemoveBadge

> RemoveUserBadgeResponse PutRemoveBadge(ctx).TenantId(tenantId).BadgeId(badgeId).UserId(userId).CommentId(commentId).BroadcastId(broadcastId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	badgeId := "badgeId_example" // string | 
	userId := "userId_example" // string |  (optional)
	commentId := "commentId_example" // string |  (optional)
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PutRemoveBadge(context.Background()).TenantId(tenantId).BadgeId(badgeId).UserId(userId).CommentId(commentId).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PutRemoveBadge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRemoveBadge`: RemoveUserBadgeResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PutRemoveBadge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutRemoveBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **badgeId** | **string** |  | 
 **userId** | **string** |  | 
 **commentId** | **string** |  | 
 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**RemoveUserBadgeResponse**](RemoveUserBadgeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutReopenThread

> APIEmptyResponse PutReopenThread(ctx).TenantId(tenantId).UrlId(urlId).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	urlId := "urlId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PutReopenThread(context.Background()).TenantId(tenantId).UrlId(urlId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PutReopenThread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutReopenThread`: APIEmptyResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PutReopenThread`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutReopenThreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **urlId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**APIEmptyResponse**](APIEmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTrustFactor

> SetUserTrustFactorResponse SetTrustFactor(ctx).TenantId(tenantId).UserId(userId).TrustFactor(trustFactor).Sso(sso).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	userId := "userId_example" // string |  (optional)
	trustFactor := "trustFactor_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.SetTrustFactor(context.Background()).TenantId(tenantId).UserId(userId).TrustFactor(trustFactor).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.SetTrustFactor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTrustFactor`: SetUserTrustFactorResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.SetTrustFactor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTrustFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **userId** | **string** |  | 
 **trustFactor** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**SetUserTrustFactorResponse**](SetUserTrustFactorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

