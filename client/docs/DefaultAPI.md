# \DefaultAPI

All URIs are relative to *https://fastcomments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDomainConfig**](DefaultAPI.md#AddDomainConfig) | **Post** /api/v1/domain-configs | 
[**AddHashTag**](DefaultAPI.md#AddHashTag) | **Post** /api/v1/hash-tags | 
[**AddHashTagsBulk**](DefaultAPI.md#AddHashTagsBulk) | **Post** /api/v1/hash-tags/bulk | 
[**AddPage**](DefaultAPI.md#AddPage) | **Post** /api/v1/pages | 
[**AddSSOUser**](DefaultAPI.md#AddSSOUser) | **Post** /api/v1/sso-users | 
[**Aggregate**](DefaultAPI.md#Aggregate) | **Post** /api/v1/aggregate | 
[**AggregateQuestionResults**](DefaultAPI.md#AggregateQuestionResults) | **Get** /api/v1/question-results-aggregation | 
[**BlockUserFromComment**](DefaultAPI.md#BlockUserFromComment) | **Post** /api/v1/comments/{id}/block | 
[**BulkAggregateQuestionResults**](DefaultAPI.md#BulkAggregateQuestionResults) | **Post** /api/v1/question-results-aggregation/bulk | 
[**ChangeTicketState**](DefaultAPI.md#ChangeTicketState) | **Patch** /api/v1/tickets/{id}/state | 
[**CombineCommentsWithQuestionResults**](DefaultAPI.md#CombineCommentsWithQuestionResults) | **Get** /api/v1/question-results-aggregation/combine/comments | 
[**CreateEmailTemplate**](DefaultAPI.md#CreateEmailTemplate) | **Post** /api/v1/email-templates | 
[**CreateFeedPost**](DefaultAPI.md#CreateFeedPost) | **Post** /api/v1/feed-posts | 
[**CreateModerator**](DefaultAPI.md#CreateModerator) | **Post** /api/v1/moderators | 
[**CreateQuestionConfig**](DefaultAPI.md#CreateQuestionConfig) | **Post** /api/v1/question-configs | 
[**CreateQuestionResult**](DefaultAPI.md#CreateQuestionResult) | **Post** /api/v1/question-results | 
[**CreateSubscription**](DefaultAPI.md#CreateSubscription) | **Post** /api/v1/subscriptions | 
[**CreateTenant**](DefaultAPI.md#CreateTenant) | **Post** /api/v1/tenants | 
[**CreateTenantPackage**](DefaultAPI.md#CreateTenantPackage) | **Post** /api/v1/tenant-packages | 
[**CreateTenantUser**](DefaultAPI.md#CreateTenantUser) | **Post** /api/v1/tenant-users | 
[**CreateTicket**](DefaultAPI.md#CreateTicket) | **Post** /api/v1/tickets | 
[**CreateUserBadge**](DefaultAPI.md#CreateUserBadge) | **Post** /api/v1/user-badges | 
[**CreateVote**](DefaultAPI.md#CreateVote) | **Post** /api/v1/votes | 
[**DeleteComment**](DefaultAPI.md#DeleteComment) | **Delete** /api/v1/comments/{id} | 
[**DeleteDomainConfig**](DefaultAPI.md#DeleteDomainConfig) | **Delete** /api/v1/domain-configs/{domain} | 
[**DeleteEmailTemplate**](DefaultAPI.md#DeleteEmailTemplate) | **Delete** /api/v1/email-templates/{id} | 
[**DeleteEmailTemplateRenderError**](DefaultAPI.md#DeleteEmailTemplateRenderError) | **Delete** /api/v1/email-templates/{id}/render-errors/{errorId} | 
[**DeleteHashTag**](DefaultAPI.md#DeleteHashTag) | **Delete** /api/v1/hash-tags/{tag} | 
[**DeleteModerator**](DefaultAPI.md#DeleteModerator) | **Delete** /api/v1/moderators/{id} | 
[**DeleteNotificationCount**](DefaultAPI.md#DeleteNotificationCount) | **Delete** /api/v1/notification-count/{id} | 
[**DeletePage**](DefaultAPI.md#DeletePage) | **Delete** /api/v1/pages/{id} | 
[**DeletePendingWebhookEvent**](DefaultAPI.md#DeletePendingWebhookEvent) | **Delete** /api/v1/pending-webhook-events/{id} | 
[**DeleteQuestionConfig**](DefaultAPI.md#DeleteQuestionConfig) | **Delete** /api/v1/question-configs/{id} | 
[**DeleteQuestionResult**](DefaultAPI.md#DeleteQuestionResult) | **Delete** /api/v1/question-results/{id} | 
[**DeleteSSOUser**](DefaultAPI.md#DeleteSSOUser) | **Delete** /api/v1/sso-users/{id} | 
[**DeleteSubscription**](DefaultAPI.md#DeleteSubscription) | **Delete** /api/v1/subscriptions/{id} | 
[**DeleteTenant**](DefaultAPI.md#DeleteTenant) | **Delete** /api/v1/tenants/{id} | 
[**DeleteTenantPackage**](DefaultAPI.md#DeleteTenantPackage) | **Delete** /api/v1/tenant-packages/{id} | 
[**DeleteTenantUser**](DefaultAPI.md#DeleteTenantUser) | **Delete** /api/v1/tenant-users/{id} | 
[**DeleteUserBadge**](DefaultAPI.md#DeleteUserBadge) | **Delete** /api/v1/user-badges/{id} | 
[**DeleteVote**](DefaultAPI.md#DeleteVote) | **Delete** /api/v1/votes/{id} | 
[**FlagComment**](DefaultAPI.md#FlagComment) | **Post** /api/v1/comments/{id}/flag | 
[**GetAuditLogs**](DefaultAPI.md#GetAuditLogs) | **Get** /api/v1/audit-logs | 
[**GetCachedNotificationCount**](DefaultAPI.md#GetCachedNotificationCount) | **Get** /api/v1/notification-count/{id} | 
[**GetComment**](DefaultAPI.md#GetComment) | **Get** /api/v1/comments/{id} | 
[**GetComments**](DefaultAPI.md#GetComments) | **Get** /api/v1/comments | 
[**GetDomainConfig**](DefaultAPI.md#GetDomainConfig) | **Get** /api/v1/domain-configs/{domain} | 
[**GetDomainConfigs**](DefaultAPI.md#GetDomainConfigs) | **Get** /api/v1/domain-configs | 
[**GetEmailTemplate**](DefaultAPI.md#GetEmailTemplate) | **Get** /api/v1/email-templates/{id} | 
[**GetEmailTemplateDefinitions**](DefaultAPI.md#GetEmailTemplateDefinitions) | **Get** /api/v1/email-templates/definitions | 
[**GetEmailTemplateRenderErrors**](DefaultAPI.md#GetEmailTemplateRenderErrors) | **Get** /api/v1/email-templates/{id}/render-errors | 
[**GetEmailTemplates**](DefaultAPI.md#GetEmailTemplates) | **Get** /api/v1/email-templates | 
[**GetFeedPosts**](DefaultAPI.md#GetFeedPosts) | **Get** /api/v1/feed-posts | 
[**GetHashTags**](DefaultAPI.md#GetHashTags) | **Get** /api/v1/hash-tags | 
[**GetModerator**](DefaultAPI.md#GetModerator) | **Get** /api/v1/moderators/{id} | 
[**GetModerators**](DefaultAPI.md#GetModerators) | **Get** /api/v1/moderators | 
[**GetNotificationCount**](DefaultAPI.md#GetNotificationCount) | **Get** /api/v1/notifications/count | 
[**GetNotifications**](DefaultAPI.md#GetNotifications) | **Get** /api/v1/notifications | 
[**GetPageByURLId**](DefaultAPI.md#GetPageByURLId) | **Get** /api/v1/pages/by-url-id | 
[**GetPages**](DefaultAPI.md#GetPages) | **Get** /api/v1/pages | 
[**GetPendingWebhookEventCount**](DefaultAPI.md#GetPendingWebhookEventCount) | **Get** /api/v1/pending-webhook-events/count | 
[**GetPendingWebhookEvents**](DefaultAPI.md#GetPendingWebhookEvents) | **Get** /api/v1/pending-webhook-events | 
[**GetQuestionConfig**](DefaultAPI.md#GetQuestionConfig) | **Get** /api/v1/question-configs/{id} | 
[**GetQuestionConfigs**](DefaultAPI.md#GetQuestionConfigs) | **Get** /api/v1/question-configs | 
[**GetQuestionResult**](DefaultAPI.md#GetQuestionResult) | **Get** /api/v1/question-results/{id} | 
[**GetQuestionResults**](DefaultAPI.md#GetQuestionResults) | **Get** /api/v1/question-results | 
[**GetSSOUserByEmail**](DefaultAPI.md#GetSSOUserByEmail) | **Get** /api/v1/sso-users/by-email/{email} | 
[**GetSSOUserById**](DefaultAPI.md#GetSSOUserById) | **Get** /api/v1/sso-users/by-id/{id} | 
[**GetSSOUsers**](DefaultAPI.md#GetSSOUsers) | **Get** /api/v1/sso-users | 
[**GetSubscriptions**](DefaultAPI.md#GetSubscriptions) | **Get** /api/v1/subscriptions | 
[**GetTenant**](DefaultAPI.md#GetTenant) | **Get** /api/v1/tenants/{id} | 
[**GetTenantDailyUsages**](DefaultAPI.md#GetTenantDailyUsages) | **Get** /api/v1/tenant-daily-usage | 
[**GetTenantPackage**](DefaultAPI.md#GetTenantPackage) | **Get** /api/v1/tenant-packages/{id} | 
[**GetTenantPackages**](DefaultAPI.md#GetTenantPackages) | **Get** /api/v1/tenant-packages | 
[**GetTenantUser**](DefaultAPI.md#GetTenantUser) | **Get** /api/v1/tenant-users/{id} | 
[**GetTenantUsers**](DefaultAPI.md#GetTenantUsers) | **Get** /api/v1/tenant-users | 
[**GetTenants**](DefaultAPI.md#GetTenants) | **Get** /api/v1/tenants | 
[**GetTicket**](DefaultAPI.md#GetTicket) | **Get** /api/v1/tickets/{id} | 
[**GetTickets**](DefaultAPI.md#GetTickets) | **Get** /api/v1/tickets | 
[**GetUser**](DefaultAPI.md#GetUser) | **Get** /api/v1/users/{id} | 
[**GetUserBadge**](DefaultAPI.md#GetUserBadge) | **Get** /api/v1/user-badges/{id} | 
[**GetUserBadgeProgressById**](DefaultAPI.md#GetUserBadgeProgressById) | **Get** /api/v1/user-badge-progress/{id} | 
[**GetUserBadgeProgressByUserId**](DefaultAPI.md#GetUserBadgeProgressByUserId) | **Get** /api/v1/user-badge-progress/user/{userId} | 
[**GetUserBadgeProgressList**](DefaultAPI.md#GetUserBadgeProgressList) | **Get** /api/v1/user-badge-progress | 
[**GetUserBadges**](DefaultAPI.md#GetUserBadges) | **Get** /api/v1/user-badges | 
[**GetVotes**](DefaultAPI.md#GetVotes) | **Get** /api/v1/votes | 
[**GetVotesForUser**](DefaultAPI.md#GetVotesForUser) | **Get** /api/v1/votes/for-user | 
[**PatchDomainConfig**](DefaultAPI.md#PatchDomainConfig) | **Patch** /api/v1/domain-configs/{domainToUpdate} | 
[**PatchHashTag**](DefaultAPI.md#PatchHashTag) | **Patch** /api/v1/hash-tags/{tag} | 
[**PatchPage**](DefaultAPI.md#PatchPage) | **Patch** /api/v1/pages/{id} | 
[**PatchSSOUser**](DefaultAPI.md#PatchSSOUser) | **Patch** /api/v1/sso-users/{id} | 
[**PutDomainConfig**](DefaultAPI.md#PutDomainConfig) | **Put** /api/v1/domain-configs/{domainToUpdate} | 
[**PutSSOUser**](DefaultAPI.md#PutSSOUser) | **Put** /api/v1/sso-users/{id} | 
[**RenderEmailTemplate**](DefaultAPI.md#RenderEmailTemplate) | **Post** /api/v1/email-templates/render | 
[**ReplaceTenantPackage**](DefaultAPI.md#ReplaceTenantPackage) | **Put** /api/v1/tenant-packages/{id} | 
[**ReplaceTenantUser**](DefaultAPI.md#ReplaceTenantUser) | **Put** /api/v1/tenant-users/{id} | 
[**SaveComment**](DefaultAPI.md#SaveComment) | **Post** /api/v1/comments | 
[**SaveCommentsBulk**](DefaultAPI.md#SaveCommentsBulk) | **Post** /api/v1/comments/bulk | 
[**SendInvite**](DefaultAPI.md#SendInvite) | **Post** /api/v1/moderators/{id}/send-invite | 
[**SendLoginLink**](DefaultAPI.md#SendLoginLink) | **Post** /api/v1/tenant-users/{id}/send-login-link | 
[**UnBlockUserFromComment**](DefaultAPI.md#UnBlockUserFromComment) | **Post** /api/v1/comments/{id}/un-block | 
[**UnFlagComment**](DefaultAPI.md#UnFlagComment) | **Post** /api/v1/comments/{id}/un-flag | 
[**UpdateComment**](DefaultAPI.md#UpdateComment) | **Patch** /api/v1/comments/{id} | 
[**UpdateEmailTemplate**](DefaultAPI.md#UpdateEmailTemplate) | **Patch** /api/v1/email-templates/{id} | 
[**UpdateFeedPost**](DefaultAPI.md#UpdateFeedPost) | **Patch** /api/v1/feed-posts/{id} | 
[**UpdateModerator**](DefaultAPI.md#UpdateModerator) | **Patch** /api/v1/moderators/{id} | 
[**UpdateNotification**](DefaultAPI.md#UpdateNotification) | **Patch** /api/v1/notifications/{id} | 
[**UpdateQuestionConfig**](DefaultAPI.md#UpdateQuestionConfig) | **Patch** /api/v1/question-configs/{id} | 
[**UpdateQuestionResult**](DefaultAPI.md#UpdateQuestionResult) | **Patch** /api/v1/question-results/{id} | 
[**UpdateSubscription**](DefaultAPI.md#UpdateSubscription) | **Patch** /api/v1/subscriptions/{id} | 
[**UpdateTenant**](DefaultAPI.md#UpdateTenant) | **Patch** /api/v1/tenants/{id} | 
[**UpdateTenantPackage**](DefaultAPI.md#UpdateTenantPackage) | **Patch** /api/v1/tenant-packages/{id} | 
[**UpdateTenantUser**](DefaultAPI.md#UpdateTenantUser) | **Patch** /api/v1/tenant-users/{id} | 
[**UpdateUserBadge**](DefaultAPI.md#UpdateUserBadge) | **Put** /api/v1/user-badges/{id} | 



## AddDomainConfig

> AddDomainConfig200Response AddDomainConfig(ctx).TenantId(tenantId).AddDomainConfigParams(addDomainConfigParams).Execute()



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
	addDomainConfigParams := *openapiclient.NewAddDomainConfigParams("Domain_example") // AddDomainConfigParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddDomainConfig(context.Background()).TenantId(tenantId).AddDomainConfigParams(addDomainConfigParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddDomainConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDomainConfig`: AddDomainConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddDomainConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDomainConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **addDomainConfigParams** | [**AddDomainConfigParams**](AddDomainConfigParams.md) |  | 

### Return type

[**AddDomainConfig200Response**](AddDomainConfig200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddHashTag

> AddHashTag200Response AddHashTag(ctx).TenantId(tenantId).CreateHashTagBody(createHashTagBody).Execute()



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
	createHashTagBody := *openapiclient.NewCreateHashTagBody("Tag_example") // CreateHashTagBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddHashTag(context.Background()).TenantId(tenantId).CreateHashTagBody(createHashTagBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddHashTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddHashTag`: AddHashTag200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddHashTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddHashTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createHashTagBody** | [**CreateHashTagBody**](CreateHashTagBody.md) |  | 

### Return type

[**AddHashTag200Response**](AddHashTag200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddHashTagsBulk

> AddHashTagsBulk200Response AddHashTagsBulk(ctx).TenantId(tenantId).BulkCreateHashTagsBody(bulkCreateHashTagsBody).Execute()



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
	bulkCreateHashTagsBody := *openapiclient.NewBulkCreateHashTagsBody([]openapiclient.BulkCreateHashTagsBodyTagsInner{*openapiclient.NewBulkCreateHashTagsBodyTagsInner("Tag_example")}) // BulkCreateHashTagsBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddHashTagsBulk(context.Background()).TenantId(tenantId).BulkCreateHashTagsBody(bulkCreateHashTagsBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddHashTagsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddHashTagsBulk`: AddHashTagsBulk200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddHashTagsBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddHashTagsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **bulkCreateHashTagsBody** | [**BulkCreateHashTagsBody**](BulkCreateHashTagsBody.md) |  | 

### Return type

[**AddHashTagsBulk200Response**](AddHashTagsBulk200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPage

> AddPageAPIResponse AddPage(ctx).TenantId(tenantId).CreateAPIPageData(createAPIPageData).Execute()



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
	createAPIPageData := *openapiclient.NewCreateAPIPageData("Title_example", "Url_example", "UrlId_example") // CreateAPIPageData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddPage(context.Background()).TenantId(tenantId).CreateAPIPageData(createAPIPageData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPage`: AddPageAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createAPIPageData** | [**CreateAPIPageData**](CreateAPIPageData.md) |  | 

### Return type

[**AddPageAPIResponse**](AddPageAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSSOUser

> AddSSOUserAPIResponse AddSSOUser(ctx).TenantId(tenantId).CreateAPISSOUserData(createAPISSOUserData).Execute()



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
	createAPISSOUserData := *openapiclient.NewCreateAPISSOUserData("Email_example", "Username_example", "Id_example") // CreateAPISSOUserData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddSSOUser(context.Background()).TenantId(tenantId).CreateAPISSOUserData(createAPISSOUserData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddSSOUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSSOUser`: AddSSOUserAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddSSOUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSSOUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createAPISSOUserData** | [**CreateAPISSOUserData**](CreateAPISSOUserData.md) |  | 

### Return type

[**AddSSOUserAPIResponse**](AddSSOUserAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Aggregate

> AggregationResponse Aggregate(ctx).TenantId(tenantId).AggregationRequest(aggregationRequest).ParentTenantId(parentTenantId).IncludeStats(includeStats).Execute()





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
	aggregationRequest := *openapiclient.NewAggregationRequest("ResourceName_example", []openapiclient.AggregationOperation{*openapiclient.NewAggregationOperation("Field_example", openapiclient.AggregationOpType("sum"))}) // AggregationRequest | 
	parentTenantId := "parentTenantId_example" // string |  (optional)
	includeStats := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Aggregate(context.Background()).TenantId(tenantId).AggregationRequest(aggregationRequest).ParentTenantId(parentTenantId).IncludeStats(includeStats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Aggregate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Aggregate`: AggregationResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Aggregate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAggregateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **aggregationRequest** | [**AggregationRequest**](AggregationRequest.md) |  | 
 **parentTenantId** | **string** |  | 
 **includeStats** | **bool** |  | 

### Return type

[**AggregationResponse**](AggregationResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AggregateQuestionResults

> AggregateQuestionResults200Response AggregateQuestionResults(ctx).TenantId(tenantId).QuestionId(questionId).QuestionIds(questionIds).UrlId(urlId).TimeBucket(timeBucket).StartDate(startDate).ForceRecalculate(forceRecalculate).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	questionId := "questionId_example" // string |  (optional)
	questionIds := []string{"Inner_example"} // []string |  (optional)
	urlId := "urlId_example" // string |  (optional)
	timeBucket := openapiclient.AggregateTimeBucket("day") // AggregateTimeBucket |  (optional)
	startDate := time.Now() // time.Time |  (optional)
	forceRecalculate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AggregateQuestionResults(context.Background()).TenantId(tenantId).QuestionId(questionId).QuestionIds(questionIds).UrlId(urlId).TimeBucket(timeBucket).StartDate(startDate).ForceRecalculate(forceRecalculate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AggregateQuestionResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AggregateQuestionResults`: AggregateQuestionResults200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AggregateQuestionResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAggregateQuestionResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **questionId** | **string** |  | 
 **questionIds** | **[]string** |  | 
 **urlId** | **string** |  | 
 **timeBucket** | [**AggregateTimeBucket**](AggregateTimeBucket.md) |  | 
 **startDate** | **time.Time** |  | 
 **forceRecalculate** | **bool** |  | 

### Return type

[**AggregateQuestionResults200Response**](AggregateQuestionResults200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockUserFromComment

> BlockFromCommentPublic200Response BlockUserFromComment(ctx, id).TenantId(tenantId).BlockFromCommentParams(blockFromCommentParams).UserId(userId).AnonUserId(anonUserId).Execute()



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
	id := "id_example" // string | 
	blockFromCommentParams := *openapiclient.NewBlockFromCommentParams() // BlockFromCommentParams | 
	userId := "userId_example" // string |  (optional)
	anonUserId := "anonUserId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.BlockUserFromComment(context.Background(), id).TenantId(tenantId).BlockFromCommentParams(blockFromCommentParams).UserId(userId).AnonUserId(anonUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BlockUserFromComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockUserFromComment`: BlockFromCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BlockUserFromComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlockUserFromCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **blockFromCommentParams** | [**BlockFromCommentParams**](BlockFromCommentParams.md) |  | 
 **userId** | **string** |  | 
 **anonUserId** | **string** |  | 

### Return type

[**BlockFromCommentPublic200Response**](BlockFromCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkAggregateQuestionResults

> BulkAggregateQuestionResults200Response BulkAggregateQuestionResults(ctx).TenantId(tenantId).BulkAggregateQuestionResultsRequest(bulkAggregateQuestionResultsRequest).ForceRecalculate(forceRecalculate).Execute()



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
	bulkAggregateQuestionResultsRequest := *openapiclient.NewBulkAggregateQuestionResultsRequest([]openapiclient.BulkAggregateQuestionItem{*openapiclient.NewBulkAggregateQuestionItem("AggId_example")}) // BulkAggregateQuestionResultsRequest | 
	forceRecalculate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.BulkAggregateQuestionResults(context.Background()).TenantId(tenantId).BulkAggregateQuestionResultsRequest(bulkAggregateQuestionResultsRequest).ForceRecalculate(forceRecalculate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkAggregateQuestionResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkAggregateQuestionResults`: BulkAggregateQuestionResults200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BulkAggregateQuestionResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkAggregateQuestionResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **bulkAggregateQuestionResultsRequest** | [**BulkAggregateQuestionResultsRequest**](BulkAggregateQuestionResultsRequest.md) |  | 
 **forceRecalculate** | **bool** |  | 

### Return type

[**BulkAggregateQuestionResults200Response**](BulkAggregateQuestionResults200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeTicketState

> ChangeTicketState200Response ChangeTicketState(ctx, id).TenantId(tenantId).UserId(userId).ChangeTicketStateBody(changeTicketStateBody).Execute()



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
	userId := "userId_example" // string | 
	id := "id_example" // string | 
	changeTicketStateBody := *openapiclient.NewChangeTicketStateBody(int32(123)) // ChangeTicketStateBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ChangeTicketState(context.Background(), id).TenantId(tenantId).UserId(userId).ChangeTicketStateBody(changeTicketStateBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ChangeTicketState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeTicketState`: ChangeTicketState200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ChangeTicketState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeTicketStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **userId** | **string** |  | 

 **changeTicketStateBody** | [**ChangeTicketStateBody**](ChangeTicketStateBody.md) |  | 

### Return type

[**ChangeTicketState200Response**](ChangeTicketState200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CombineCommentsWithQuestionResults

> CombineCommentsWithQuestionResults200Response CombineCommentsWithQuestionResults(ctx).TenantId(tenantId).QuestionId(questionId).QuestionIds(questionIds).UrlId(urlId).StartDate(startDate).ForceRecalculate(forceRecalculate).MinValue(minValue).MaxValue(maxValue).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	questionId := "questionId_example" // string |  (optional)
	questionIds := []string{"Inner_example"} // []string |  (optional)
	urlId := "urlId_example" // string |  (optional)
	startDate := time.Now() // time.Time |  (optional)
	forceRecalculate := true // bool |  (optional)
	minValue := float64(1.2) // float64 |  (optional)
	maxValue := float64(1.2) // float64 |  (optional)
	limit := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CombineCommentsWithQuestionResults(context.Background()).TenantId(tenantId).QuestionId(questionId).QuestionIds(questionIds).UrlId(urlId).StartDate(startDate).ForceRecalculate(forceRecalculate).MinValue(minValue).MaxValue(maxValue).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CombineCommentsWithQuestionResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CombineCommentsWithQuestionResults`: CombineCommentsWithQuestionResults200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CombineCommentsWithQuestionResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCombineCommentsWithQuestionResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **questionId** | **string** |  | 
 **questionIds** | **[]string** |  | 
 **urlId** | **string** |  | 
 **startDate** | **time.Time** |  | 
 **forceRecalculate** | **bool** |  | 
 **minValue** | **float64** |  | 
 **maxValue** | **float64** |  | 
 **limit** | **float64** |  | 

### Return type

[**CombineCommentsWithQuestionResults200Response**](CombineCommentsWithQuestionResults200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEmailTemplate

> CreateEmailTemplate200Response CreateEmailTemplate(ctx).TenantId(tenantId).CreateEmailTemplateBody(createEmailTemplateBody).Execute()



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
	createEmailTemplateBody := *openapiclient.NewCreateEmailTemplateBody("EmailTemplateId_example", "DisplayName_example", "Ejs_example") // CreateEmailTemplateBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateEmailTemplate(context.Background()).TenantId(tenantId).CreateEmailTemplateBody(createEmailTemplateBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmailTemplate`: CreateEmailTemplate200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateEmailTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createEmailTemplateBody** | [**CreateEmailTemplateBody**](CreateEmailTemplateBody.md) |  | 

### Return type

[**CreateEmailTemplate200Response**](CreateEmailTemplate200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFeedPost

> CreateFeedPost200Response CreateFeedPost(ctx).TenantId(tenantId).CreateFeedPostParams(createFeedPostParams).BroadcastId(broadcastId).IsLive(isLive).DoSpamCheck(doSpamCheck).SkipDupCheck(skipDupCheck).Execute()



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
	isLive := true // bool |  (optional)
	doSpamCheck := true // bool |  (optional)
	skipDupCheck := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateFeedPost(context.Background()).TenantId(tenantId).CreateFeedPostParams(createFeedPostParams).BroadcastId(broadcastId).IsLive(isLive).DoSpamCheck(doSpamCheck).SkipDupCheck(skipDupCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateFeedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFeedPost`: CreateFeedPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateFeedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createFeedPostParams** | [**CreateFeedPostParams**](CreateFeedPostParams.md) |  | 
 **broadcastId** | **string** |  | 
 **isLive** | **bool** |  | 
 **doSpamCheck** | **bool** |  | 
 **skipDupCheck** | **bool** |  | 

### Return type

[**CreateFeedPost200Response**](CreateFeedPost200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateModerator

> CreateModerator200Response CreateModerator(ctx).TenantId(tenantId).CreateModeratorBody(createModeratorBody).Execute()



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
	createModeratorBody := *openapiclient.NewCreateModeratorBody("Name_example", "Email_example") // CreateModeratorBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateModerator(context.Background()).TenantId(tenantId).CreateModeratorBody(createModeratorBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateModerator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateModerator`: CreateModerator200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateModerator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateModeratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createModeratorBody** | [**CreateModeratorBody**](CreateModeratorBody.md) |  | 

### Return type

[**CreateModerator200Response**](CreateModerator200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuestionConfig

> CreateQuestionConfig200Response CreateQuestionConfig(ctx).TenantId(tenantId).CreateQuestionConfigBody(createQuestionConfigBody).Execute()



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
	createQuestionConfigBody := *openapiclient.NewCreateQuestionConfigBody("Name_example", "Question_example", "Type_example", float64(123)) // CreateQuestionConfigBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateQuestionConfig(context.Background()).TenantId(tenantId).CreateQuestionConfigBody(createQuestionConfigBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateQuestionConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQuestionConfig`: CreateQuestionConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateQuestionConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuestionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createQuestionConfigBody** | [**CreateQuestionConfigBody**](CreateQuestionConfigBody.md) |  | 

### Return type

[**CreateQuestionConfig200Response**](CreateQuestionConfig200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuestionResult

> CreateQuestionResult200Response CreateQuestionResult(ctx).TenantId(tenantId).CreateQuestionResultBody(createQuestionResultBody).Execute()



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
	createQuestionResultBody := *openapiclient.NewCreateQuestionResultBody("UrlId_example", float64(123), "QuestionId_example") // CreateQuestionResultBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateQuestionResult(context.Background()).TenantId(tenantId).CreateQuestionResultBody(createQuestionResultBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateQuestionResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQuestionResult`: CreateQuestionResult200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateQuestionResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuestionResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createQuestionResultBody** | [**CreateQuestionResultBody**](CreateQuestionResultBody.md) |  | 

### Return type

[**CreateQuestionResult200Response**](CreateQuestionResult200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscription

> CreateSubscriptionAPIResponse CreateSubscription(ctx).TenantId(tenantId).CreateAPIUserSubscriptionData(createAPIUserSubscriptionData).Execute()



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
	createAPIUserSubscriptionData := *openapiclient.NewCreateAPIUserSubscriptionData("UrlId_example") // CreateAPIUserSubscriptionData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateSubscription(context.Background()).TenantId(tenantId).CreateAPIUserSubscriptionData(createAPIUserSubscriptionData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubscription`: CreateSubscriptionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createAPIUserSubscriptionData** | [**CreateAPIUserSubscriptionData**](CreateAPIUserSubscriptionData.md) |  | 

### Return type

[**CreateSubscriptionAPIResponse**](CreateSubscriptionAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTenant

> CreateTenant200Response CreateTenant(ctx).TenantId(tenantId).CreateTenantBody(createTenantBody).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	createTenantBody := *openapiclient.NewCreateTenantBody("Name_example", []openapiclient.APIDomainConfiguration{*openapiclient.NewAPIDomainConfiguration("Id_example", "Domain_example", time.Now())}) // CreateTenantBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateTenant(context.Background()).TenantId(tenantId).CreateTenantBody(createTenantBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenant`: CreateTenant200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createTenantBody** | [**CreateTenantBody**](CreateTenantBody.md) |  | 

### Return type

[**CreateTenant200Response**](CreateTenant200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTenantPackage

> CreateTenantPackage200Response CreateTenantPackage(ctx).TenantId(tenantId).CreateTenantPackageBody(createTenantPackageBody).Execute()



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
	createTenantPackageBody := *openapiclient.NewCreateTenantPackageBody("Name_example", float64(123), float64(123), float64(123), float64(123), float64(123), float64(123), float64(123), float64(123), false, "ForWhoText_example", []string{"FeatureTaglines_example"}, false) // CreateTenantPackageBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateTenantPackage(context.Background()).TenantId(tenantId).CreateTenantPackageBody(createTenantPackageBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTenantPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenantPackage`: CreateTenantPackage200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTenantPackage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createTenantPackageBody** | [**CreateTenantPackageBody**](CreateTenantPackageBody.md) |  | 

### Return type

[**CreateTenantPackage200Response**](CreateTenantPackage200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTenantUser

> CreateTenantUser200Response CreateTenantUser(ctx).TenantId(tenantId).CreateTenantUserBody(createTenantUserBody).Execute()



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
	createTenantUserBody := *openapiclient.NewCreateTenantUserBody("Username_example", "Email_example") // CreateTenantUserBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateTenantUser(context.Background()).TenantId(tenantId).CreateTenantUserBody(createTenantUserBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTenantUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenantUser`: CreateTenantUser200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTenantUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createTenantUserBody** | [**CreateTenantUserBody**](CreateTenantUserBody.md) |  | 

### Return type

[**CreateTenantUser200Response**](CreateTenantUser200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTicket

> CreateTicket200Response CreateTicket(ctx).TenantId(tenantId).UserId(userId).CreateTicketBody(createTicketBody).Execute()



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
	userId := "userId_example" // string | 
	createTicketBody := *openapiclient.NewCreateTicketBody("Subject_example") // CreateTicketBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateTicket(context.Background()).TenantId(tenantId).UserId(userId).CreateTicketBody(createTicketBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTicket`: CreateTicket200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTicket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **userId** | **string** |  | 
 **createTicketBody** | [**CreateTicketBody**](CreateTicketBody.md) |  | 

### Return type

[**CreateTicket200Response**](CreateTicket200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserBadge

> CreateUserBadge200Response CreateUserBadge(ctx).TenantId(tenantId).CreateUserBadgeParams(createUserBadgeParams).Execute()



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
	createUserBadgeParams := *openapiclient.NewCreateUserBadgeParams("UserId_example", "BadgeId_example") // CreateUserBadgeParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateUserBadge(context.Background()).TenantId(tenantId).CreateUserBadgeParams(createUserBadgeParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateUserBadge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserBadge`: CreateUserBadge200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateUserBadge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createUserBadgeParams** | [**CreateUserBadgeParams**](CreateUserBadgeParams.md) |  | 

### Return type

[**CreateUserBadge200Response**](CreateUserBadge200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVote

> VoteComment200Response CreateVote(ctx).TenantId(tenantId).CommentId(commentId).Direction(direction).UserId(userId).AnonUserId(anonUserId).Execute()



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
	direction := "direction_example" // string | 
	userId := "userId_example" // string |  (optional)
	anonUserId := "anonUserId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateVote(context.Background()).TenantId(tenantId).CommentId(commentId).Direction(direction).UserId(userId).AnonUserId(anonUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateVote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVote`: VoteComment200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateVote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **commentId** | **string** |  | 
 **direction** | **string** |  | 
 **userId** | **string** |  | 
 **anonUserId** | **string** |  | 

### Return type

[**VoteComment200Response**](VoteComment200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComment

> DeleteComment200Response DeleteComment(ctx, id).TenantId(tenantId).ContextUserId(contextUserId).IsLive(isLive).Execute()



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
	id := "id_example" // string | 
	contextUserId := "contextUserId_example" // string |  (optional)
	isLive := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteComment(context.Background(), id).TenantId(tenantId).ContextUserId(contextUserId).IsLive(isLive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteComment`: DeleteComment200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **contextUserId** | **string** |  | 
 **isLive** | **bool** |  | 

### Return type

[**DeleteComment200Response**](DeleteComment200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDomainConfig

> DeleteDomainConfig200Response DeleteDomainConfig(ctx, domain).TenantId(tenantId).Execute()



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
	domain := "domain_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteDomainConfig(context.Background(), domain).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteDomainConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDomainConfig`: DeleteDomainConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteDomainConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**DeleteDomainConfig200Response**](DeleteDomainConfig200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmailTemplate

> FlagCommentPublic200Response DeleteEmailTemplate(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteEmailTemplate(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEmailTemplate`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmailTemplateRenderError

> FlagCommentPublic200Response DeleteEmailTemplateRenderError(ctx, id, errorId).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 
	errorId := "errorId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteEmailTemplateRenderError(context.Background(), id, errorId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteEmailTemplateRenderError``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEmailTemplateRenderError`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteEmailTemplateRenderError`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**errorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailTemplateRenderErrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHashTag

> FlagCommentPublic200Response DeleteHashTag(ctx, tag).TenantId(tenantId).DeleteHashTagRequest(deleteHashTagRequest).Execute()



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
	tag := "tag_example" // string | 
	tenantId := "tenantId_example" // string |  (optional)
	deleteHashTagRequest := *openapiclient.NewDeleteHashTagRequest() // DeleteHashTagRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteHashTag(context.Background(), tag).TenantId(tenantId).DeleteHashTagRequest(deleteHashTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteHashTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteHashTag`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteHashTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHashTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **deleteHashTagRequest** | [**DeleteHashTagRequest**](DeleteHashTagRequest.md) |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteModerator

> FlagCommentPublic200Response DeleteModerator(ctx, id).TenantId(tenantId).SendEmail(sendEmail).Execute()



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
	id := "id_example" // string | 
	sendEmail := "sendEmail_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteModerator(context.Background(), id).TenantId(tenantId).SendEmail(sendEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteModerator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteModerator`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteModerator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteModeratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **sendEmail** | **string** |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationCount

> FlagCommentPublic200Response DeleteNotificationCount(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteNotificationCount(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteNotificationCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotificationCount`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteNotificationCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePage

> DeletePageAPIResponse DeletePage(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeletePage(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeletePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePage`: DeletePageAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeletePage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**DeletePageAPIResponse**](DeletePageAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePendingWebhookEvent

> FlagCommentPublic200Response DeletePendingWebhookEvent(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeletePendingWebhookEvent(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeletePendingWebhookEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePendingWebhookEvent`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeletePendingWebhookEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePendingWebhookEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQuestionConfig

> FlagCommentPublic200Response DeleteQuestionConfig(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteQuestionConfig(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteQuestionConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteQuestionConfig`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteQuestionConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQuestionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQuestionResult

> FlagCommentPublic200Response DeleteQuestionResult(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteQuestionResult(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteQuestionResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteQuestionResult`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteQuestionResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQuestionResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSSOUser

> DeleteSSOUserAPIResponse DeleteSSOUser(ctx, id).TenantId(tenantId).DeleteComments(deleteComments).CommentDeleteMode(commentDeleteMode).Execute()



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
	id := "id_example" // string | 
	deleteComments := true // bool |  (optional)
	commentDeleteMode := "commentDeleteMode_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteSSOUser(context.Background(), id).TenantId(tenantId).DeleteComments(deleteComments).CommentDeleteMode(commentDeleteMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteSSOUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSSOUser`: DeleteSSOUserAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteSSOUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSOUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **deleteComments** | **bool** |  | 
 **commentDeleteMode** | **string** |  | 

### Return type

[**DeleteSSOUserAPIResponse**](DeleteSSOUserAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> DeleteSubscriptionAPIResponse DeleteSubscription(ctx, id).TenantId(tenantId).UserId(userId).Execute()



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
	id := "id_example" // string | 
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteSubscription(context.Background(), id).TenantId(tenantId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSubscription`: DeleteSubscriptionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **userId** | **string** |  | 

### Return type

[**DeleteSubscriptionAPIResponse**](DeleteSubscriptionAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenant

> FlagCommentPublic200Response DeleteTenant(ctx, id).TenantId(tenantId).Sure(sure).Execute()



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
	id := "id_example" // string | 
	sure := "sure_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteTenant(context.Background(), id).TenantId(tenantId).Sure(sure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenant`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **sure** | **string** |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenantPackage

> FlagCommentPublic200Response DeleteTenantPackage(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteTenantPackage(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteTenantPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenantPackage`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteTenantPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenantUser

> FlagCommentPublic200Response DeleteTenantUser(ctx, id).TenantId(tenantId).DeleteComments(deleteComments).CommentDeleteMode(commentDeleteMode).Execute()



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
	id := "id_example" // string | 
	deleteComments := "deleteComments_example" // string |  (optional)
	commentDeleteMode := "commentDeleteMode_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteTenantUser(context.Background(), id).TenantId(tenantId).DeleteComments(deleteComments).CommentDeleteMode(commentDeleteMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteTenantUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenantUser`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteTenantUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **deleteComments** | **string** |  | 
 **commentDeleteMode** | **string** |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserBadge

> UpdateUserBadge200Response DeleteUserBadge(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteUserBadge(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteUserBadge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserBadge`: UpdateUserBadge200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteUserBadge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**UpdateUserBadge200Response**](UpdateUserBadge200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVote

> DeleteCommentVote200Response DeleteVote(ctx, id).TenantId(tenantId).EditKey(editKey).Execute()



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
	id := "id_example" // string | 
	editKey := "editKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteVote(context.Background(), id).TenantId(tenantId).EditKey(editKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteVote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVote`: DeleteCommentVote200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteVote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **editKey** | **string** |  | 

### Return type

[**DeleteCommentVote200Response**](DeleteCommentVote200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlagComment

> FlagComment200Response FlagComment(ctx, id).TenantId(tenantId).UserId(userId).AnonUserId(anonUserId).Execute()



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
	id := "id_example" // string | 
	userId := "userId_example" // string |  (optional)
	anonUserId := "anonUserId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FlagComment(context.Background(), id).TenantId(tenantId).UserId(userId).AnonUserId(anonUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FlagComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlagComment`: FlagComment200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FlagComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlagCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **userId** | **string** |  | 
 **anonUserId** | **string** |  | 

### Return type

[**FlagComment200Response**](FlagComment200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditLogs

> GetAuditLogs200Response GetAuditLogs(ctx).TenantId(tenantId).Limit(limit).Skip(skip).Order(order).After(after).Before(before).Execute()



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
	limit := float64(1.2) // float64 |  (optional)
	skip := float64(1.2) // float64 |  (optional)
	order := openapiclient.SORT_DIR("ASC") // SORTDIR |  (optional)
	after := float64(1.2) // float64 |  (optional)
	before := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetAuditLogs(context.Background()).TenantId(tenantId).Limit(limit).Skip(skip).Order(order).After(after).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogs`: GetAuditLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **limit** | **float64** |  | 
 **skip** | **float64** |  | 
 **order** | [**SORTDIR**](SORTDIR.md) |  | 
 **after** | **float64** |  | 
 **before** | **float64** |  | 

### Return type

[**GetAuditLogs200Response**](GetAuditLogs200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCachedNotificationCount

> GetCachedNotificationCount200Response GetCachedNotificationCount(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCachedNotificationCount(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCachedNotificationCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCachedNotificationCount`: GetCachedNotificationCount200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCachedNotificationCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCachedNotificationCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetCachedNotificationCount200Response**](GetCachedNotificationCount200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComment

> GetComment200Response GetComment(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetComment(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComment`: GetComment200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetComment200Response**](GetComment200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComments

> GetComments200Response GetComments(ctx).TenantId(tenantId).Page(page).Limit(limit).Skip(skip).AsTree(asTree).SkipChildren(skipChildren).LimitChildren(limitChildren).MaxTreeDepth(maxTreeDepth).UrlId(urlId).UserId(userId).AnonUserId(anonUserId).ContextUserId(contextUserId).HashTag(hashTag).ParentId(parentId).Direction(direction).Execute()



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
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	skip := int32(56) // int32 |  (optional)
	asTree := true // bool |  (optional)
	skipChildren := int32(56) // int32 |  (optional)
	limitChildren := int32(56) // int32 |  (optional)
	maxTreeDepth := int32(56) // int32 |  (optional)
	urlId := "urlId_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	anonUserId := "anonUserId_example" // string |  (optional)
	contextUserId := "contextUserId_example" // string |  (optional)
	hashTag := "hashTag_example" // string |  (optional)
	parentId := "parentId_example" // string |  (optional)
	direction := openapiclient.SortDirections("OF") // SortDirections |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetComments(context.Background()).TenantId(tenantId).Page(page).Limit(limit).Skip(skip).AsTree(asTree).SkipChildren(skipChildren).LimitChildren(limitChildren).MaxTreeDepth(maxTreeDepth).UrlId(urlId).UserId(userId).AnonUserId(anonUserId).ContextUserId(contextUserId).HashTag(hashTag).ParentId(parentId).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComments`: GetComments200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **page** | **int32** |  | 
 **limit** | **int32** |  | 
 **skip** | **int32** |  | 
 **asTree** | **bool** |  | 
 **skipChildren** | **int32** |  | 
 **limitChildren** | **int32** |  | 
 **maxTreeDepth** | **int32** |  | 
 **urlId** | **string** |  | 
 **userId** | **string** |  | 
 **anonUserId** | **string** |  | 
 **contextUserId** | **string** |  | 
 **hashTag** | **string** |  | 
 **parentId** | **string** |  | 
 **direction** | [**SortDirections**](SortDirections.md) |  | 

### Return type

[**GetComments200Response**](GetComments200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainConfig

> GetDomainConfig200Response GetDomainConfig(ctx, domain).TenantId(tenantId).Execute()



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
	domain := "domain_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetDomainConfig(context.Background(), domain).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetDomainConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainConfig`: GetDomainConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetDomainConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetDomainConfig200Response**](GetDomainConfig200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainConfigs

> GetDomainConfigs200Response GetDomainConfigs(ctx).TenantId(tenantId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetDomainConfigs(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetDomainConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainConfigs`: GetDomainConfigs200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetDomainConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**GetDomainConfigs200Response**](GetDomainConfigs200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailTemplate

> GetEmailTemplate200Response GetEmailTemplate(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetEmailTemplate(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailTemplate`: GetEmailTemplate200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetEmailTemplate200Response**](GetEmailTemplate200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailTemplateDefinitions

> GetEmailTemplateDefinitions200Response GetEmailTemplateDefinitions(ctx).TenantId(tenantId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetEmailTemplateDefinitions(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEmailTemplateDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailTemplateDefinitions`: GetEmailTemplateDefinitions200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEmailTemplateDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailTemplateDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**GetEmailTemplateDefinitions200Response**](GetEmailTemplateDefinitions200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailTemplateRenderErrors

> GetEmailTemplateRenderErrors200Response GetEmailTemplateRenderErrors(ctx, id).TenantId(tenantId).Skip(skip).Execute()



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
	id := "id_example" // string | 
	skip := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetEmailTemplateRenderErrors(context.Background(), id).TenantId(tenantId).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEmailTemplateRenderErrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailTemplateRenderErrors`: GetEmailTemplateRenderErrors200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEmailTemplateRenderErrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailTemplateRenderErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **skip** | **float64** |  | 

### Return type

[**GetEmailTemplateRenderErrors200Response**](GetEmailTemplateRenderErrors200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailTemplates

> GetEmailTemplates200Response GetEmailTemplates(ctx).TenantId(tenantId).Skip(skip).Execute()



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
	skip := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetEmailTemplates(context.Background()).TenantId(tenantId).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEmailTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailTemplates`: GetEmailTemplates200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEmailTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **skip** | **float64** |  | 

### Return type

[**GetEmailTemplates200Response**](GetEmailTemplates200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeedPosts

> GetFeedPosts200Response GetFeedPosts(ctx).TenantId(tenantId).AfterId(afterId).Limit(limit).Tags(tags).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetFeedPosts(context.Background()).TenantId(tenantId).AfterId(afterId).Limit(limit).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetFeedPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedPosts`: GetFeedPosts200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetFeedPosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **afterId** | **string** |  | 
 **limit** | **int32** |  | 
 **tags** | **[]string** |  | 

### Return type

[**GetFeedPosts200Response**](GetFeedPosts200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHashTags

> GetHashTags200Response GetHashTags(ctx).TenantId(tenantId).Page(page).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetHashTags(context.Background()).TenantId(tenantId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetHashTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHashTags`: GetHashTags200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetHashTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHashTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **page** | **float64** |  | 

### Return type

[**GetHashTags200Response**](GetHashTags200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModerator

> GetModerator200Response GetModerator(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetModerator(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetModerator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModerator`: GetModerator200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetModerator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModeratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetModerator200Response**](GetModerator200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModerators

> GetModerators200Response GetModerators(ctx).TenantId(tenantId).Skip(skip).Execute()



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
	skip := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetModerators(context.Background()).TenantId(tenantId).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetModerators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModerators`: GetModerators200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetModerators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetModeratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **skip** | **float64** |  | 

### Return type

[**GetModerators200Response**](GetModerators200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationCount

> GetNotificationCount200Response GetNotificationCount(ctx).TenantId(tenantId).UserId(userId).UrlId(urlId).FromCommentId(fromCommentId).Viewed(viewed).Type_(type_).Execute()



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
	urlId := "urlId_example" // string |  (optional)
	fromCommentId := "fromCommentId_example" // string |  (optional)
	viewed := true // bool |  (optional)
	type_ := "type__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetNotificationCount(context.Background()).TenantId(tenantId).UserId(userId).UrlId(urlId).FromCommentId(fromCommentId).Viewed(viewed).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetNotificationCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationCount`: GetNotificationCount200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetNotificationCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **userId** | **string** |  | 
 **urlId** | **string** |  | 
 **fromCommentId** | **string** |  | 
 **viewed** | **bool** |  | 
 **type_** | **string** |  | 

### Return type

[**GetNotificationCount200Response**](GetNotificationCount200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotifications

> GetNotifications200Response GetNotifications(ctx).TenantId(tenantId).UserId(userId).UrlId(urlId).FromCommentId(fromCommentId).Viewed(viewed).Type_(type_).Skip(skip).Execute()



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
	urlId := "urlId_example" // string |  (optional)
	fromCommentId := "fromCommentId_example" // string |  (optional)
	viewed := true // bool |  (optional)
	type_ := "type__example" // string |  (optional)
	skip := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetNotifications(context.Background()).TenantId(tenantId).UserId(userId).UrlId(urlId).FromCommentId(fromCommentId).Viewed(viewed).Type_(type_).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotifications`: GetNotifications200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **userId** | **string** |  | 
 **urlId** | **string** |  | 
 **fromCommentId** | **string** |  | 
 **viewed** | **bool** |  | 
 **type_** | **string** |  | 
 **skip** | **float64** |  | 

### Return type

[**GetNotifications200Response**](GetNotifications200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageByURLId

> GetPageByURLIdAPIResponse GetPageByURLId(ctx).TenantId(tenantId).UrlId(urlId).Execute()



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
	resp, r, err := apiClient.DefaultAPI.GetPageByURLId(context.Background()).TenantId(tenantId).UrlId(urlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPageByURLId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageByURLId`: GetPageByURLIdAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPageByURLId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPageByURLIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **urlId** | **string** |  | 

### Return type

[**GetPageByURLIdAPIResponse**](GetPageByURLIdAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPages

> GetPagesAPIResponse GetPages(ctx).TenantId(tenantId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPages(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPages`: GetPagesAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**GetPagesAPIResponse**](GetPagesAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPendingWebhookEventCount

> GetPendingWebhookEventCount200Response GetPendingWebhookEventCount(ctx).TenantId(tenantId).CommentId(commentId).ExternalId(externalId).EventType(eventType).Type_(type_).Domain(domain).AttemptCountGT(attemptCountGT).Execute()



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
	externalId := "externalId_example" // string |  (optional)
	eventType := "eventType_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)
	domain := "domain_example" // string |  (optional)
	attemptCountGT := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPendingWebhookEventCount(context.Background()).TenantId(tenantId).CommentId(commentId).ExternalId(externalId).EventType(eventType).Type_(type_).Domain(domain).AttemptCountGT(attemptCountGT).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPendingWebhookEventCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPendingWebhookEventCount`: GetPendingWebhookEventCount200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPendingWebhookEventCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingWebhookEventCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **commentId** | **string** |  | 
 **externalId** | **string** |  | 
 **eventType** | **string** |  | 
 **type_** | **string** |  | 
 **domain** | **string** |  | 
 **attemptCountGT** | **float64** |  | 

### Return type

[**GetPendingWebhookEventCount200Response**](GetPendingWebhookEventCount200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPendingWebhookEvents

> GetPendingWebhookEvents200Response GetPendingWebhookEvents(ctx).TenantId(tenantId).CommentId(commentId).ExternalId(externalId).EventType(eventType).Type_(type_).Domain(domain).AttemptCountGT(attemptCountGT).Skip(skip).Execute()



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
	externalId := "externalId_example" // string |  (optional)
	eventType := "eventType_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)
	domain := "domain_example" // string |  (optional)
	attemptCountGT := float64(1.2) // float64 |  (optional)
	skip := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPendingWebhookEvents(context.Background()).TenantId(tenantId).CommentId(commentId).ExternalId(externalId).EventType(eventType).Type_(type_).Domain(domain).AttemptCountGT(attemptCountGT).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPendingWebhookEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPendingWebhookEvents`: GetPendingWebhookEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPendingWebhookEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingWebhookEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **commentId** | **string** |  | 
 **externalId** | **string** |  | 
 **eventType** | **string** |  | 
 **type_** | **string** |  | 
 **domain** | **string** |  | 
 **attemptCountGT** | **float64** |  | 
 **skip** | **float64** |  | 

### Return type

[**GetPendingWebhookEvents200Response**](GetPendingWebhookEvents200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuestionConfig

> GetQuestionConfig200Response GetQuestionConfig(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetQuestionConfig(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetQuestionConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuestionConfig`: GetQuestionConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetQuestionConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuestionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetQuestionConfig200Response**](GetQuestionConfig200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuestionConfigs

> GetQuestionConfigs200Response GetQuestionConfigs(ctx).TenantId(tenantId).Skip(skip).Execute()



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
	skip := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetQuestionConfigs(context.Background()).TenantId(tenantId).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetQuestionConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuestionConfigs`: GetQuestionConfigs200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetQuestionConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuestionConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **skip** | **float64** |  | 

### Return type

[**GetQuestionConfigs200Response**](GetQuestionConfigs200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuestionResult

> GetQuestionResult200Response GetQuestionResult(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetQuestionResult(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetQuestionResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuestionResult`: GetQuestionResult200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetQuestionResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuestionResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetQuestionResult200Response**](GetQuestionResult200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuestionResults

> GetQuestionResults200Response GetQuestionResults(ctx).TenantId(tenantId).UrlId(urlId).UserId(userId).StartDate(startDate).QuestionId(questionId).QuestionIds(questionIds).Skip(skip).Execute()



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
	urlId := "urlId_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	startDate := "startDate_example" // string |  (optional)
	questionId := "questionId_example" // string |  (optional)
	questionIds := "questionIds_example" // string |  (optional)
	skip := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetQuestionResults(context.Background()).TenantId(tenantId).UrlId(urlId).UserId(userId).StartDate(startDate).QuestionId(questionId).QuestionIds(questionIds).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetQuestionResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuestionResults`: GetQuestionResults200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetQuestionResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuestionResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **urlId** | **string** |  | 
 **userId** | **string** |  | 
 **startDate** | **string** |  | 
 **questionId** | **string** |  | 
 **questionIds** | **string** |  | 
 **skip** | **float64** |  | 

### Return type

[**GetQuestionResults200Response**](GetQuestionResults200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSOUserByEmail

> GetSSOUserByEmailAPIResponse GetSSOUserByEmail(ctx, email).TenantId(tenantId).Execute()



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
	email := "email_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSSOUserByEmail(context.Background(), email).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSSOUserByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSSOUserByEmail`: GetSSOUserByEmailAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSSOUserByEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSOUserByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetSSOUserByEmailAPIResponse**](GetSSOUserByEmailAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSOUserById

> GetSSOUserByIdAPIResponse GetSSOUserById(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSSOUserById(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSSOUserById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSSOUserById`: GetSSOUserByIdAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSSOUserById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSOUserByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetSSOUserByIdAPIResponse**](GetSSOUserByIdAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSOUsers

> GetSSOUsers200Response GetSSOUsers(ctx).TenantId(tenantId).Skip(skip).Execute()



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
	skip := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSSOUsers(context.Background()).TenantId(tenantId).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSSOUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSSOUsers`: GetSSOUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSSOUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSSOUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **skip** | **int32** |  | 

### Return type

[**GetSSOUsers200Response**](GetSSOUsers200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptions

> GetSubscriptionsAPIResponse GetSubscriptions(ctx).TenantId(tenantId).UserId(userId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSubscriptions(context.Background()).TenantId(tenantId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptions`: GetSubscriptionsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **userId** | **string** |  | 

### Return type

[**GetSubscriptionsAPIResponse**](GetSubscriptionsAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenant

> GetTenant200Response GetTenant(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTenant(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenant`: GetTenant200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetTenant200Response**](GetTenant200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantDailyUsages

> GetTenantDailyUsages200Response GetTenantDailyUsages(ctx).TenantId(tenantId).YearNumber(yearNumber).MonthNumber(monthNumber).DayNumber(dayNumber).Skip(skip).Execute()



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
	yearNumber := float64(1.2) // float64 |  (optional)
	monthNumber := float64(1.2) // float64 |  (optional)
	dayNumber := float64(1.2) // float64 |  (optional)
	skip := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTenantDailyUsages(context.Background()).TenantId(tenantId).YearNumber(yearNumber).MonthNumber(monthNumber).DayNumber(dayNumber).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTenantDailyUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantDailyUsages`: GetTenantDailyUsages200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTenantDailyUsages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantDailyUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **yearNumber** | **float64** |  | 
 **monthNumber** | **float64** |  | 
 **dayNumber** | **float64** |  | 
 **skip** | **float64** |  | 

### Return type

[**GetTenantDailyUsages200Response**](GetTenantDailyUsages200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantPackage

> GetTenantPackage200Response GetTenantPackage(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTenantPackage(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTenantPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantPackage`: GetTenantPackage200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTenantPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetTenantPackage200Response**](GetTenantPackage200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantPackages

> GetTenantPackages200Response GetTenantPackages(ctx).TenantId(tenantId).Skip(skip).Execute()



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
	skip := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTenantPackages(context.Background()).TenantId(tenantId).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTenantPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantPackages`: GetTenantPackages200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTenantPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **skip** | **float64** |  | 

### Return type

[**GetTenantPackages200Response**](GetTenantPackages200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantUser

> GetTenantUser200Response GetTenantUser(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTenantUser(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTenantUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantUser`: GetTenantUser200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTenantUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetTenantUser200Response**](GetTenantUser200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantUsers

> GetTenantUsers200Response GetTenantUsers(ctx).TenantId(tenantId).Skip(skip).Execute()



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
	skip := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTenantUsers(context.Background()).TenantId(tenantId).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTenantUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantUsers`: GetTenantUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTenantUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **skip** | **float64** |  | 

### Return type

[**GetTenantUsers200Response**](GetTenantUsers200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenants

> GetTenants200Response GetTenants(ctx).TenantId(tenantId).Meta(meta).Skip(skip).Execute()



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
	meta := "meta_example" // string |  (optional)
	skip := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTenants(context.Background()).TenantId(tenantId).Meta(meta).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTenants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenants`: GetTenants200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTenants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **meta** | **string** |  | 
 **skip** | **float64** |  | 

### Return type

[**GetTenants200Response**](GetTenants200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicket

> GetTicket200Response GetTicket(ctx, id).TenantId(tenantId).UserId(userId).Execute()



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
	id := "id_example" // string | 
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTicket(context.Background(), id).TenantId(tenantId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTicket`: GetTicket200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **userId** | **string** |  | 

### Return type

[**GetTicket200Response**](GetTicket200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTickets

> GetTickets200Response GetTickets(ctx).TenantId(tenantId).UserId(userId).State(state).Skip(skip).Limit(limit).Execute()



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
	state := float64(1.2) // float64 |  (optional)
	skip := float64(1.2) // float64 |  (optional)
	limit := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTickets(context.Background()).TenantId(tenantId).UserId(userId).State(state).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTickets`: GetTickets200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTickets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **userId** | **string** |  | 
 **state** | **float64** |  | 
 **skip** | **float64** |  | 
 **limit** | **float64** |  | 

### Return type

[**GetTickets200Response**](GetTickets200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> GetUser200Response GetUser(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetUser(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: GetUser200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetUser200Response**](GetUser200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBadge

> GetUserBadge200Response GetUserBadge(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetUserBadge(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetUserBadge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBadge`: GetUserBadge200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetUserBadge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetUserBadge200Response**](GetUserBadge200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBadgeProgressById

> GetUserBadgeProgressById200Response GetUserBadgeProgressById(ctx, id).TenantId(tenantId).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetUserBadgeProgressById(context.Background(), id).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetUserBadgeProgressById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBadgeProgressById`: GetUserBadgeProgressById200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetUserBadgeProgressById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBadgeProgressByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetUserBadgeProgressById200Response**](GetUserBadgeProgressById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBadgeProgressByUserId

> GetUserBadgeProgressById200Response GetUserBadgeProgressByUserId(ctx, userId).TenantId(tenantId).Execute()



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetUserBadgeProgressByUserId(context.Background(), userId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetUserBadgeProgressByUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBadgeProgressByUserId`: GetUserBadgeProgressById200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetUserBadgeProgressByUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBadgeProgressByUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**GetUserBadgeProgressById200Response**](GetUserBadgeProgressById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBadgeProgressList

> GetUserBadgeProgressList200Response GetUserBadgeProgressList(ctx).TenantId(tenantId).UserId(userId).Limit(limit).Skip(skip).Execute()



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
	limit := float64(1.2) // float64 |  (optional)
	skip := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetUserBadgeProgressList(context.Background()).TenantId(tenantId).UserId(userId).Limit(limit).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetUserBadgeProgressList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBadgeProgressList`: GetUserBadgeProgressList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetUserBadgeProgressList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBadgeProgressListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **userId** | **string** |  | 
 **limit** | **float64** |  | 
 **skip** | **float64** |  | 

### Return type

[**GetUserBadgeProgressList200Response**](GetUserBadgeProgressList200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBadges

> GetUserBadges200Response GetUserBadges(ctx).TenantId(tenantId).UserId(userId).BadgeId(badgeId).Type_(type_).DisplayedOnComments(displayedOnComments).Limit(limit).Skip(skip).Execute()



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
	badgeId := "badgeId_example" // string |  (optional)
	type_ := float64(1.2) // float64 |  (optional)
	displayedOnComments := true // bool |  (optional)
	limit := float64(1.2) // float64 |  (optional)
	skip := float64(1.2) // float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetUserBadges(context.Background()).TenantId(tenantId).UserId(userId).BadgeId(badgeId).Type_(type_).DisplayedOnComments(displayedOnComments).Limit(limit).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetUserBadges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBadges`: GetUserBadges200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetUserBadges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBadgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **userId** | **string** |  | 
 **badgeId** | **string** |  | 
 **type_** | **float64** |  | 
 **displayedOnComments** | **bool** |  | 
 **limit** | **float64** |  | 
 **skip** | **float64** |  | 

### Return type

[**GetUserBadges200Response**](GetUserBadges200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVotes

> GetVotes200Response GetVotes(ctx).TenantId(tenantId).UrlId(urlId).Execute()



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
	resp, r, err := apiClient.DefaultAPI.GetVotes(context.Background()).TenantId(tenantId).UrlId(urlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetVotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVotes`: GetVotes200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetVotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **urlId** | **string** |  | 

### Return type

[**GetVotes200Response**](GetVotes200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVotesForUser

> GetVotesForUser200Response GetVotesForUser(ctx).TenantId(tenantId).UrlId(urlId).UserId(userId).AnonUserId(anonUserId).Execute()



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
	userId := "userId_example" // string |  (optional)
	anonUserId := "anonUserId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetVotesForUser(context.Background()).TenantId(tenantId).UrlId(urlId).UserId(userId).AnonUserId(anonUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetVotesForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVotesForUser`: GetVotesForUser200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetVotesForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVotesForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **urlId** | **string** |  | 
 **userId** | **string** |  | 
 **anonUserId** | **string** |  | 

### Return type

[**GetVotesForUser200Response**](GetVotesForUser200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDomainConfig

> GetDomainConfig200Response PatchDomainConfig(ctx, domainToUpdate).TenantId(tenantId).PatchDomainConfigParams(patchDomainConfigParams).Execute()



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
	domainToUpdate := "domainToUpdate_example" // string | 
	patchDomainConfigParams := *openapiclient.NewPatchDomainConfigParams() // PatchDomainConfigParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PatchDomainConfig(context.Background(), domainToUpdate).TenantId(tenantId).PatchDomainConfigParams(patchDomainConfigParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PatchDomainConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDomainConfig`: GetDomainConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PatchDomainConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainToUpdate** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDomainConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **patchDomainConfigParams** | [**PatchDomainConfigParams**](PatchDomainConfigParams.md) |  | 

### Return type

[**GetDomainConfig200Response**](GetDomainConfig200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchHashTag

> PatchHashTag200Response PatchHashTag(ctx, tag).TenantId(tenantId).UpdateHashTagBody(updateHashTagBody).Execute()



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
	tag := "tag_example" // string | 
	tenantId := "tenantId_example" // string |  (optional)
	updateHashTagBody := *openapiclient.NewUpdateHashTagBody() // UpdateHashTagBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PatchHashTag(context.Background(), tag).TenantId(tenantId).UpdateHashTagBody(updateHashTagBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PatchHashTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchHashTag`: PatchHashTag200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PatchHashTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchHashTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **updateHashTagBody** | [**UpdateHashTagBody**](UpdateHashTagBody.md) |  | 

### Return type

[**PatchHashTag200Response**](PatchHashTag200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPage

> PatchPageAPIResponse PatchPage(ctx, id).TenantId(tenantId).UpdateAPIPageData(updateAPIPageData).Execute()



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
	id := "id_example" // string | 
	updateAPIPageData := *openapiclient.NewUpdateAPIPageData() // UpdateAPIPageData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PatchPage(context.Background(), id).TenantId(tenantId).UpdateAPIPageData(updateAPIPageData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PatchPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPage`: PatchPageAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PatchPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updateAPIPageData** | [**UpdateAPIPageData**](UpdateAPIPageData.md) |  | 

### Return type

[**PatchPageAPIResponse**](PatchPageAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSSOUser

> PatchSSOUserAPIResponse PatchSSOUser(ctx, id).TenantId(tenantId).UpdateAPISSOUserData(updateAPISSOUserData).UpdateComments(updateComments).Execute()



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
	id := "id_example" // string | 
	updateAPISSOUserData := *openapiclient.NewUpdateAPISSOUserData() // UpdateAPISSOUserData | 
	updateComments := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PatchSSOUser(context.Background(), id).TenantId(tenantId).UpdateAPISSOUserData(updateAPISSOUserData).UpdateComments(updateComments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PatchSSOUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSSOUser`: PatchSSOUserAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PatchSSOUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSSOUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updateAPISSOUserData** | [**UpdateAPISSOUserData**](UpdateAPISSOUserData.md) |  | 
 **updateComments** | **bool** |  | 

### Return type

[**PatchSSOUserAPIResponse**](PatchSSOUserAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDomainConfig

> GetDomainConfig200Response PutDomainConfig(ctx, domainToUpdate).TenantId(tenantId).UpdateDomainConfigParams(updateDomainConfigParams).Execute()



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
	domainToUpdate := "domainToUpdate_example" // string | 
	updateDomainConfigParams := *openapiclient.NewUpdateDomainConfigParams("Domain_example") // UpdateDomainConfigParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutDomainConfig(context.Background(), domainToUpdate).TenantId(tenantId).UpdateDomainConfigParams(updateDomainConfigParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutDomainConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDomainConfig`: GetDomainConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutDomainConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainToUpdate** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDomainConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updateDomainConfigParams** | [**UpdateDomainConfigParams**](UpdateDomainConfigParams.md) |  | 

### Return type

[**GetDomainConfig200Response**](GetDomainConfig200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSSOUser

> PutSSOUserAPIResponse PutSSOUser(ctx, id).TenantId(tenantId).UpdateAPISSOUserData(updateAPISSOUserData).UpdateComments(updateComments).Execute()



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
	id := "id_example" // string | 
	updateAPISSOUserData := *openapiclient.NewUpdateAPISSOUserData() // UpdateAPISSOUserData | 
	updateComments := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutSSOUser(context.Background(), id).TenantId(tenantId).UpdateAPISSOUserData(updateAPISSOUserData).UpdateComments(updateComments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutSSOUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSSOUser`: PutSSOUserAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutSSOUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSSOUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updateAPISSOUserData** | [**UpdateAPISSOUserData**](UpdateAPISSOUserData.md) |  | 
 **updateComments** | **bool** |  | 

### Return type

[**PutSSOUserAPIResponse**](PutSSOUserAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderEmailTemplate

> RenderEmailTemplate200Response RenderEmailTemplate(ctx).TenantId(tenantId).RenderEmailTemplateBody(renderEmailTemplateBody).Locale(locale).Execute()



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
	renderEmailTemplateBody := *openapiclient.NewRenderEmailTemplateBody("EmailTemplateId_example", "Ejs_example") // RenderEmailTemplateBody | 
	locale := "locale_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RenderEmailTemplate(context.Background()).TenantId(tenantId).RenderEmailTemplateBody(renderEmailTemplateBody).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RenderEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderEmailTemplate`: RenderEmailTemplate200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RenderEmailTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenderEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **renderEmailTemplateBody** | [**RenderEmailTemplateBody**](RenderEmailTemplateBody.md) |  | 
 **locale** | **string** |  | 

### Return type

[**RenderEmailTemplate200Response**](RenderEmailTemplate200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceTenantPackage

> FlagCommentPublic200Response ReplaceTenantPackage(ctx, id).TenantId(tenantId).ReplaceTenantPackageBody(replaceTenantPackageBody).Execute()



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
	id := "id_example" // string | 
	replaceTenantPackageBody := *openapiclient.NewReplaceTenantPackageBody("Name_example", float64(123), float64(123), float64(123), float64(123), float64(123), float64(123), float64(123), float64(123), float64(123), float64(123), false, "ForWhoText_example", []string{"FeatureTaglines_example"}, false) // ReplaceTenantPackageBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ReplaceTenantPackage(context.Background(), id).TenantId(tenantId).ReplaceTenantPackageBody(replaceTenantPackageBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ReplaceTenantPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceTenantPackage`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ReplaceTenantPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceTenantPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **replaceTenantPackageBody** | [**ReplaceTenantPackageBody**](ReplaceTenantPackageBody.md) |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceTenantUser

> FlagCommentPublic200Response ReplaceTenantUser(ctx, id).TenantId(tenantId).ReplaceTenantUserBody(replaceTenantUserBody).UpdateComments(updateComments).Execute()



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
	id := "id_example" // string | 
	replaceTenantUserBody := *openapiclient.NewReplaceTenantUserBody("Username_example", "Email_example") // ReplaceTenantUserBody | 
	updateComments := "updateComments_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ReplaceTenantUser(context.Background(), id).TenantId(tenantId).ReplaceTenantUserBody(replaceTenantUserBody).UpdateComments(updateComments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ReplaceTenantUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceTenantUser`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ReplaceTenantUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceTenantUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **replaceTenantUserBody** | [**ReplaceTenantUserBody**](ReplaceTenantUserBody.md) |  | 
 **updateComments** | **string** |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveComment

> SaveComment200Response SaveComment(ctx).TenantId(tenantId).CreateCommentParams(createCommentParams).IsLive(isLive).DoSpamCheck(doSpamCheck).SendEmails(sendEmails).PopulateNotifications(populateNotifications).Execute()



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
	createCommentParams := *openapiclient.NewCreateCommentParams("CommenterName_example", "Comment_example", "Url_example", "UrlId_example", "Locale_example") // CreateCommentParams | 
	isLive := true // bool |  (optional)
	doSpamCheck := true // bool |  (optional)
	sendEmails := true // bool |  (optional)
	populateNotifications := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SaveComment(context.Background()).TenantId(tenantId).CreateCommentParams(createCommentParams).IsLive(isLive).DoSpamCheck(doSpamCheck).SendEmails(sendEmails).PopulateNotifications(populateNotifications).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SaveComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveComment`: SaveComment200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SaveComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createCommentParams** | [**CreateCommentParams**](CreateCommentParams.md) |  | 
 **isLive** | **bool** |  | 
 **doSpamCheck** | **bool** |  | 
 **sendEmails** | **bool** |  | 
 **populateNotifications** | **bool** |  | 

### Return type

[**SaveComment200Response**](SaveComment200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveCommentsBulk

> []SaveComment200Response SaveCommentsBulk(ctx).TenantId(tenantId).CreateCommentParams(createCommentParams).IsLive(isLive).DoSpamCheck(doSpamCheck).SendEmails(sendEmails).PopulateNotifications(populateNotifications).Execute()



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
	createCommentParams := []openapiclient.CreateCommentParams{*openapiclient.NewCreateCommentParams("CommenterName_example", "Comment_example", "Url_example", "UrlId_example", "Locale_example")} // []CreateCommentParams | 
	isLive := true // bool |  (optional)
	doSpamCheck := true // bool |  (optional)
	sendEmails := true // bool |  (optional)
	populateNotifications := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SaveCommentsBulk(context.Background()).TenantId(tenantId).CreateCommentParams(createCommentParams).IsLive(isLive).DoSpamCheck(doSpamCheck).SendEmails(sendEmails).PopulateNotifications(populateNotifications).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SaveCommentsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveCommentsBulk`: []SaveComment200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SaveCommentsBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveCommentsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createCommentParams** | [**[]CreateCommentParams**](CreateCommentParams.md) |  | 
 **isLive** | **bool** |  | 
 **doSpamCheck** | **bool** |  | 
 **sendEmails** | **bool** |  | 
 **populateNotifications** | **bool** |  | 

### Return type

[**[]SaveComment200Response**](SaveComment200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendInvite

> FlagCommentPublic200Response SendInvite(ctx, id).TenantId(tenantId).FromName(fromName).Execute()



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
	id := "id_example" // string | 
	fromName := "fromName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SendInvite(context.Background(), id).TenantId(tenantId).FromName(fromName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SendInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendInvite`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SendInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **fromName** | **string** |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendLoginLink

> FlagCommentPublic200Response SendLoginLink(ctx, id).TenantId(tenantId).RedirectURL(redirectURL).Execute()



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
	id := "id_example" // string | 
	redirectURL := "redirectURL_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SendLoginLink(context.Background(), id).TenantId(tenantId).RedirectURL(redirectURL).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SendLoginLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendLoginLink`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SendLoginLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendLoginLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **redirectURL** | **string** |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnBlockUserFromComment

> UnBlockCommentPublic200Response UnBlockUserFromComment(ctx, id).TenantId(tenantId).UnBlockFromCommentParams(unBlockFromCommentParams).UserId(userId).AnonUserId(anonUserId).Execute()



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
	id := "id_example" // string | 
	unBlockFromCommentParams := *openapiclient.NewUnBlockFromCommentParams() // UnBlockFromCommentParams | 
	userId := "userId_example" // string |  (optional)
	anonUserId := "anonUserId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UnBlockUserFromComment(context.Background(), id).TenantId(tenantId).UnBlockFromCommentParams(unBlockFromCommentParams).UserId(userId).AnonUserId(anonUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UnBlockUserFromComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnBlockUserFromComment`: UnBlockCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UnBlockUserFromComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnBlockUserFromCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **unBlockFromCommentParams** | [**UnBlockFromCommentParams**](UnBlockFromCommentParams.md) |  | 
 **userId** | **string** |  | 
 **anonUserId** | **string** |  | 

### Return type

[**UnBlockCommentPublic200Response**](UnBlockCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnFlagComment

> FlagComment200Response UnFlagComment(ctx, id).TenantId(tenantId).UserId(userId).AnonUserId(anonUserId).Execute()



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
	id := "id_example" // string | 
	userId := "userId_example" // string |  (optional)
	anonUserId := "anonUserId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UnFlagComment(context.Background(), id).TenantId(tenantId).UserId(userId).AnonUserId(anonUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UnFlagComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnFlagComment`: FlagComment200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UnFlagComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnFlagCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **userId** | **string** |  | 
 **anonUserId** | **string** |  | 

### Return type

[**FlagComment200Response**](FlagComment200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComment

> FlagCommentPublic200Response UpdateComment(ctx, id).TenantId(tenantId).UpdatableCommentParams(updatableCommentParams).ContextUserId(contextUserId).DoSpamCheck(doSpamCheck).IsLive(isLive).Execute()



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
	id := "id_example" // string | 
	updatableCommentParams := *openapiclient.NewUpdatableCommentParams() // UpdatableCommentParams | 
	contextUserId := "contextUserId_example" // string |  (optional)
	doSpamCheck := true // bool |  (optional)
	isLive := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateComment(context.Background(), id).TenantId(tenantId).UpdatableCommentParams(updatableCommentParams).ContextUserId(contextUserId).DoSpamCheck(doSpamCheck).IsLive(isLive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateComment`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updatableCommentParams** | [**UpdatableCommentParams**](UpdatableCommentParams.md) |  | 
 **contextUserId** | **string** |  | 
 **doSpamCheck** | **bool** |  | 
 **isLive** | **bool** |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailTemplate

> FlagCommentPublic200Response UpdateEmailTemplate(ctx, id).TenantId(tenantId).UpdateEmailTemplateBody(updateEmailTemplateBody).Execute()



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
	id := "id_example" // string | 
	updateEmailTemplateBody := *openapiclient.NewUpdateEmailTemplateBody() // UpdateEmailTemplateBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateEmailTemplate(context.Background(), id).TenantId(tenantId).UpdateEmailTemplateBody(updateEmailTemplateBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEmailTemplate`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updateEmailTemplateBody** | [**UpdateEmailTemplateBody**](UpdateEmailTemplateBody.md) |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeedPost

> FlagCommentPublic200Response UpdateFeedPost(ctx, id).TenantId(tenantId).FeedPost(feedPost).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fastcomments/fastcomments-go/client"
)

func main() {
	tenantId := "tenantId_example" // string | 
	id := "id_example" // string | 
	feedPost := *openapiclient.NewFeedPost("Id_example", "TenantId_example", time.Now()) // FeedPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateFeedPost(context.Background(), id).TenantId(tenantId).FeedPost(feedPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateFeedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFeedPost`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateFeedPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **feedPost** | [**FeedPost**](FeedPost.md) |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateModerator

> FlagCommentPublic200Response UpdateModerator(ctx, id).TenantId(tenantId).UpdateModeratorBody(updateModeratorBody).Execute()



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
	id := "id_example" // string | 
	updateModeratorBody := *openapiclient.NewUpdateModeratorBody() // UpdateModeratorBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateModerator(context.Background(), id).TenantId(tenantId).UpdateModeratorBody(updateModeratorBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateModerator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateModerator`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateModerator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateModeratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updateModeratorBody** | [**UpdateModeratorBody**](UpdateModeratorBody.md) |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotification

> FlagCommentPublic200Response UpdateNotification(ctx, id).TenantId(tenantId).UpdateNotificationBody(updateNotificationBody).UserId(userId).Execute()



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
	id := "id_example" // string | 
	updateNotificationBody := *openapiclient.NewUpdateNotificationBody() // UpdateNotificationBody | 
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateNotification(context.Background(), id).TenantId(tenantId).UpdateNotificationBody(updateNotificationBody).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotification`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updateNotificationBody** | [**UpdateNotificationBody**](UpdateNotificationBody.md) |  | 
 **userId** | **string** |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuestionConfig

> FlagCommentPublic200Response UpdateQuestionConfig(ctx, id).TenantId(tenantId).UpdateQuestionConfigBody(updateQuestionConfigBody).Execute()



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
	id := "id_example" // string | 
	updateQuestionConfigBody := *openapiclient.NewUpdateQuestionConfigBody() // UpdateQuestionConfigBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateQuestionConfig(context.Background(), id).TenantId(tenantId).UpdateQuestionConfigBody(updateQuestionConfigBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateQuestionConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateQuestionConfig`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateQuestionConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuestionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updateQuestionConfigBody** | [**UpdateQuestionConfigBody**](UpdateQuestionConfigBody.md) |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuestionResult

> FlagCommentPublic200Response UpdateQuestionResult(ctx, id).TenantId(tenantId).UpdateQuestionResultBody(updateQuestionResultBody).Execute()



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
	id := "id_example" // string | 
	updateQuestionResultBody := *openapiclient.NewUpdateQuestionResultBody() // UpdateQuestionResultBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateQuestionResult(context.Background(), id).TenantId(tenantId).UpdateQuestionResultBody(updateQuestionResultBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateQuestionResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateQuestionResult`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateQuestionResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuestionResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updateQuestionResultBody** | [**UpdateQuestionResultBody**](UpdateQuestionResultBody.md) |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> UpdateSubscriptionAPIResponse UpdateSubscription(ctx, id).TenantId(tenantId).UpdateAPIUserSubscriptionData(updateAPIUserSubscriptionData).UserId(userId).Execute()



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
	id := "id_example" // string | 
	updateAPIUserSubscriptionData := *openapiclient.NewUpdateAPIUserSubscriptionData() // UpdateAPIUserSubscriptionData | 
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateSubscription(context.Background(), id).TenantId(tenantId).UpdateAPIUserSubscriptionData(updateAPIUserSubscriptionData).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubscription`: UpdateSubscriptionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updateAPIUserSubscriptionData** | [**UpdateAPIUserSubscriptionData**](UpdateAPIUserSubscriptionData.md) |  | 
 **userId** | **string** |  | 

### Return type

[**UpdateSubscriptionAPIResponse**](UpdateSubscriptionAPIResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenant

> FlagCommentPublic200Response UpdateTenant(ctx, id).TenantId(tenantId).UpdateTenantBody(updateTenantBody).Execute()



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
	id := "id_example" // string | 
	updateTenantBody := *openapiclient.NewUpdateTenantBody() // UpdateTenantBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateTenant(context.Background(), id).TenantId(tenantId).UpdateTenantBody(updateTenantBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTenant`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updateTenantBody** | [**UpdateTenantBody**](UpdateTenantBody.md) |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantPackage

> FlagCommentPublic200Response UpdateTenantPackage(ctx, id).TenantId(tenantId).UpdateTenantPackageBody(updateTenantPackageBody).Execute()



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
	id := "id_example" // string | 
	updateTenantPackageBody := *openapiclient.NewUpdateTenantPackageBody() // UpdateTenantPackageBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateTenantPackage(context.Background(), id).TenantId(tenantId).UpdateTenantPackageBody(updateTenantPackageBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateTenantPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTenantPackage`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateTenantPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updateTenantPackageBody** | [**UpdateTenantPackageBody**](UpdateTenantPackageBody.md) |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantUser

> FlagCommentPublic200Response UpdateTenantUser(ctx, id).TenantId(tenantId).UpdateTenantUserBody(updateTenantUserBody).UpdateComments(updateComments).Execute()



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
	id := "id_example" // string | 
	updateTenantUserBody := *openapiclient.NewUpdateTenantUserBody() // UpdateTenantUserBody | 
	updateComments := "updateComments_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateTenantUser(context.Background(), id).TenantId(tenantId).UpdateTenantUserBody(updateTenantUserBody).UpdateComments(updateComments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateTenantUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTenantUser`: FlagCommentPublic200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateTenantUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updateTenantUserBody** | [**UpdateTenantUserBody**](UpdateTenantUserBody.md) |  | 
 **updateComments** | **string** |  | 

### Return type

[**FlagCommentPublic200Response**](FlagCommentPublic200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserBadge

> UpdateUserBadge200Response UpdateUserBadge(ctx, id).TenantId(tenantId).UpdateUserBadgeParams(updateUserBadgeParams).Execute()



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
	id := "id_example" // string | 
	updateUserBadgeParams := *openapiclient.NewUpdateUserBadgeParams() // UpdateUserBadgeParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateUserBadge(context.Background(), id).TenantId(tenantId).UpdateUserBadgeParams(updateUserBadgeParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateUserBadge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserBadge`: UpdateUserBadge200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateUserBadge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **updateUserBadgeParams** | [**UpdateUserBadgeParams**](UpdateUserBadgeParams.md) |  | 

### Return type

[**UpdateUserBadge200Response**](UpdateUserBadge200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

