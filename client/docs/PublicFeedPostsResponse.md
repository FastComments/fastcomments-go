# PublicFeedPostsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ImportedAPIStatusSUCCESS**](ImportedAPIStatusSUCCESS.md) |  | 
**FeedPosts** | [**[]FeedPost**](FeedPost.md) |  | 
**User** | Pointer to [**NullableUserSessionInfo**](UserSessionInfo.md) |  | [optional] 
**UrlIdWS** | Pointer to **string** |  | [optional] 
**UserIdWS** | Pointer to **string** |  | [optional] 
**TenantIdWS** | Pointer to **string** |  | [optional] 
**MyReacts** | Pointer to **map[string]map[string]bool** |  | [optional] 

## Methods

### NewPublicFeedPostsResponse

`func NewPublicFeedPostsResponse(status ImportedAPIStatusSUCCESS, feedPosts []FeedPost, ) *PublicFeedPostsResponse`

NewPublicFeedPostsResponse instantiates a new PublicFeedPostsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicFeedPostsResponseWithDefaults

`func NewPublicFeedPostsResponseWithDefaults() *PublicFeedPostsResponse`

NewPublicFeedPostsResponseWithDefaults instantiates a new PublicFeedPostsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PublicFeedPostsResponse) GetStatus() ImportedAPIStatusSUCCESS`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicFeedPostsResponse) GetStatusOk() (*ImportedAPIStatusSUCCESS, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicFeedPostsResponse) SetStatus(v ImportedAPIStatusSUCCESS)`

SetStatus sets Status field to given value.


### GetFeedPosts

`func (o *PublicFeedPostsResponse) GetFeedPosts() []FeedPost`

GetFeedPosts returns the FeedPosts field if non-nil, zero value otherwise.

### GetFeedPostsOk

`func (o *PublicFeedPostsResponse) GetFeedPostsOk() (*[]FeedPost, bool)`

GetFeedPostsOk returns a tuple with the FeedPosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedPosts

`func (o *PublicFeedPostsResponse) SetFeedPosts(v []FeedPost)`

SetFeedPosts sets FeedPosts field to given value.


### GetUser

`func (o *PublicFeedPostsResponse) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PublicFeedPostsResponse) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PublicFeedPostsResponse) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.

### HasUser

`func (o *PublicFeedPostsResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PublicFeedPostsResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PublicFeedPostsResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUrlIdWS

`func (o *PublicFeedPostsResponse) GetUrlIdWS() string`

GetUrlIdWS returns the UrlIdWS field if non-nil, zero value otherwise.

### GetUrlIdWSOk

`func (o *PublicFeedPostsResponse) GetUrlIdWSOk() (*string, bool)`

GetUrlIdWSOk returns a tuple with the UrlIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdWS

`func (o *PublicFeedPostsResponse) SetUrlIdWS(v string)`

SetUrlIdWS sets UrlIdWS field to given value.

### HasUrlIdWS

`func (o *PublicFeedPostsResponse) HasUrlIdWS() bool`

HasUrlIdWS returns a boolean if a field has been set.

### GetUserIdWS

`func (o *PublicFeedPostsResponse) GetUserIdWS() string`

GetUserIdWS returns the UserIdWS field if non-nil, zero value otherwise.

### GetUserIdWSOk

`func (o *PublicFeedPostsResponse) GetUserIdWSOk() (*string, bool)`

GetUserIdWSOk returns a tuple with the UserIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdWS

`func (o *PublicFeedPostsResponse) SetUserIdWS(v string)`

SetUserIdWS sets UserIdWS field to given value.

### HasUserIdWS

`func (o *PublicFeedPostsResponse) HasUserIdWS() bool`

HasUserIdWS returns a boolean if a field has been set.

### GetTenantIdWS

`func (o *PublicFeedPostsResponse) GetTenantIdWS() string`

GetTenantIdWS returns the TenantIdWS field if non-nil, zero value otherwise.

### GetTenantIdWSOk

`func (o *PublicFeedPostsResponse) GetTenantIdWSOk() (*string, bool)`

GetTenantIdWSOk returns a tuple with the TenantIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdWS

`func (o *PublicFeedPostsResponse) SetTenantIdWS(v string)`

SetTenantIdWS sets TenantIdWS field to given value.

### HasTenantIdWS

`func (o *PublicFeedPostsResponse) HasTenantIdWS() bool`

HasTenantIdWS returns a boolean if a field has been set.

### GetMyReacts

`func (o *PublicFeedPostsResponse) GetMyReacts() map[string]map[string]bool`

GetMyReacts returns the MyReacts field if non-nil, zero value otherwise.

### GetMyReactsOk

`func (o *PublicFeedPostsResponse) GetMyReactsOk() (*map[string]map[string]bool, bool)`

GetMyReactsOk returns a tuple with the MyReacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyReacts

`func (o *PublicFeedPostsResponse) SetMyReacts(v map[string]map[string]bool)`

SetMyReacts sets MyReacts field to given value.

### HasMyReacts

`func (o *PublicFeedPostsResponse) HasMyReacts() bool`

HasMyReacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


