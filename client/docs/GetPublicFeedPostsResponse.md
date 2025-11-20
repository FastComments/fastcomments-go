# GetPublicFeedPostsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**FeedPosts** | [**[]FeedPost**](FeedPost.md) |  | 
**User** | Pointer to [**NullableUserSessionInfo**](UserSessionInfo.md) |  | [optional] 

## Methods

### NewGetPublicFeedPostsResponse

`func NewGetPublicFeedPostsResponse(status APIStatus, feedPosts []FeedPost, ) *GetPublicFeedPostsResponse`

NewGetPublicFeedPostsResponse instantiates a new GetPublicFeedPostsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicFeedPostsResponseWithDefaults

`func NewGetPublicFeedPostsResponseWithDefaults() *GetPublicFeedPostsResponse`

NewGetPublicFeedPostsResponseWithDefaults instantiates a new GetPublicFeedPostsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetPublicFeedPostsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPublicFeedPostsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPublicFeedPostsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetFeedPosts

`func (o *GetPublicFeedPostsResponse) GetFeedPosts() []FeedPost`

GetFeedPosts returns the FeedPosts field if non-nil, zero value otherwise.

### GetFeedPostsOk

`func (o *GetPublicFeedPostsResponse) GetFeedPostsOk() (*[]FeedPost, bool)`

GetFeedPostsOk returns a tuple with the FeedPosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedPosts

`func (o *GetPublicFeedPostsResponse) SetFeedPosts(v []FeedPost)`

SetFeedPosts sets FeedPosts field to given value.


### GetUser

`func (o *GetPublicFeedPostsResponse) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetPublicFeedPostsResponse) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetPublicFeedPostsResponse) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.

### HasUser

`func (o *GetPublicFeedPostsResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *GetPublicFeedPostsResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GetPublicFeedPostsResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


