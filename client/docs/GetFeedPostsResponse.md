# GetFeedPostsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**FeedPosts** | [**[]FeedPost**](FeedPost.md) |  | 

## Methods

### NewGetFeedPostsResponse

`func NewGetFeedPostsResponse(status APIStatus, feedPosts []FeedPost, ) *GetFeedPostsResponse`

NewGetFeedPostsResponse instantiates a new GetFeedPostsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFeedPostsResponseWithDefaults

`func NewGetFeedPostsResponseWithDefaults() *GetFeedPostsResponse`

NewGetFeedPostsResponseWithDefaults instantiates a new GetFeedPostsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetFeedPostsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetFeedPostsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetFeedPostsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetFeedPosts

`func (o *GetFeedPostsResponse) GetFeedPosts() []FeedPost`

GetFeedPosts returns the FeedPosts field if non-nil, zero value otherwise.

### GetFeedPostsOk

`func (o *GetFeedPostsResponse) GetFeedPostsOk() (*[]FeedPost, bool)`

GetFeedPostsOk returns a tuple with the FeedPosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedPosts

`func (o *GetFeedPostsResponse) SetFeedPosts(v []FeedPost)`

SetFeedPosts sets FeedPosts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


