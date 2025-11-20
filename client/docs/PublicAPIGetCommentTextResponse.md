# PublicAPIGetCommentTextResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**CommentText** | **string** |  | 
**SanitizedCommentText** | **string** |  | 

## Methods

### NewPublicAPIGetCommentTextResponse

`func NewPublicAPIGetCommentTextResponse(status APIStatus, commentText string, sanitizedCommentText string, ) *PublicAPIGetCommentTextResponse`

NewPublicAPIGetCommentTextResponse instantiates a new PublicAPIGetCommentTextResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAPIGetCommentTextResponseWithDefaults

`func NewPublicAPIGetCommentTextResponseWithDefaults() *PublicAPIGetCommentTextResponse`

NewPublicAPIGetCommentTextResponseWithDefaults instantiates a new PublicAPIGetCommentTextResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PublicAPIGetCommentTextResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicAPIGetCommentTextResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicAPIGetCommentTextResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetCommentText

`func (o *PublicAPIGetCommentTextResponse) GetCommentText() string`

GetCommentText returns the CommentText field if non-nil, zero value otherwise.

### GetCommentTextOk

`func (o *PublicAPIGetCommentTextResponse) GetCommentTextOk() (*string, bool)`

GetCommentTextOk returns a tuple with the CommentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentText

`func (o *PublicAPIGetCommentTextResponse) SetCommentText(v string)`

SetCommentText sets CommentText field to given value.


### GetSanitizedCommentText

`func (o *PublicAPIGetCommentTextResponse) GetSanitizedCommentText() string`

GetSanitizedCommentText returns the SanitizedCommentText field if non-nil, zero value otherwise.

### GetSanitizedCommentTextOk

`func (o *PublicAPIGetCommentTextResponse) GetSanitizedCommentTextOk() (*string, bool)`

GetSanitizedCommentTextOk returns a tuple with the SanitizedCommentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanitizedCommentText

`func (o *PublicAPIGetCommentTextResponse) SetSanitizedCommentText(v string)`

SetSanitizedCommentText sets SanitizedCommentText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


