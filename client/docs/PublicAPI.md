# \PublicAPI

All URIs are relative to *https://fastcomments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockFromCommentPublic**](PublicAPI.md#BlockFromCommentPublic) | **Post** /block-from-comment/{commentId} | 
[**CheckedCommentsForBlocked**](PublicAPI.md#CheckedCommentsForBlocked) | **Get** /check-blocked-comments | 
[**CreateCommentPublic**](PublicAPI.md#CreateCommentPublic) | **Post** /comments/{tenantId} | 
[**CreateFeedPostPublic**](PublicAPI.md#CreateFeedPostPublic) | **Post** /feed-posts/{tenantId} | 
[**CreateV1PageReact**](PublicAPI.md#CreateV1PageReact) | **Post** /page-reacts/v1/likes/{tenantId} | 
[**CreateV2PageReact**](PublicAPI.md#CreateV2PageReact) | **Post** /page-reacts/v2/{tenantId} | 
[**DeleteCommentPublic**](PublicAPI.md#DeleteCommentPublic) | **Delete** /comments/{tenantId}/{commentId} | 
[**DeleteCommentVote**](PublicAPI.md#DeleteCommentVote) | **Delete** /comments/{tenantId}/{commentId}/vote/{voteId} | 
[**DeleteFeedPostPublic**](PublicAPI.md#DeleteFeedPostPublic) | **Delete** /feed-posts/{tenantId}/{postId} | 
[**DeleteV1PageReact**](PublicAPI.md#DeleteV1PageReact) | **Delete** /page-reacts/v1/likes/{tenantId} | 
[**DeleteV2PageReact**](PublicAPI.md#DeleteV2PageReact) | **Delete** /page-reacts/v2/{tenantId} | 
[**FlagCommentPublic**](PublicAPI.md#FlagCommentPublic) | **Post** /flag-comment/{commentId} | 
[**GetCommentText**](PublicAPI.md#GetCommentText) | **Get** /comments/{tenantId}/{commentId}/text | 
[**GetCommentVoteUserNames**](PublicAPI.md#GetCommentVoteUserNames) | **Get** /comments/{tenantId}/{commentId}/votes | 
[**GetCommentsForUser**](PublicAPI.md#GetCommentsForUser) | **Get** /comments-for-user | 
[**GetCommentsPublic**](PublicAPI.md#GetCommentsPublic) | **Get** /comments/{tenantId} | 
[**GetEventLog**](PublicAPI.md#GetEventLog) | **Get** /event-log/{tenantId} | 
[**GetFeedPostsPublic**](PublicAPI.md#GetFeedPostsPublic) | **Get** /feed-posts/{tenantId} | 
[**GetFeedPostsStats**](PublicAPI.md#GetFeedPostsStats) | **Get** /feed-posts/{tenantId}/stats | 
[**GetGifLarge**](PublicAPI.md#GetGifLarge) | **Get** /gifs/get-large/{tenantId} | 
[**GetGifsSearch**](PublicAPI.md#GetGifsSearch) | **Get** /gifs/search/{tenantId} | 
[**GetGifsTrending**](PublicAPI.md#GetGifsTrending) | **Get** /gifs/trending/{tenantId} | 
[**GetGlobalEventLog**](PublicAPI.md#GetGlobalEventLog) | **Get** /event-log/global/{tenantId} | 
[**GetOfflineUsers**](PublicAPI.md#GetOfflineUsers) | **Get** /pages/{tenantId}/users/offline | 
[**GetOnlineUsers**](PublicAPI.md#GetOnlineUsers) | **Get** /pages/{tenantId}/users/online | 
[**GetPagesPublic**](PublicAPI.md#GetPagesPublic) | **Get** /pages/{tenantId} | 
[**GetTranslations**](PublicAPI.md#GetTranslations) | **Get** /translations/{namespace}/{component} | 
[**GetUserNotificationCount**](PublicAPI.md#GetUserNotificationCount) | **Get** /user-notifications/get-count | 
[**GetUserNotifications**](PublicAPI.md#GetUserNotifications) | **Get** /user-notifications | 
[**GetUserPresenceStatuses**](PublicAPI.md#GetUserPresenceStatuses) | **Get** /user-presence-status | 
[**GetUserReactsPublic**](PublicAPI.md#GetUserReactsPublic) | **Get** /feed-posts/{tenantId}/user-reacts | 
[**GetUsersInfo**](PublicAPI.md#GetUsersInfo) | **Get** /pages/{tenantId}/users/info | 
[**GetV1PageLikes**](PublicAPI.md#GetV1PageLikes) | **Get** /page-reacts/v1/likes/{tenantId} | 
[**GetV2PageReactUsers**](PublicAPI.md#GetV2PageReactUsers) | **Get** /page-reacts/v2/{tenantId}/list | 
[**GetV2PageReacts**](PublicAPI.md#GetV2PageReacts) | **Get** /page-reacts/v2/{tenantId} | 
[**LockComment**](PublicAPI.md#LockComment) | **Post** /comments/{tenantId}/{commentId}/lock | 
[**LogoutPublic**](PublicAPI.md#LogoutPublic) | **Put** /auth/logout | 
[**PinComment**](PublicAPI.md#PinComment) | **Post** /comments/{tenantId}/{commentId}/pin | 
[**ReactFeedPostPublic**](PublicAPI.md#ReactFeedPostPublic) | **Post** /feed-posts/{tenantId}/react/{postId} | 
[**ResetUserNotificationCount**](PublicAPI.md#ResetUserNotificationCount) | **Post** /user-notifications/reset-count | 
[**ResetUserNotifications**](PublicAPI.md#ResetUserNotifications) | **Post** /user-notifications/reset | 
[**SearchUsers**](PublicAPI.md#SearchUsers) | **Get** /user-search/{tenantId} | 
[**SetCommentText**](PublicAPI.md#SetCommentText) | **Post** /comments/{tenantId}/{commentId}/update-text | 
[**UnBlockCommentPublic**](PublicAPI.md#UnBlockCommentPublic) | **Delete** /block-from-comment/{commentId} | 
[**UnLockComment**](PublicAPI.md#UnLockComment) | **Post** /comments/{tenantId}/{commentId}/unlock | 
[**UnPinComment**](PublicAPI.md#UnPinComment) | **Post** /comments/{tenantId}/{commentId}/unpin | 
[**UpdateFeedPostPublic**](PublicAPI.md#UpdateFeedPostPublic) | **Put** /feed-posts/{tenantId}/{postId} | 
[**UpdateUserNotificationCommentSubscriptionStatus**](PublicAPI.md#UpdateUserNotificationCommentSubscriptionStatus) | **Post** /user-notifications/{notificationId}/mark-opted/{optedInOrOut} | 
[**UpdateUserNotificationPageSubscriptionStatus**](PublicAPI.md#UpdateUserNotificationPageSubscriptionStatus) | **Post** /user-notifications/set-subscription-state/{subscribedOrUnsubscribed} | 
[**UpdateUserNotificationStatus**](PublicAPI.md#UpdateUserNotificationStatus) | **Post** /user-notifications/{notificationId}/mark/{newStatus} | 
[**UploadImage**](PublicAPI.md#UploadImage) | **Post** /upload-image/{tenantId} | 
[**VoteComment**](PublicAPI.md#VoteComment) | **Post** /comments/{tenantId}/{commentId}/vote | 



## BlockFromCommentPublic

> BlockFromCommentPublicResponse BlockFromCommentPublic(ctx, commentId).TenantId(tenantId).PublicBlockFromCommentParams(publicBlockFromCommentParams).Sso(sso).Execute()



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
	publicBlockFromCommentParams := *openapiclient.NewPublicBlockFromCommentParams([]string{"CommentIds_example"}) // PublicBlockFromCommentParams | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.BlockFromCommentPublic(context.Background(), commentId).TenantId(tenantId).PublicBlockFromCommentParams(publicBlockFromCommentParams).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.BlockFromCommentPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockFromCommentPublic`: BlockFromCommentPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.BlockFromCommentPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlockFromCommentPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **publicBlockFromCommentParams** | [**PublicBlockFromCommentParams**](PublicBlockFromCommentParams.md) |  | 
 **sso** | **string** |  | 

### Return type

[**BlockFromCommentPublicResponse**](BlockFromCommentPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckedCommentsForBlocked

> CheckedCommentsForBlockedResponse CheckedCommentsForBlocked(ctx).TenantId(tenantId).CommentIds(commentIds).Sso(sso).Execute()



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
	commentIds := "commentIds_example" // string | A comma separated list of comment ids.
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.CheckedCommentsForBlocked(context.Background()).TenantId(tenantId).CommentIds(commentIds).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.CheckedCommentsForBlocked``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckedCommentsForBlocked`: CheckedCommentsForBlockedResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.CheckedCommentsForBlocked`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckedCommentsForBlockedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **commentIds** | **string** | A comma separated list of comment ids. | 
 **sso** | **string** |  | 

### Return type

[**CheckedCommentsForBlockedResponse**](CheckedCommentsForBlockedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCommentPublic

> CreateCommentPublicResponse CreateCommentPublic(ctx, tenantId).UrlId(urlId).BroadcastId(broadcastId).CommentData(commentData).SessionId(sessionId).Sso(sso).Execute()



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
	broadcastId := "broadcastId_example" // string | 
	commentData := *openapiclient.NewCommentData("CommenterName_example", "Comment_example", "Url_example", "UrlId_example") // CommentData | 
	sessionId := "sessionId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.CreateCommentPublic(context.Background(), tenantId).UrlId(urlId).BroadcastId(broadcastId).CommentData(commentData).SessionId(sessionId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.CreateCommentPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCommentPublic`: CreateCommentPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.CreateCommentPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlId** | **string** |  | 
 **broadcastId** | **string** |  | 
 **commentData** | [**CommentData**](CommentData.md) |  | 
 **sessionId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**CreateCommentPublicResponse**](CreateCommentPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFeedPostPublic

> CreateFeedPostPublicResponse CreateFeedPostPublic(ctx, tenantId).CreateFeedPostParams(createFeedPostParams).BroadcastId(broadcastId).Sso(sso).Execute()



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
	createFeedPostParams := *openapiclient.NewCreateFeedPostParams() // CreateFeedPostParams | 
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.CreateFeedPostPublic(context.Background(), tenantId).CreateFeedPostParams(createFeedPostParams).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.CreateFeedPostPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFeedPostPublic`: CreateFeedPostPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.CreateFeedPostPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeedPostPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFeedPostParams** | [**CreateFeedPostParams**](CreateFeedPostParams.md) |  | 
 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**CreateFeedPostPublicResponse**](CreateFeedPostPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateV1PageReact

> CreateV1PageReactResponse CreateV1PageReact(ctx, tenantId).UrlId(urlId).Title(title).Execute()



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
	title := "title_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.CreateV1PageReact(context.Background(), tenantId).UrlId(urlId).Title(title).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.CreateV1PageReact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateV1PageReact`: CreateV1PageReactResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.CreateV1PageReact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateV1PageReactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlId** | **string** |  | 
 **title** | **string** |  | 

### Return type

[**CreateV1PageReactResponse**](CreateV1PageReactResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateV2PageReact

> CreateV2PageReactResponse CreateV2PageReact(ctx, tenantId).UrlId(urlId).Id(id).Title(title).Execute()



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
	id := "id_example" // string | 
	title := "title_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.CreateV2PageReact(context.Background(), tenantId).UrlId(urlId).Id(id).Title(title).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.CreateV2PageReact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateV2PageReact`: CreateV2PageReactResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.CreateV2PageReact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateV2PageReactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlId** | **string** |  | 
 **id** | **string** |  | 
 **title** | **string** |  | 

### Return type

[**CreateV2PageReactResponse**](CreateV2PageReactResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommentPublic

> DeleteCommentPublicResponse DeleteCommentPublic(ctx, tenantId, commentId).BroadcastId(broadcastId).EditKey(editKey).Sso(sso).Execute()



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
	broadcastId := "broadcastId_example" // string | 
	editKey := "editKey_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.DeleteCommentPublic(context.Background(), tenantId, commentId).BroadcastId(broadcastId).EditKey(editKey).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.DeleteCommentPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCommentPublic`: DeleteCommentPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.DeleteCommentPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **broadcastId** | **string** |  | 
 **editKey** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**DeleteCommentPublicResponse**](DeleteCommentPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommentVote

> DeleteCommentVoteResponse DeleteCommentVote(ctx, tenantId, commentId, voteId).UrlId(urlId).BroadcastId(broadcastId).EditKey(editKey).Sso(sso).Execute()



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
	urlId := "urlId_example" // string | 
	broadcastId := "broadcastId_example" // string | 
	editKey := "editKey_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.DeleteCommentVote(context.Background(), tenantId, commentId, voteId).UrlId(urlId).BroadcastId(broadcastId).EditKey(editKey).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.DeleteCommentVote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCommentVote`: DeleteCommentVoteResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.DeleteCommentVote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**commentId** | **string** |  | 
**voteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentVoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **urlId** | **string** |  | 
 **broadcastId** | **string** |  | 
 **editKey** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**DeleteCommentVoteResponse**](DeleteCommentVoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFeedPostPublic

> DeleteFeedPostPublicResponse DeleteFeedPostPublic(ctx, tenantId, postId).BroadcastId(broadcastId).Sso(sso).Execute()



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
	postId := "postId_example" // string | 
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.DeleteFeedPostPublic(context.Background(), tenantId, postId).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.DeleteFeedPostPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFeedPostPublic`: DeleteFeedPostPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.DeleteFeedPostPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**postId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeedPostPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**DeleteFeedPostPublicResponse**](DeleteFeedPostPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteV1PageReact

> DeleteV1PageReactResponse DeleteV1PageReact(ctx, tenantId).UrlId(urlId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.DeleteV1PageReact(context.Background(), tenantId).UrlId(urlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.DeleteV1PageReact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1PageReact`: DeleteV1PageReactResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.DeleteV1PageReact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1PageReactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlId** | **string** |  | 

### Return type

[**DeleteV1PageReactResponse**](DeleteV1PageReactResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteV2PageReact

> DeleteV2PageReactResponse DeleteV2PageReact(ctx, tenantId).UrlId(urlId).Id(id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.DeleteV2PageReact(context.Background(), tenantId).UrlId(urlId).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.DeleteV2PageReact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV2PageReact`: DeleteV2PageReactResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.DeleteV2PageReact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV2PageReactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlId** | **string** |  | 
 **id** | **string** |  | 

### Return type

[**DeleteV2PageReactResponse**](DeleteV2PageReactResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlagCommentPublic

> FlagCommentPublicResponse FlagCommentPublic(ctx, commentId).TenantId(tenantId).IsFlagged(isFlagged).Sso(sso).Execute()



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
	isFlagged := true // bool | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.FlagCommentPublic(context.Background(), commentId).TenantId(tenantId).IsFlagged(isFlagged).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.FlagCommentPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlagCommentPublic`: FlagCommentPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.FlagCommentPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlagCommentPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **isFlagged** | **bool** |  | 
 **sso** | **string** |  | 

### Return type

[**FlagCommentPublicResponse**](FlagCommentPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentText

> GetCommentTextResponse1 GetCommentText(ctx, tenantId, commentId).EditKey(editKey).Sso(sso).Execute()



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
	editKey := "editKey_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetCommentText(context.Background(), tenantId, commentId).EditKey(editKey).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetCommentText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentText`: GetCommentTextResponse1
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetCommentText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **editKey** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetCommentTextResponse1**](GetCommentTextResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentVoteUserNames

> GetCommentVoteUserNamesResponse GetCommentVoteUserNames(ctx, tenantId, commentId).Dir(dir).Sso(sso).Execute()



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
	dir := int32(56) // int32 | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetCommentVoteUserNames(context.Background(), tenantId, commentId).Dir(dir).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetCommentVoteUserNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentVoteUserNames`: GetCommentVoteUserNamesResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetCommentVoteUserNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentVoteUserNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dir** | **int32** |  | 
 **sso** | **string** |  | 

### Return type

[**GetCommentVoteUserNamesResponse**](GetCommentVoteUserNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentsForUser

> GetCommentsForUserResponse1 GetCommentsForUser(ctx).UserId(userId).Direction(direction).RepliesToUserId(repliesToUserId).Page(page).Includei10n(includei10n).Locale(locale).IsCrawler(isCrawler).Execute()



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
	direction := openapiclient.SortDirections("OF") // SortDirections |  (optional)
	repliesToUserId := "repliesToUserId_example" // string |  (optional)
	page := float64(1.2) // float64 |  (optional)
	includei10n := true // bool |  (optional)
	locale := "locale_example" // string |  (optional)
	isCrawler := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetCommentsForUser(context.Background()).UserId(userId).Direction(direction).RepliesToUserId(repliesToUserId).Page(page).Includei10n(includei10n).Locale(locale).IsCrawler(isCrawler).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetCommentsForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentsForUser`: GetCommentsForUserResponse1
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetCommentsForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 
 **direction** | [**SortDirections**](SortDirections.md) |  | 
 **repliesToUserId** | **string** |  | 
 **page** | **float64** |  | 
 **includei10n** | **bool** |  | 
 **locale** | **string** |  | 
 **isCrawler** | **bool** |  | 

### Return type

[**GetCommentsForUserResponse1**](GetCommentsForUserResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentsPublic

> GetCommentsPublicResponse GetCommentsPublic(ctx, tenantId).UrlId(urlId).Page(page).Direction(direction).Sso(sso).Skip(skip).SkipChildren(skipChildren).Limit(limit).LimitChildren(limitChildren).CountChildren(countChildren).FetchPageForCommentId(fetchPageForCommentId).IncludeConfig(includeConfig).CountAll(countAll).Includei10n(includei10n).Locale(locale).Modules(modules).IsCrawler(isCrawler).IncludeNotificationCount(includeNotificationCount).AsTree(asTree).MaxTreeDepth(maxTreeDepth).UseFullTranslationIds(useFullTranslationIds).ParentId(parentId).SearchText(searchText).HashTags(hashTags).UserId(userId).CustomConfigStr(customConfigStr).AfterCommentId(afterCommentId).BeforeCommentId(beforeCommentId).Execute()





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
	page := int32(56) // int32 |  (optional)
	direction := openapiclient.SortDirections("OF") // SortDirections |  (optional)
	sso := "sso_example" // string |  (optional)
	skip := int32(56) // int32 |  (optional)
	skipChildren := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	limitChildren := int32(56) // int32 |  (optional)
	countChildren := true // bool |  (optional)
	fetchPageForCommentId := "fetchPageForCommentId_example" // string |  (optional)
	includeConfig := true // bool |  (optional)
	countAll := true // bool |  (optional)
	includei10n := true // bool |  (optional)
	locale := "locale_example" // string |  (optional)
	modules := "modules_example" // string |  (optional)
	isCrawler := true // bool |  (optional)
	includeNotificationCount := true // bool |  (optional)
	asTree := true // bool |  (optional)
	maxTreeDepth := int32(56) // int32 |  (optional)
	useFullTranslationIds := true // bool |  (optional)
	parentId := "parentId_example" // string |  (optional)
	searchText := "searchText_example" // string |  (optional)
	hashTags := []string{"Inner_example"} // []string |  (optional)
	userId := "userId_example" // string |  (optional)
	customConfigStr := "customConfigStr_example" // string |  (optional)
	afterCommentId := "afterCommentId_example" // string |  (optional)
	beforeCommentId := "beforeCommentId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetCommentsPublic(context.Background(), tenantId).UrlId(urlId).Page(page).Direction(direction).Sso(sso).Skip(skip).SkipChildren(skipChildren).Limit(limit).LimitChildren(limitChildren).CountChildren(countChildren).FetchPageForCommentId(fetchPageForCommentId).IncludeConfig(includeConfig).CountAll(countAll).Includei10n(includei10n).Locale(locale).Modules(modules).IsCrawler(isCrawler).IncludeNotificationCount(includeNotificationCount).AsTree(asTree).MaxTreeDepth(maxTreeDepth).UseFullTranslationIds(useFullTranslationIds).ParentId(parentId).SearchText(searchText).HashTags(hashTags).UserId(userId).CustomConfigStr(customConfigStr).AfterCommentId(afterCommentId).BeforeCommentId(beforeCommentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetCommentsPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentsPublic`: GetCommentsPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetCommentsPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentsPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlId** | **string** |  | 
 **page** | **int32** |  | 
 **direction** | [**SortDirections**](SortDirections.md) |  | 
 **sso** | **string** |  | 
 **skip** | **int32** |  | 
 **skipChildren** | **int32** |  | 
 **limit** | **int32** |  | 
 **limitChildren** | **int32** |  | 
 **countChildren** | **bool** |  | 
 **fetchPageForCommentId** | **string** |  | 
 **includeConfig** | **bool** |  | 
 **countAll** | **bool** |  | 
 **includei10n** | **bool** |  | 
 **locale** | **string** |  | 
 **modules** | **string** |  | 
 **isCrawler** | **bool** |  | 
 **includeNotificationCount** | **bool** |  | 
 **asTree** | **bool** |  | 
 **maxTreeDepth** | **int32** |  | 
 **useFullTranslationIds** | **bool** |  | 
 **parentId** | **string** |  | 
 **searchText** | **string** |  | 
 **hashTags** | **[]string** |  | 
 **userId** | **string** |  | 
 **customConfigStr** | **string** |  | 
 **afterCommentId** | **string** |  | 
 **beforeCommentId** | **string** |  | 

### Return type

[**GetCommentsPublicResponse**](GetCommentsPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventLog

> GetEventLogResponse1 GetEventLog(ctx, tenantId).UrlId(urlId).UserIdWS(userIdWS).StartTime(startTime).EndTime(endTime).Execute()





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
	userIdWS := "userIdWS_example" // string | 
	startTime := int64(789) // int64 | 
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetEventLog(context.Background(), tenantId).UrlId(urlId).UserIdWS(userIdWS).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetEventLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventLog`: GetEventLogResponse1
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetEventLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlId** | **string** |  | 
 **userIdWS** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**GetEventLogResponse1**](GetEventLogResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeedPostsPublic

> GetFeedPostsPublicResponse GetFeedPostsPublic(ctx, tenantId).AfterId(afterId).Limit(limit).Tags(tags).Sso(sso).IsCrawler(isCrawler).IncludeUserInfo(includeUserInfo).Execute()





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
	afterId := "afterId_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	tags := []string{"Inner_example"} // []string |  (optional)
	sso := "sso_example" // string |  (optional)
	isCrawler := true // bool |  (optional)
	includeUserInfo := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetFeedPostsPublic(context.Background(), tenantId).AfterId(afterId).Limit(limit).Tags(tags).Sso(sso).IsCrawler(isCrawler).IncludeUserInfo(includeUserInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetFeedPostsPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedPostsPublic`: GetFeedPostsPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetFeedPostsPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedPostsPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afterId** | **string** |  | 
 **limit** | **int32** |  | 
 **tags** | **[]string** |  | 
 **sso** | **string** |  | 
 **isCrawler** | **bool** |  | 
 **includeUserInfo** | **bool** |  | 

### Return type

[**GetFeedPostsPublicResponse**](GetFeedPostsPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeedPostsStats

> GetFeedPostsStatsResponse GetFeedPostsStats(ctx, tenantId).PostIds(postIds).Sso(sso).Execute()



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
	postIds := []string{"Inner_example"} // []string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetFeedPostsStats(context.Background(), tenantId).PostIds(postIds).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetFeedPostsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedPostsStats`: GetFeedPostsStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetFeedPostsStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedPostsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postIds** | **[]string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetFeedPostsStatsResponse**](GetFeedPostsStatsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGifLarge

> GetGifLargeResponse GetGifLarge(ctx, tenantId).LargeInternalURLSanitized(largeInternalURLSanitized).Execute()



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
	largeInternalURLSanitized := "largeInternalURLSanitized_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetGifLarge(context.Background(), tenantId).LargeInternalURLSanitized(largeInternalURLSanitized).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetGifLarge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGifLarge`: GetGifLargeResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetGifLarge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGifLargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **largeInternalURLSanitized** | **string** |  | 

### Return type

[**GetGifLargeResponse**](GetGifLargeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGifsSearch

> GetGifsSearchResponse GetGifsSearch(ctx, tenantId).Search(search).Locale(locale).Rating(rating).Page(page).Execute()



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
	search := "search_example" // string | 
	locale := "locale_example" // string |  (optional)
	rating := "rating_example" // string |  (optional)
	page := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetGifsSearch(context.Background(), tenantId).Search(search).Locale(locale).Rating(rating).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetGifsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGifsSearch`: GetGifsSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetGifsSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGifsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 
 **locale** | **string** |  | 
 **rating** | **string** |  | 
 **page** | **float64** |  | 

### Return type

[**GetGifsSearchResponse**](GetGifsSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGifsTrending

> GetGifsTrendingResponse GetGifsTrending(ctx, tenantId).Locale(locale).Rating(rating).Page(page).Execute()



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
	locale := "locale_example" // string |  (optional)
	rating := "rating_example" // string |  (optional)
	page := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetGifsTrending(context.Background(), tenantId).Locale(locale).Rating(rating).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetGifsTrending``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGifsTrending`: GetGifsTrendingResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetGifsTrending`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGifsTrendingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locale** | **string** |  | 
 **rating** | **string** |  | 
 **page** | **float64** |  | 

### Return type

[**GetGifsTrendingResponse**](GetGifsTrendingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalEventLog

> GetGlobalEventLogResponse GetGlobalEventLog(ctx, tenantId).UrlId(urlId).UserIdWS(userIdWS).StartTime(startTime).EndTime(endTime).Execute()





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
	userIdWS := "userIdWS_example" // string | 
	startTime := int64(789) // int64 | 
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetGlobalEventLog(context.Background(), tenantId).UrlId(urlId).UserIdWS(userIdWS).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetGlobalEventLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalEventLog`: GetGlobalEventLogResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetGlobalEventLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalEventLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlId** | **string** |  | 
 **userIdWS** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**GetGlobalEventLogResponse**](GetGlobalEventLogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOfflineUsers

> GetOfflineUsersResponse GetOfflineUsers(ctx, tenantId).UrlId(urlId).AfterName(afterName).AfterUserId(afterUserId).Execute()





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
	urlId := "urlId_example" // string | Page URL identifier (cleaned server-side).
	afterName := "afterName_example" // string | Cursor: pass nextAfterName from the previous response. (optional)
	afterUserId := "afterUserId_example" // string | Cursor tiebreaker: pass nextAfterUserId from the previous response. Required when afterName is set so name-ties don't drop entries. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetOfflineUsers(context.Background(), tenantId).UrlId(urlId).AfterName(afterName).AfterUserId(afterUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetOfflineUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOfflineUsers`: GetOfflineUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetOfflineUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfflineUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlId** | **string** | Page URL identifier (cleaned server-side). | 
 **afterName** | **string** | Cursor: pass nextAfterName from the previous response. | 
 **afterUserId** | **string** | Cursor tiebreaker: pass nextAfterUserId from the previous response. Required when afterName is set so name-ties don&#39;t drop entries. | 

### Return type

[**GetOfflineUsersResponse**](GetOfflineUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnlineUsers

> GetOnlineUsersResponse GetOnlineUsers(ctx, tenantId).UrlId(urlId).AfterName(afterName).AfterUserId(afterUserId).Execute()





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
	urlId := "urlId_example" // string | Page URL identifier (cleaned server-side).
	afterName := "afterName_example" // string | Cursor: pass nextAfterName from the previous response. (optional)
	afterUserId := "afterUserId_example" // string | Cursor tiebreaker: pass nextAfterUserId from the previous response. Required when afterName is set so name-ties don't drop entries. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetOnlineUsers(context.Background(), tenantId).UrlId(urlId).AfterName(afterName).AfterUserId(afterUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetOnlineUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnlineUsers`: GetOnlineUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetOnlineUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnlineUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlId** | **string** | Page URL identifier (cleaned server-side). | 
 **afterName** | **string** | Cursor: pass nextAfterName from the previous response. | 
 **afterUserId** | **string** | Cursor tiebreaker: pass nextAfterUserId from the previous response. Required when afterName is set so name-ties don&#39;t drop entries. | 

### Return type

[**GetOnlineUsersResponse**](GetOnlineUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPublic

> GetPagesPublicResponse GetPagesPublic(ctx, tenantId).Cursor(cursor).Limit(limit).Q(q).SortBy(sortBy).HasComments(hasComments).Execute()





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
	cursor := "cursor_example" // string | Opaque pagination cursor returned as `nextCursor` from a prior request. Tied to the same `sortBy`. (optional)
	limit := int32(56) // int32 | 1..200, default 50 (optional)
	q := "q_example" // string | Optional case-insensitive title prefix filter. (optional)
	sortBy := openapiclient.PagesSortBy("updatedAt") // PagesSortBy | Sort order. `updatedAt` (default, newest first), `commentCount` (most comments first), or `title` (alphabetical). (optional)
	hasComments := true // bool | If true, only return pages with at least one comment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetPagesPublic(context.Background(), tenantId).Cursor(cursor).Limit(limit).Q(q).SortBy(sortBy).HasComments(hasComments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetPagesPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPublic`: GetPagesPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetPagesPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor returned as &#x60;nextCursor&#x60; from a prior request. Tied to the same &#x60;sortBy&#x60;. | 
 **limit** | **int32** | 1..200, default 50 | 
 **q** | **string** | Optional case-insensitive title prefix filter. | 
 **sortBy** | [**PagesSortBy**](PagesSortBy.md) | Sort order. &#x60;updatedAt&#x60; (default, newest first), &#x60;commentCount&#x60; (most comments first), or &#x60;title&#x60; (alphabetical). | 
 **hasComments** | **bool** | If true, only return pages with at least one comment. | 

### Return type

[**GetPagesPublicResponse**](GetPagesPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTranslations

> GetTranslationsResponse1 GetTranslations(ctx, namespace, component).Locale(locale).UseFullTranslationIds(useFullTranslationIds).Execute()



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
	namespace := "namespace_example" // string | 
	component := "component_example" // string | 
	locale := "locale_example" // string |  (optional)
	useFullTranslationIds := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetTranslations(context.Background(), namespace, component).Locale(locale).UseFullTranslationIds(useFullTranslationIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetTranslations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTranslations`: GetTranslationsResponse1
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetTranslations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**component** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTranslationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locale** | **string** |  | 
 **useFullTranslationIds** | **bool** |  | 

### Return type

[**GetTranslationsResponse1**](GetTranslationsResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserNotificationCount

> GetUserNotificationCountResponse1 GetUserNotificationCount(ctx).TenantId(tenantId).Sso(sso).Execute()



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
	resp, r, err := apiClient.PublicAPI.GetUserNotificationCount(context.Background()).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetUserNotificationCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserNotificationCount`: GetUserNotificationCountResponse1
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetUserNotificationCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserNotificationCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetUserNotificationCountResponse1**](GetUserNotificationCountResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserNotifications

> GetUserNotificationsResponse GetUserNotifications(ctx).TenantId(tenantId).UrlId(urlId).PageSize(pageSize).AfterId(afterId).IncludeContext(includeContext).AfterCreatedAt(afterCreatedAt).UnreadOnly(unreadOnly).DmOnly(dmOnly).NoDm(noDm).IncludeTranslations(includeTranslations).IncludeTenantNotifications(includeTenantNotifications).Sso(sso).Execute()



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
	urlId := "urlId_example" // string | Used to determine whether the current page is subscribed. (optional)
	pageSize := int32(56) // int32 |  (optional)
	afterId := "afterId_example" // string |  (optional)
	includeContext := true // bool |  (optional)
	afterCreatedAt := int64(789) // int64 |  (optional)
	unreadOnly := true // bool |  (optional)
	dmOnly := true // bool |  (optional)
	noDm := true // bool |  (optional)
	includeTranslations := true // bool |  (optional)
	includeTenantNotifications := true // bool |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetUserNotifications(context.Background()).TenantId(tenantId).UrlId(urlId).PageSize(pageSize).AfterId(afterId).IncludeContext(includeContext).AfterCreatedAt(afterCreatedAt).UnreadOnly(unreadOnly).DmOnly(dmOnly).NoDm(noDm).IncludeTranslations(includeTranslations).IncludeTenantNotifications(includeTenantNotifications).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetUserNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserNotifications`: GetUserNotificationsResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetUserNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **urlId** | **string** | Used to determine whether the current page is subscribed. | 
 **pageSize** | **int32** |  | 
 **afterId** | **string** |  | 
 **includeContext** | **bool** |  | 
 **afterCreatedAt** | **int64** |  | 
 **unreadOnly** | **bool** |  | 
 **dmOnly** | **bool** |  | 
 **noDm** | **bool** |  | 
 **includeTranslations** | **bool** |  | 
 **includeTenantNotifications** | **bool** |  | 
 **sso** | **string** |  | 

### Return type

[**GetUserNotificationsResponse**](GetUserNotificationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPresenceStatuses

> GetUserPresenceStatusesResponse1 GetUserPresenceStatuses(ctx).TenantId(tenantId).UrlIdWS(urlIdWS).UserIds(userIds).Execute()



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
	urlIdWS := "urlIdWS_example" // string | 
	userIds := "userIds_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetUserPresenceStatuses(context.Background()).TenantId(tenantId).UrlIdWS(urlIdWS).UserIds(userIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetUserPresenceStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPresenceStatuses`: GetUserPresenceStatusesResponse1
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetUserPresenceStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPresenceStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **urlIdWS** | **string** |  | 
 **userIds** | **string** |  | 

### Return type

[**GetUserPresenceStatusesResponse1**](GetUserPresenceStatusesResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserReactsPublic

> GetUserReactsPublicResponse GetUserReactsPublic(ctx, tenantId).PostIds(postIds).Sso(sso).Execute()



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
	postIds := []string{"Inner_example"} // []string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetUserReactsPublic(context.Background(), tenantId).PostIds(postIds).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetUserReactsPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserReactsPublic`: GetUserReactsPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetUserReactsPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserReactsPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postIds** | **[]string** |  | 
 **sso** | **string** |  | 

### Return type

[**GetUserReactsPublicResponse**](GetUserReactsPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersInfo

> GetUsersInfoResponse GetUsersInfo(ctx, tenantId).Ids(ids).Execute()





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
	ids := "ids_example" // string | Comma-delimited userIds.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetUsersInfo(context.Background(), tenantId).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetUsersInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersInfo`: GetUsersInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetUsersInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **string** | Comma-delimited userIds. | 

### Return type

[**GetUsersInfoResponse**](GetUsersInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1PageLikes

> GetV1PageLikesResponse GetV1PageLikes(ctx, tenantId).UrlId(urlId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetV1PageLikes(context.Background(), tenantId).UrlId(urlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetV1PageLikes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1PageLikes`: GetV1PageLikesResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetV1PageLikes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1PageLikesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlId** | **string** |  | 

### Return type

[**GetV1PageLikesResponse**](GetV1PageLikesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV2PageReactUsers

> GetV2PageReactUsersResponse1 GetV2PageReactUsers(ctx, tenantId).UrlId(urlId).Id(id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetV2PageReactUsers(context.Background(), tenantId).UrlId(urlId).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetV2PageReactUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV2PageReactUsers`: GetV2PageReactUsersResponse1
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetV2PageReactUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV2PageReactUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlId** | **string** |  | 
 **id** | **string** |  | 

### Return type

[**GetV2PageReactUsersResponse1**](GetV2PageReactUsersResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV2PageReacts

> GetV2PageReactsResponse GetV2PageReacts(ctx, tenantId).UrlId(urlId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.GetV2PageReacts(context.Background(), tenantId).UrlId(urlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetV2PageReacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV2PageReacts`: GetV2PageReactsResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetV2PageReacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV2PageReactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlId** | **string** |  | 

### Return type

[**GetV2PageReactsResponse**](GetV2PageReactsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockComment

> LockCommentResponse LockComment(ctx, tenantId, commentId).BroadcastId(broadcastId).Sso(sso).Execute()



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
	broadcastId := "broadcastId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.LockComment(context.Background(), tenantId, commentId).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.LockComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LockComment`: LockCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.LockComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**LockCommentResponse**](LockCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogoutPublic

> APIEmptyResponse LogoutPublic(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.LogoutPublic(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.LogoutPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogoutPublic`: APIEmptyResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.LogoutPublic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutPublicRequest struct via the builder pattern


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


## PinComment

> PinCommentResponse PinComment(ctx, tenantId, commentId).BroadcastId(broadcastId).Sso(sso).Execute()



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
	broadcastId := "broadcastId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.PinComment(context.Background(), tenantId, commentId).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.PinComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PinComment`: PinCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.PinComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPinCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**PinCommentResponse**](PinCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactFeedPostPublic

> ReactFeedPostPublicResponse ReactFeedPostPublic(ctx, tenantId, postId).ReactBodyParams(reactBodyParams).IsUndo(isUndo).BroadcastId(broadcastId).Sso(sso).Execute()



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
	postId := "postId_example" // string | 
	reactBodyParams := *openapiclient.NewReactBodyParams() // ReactBodyParams | 
	isUndo := true // bool |  (optional)
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.ReactFeedPostPublic(context.Background(), tenantId, postId).ReactBodyParams(reactBodyParams).IsUndo(isUndo).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.ReactFeedPostPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReactFeedPostPublic`: ReactFeedPostPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.ReactFeedPostPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**postId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactFeedPostPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reactBodyParams** | [**ReactBodyParams**](ReactBodyParams.md) |  | 
 **isUndo** | **bool** |  | 
 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**ReactFeedPostPublicResponse**](ReactFeedPostPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetUserNotificationCount

> ResetUserNotificationCountResponse ResetUserNotificationCount(ctx).TenantId(tenantId).Sso(sso).Execute()



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
	resp, r, err := apiClient.PublicAPI.ResetUserNotificationCount(context.Background()).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.ResetUserNotificationCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetUserNotificationCount`: ResetUserNotificationCountResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.ResetUserNotificationCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetUserNotificationCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**ResetUserNotificationCountResponse**](ResetUserNotificationCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetUserNotifications

> ResetUserNotificationsResponse1 ResetUserNotifications(ctx).TenantId(tenantId).AfterId(afterId).AfterCreatedAt(afterCreatedAt).UnreadOnly(unreadOnly).DmOnly(dmOnly).NoDm(noDm).Sso(sso).Execute()



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
	afterId := "afterId_example" // string |  (optional)
	afterCreatedAt := int64(789) // int64 |  (optional)
	unreadOnly := true // bool |  (optional)
	dmOnly := true // bool |  (optional)
	noDm := true // bool |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.ResetUserNotifications(context.Background()).TenantId(tenantId).AfterId(afterId).AfterCreatedAt(afterCreatedAt).UnreadOnly(unreadOnly).DmOnly(dmOnly).NoDm(noDm).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.ResetUserNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetUserNotifications`: ResetUserNotificationsResponse1
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.ResetUserNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetUserNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **afterId** | **string** |  | 
 **afterCreatedAt** | **int64** |  | 
 **unreadOnly** | **bool** |  | 
 **dmOnly** | **bool** |  | 
 **noDm** | **bool** |  | 
 **sso** | **string** |  | 

### Return type

[**ResetUserNotificationsResponse1**](ResetUserNotificationsResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUsers

> SearchUsersResponse1 SearchUsers(ctx, tenantId).UrlId(urlId).UsernameStartsWith(usernameStartsWith).MentionGroupIds(mentionGroupIds).Sso(sso).SearchSection(searchSection).Execute()



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
	usernameStartsWith := "usernameStartsWith_example" // string |  (optional)
	mentionGroupIds := []string{"Inner_example"} // []string |  (optional)
	sso := "sso_example" // string |  (optional)
	searchSection := "searchSection_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.SearchUsers(context.Background(), tenantId).UrlId(urlId).UsernameStartsWith(usernameStartsWith).MentionGroupIds(mentionGroupIds).Sso(sso).SearchSection(searchSection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.SearchUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUsers`: SearchUsersResponse1
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.SearchUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urlId** | **string** |  | 
 **usernameStartsWith** | **string** |  | 
 **mentionGroupIds** | **[]string** |  | 
 **sso** | **string** |  | 
 **searchSection** | **string** |  | 

### Return type

[**SearchUsersResponse1**](SearchUsersResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCommentText

> SetCommentTextResponse1 SetCommentText(ctx, tenantId, commentId).BroadcastId(broadcastId).CommentTextUpdateRequest(commentTextUpdateRequest).EditKey(editKey).Sso(sso).Execute()



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
	broadcastId := "broadcastId_example" // string | 
	commentTextUpdateRequest := *openapiclient.NewCommentTextUpdateRequest("Comment_example") // CommentTextUpdateRequest | 
	editKey := "editKey_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.SetCommentText(context.Background(), tenantId, commentId).BroadcastId(broadcastId).CommentTextUpdateRequest(commentTextUpdateRequest).EditKey(editKey).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.SetCommentText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetCommentText`: SetCommentTextResponse1
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.SetCommentText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCommentTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **broadcastId** | **string** |  | 
 **commentTextUpdateRequest** | [**CommentTextUpdateRequest**](CommentTextUpdateRequest.md) |  | 
 **editKey** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**SetCommentTextResponse1**](SetCommentTextResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnBlockCommentPublic

> UnBlockCommentPublicResponse UnBlockCommentPublic(ctx, commentId).TenantId(tenantId).PublicBlockFromCommentParams(publicBlockFromCommentParams).Sso(sso).Execute()



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
	publicBlockFromCommentParams := *openapiclient.NewPublicBlockFromCommentParams([]string{"CommentIds_example"}) // PublicBlockFromCommentParams | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.UnBlockCommentPublic(context.Background(), commentId).TenantId(tenantId).PublicBlockFromCommentParams(publicBlockFromCommentParams).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.UnBlockCommentPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnBlockCommentPublic`: UnBlockCommentPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.UnBlockCommentPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnBlockCommentPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **publicBlockFromCommentParams** | [**PublicBlockFromCommentParams**](PublicBlockFromCommentParams.md) |  | 
 **sso** | **string** |  | 

### Return type

[**UnBlockCommentPublicResponse**](UnBlockCommentPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnLockComment

> UnLockCommentResponse UnLockComment(ctx, tenantId, commentId).BroadcastId(broadcastId).Sso(sso).Execute()



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
	broadcastId := "broadcastId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.UnLockComment(context.Background(), tenantId, commentId).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.UnLockComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnLockComment`: UnLockCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.UnLockComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnLockCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**UnLockCommentResponse**](UnLockCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnPinComment

> UnPinCommentResponse UnPinComment(ctx, tenantId, commentId).BroadcastId(broadcastId).Sso(sso).Execute()



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
	broadcastId := "broadcastId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.UnPinComment(context.Background(), tenantId, commentId).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.UnPinComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnPinComment`: UnPinCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.UnPinComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnPinCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**UnPinCommentResponse**](UnPinCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeedPostPublic

> UpdateFeedPostPublicResponse UpdateFeedPostPublic(ctx, tenantId, postId).UpdateFeedPostParams(updateFeedPostParams).BroadcastId(broadcastId).Sso(sso).Execute()



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
	postId := "postId_example" // string | 
	updateFeedPostParams := *openapiclient.NewUpdateFeedPostParams() // UpdateFeedPostParams | 
	broadcastId := "broadcastId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.UpdateFeedPostPublic(context.Background(), tenantId, postId).UpdateFeedPostParams(updateFeedPostParams).BroadcastId(broadcastId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.UpdateFeedPostPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFeedPostPublic`: UpdateFeedPostPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.UpdateFeedPostPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**postId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeedPostPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFeedPostParams** | [**UpdateFeedPostParams**](UpdateFeedPostParams.md) |  | 
 **broadcastId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**UpdateFeedPostPublicResponse**](UpdateFeedPostPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserNotificationCommentSubscriptionStatus

> UpdateUserNotificationCommentSubscriptionStatusResponse UpdateUserNotificationCommentSubscriptionStatus(ctx, notificationId, optedInOrOut).TenantId(tenantId).CommentId(commentId).Sso(sso).Execute()





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
	notificationId := "notificationId_example" // string | 
	optedInOrOut := "optedInOrOut_example" // string | 
	commentId := "commentId_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.UpdateUserNotificationCommentSubscriptionStatus(context.Background(), notificationId, optedInOrOut).TenantId(tenantId).CommentId(commentId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.UpdateUserNotificationCommentSubscriptionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserNotificationCommentSubscriptionStatus`: UpdateUserNotificationCommentSubscriptionStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.UpdateUserNotificationCommentSubscriptionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** |  | 
**optedInOrOut** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserNotificationCommentSubscriptionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **commentId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**UpdateUserNotificationCommentSubscriptionStatusResponse**](UpdateUserNotificationCommentSubscriptionStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserNotificationPageSubscriptionStatus

> UpdateUserNotificationPageSubscriptionStatusResponse UpdateUserNotificationPageSubscriptionStatus(ctx, subscribedOrUnsubscribed).TenantId(tenantId).UrlId(urlId).Url(url).PageTitle(pageTitle).Sso(sso).Execute()





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
	url := "url_example" // string | 
	pageTitle := "pageTitle_example" // string | 
	subscribedOrUnsubscribed := "subscribedOrUnsubscribed_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.UpdateUserNotificationPageSubscriptionStatus(context.Background(), subscribedOrUnsubscribed).TenantId(tenantId).UrlId(urlId).Url(url).PageTitle(pageTitle).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.UpdateUserNotificationPageSubscriptionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserNotificationPageSubscriptionStatus`: UpdateUserNotificationPageSubscriptionStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.UpdateUserNotificationPageSubscriptionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscribedOrUnsubscribed** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserNotificationPageSubscriptionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **urlId** | **string** |  | 
 **url** | **string** |  | 
 **pageTitle** | **string** |  | 

 **sso** | **string** |  | 

### Return type

[**UpdateUserNotificationPageSubscriptionStatusResponse**](UpdateUserNotificationPageSubscriptionStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserNotificationStatus

> UpdateUserNotificationStatusResponse UpdateUserNotificationStatus(ctx, notificationId, newStatus).TenantId(tenantId).Sso(sso).Execute()



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
	notificationId := "notificationId_example" // string | 
	newStatus := "newStatus_example" // string | 
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.UpdateUserNotificationStatus(context.Background(), notificationId, newStatus).TenantId(tenantId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.UpdateUserNotificationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserNotificationStatus`: UpdateUserNotificationStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.UpdateUserNotificationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** |  | 
**newStatus** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserNotificationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **sso** | **string** |  | 

### Return type

[**UpdateUserNotificationStatusResponse**](UpdateUserNotificationStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadImage

> UploadImageResponse UploadImage(ctx, tenantId).File(file).SizePreset(sizePreset).UrlId(urlId).Execute()





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
	file := os.NewFile(1234, "some_file") // *os.File | 
	sizePreset := openapiclient.SizePreset("Default") // SizePreset | Size preset: \"Default\" (1000x1000px) or \"CrossPlatform\" (creates sizes for popular devices) (optional)
	urlId := "urlId_example" // string | Page id that upload is happening from, to configure (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.UploadImage(context.Background(), tenantId).File(file).SizePreset(sizePreset).UrlId(urlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.UploadImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadImage`: UploadImageResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.UploadImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 
 **sizePreset** | [**SizePreset**](SizePreset.md) | Size preset: \&quot;Default\&quot; (1000x1000px) or \&quot;CrossPlatform\&quot; (creates sizes for popular devices) | 
 **urlId** | **string** | Page id that upload is happening from, to configure | 

### Return type

[**UploadImageResponse**](UploadImageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoteComment

> VoteCommentResponse VoteComment(ctx, tenantId, commentId).UrlId(urlId).BroadcastId(broadcastId).VoteBodyParams(voteBodyParams).SessionId(sessionId).Sso(sso).Execute()



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
	urlId := "urlId_example" // string | 
	broadcastId := "broadcastId_example" // string | 
	voteBodyParams := *openapiclient.NewVoteBodyParams("CommenterEmail_example", "CommenterName_example", "VoteDir_example", "Url_example") // VoteBodyParams | 
	sessionId := "sessionId_example" // string |  (optional)
	sso := "sso_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.VoteComment(context.Background(), tenantId, commentId).UrlId(urlId).BroadcastId(broadcastId).VoteBodyParams(voteBodyParams).SessionId(sessionId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.VoteComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoteComment`: VoteCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.VoteComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoteCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **urlId** | **string** |  | 
 **broadcastId** | **string** |  | 
 **voteBodyParams** | [**VoteBodyParams**](VoteBodyParams.md) |  | 
 **sessionId** | **string** |  | 
 **sso** | **string** |  | 

### Return type

[**VoteCommentResponse**](VoteCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

