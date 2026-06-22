# \ModerationAPI

All URIs are relative to *https://fastcomments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteModerationVote**](ModerationAPI.md#DeleteModerationVote) | **Delete** /auth/my-account/moderate-comments/vote/{commentId}/{voteId} | 
[**GetApiComments**](ModerationAPI.md#GetApiComments) | **Get** /auth/my-account/moderate-comments/api/comments | 
[**GetApiExportStatus**](ModerationAPI.md#GetApiExportStatus) | **Get** /auth/my-account/moderate-comments/api/export/status | 
[**GetApiIds**](ModerationAPI.md#GetApiIds) | **Get** /auth/my-account/moderate-comments/api/ids | 
[**GetBanUsersFromComment**](ModerationAPI.md#GetBanUsersFromComment) | **Get** /auth/my-account/moderate-comments/ban-users/from-comment/{commentId} | 
[**GetCommentBanStatus**](ModerationAPI.md#GetCommentBanStatus) | **Get** /auth/my-account/moderate-comments/get-comment-ban-status/{commentId} | 
[**GetCommentChildren**](ModerationAPI.md#GetCommentChildren) | **Get** /auth/my-account/moderate-comments/comment-children/{commentId} | 
[**GetCount**](ModerationAPI.md#GetCount) | **Get** /auth/my-account/moderate-comments/count | 
[**GetCounts**](ModerationAPI.md#GetCounts) | **Get** /auth/my-account/moderate-comments/banned-users/counts | 
[**GetLogs**](ModerationAPI.md#GetLogs) | **Get** /auth/my-account/moderate-comments/logs/{commentId} | 
[**GetManualBadges**](ModerationAPI.md#GetManualBadges) | **Get** /auth/my-account/moderate-comments/get-manual-badges | 
[**GetManualBadgesForUser**](ModerationAPI.md#GetManualBadgesForUser) | **Get** /auth/my-account/moderate-comments/get-manual-badges-for-user | 
[**GetModerationComment**](ModerationAPI.md#GetModerationComment) | **Get** /auth/my-account/moderate-comments/comment/{commentId} | 
[**GetModerationCommentText**](ModerationAPI.md#GetModerationCommentText) | **Get** /auth/my-account/moderate-comments/get-comment-text/{commentId} | 
[**GetPreBanSummary**](ModerationAPI.md#GetPreBanSummary) | **Get** /auth/my-account/moderate-comments/pre-ban-summary/{commentId} | 
[**GetSearchCommentsSummary**](ModerationAPI.md#GetSearchCommentsSummary) | **Get** /auth/my-account/moderate-comments/search/comments/summary | 
[**GetSearchPages**](ModerationAPI.md#GetSearchPages) | **Get** /auth/my-account/moderate-comments/search/pages | 
[**GetSearchSites**](ModerationAPI.md#GetSearchSites) | **Get** /auth/my-account/moderate-comments/search/sites | 
[**GetSearchSuggest**](ModerationAPI.md#GetSearchSuggest) | **Get** /auth/my-account/moderate-comments/search/suggest | 
[**GetSearchUsers**](ModerationAPI.md#GetSearchUsers) | **Get** /auth/my-account/moderate-comments/search/users | 
[**GetTrustFactor**](ModerationAPI.md#GetTrustFactor) | **Get** /auth/my-account/moderate-comments/get-trust-factor | 
[**GetUserBanPreference**](ModerationAPI.md#GetUserBanPreference) | **Get** /auth/my-account/moderate-comments/user-ban-preference | 
[**GetUserInternalProfile**](ModerationAPI.md#GetUserInternalProfile) | **Get** /auth/my-account/moderate-comments/get-user-internal-profile | 
[**PostAdjustCommentVotes**](ModerationAPI.md#PostAdjustCommentVotes) | **Post** /auth/my-account/moderate-comments/adjust-comment-votes/{commentId} | 
[**PostApiExport**](ModerationAPI.md#PostApiExport) | **Post** /auth/my-account/moderate-comments/api/export | 
[**PostBanUserFromComment**](ModerationAPI.md#PostBanUserFromComment) | **Post** /auth/my-account/moderate-comments/ban-user/from-comment/{commentId} | 
[**PostBanUserUndo**](ModerationAPI.md#PostBanUserUndo) | **Post** /auth/my-account/moderate-comments/ban-user/undo | 
[**PostBulkPreBanSummary**](ModerationAPI.md#PostBulkPreBanSummary) | **Post** /auth/my-account/moderate-comments/bulk-pre-ban-summary | 
[**PostCommentsByIds**](ModerationAPI.md#PostCommentsByIds) | **Post** /auth/my-account/moderate-comments/comments-by-ids | 
[**PostFlagComment**](ModerationAPI.md#PostFlagComment) | **Post** /auth/my-account/moderate-comments/flag-comment/{commentId} | 
[**PostRemoveComment**](ModerationAPI.md#PostRemoveComment) | **Post** /auth/my-account/moderate-comments/remove-comment/{commentId} | 
[**PostRestoreDeletedComment**](ModerationAPI.md#PostRestoreDeletedComment) | **Post** /auth/my-account/moderate-comments/restore-deleted-comment/{commentId} | 
[**PostSetCommentApprovalStatus**](ModerationAPI.md#PostSetCommentApprovalStatus) | **Post** /auth/my-account/moderate-comments/set-comment-approval-status/{commentId} | 
[**PostSetCommentReviewStatus**](ModerationAPI.md#PostSetCommentReviewStatus) | **Post** /auth/my-account/moderate-comments/set-comment-review-status/{commentId} | 
[**PostSetCommentSpamStatus**](ModerationAPI.md#PostSetCommentSpamStatus) | **Post** /auth/my-account/moderate-comments/set-comment-spam-status/{commentId} | 
[**PostSetCommentText**](ModerationAPI.md#PostSetCommentText) | **Post** /auth/my-account/moderate-comments/set-comment-text/{commentId} | 
[**PostUnFlagComment**](ModerationAPI.md#PostUnFlagComment) | **Post** /auth/my-account/moderate-comments/un-flag-comment/{commentId} | 
[**PostVote**](ModerationAPI.md#PostVote) | **Post** /auth/my-account/moderate-comments/vote/{commentId} | 
[**PutAwardBadge**](ModerationAPI.md#PutAwardBadge) | **Put** /auth/my-account/moderate-comments/award-badge | 
[**PutCloseThread**](ModerationAPI.md#PutCloseThread) | **Put** /auth/my-account/moderate-comments/close-thread | 
[**PutRemoveBadge**](ModerationAPI.md#PutRemoveBadge) | **Put** /auth/my-account/moderate-comments/remove-badge | 
[**PutReopenThread**](ModerationAPI.md#PutReopenThread) | **Put** /auth/my-account/moderate-comments/reopen-thread | 
[**SetTrustFactor**](ModerationAPI.md#SetTrustFactor) | **Put** /auth/my-account/moderate-comments/set-trust-factor | 



## DeleteModerationVote

> DeleteModerationVoteResponse DeleteModerationVote(ctx, commentId, voteId).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	voteId := "voteId_example" // string | 
	broadcastId := "broadcastId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.DeleteModerationVote(context.Background(), commentId, voteId).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.DeleteModerationVote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteModerationVote`: DeleteModerationVoteResponse
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


 **broadcastId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**DeleteModerationVoteResponse**](DeleteModerationVoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiComments

> GetApiCommentsResponse GetApiComments(ctx).Page(page).Count(count).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filters(filters).SearchFilters(searchFilters).Sorts(sorts).Demo(demo).TenantId(tenantId).Sso(sso).Execute()



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
	page := float64(1.2) // float64 |  (optional)
	count := float64(1.2) // float64 |  (optional)
	textSearch := "textSearch_example" // string |  (optional)
	byIPFromComment := "byIPFromComment_example" // string |  (optional)
	filters := "filters_example" // string |  (optional)
	searchFilters := "searchFilters_example" // string |  (optional)
	sorts := "sorts_example" // string |  (optional)
	demo := true // bool |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetApiComments(context.Background()).Page(page).Count(count).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filters(filters).SearchFilters(searchFilters).Sorts(sorts).Demo(demo).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetApiComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiComments`: GetApiCommentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetApiComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float64** |  | 
 **count** | **float64** |  | 
 **textSearch** | **string** |  | 
 **byIPFromComment** | **string** |  | 
 **filters** | **string** |  | 
 **searchFilters** | **string** |  | 
 **sorts** | **string** |  | 
 **demo** | **bool** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetApiCommentsResponse**](GetApiCommentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiExportStatus

> GetApiExportStatusResponse GetApiExportStatus(ctx).BatchJobId(batchJobId).TenantId(tenantId).Sso(sso).Execute()



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
	batchJobId := "batchJobId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetApiExportStatus(context.Background()).BatchJobId(batchJobId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetApiExportStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiExportStatus`: GetApiExportStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetApiExportStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiExportStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchJobId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetApiExportStatusResponse**](GetApiExportStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiIds

> GetApiIdsResponse GetApiIds(ctx).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filters(filters).SearchFilters(searchFilters).AfterId(afterId).Demo(demo).TenantId(tenantId).Sso(sso).Execute()



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
	textSearch := "textSearch_example" // string |  (optional)
	byIPFromComment := "byIPFromComment_example" // string |  (optional)
	filters := "filters_example" // string |  (optional)
	searchFilters := "searchFilters_example" // string |  (optional)
	afterId := "afterId_example" // string |  (optional)
	demo := true // bool |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetApiIds(context.Background()).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filters(filters).SearchFilters(searchFilters).AfterId(afterId).Demo(demo).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetApiIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiIds`: GetApiIdsResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetApiIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textSearch** | **string** |  | 
 **byIPFromComment** | **string** |  | 
 **filters** | **string** |  | 
 **searchFilters** | **string** |  | 
 **afterId** | **string** |  | 
 **demo** | **bool** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetApiIdsResponse**](GetApiIdsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBanUsersFromComment

> GetBanUsersFromCommentResponse GetBanUsersFromComment(ctx, commentId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetBanUsersFromComment(context.Background(), commentId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetBanUsersFromComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBanUsersFromComment`: GetBanUsersFromCommentResponse
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

[**GetBanUsersFromCommentResponse**](GetBanUsersFromCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentBanStatus

> GetCommentBanStatusResponse1 GetCommentBanStatus(ctx, commentId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetCommentBanStatus(context.Background(), commentId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetCommentBanStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentBanStatus`: GetCommentBanStatusResponse1
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

[**GetCommentBanStatusResponse1**](GetCommentBanStatusResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentChildren

> GetCommentChildrenResponse GetCommentChildren(ctx, commentId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetCommentChildren(context.Background(), commentId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetCommentChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentChildren`: GetCommentChildrenResponse
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

[**GetCommentChildrenResponse**](GetCommentChildrenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCount

> GetCountResponse GetCount(ctx).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filter(filter).SearchFilters(searchFilters).Demo(demo).TenantId(tenantId).Sso(sso).Execute()



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
	textSearch := "textSearch_example" // string |  (optional)
	byIPFromComment := "byIPFromComment_example" // string |  (optional)
	filter := "filter_example" // string |  (optional)
	searchFilters := "searchFilters_example" // string |  (optional)
	demo := true // bool |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetCount(context.Background()).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filter(filter).SearchFilters(searchFilters).Demo(demo).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCount`: GetCountResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textSearch** | **string** |  | 
 **byIPFromComment** | **string** |  | 
 **filter** | **string** |  | 
 **searchFilters** | **string** |  | 
 **demo** | **bool** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetCountResponse**](GetCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCounts

> GetCountsResponse GetCounts(ctx).TenantId(tenantId).Sso(sso).Execute()



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
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetCounts(context.Background()).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCounts`: GetCountsResponse
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

[**GetCountsResponse**](GetCountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogs

> GetLogsResponse GetLogs(ctx, commentId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetLogs(context.Background(), commentId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogs`: GetLogsResponse
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

[**GetLogsResponse**](GetLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualBadges

> GetManualBadgesResponse GetManualBadges(ctx).TenantId(tenantId).Sso(sso).Execute()



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
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetManualBadges(context.Background()).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetManualBadges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualBadges`: GetManualBadgesResponse
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

[**GetManualBadgesResponse**](GetManualBadgesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualBadgesForUser

> GetManualBadgesForUserResponse GetManualBadgesForUser(ctx).BadgesUserId(badgesUserId).CommentId(commentId).TenantId(tenantId).Sso(sso).Execute()



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
	badgesUserId := "badgesUserId_example" // string |  (optional)
	commentId := "commentId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetManualBadgesForUser(context.Background()).BadgesUserId(badgesUserId).CommentId(commentId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetManualBadgesForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualBadgesForUser`: GetManualBadgesForUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetManualBadgesForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManualBadgesForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **badgesUserId** | **string** |  | 
 **commentId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetManualBadgesForUserResponse**](GetManualBadgesForUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModerationComment

> GetModerationCommentResponse GetModerationComment(ctx, commentId).IncludeEmail(includeEmail).IncludeIP(includeIP).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	includeEmail := true // bool |  (optional)
	includeIP := true // bool |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetModerationComment(context.Background(), commentId).IncludeEmail(includeEmail).IncludeIP(includeIP).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetModerationComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModerationComment`: GetModerationCommentResponse
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

 **includeEmail** | **bool** |  | 
 **includeIP** | **bool** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetModerationCommentResponse**](GetModerationCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModerationCommentText

> GetModerationCommentTextResponse GetModerationCommentText(ctx, commentId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetModerationCommentText(context.Background(), commentId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetModerationCommentText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModerationCommentText`: GetModerationCommentTextResponse
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

[**GetModerationCommentTextResponse**](GetModerationCommentTextResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPreBanSummary

> GetPreBanSummaryResponse GetPreBanSummary(ctx, commentId).IncludeByUserIdAndEmail(includeByUserIdAndEmail).IncludeByIP(includeByIP).IncludeByEmailDomain(includeByEmailDomain).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	includeByUserIdAndEmail := true // bool |  (optional)
	includeByIP := true // bool |  (optional)
	includeByEmailDomain := true // bool |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetPreBanSummary(context.Background(), commentId).IncludeByUserIdAndEmail(includeByUserIdAndEmail).IncludeByIP(includeByIP).IncludeByEmailDomain(includeByEmailDomain).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetPreBanSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPreBanSummary`: GetPreBanSummaryResponse
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

 **includeByUserIdAndEmail** | **bool** |  | 
 **includeByIP** | **bool** |  | 
 **includeByEmailDomain** | **bool** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetPreBanSummaryResponse**](GetPreBanSummaryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchCommentsSummary

> GetSearchCommentsSummaryResponse GetSearchCommentsSummary(ctx).Value(value).Filters(filters).SearchFilters(searchFilters).TenantId(tenantId).Sso(sso).Execute()



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
	value := "value_example" // string |  (optional)
	filters := "filters_example" // string |  (optional)
	searchFilters := "searchFilters_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetSearchCommentsSummary(context.Background()).Value(value).Filters(filters).SearchFilters(searchFilters).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetSearchCommentsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchCommentsSummary`: GetSearchCommentsSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetSearchCommentsSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchCommentsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **value** | **string** |  | 
 **filters** | **string** |  | 
 **searchFilters** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetSearchCommentsSummaryResponse**](GetSearchCommentsSummaryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchPages

> GetSearchPagesResponse GetSearchPages(ctx).Value(value).TenantId(tenantId).Sso(sso).Execute()



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
	value := "value_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetSearchPages(context.Background()).Value(value).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetSearchPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchPages`: GetSearchPagesResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetSearchPages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **value** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetSearchPagesResponse**](GetSearchPagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchSites

> GetSearchSitesResponse GetSearchSites(ctx).Value(value).TenantId(tenantId).Sso(sso).Execute()



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
	value := "value_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetSearchSites(context.Background()).Value(value).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetSearchSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchSites`: GetSearchSitesResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetSearchSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **value** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetSearchSitesResponse**](GetSearchSitesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchSuggest

> GetSearchSuggestResponse GetSearchSuggest(ctx).TextSearch(textSearch).TenantId(tenantId).Sso(sso).Execute()



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
	textSearch := "textSearch_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetSearchSuggest(context.Background()).TextSearch(textSearch).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetSearchSuggest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchSuggest`: GetSearchSuggestResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetSearchSuggest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchSuggestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textSearch** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetSearchSuggestResponse**](GetSearchSuggestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchUsers

> GetSearchUsersResponse GetSearchUsers(ctx).Value(value).TenantId(tenantId).Sso(sso).Execute()



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
	value := "value_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetSearchUsers(context.Background()).Value(value).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetSearchUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchUsers`: GetSearchUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetSearchUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **value** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetSearchUsersResponse**](GetSearchUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustFactor

> GetTrustFactorResponse GetTrustFactor(ctx).UserId(userId).TenantId(tenantId).Sso(sso).Execute()



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
	userId := "userId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetTrustFactor(context.Background()).UserId(userId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetTrustFactor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustFactor`: GetTrustFactorResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetTrustFactor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetTrustFactorResponse**](GetTrustFactorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBanPreference

> GetUserBanPreferenceResponse GetUserBanPreference(ctx).TenantId(tenantId).Sso(sso).Execute()



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
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetUserBanPreference(context.Background()).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetUserBanPreference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBanPreference`: GetUserBanPreferenceResponse
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

[**GetUserBanPreferenceResponse**](GetUserBanPreferenceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserInternalProfile

> GetUserInternalProfileResponse1 GetUserInternalProfile(ctx).CommentId(commentId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.GetUserInternalProfile(context.Background()).CommentId(commentId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.GetUserInternalProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInternalProfile`: GetUserInternalProfileResponse1
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.GetUserInternalProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInternalProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commentId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetUserInternalProfileResponse1**](GetUserInternalProfileResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdjustCommentVotes

> PostAdjustCommentVotesResponse PostAdjustCommentVotes(ctx, commentId).AdjustCommentVotesParams(adjustCommentVotesParams).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	adjustCommentVotesParams := *openapiclient.NewAdjustCommentVotesParams(float64(123)) // AdjustCommentVotesParams | 
	broadcastId := "broadcastId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostAdjustCommentVotes(context.Background(), commentId).AdjustCommentVotesParams(adjustCommentVotesParams).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostAdjustCommentVotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdjustCommentVotes`: PostAdjustCommentVotesResponse
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

 **adjustCommentVotesParams** | [**AdjustCommentVotesParams**](AdjustCommentVotesParams.md) |  | 
 **broadcastId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostAdjustCommentVotesResponse**](PostAdjustCommentVotesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiExport

> PostApiExportResponse PostApiExport(ctx).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filters(filters).SearchFilters(searchFilters).Sorts(sorts).TenantId(tenantId).Sso(sso).Execute()



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
	textSearch := "textSearch_example" // string |  (optional)
	byIPFromComment := "byIPFromComment_example" // string |  (optional)
	filters := "filters_example" // string |  (optional)
	searchFilters := "searchFilters_example" // string |  (optional)
	sorts := "sorts_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostApiExport(context.Background()).TextSearch(textSearch).ByIPFromComment(byIPFromComment).Filters(filters).SearchFilters(searchFilters).Sorts(sorts).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostApiExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiExport`: PostApiExportResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostApiExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textSearch** | **string** |  | 
 **byIPFromComment** | **string** |  | 
 **filters** | **string** |  | 
 **searchFilters** | **string** |  | 
 **sorts** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostApiExportResponse**](PostApiExportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBanUserFromComment

> PostBanUserFromCommentResponse PostBanUserFromComment(ctx, commentId).BanEmail(banEmail).BanEmailDomain(banEmailDomain).BanIP(banIP).DeleteAllUsersComments(deleteAllUsersComments).BannedUntil(bannedUntil).IsShadowBan(isShadowBan).UpdateId(updateId).BanReason(banReason).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	banEmail := true // bool |  (optional)
	banEmailDomain := true // bool |  (optional)
	banIP := true // bool |  (optional)
	deleteAllUsersComments := true // bool |  (optional)
	bannedUntil := "bannedUntil_example" // string |  (optional)
	isShadowBan := true // bool |  (optional)
	updateId := "updateId_example" // string |  (optional)
	banReason := "banReason_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostBanUserFromComment(context.Background(), commentId).BanEmail(banEmail).BanEmailDomain(banEmailDomain).BanIP(banIP).DeleteAllUsersComments(deleteAllUsersComments).BannedUntil(bannedUntil).IsShadowBan(isShadowBan).UpdateId(updateId).BanReason(banReason).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostBanUserFromComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBanUserFromComment`: PostBanUserFromCommentResponse
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

 **banEmail** | **bool** |  | 
 **banEmailDomain** | **bool** |  | 
 **banIP** | **bool** |  | 
 **deleteAllUsersComments** | **bool** |  | 
 **bannedUntil** | **string** |  | 
 **isShadowBan** | **bool** |  | 
 **updateId** | **string** |  | 
 **banReason** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostBanUserFromCommentResponse**](PostBanUserFromCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBanUserUndo

> PostBanUserUndoResponse PostBanUserUndo(ctx).BanUserUndoParams(banUserUndoParams).TenantId(tenantId).Sso(sso).Execute()



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
	banUserUndoParams := *openapiclient.NewBanUserUndoParams(*openapiclient.NewAPIBanUserChangeLog()) // BanUserUndoParams | 
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostBanUserUndo(context.Background()).BanUserUndoParams(banUserUndoParams).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostBanUserUndo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBanUserUndo`: PostBanUserUndoResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostBanUserUndo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBanUserUndoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **banUserUndoParams** | [**BanUserUndoParams**](BanUserUndoParams.md) |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostBanUserUndoResponse**](PostBanUserUndoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBulkPreBanSummary

> PostBulkPreBanSummaryResponse PostBulkPreBanSummary(ctx).BulkPreBanParams(bulkPreBanParams).IncludeByUserIdAndEmail(includeByUserIdAndEmail).IncludeByIP(includeByIP).IncludeByEmailDomain(includeByEmailDomain).TenantId(tenantId).Sso(sso).Execute()



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
	bulkPreBanParams := *openapiclient.NewBulkPreBanParams([]string{"CommentIds_example"}) // BulkPreBanParams | 
	includeByUserIdAndEmail := true // bool |  (optional)
	includeByIP := true // bool |  (optional)
	includeByEmailDomain := true // bool |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostBulkPreBanSummary(context.Background()).BulkPreBanParams(bulkPreBanParams).IncludeByUserIdAndEmail(includeByUserIdAndEmail).IncludeByIP(includeByIP).IncludeByEmailDomain(includeByEmailDomain).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostBulkPreBanSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBulkPreBanSummary`: PostBulkPreBanSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostBulkPreBanSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBulkPreBanSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkPreBanParams** | [**BulkPreBanParams**](BulkPreBanParams.md) |  | 
 **includeByUserIdAndEmail** | **bool** |  | 
 **includeByIP** | **bool** |  | 
 **includeByEmailDomain** | **bool** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostBulkPreBanSummaryResponse**](PostBulkPreBanSummaryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommentsByIds

> PostCommentsByIdsResponse PostCommentsByIds(ctx).CommentsByIdsParams(commentsByIdsParams).TenantId(tenantId).Sso(sso).Execute()



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
	commentsByIdsParams := *openapiclient.NewCommentsByIdsParams([]string{"Ids_example"}) // CommentsByIdsParams | 
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostCommentsByIds(context.Background()).CommentsByIdsParams(commentsByIdsParams).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostCommentsByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCommentsByIds`: PostCommentsByIdsResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PostCommentsByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCommentsByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commentsByIdsParams** | [**CommentsByIdsParams**](CommentsByIdsParams.md) |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostCommentsByIdsResponse**](PostCommentsByIdsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFlagComment

> PostFlagCommentResponse PostFlagComment(ctx, commentId).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	broadcastId := "broadcastId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostFlagComment(context.Background(), commentId).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostFlagComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFlagComment`: PostFlagCommentResponse
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

 **broadcastId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostFlagCommentResponse**](PostFlagCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRemoveComment

> PostRemoveCommentResponse PostRemoveComment(ctx, commentId).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	broadcastId := "broadcastId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostRemoveComment(context.Background(), commentId).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostRemoveComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRemoveComment`: PostRemoveCommentResponse
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

 **broadcastId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostRemoveCommentResponse**](PostRemoveCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRestoreDeletedComment

> PostRestoreDeletedCommentResponse PostRestoreDeletedComment(ctx, commentId).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	broadcastId := "broadcastId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostRestoreDeletedComment(context.Background(), commentId).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostRestoreDeletedComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRestoreDeletedComment`: PostRestoreDeletedCommentResponse
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

 **broadcastId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostRestoreDeletedCommentResponse**](PostRestoreDeletedCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSetCommentApprovalStatus

> PostSetCommentApprovalStatusResponse PostSetCommentApprovalStatus(ctx, commentId).Approved(approved).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	approved := true // bool |  (optional)
	broadcastId := "broadcastId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostSetCommentApprovalStatus(context.Background(), commentId).Approved(approved).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostSetCommentApprovalStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSetCommentApprovalStatus`: PostSetCommentApprovalStatusResponse
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

 **approved** | **bool** |  | 
 **broadcastId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostSetCommentApprovalStatusResponse**](PostSetCommentApprovalStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSetCommentReviewStatus

> PostSetCommentReviewStatusResponse PostSetCommentReviewStatus(ctx, commentId).Reviewed(reviewed).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	reviewed := true // bool |  (optional)
	broadcastId := "broadcastId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostSetCommentReviewStatus(context.Background(), commentId).Reviewed(reviewed).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostSetCommentReviewStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSetCommentReviewStatus`: PostSetCommentReviewStatusResponse
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

 **reviewed** | **bool** |  | 
 **broadcastId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostSetCommentReviewStatusResponse**](PostSetCommentReviewStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSetCommentSpamStatus

> PostSetCommentSpamStatusResponse PostSetCommentSpamStatus(ctx, commentId).Spam(spam).PermNotSpam(permNotSpam).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	spam := true // bool |  (optional)
	permNotSpam := true // bool |  (optional)
	broadcastId := "broadcastId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostSetCommentSpamStatus(context.Background(), commentId).Spam(spam).PermNotSpam(permNotSpam).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostSetCommentSpamStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSetCommentSpamStatus`: PostSetCommentSpamStatusResponse
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

 **spam** | **bool** |  | 
 **permNotSpam** | **bool** |  | 
 **broadcastId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostSetCommentSpamStatusResponse**](PostSetCommentSpamStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSetCommentText

> PostSetCommentTextResponse PostSetCommentText(ctx, commentId).SetCommentTextParams(setCommentTextParams).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	setCommentTextParams := *openapiclient.NewSetCommentTextParams("Comment_example") // SetCommentTextParams | 
	broadcastId := "broadcastId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostSetCommentText(context.Background(), commentId).SetCommentTextParams(setCommentTextParams).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostSetCommentText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSetCommentText`: PostSetCommentTextResponse
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

 **setCommentTextParams** | [**SetCommentTextParams**](SetCommentTextParams.md) |  | 
 **broadcastId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostSetCommentTextResponse**](PostSetCommentTextResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUnFlagComment

> PostUnFlagCommentResponse PostUnFlagComment(ctx, commentId).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	broadcastId := "broadcastId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostUnFlagComment(context.Background(), commentId).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostUnFlagComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUnFlagComment`: PostUnFlagCommentResponse
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

 **broadcastId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostUnFlagCommentResponse**](PostUnFlagCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVote

> PostVoteResponse PostVote(ctx, commentId).Direction(direction).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()



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
	commentId := "commentId_example" // string | 
	direction := "direction_example" // string |  (optional)
	broadcastId := "broadcastId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PostVote(context.Background(), commentId).Direction(direction).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PostVote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVote`: PostVoteResponse
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

 **direction** | **string** |  | 
 **broadcastId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PostVoteResponse**](PostVoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAwardBadge

> PutAwardBadgeResponse PutAwardBadge(ctx).BadgeId(badgeId).UserId(userId).CommentId(commentId).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()



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
	badgeId := "badgeId_example" // string | 
	userId := "userId_example" // string |  (optional)
	commentId := "commentId_example" // string |  (optional)
	broadcastId := "broadcastId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PutAwardBadge(context.Background()).BadgeId(badgeId).UserId(userId).CommentId(commentId).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PutAwardBadge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAwardBadge`: PutAwardBadgeResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PutAwardBadge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAwardBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **badgeId** | **string** |  | 
 **userId** | **string** |  | 
 **commentId** | **string** |  | 
 **broadcastId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PutAwardBadgeResponse**](PutAwardBadgeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCloseThread

> PutCloseThreadResponse PutCloseThread(ctx).UrlId(urlId).TenantId(tenantId).Sso(sso).Execute()



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
	urlId := "urlId_example" // string | 
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PutCloseThread(context.Background()).UrlId(urlId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PutCloseThread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCloseThread`: PutCloseThreadResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PutCloseThread`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCloseThreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **urlId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PutCloseThreadResponse**](PutCloseThreadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRemoveBadge

> PutRemoveBadgeResponse PutRemoveBadge(ctx).BadgeId(badgeId).UserId(userId).CommentId(commentId).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()



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
	badgeId := "badgeId_example" // string | 
	userId := "userId_example" // string |  (optional)
	commentId := "commentId_example" // string |  (optional)
	broadcastId := "broadcastId_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PutRemoveBadge(context.Background()).BadgeId(badgeId).UserId(userId).CommentId(commentId).BroadcastId(broadcastId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PutRemoveBadge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRemoveBadge`: PutRemoveBadgeResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PutRemoveBadge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutRemoveBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **badgeId** | **string** |  | 
 **userId** | **string** |  | 
 **commentId** | **string** |  | 
 **broadcastId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PutRemoveBadgeResponse**](PutRemoveBadgeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutReopenThread

> PutReopenThreadResponse PutReopenThread(ctx).UrlId(urlId).TenantId(tenantId).Sso(sso).Execute()



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
	urlId := "urlId_example" // string | 
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.PutReopenThread(context.Background()).UrlId(urlId).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.PutReopenThread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutReopenThread`: PutReopenThreadResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.PutReopenThread`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutReopenThreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **urlId** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PutReopenThreadResponse**](PutReopenThreadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTrustFactor

> SetTrustFactorResponse SetTrustFactor(ctx).UserId(userId).TrustFactor(trustFactor).TenantId(tenantId).Sso(sso).Execute()



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
	userId := "userId_example" // string |  (optional)
	trustFactor := "trustFactor_example" // string |  (optional)
	tenantId := "tenantId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationAPI.SetTrustFactor(context.Background()).UserId(userId).TrustFactor(trustFactor).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationAPI.SetTrustFactor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTrustFactor`: SetTrustFactorResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationAPI.SetTrustFactor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTrustFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 
 **trustFactor** | **string** |  | 
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**SetTrustFactorResponse**](SetTrustFactorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

