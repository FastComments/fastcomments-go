# ModerationCommentSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommentCount** | **int32** |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewModerationCommentSearchResponse

`func NewModerationCommentSearchResponse(commentCount int32, status APIStatus, ) *ModerationCommentSearchResponse`

NewModerationCommentSearchResponse instantiates a new ModerationCommentSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationCommentSearchResponseWithDefaults

`func NewModerationCommentSearchResponseWithDefaults() *ModerationCommentSearchResponse`

NewModerationCommentSearchResponseWithDefaults instantiates a new ModerationCommentSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentCount

`func (o *ModerationCommentSearchResponse) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *ModerationCommentSearchResponse) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *ModerationCommentSearchResponse) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.


### GetStatus

`func (o *ModerationCommentSearchResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModerationCommentSearchResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModerationCommentSearchResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


