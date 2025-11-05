# CreateFeedPostsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**FeedPost** | [**FeedPost**](FeedPost.md) |  | 

## Methods

### NewCreateFeedPostsResponse

`func NewCreateFeedPostsResponse(status APIStatus, feedPost FeedPost, ) *CreateFeedPostsResponse`

NewCreateFeedPostsResponse instantiates a new CreateFeedPostsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFeedPostsResponseWithDefaults

`func NewCreateFeedPostsResponseWithDefaults() *CreateFeedPostsResponse`

NewCreateFeedPostsResponseWithDefaults instantiates a new CreateFeedPostsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateFeedPostsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateFeedPostsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateFeedPostsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetFeedPost

`func (o *CreateFeedPostsResponse) GetFeedPost() FeedPost`

GetFeedPost returns the FeedPost field if non-nil, zero value otherwise.

### GetFeedPostOk

`func (o *CreateFeedPostsResponse) GetFeedPostOk() (*FeedPost, bool)`

GetFeedPostOk returns a tuple with the FeedPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedPost

`func (o *CreateFeedPostsResponse) SetFeedPost(v FeedPost)`

SetFeedPost sets FeedPost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


