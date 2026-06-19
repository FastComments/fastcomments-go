# ModerationAPICommentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | [**ModerationAPIComment**](ModerationAPIComment.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewModerationAPICommentResponse

`func NewModerationAPICommentResponse(comment ModerationAPIComment, status APIStatus, ) *ModerationAPICommentResponse`

NewModerationAPICommentResponse instantiates a new ModerationAPICommentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationAPICommentResponseWithDefaults

`func NewModerationAPICommentResponseWithDefaults() *ModerationAPICommentResponse`

NewModerationAPICommentResponseWithDefaults instantiates a new ModerationAPICommentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ModerationAPICommentResponse) GetComment() ModerationAPIComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModerationAPICommentResponse) GetCommentOk() (*ModerationAPIComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModerationAPICommentResponse) SetComment(v ModerationAPIComment)`

SetComment sets Comment field to given value.


### GetStatus

`func (o *ModerationAPICommentResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModerationAPICommentResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModerationAPICommentResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


