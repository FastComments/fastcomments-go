# UnblockSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**CommentStatuses** | **map[string]bool** | Construct a type with a set of properties K of type T | 

## Methods

### NewUnblockSuccess

`func NewUnblockSuccess(status APIStatus, commentStatuses map[string]bool, ) *UnblockSuccess`

NewUnblockSuccess instantiates a new UnblockSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnblockSuccessWithDefaults

`func NewUnblockSuccessWithDefaults() *UnblockSuccess`

NewUnblockSuccessWithDefaults instantiates a new UnblockSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UnblockSuccess) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnblockSuccess) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnblockSuccess) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetCommentStatuses

`func (o *UnblockSuccess) GetCommentStatuses() map[string]bool`

GetCommentStatuses returns the CommentStatuses field if non-nil, zero value otherwise.

### GetCommentStatusesOk

`func (o *UnblockSuccess) GetCommentStatusesOk() (*map[string]bool, bool)`

GetCommentStatusesOk returns a tuple with the CommentStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentStatuses

`func (o *UnblockSuccess) SetCommentStatuses(v map[string]bool)`

SetCommentStatuses sets CommentStatuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


