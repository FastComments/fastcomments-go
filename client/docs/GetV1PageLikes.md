# GetV1PageLikes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UrlIdWS** | **string** |  | 
**DidLike** | **bool** |  | 
**CommentCount** | **int32** |  | 
**LikeCount** | **int32** |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewGetV1PageLikes

`func NewGetV1PageLikes(urlIdWS string, didLike bool, commentCount int32, likeCount int32, status APIStatus, ) *GetV1PageLikes`

NewGetV1PageLikes instantiates a new GetV1PageLikes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetV1PageLikesWithDefaults

`func NewGetV1PageLikesWithDefaults() *GetV1PageLikes`

NewGetV1PageLikesWithDefaults instantiates a new GetV1PageLikes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrlIdWS

`func (o *GetV1PageLikes) GetUrlIdWS() string`

GetUrlIdWS returns the UrlIdWS field if non-nil, zero value otherwise.

### GetUrlIdWSOk

`func (o *GetV1PageLikes) GetUrlIdWSOk() (*string, bool)`

GetUrlIdWSOk returns a tuple with the UrlIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdWS

`func (o *GetV1PageLikes) SetUrlIdWS(v string)`

SetUrlIdWS sets UrlIdWS field to given value.


### GetDidLike

`func (o *GetV1PageLikes) GetDidLike() bool`

GetDidLike returns the DidLike field if non-nil, zero value otherwise.

### GetDidLikeOk

`func (o *GetV1PageLikes) GetDidLikeOk() (*bool, bool)`

GetDidLikeOk returns a tuple with the DidLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDidLike

`func (o *GetV1PageLikes) SetDidLike(v bool)`

SetDidLike sets DidLike field to given value.


### GetCommentCount

`func (o *GetV1PageLikes) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *GetV1PageLikes) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *GetV1PageLikes) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.


### GetLikeCount

`func (o *GetV1PageLikes) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *GetV1PageLikes) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *GetV1PageLikes) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.


### GetStatus

`func (o *GetV1PageLikes) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetV1PageLikes) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetV1PageLikes) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


