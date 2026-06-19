# SetCommentApprovedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DidResetFlaggedCount** | Pointer to **bool** |  | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewSetCommentApprovedResponse

`func NewSetCommentApprovedResponse(status APIStatus, ) *SetCommentApprovedResponse`

NewSetCommentApprovedResponse instantiates a new SetCommentApprovedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetCommentApprovedResponseWithDefaults

`func NewSetCommentApprovedResponseWithDefaults() *SetCommentApprovedResponse`

NewSetCommentApprovedResponseWithDefaults instantiates a new SetCommentApprovedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDidResetFlaggedCount

`func (o *SetCommentApprovedResponse) GetDidResetFlaggedCount() bool`

GetDidResetFlaggedCount returns the DidResetFlaggedCount field if non-nil, zero value otherwise.

### GetDidResetFlaggedCountOk

`func (o *SetCommentApprovedResponse) GetDidResetFlaggedCountOk() (*bool, bool)`

GetDidResetFlaggedCountOk returns a tuple with the DidResetFlaggedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDidResetFlaggedCount

`func (o *SetCommentApprovedResponse) SetDidResetFlaggedCount(v bool)`

SetDidResetFlaggedCount sets DidResetFlaggedCount field to given value.

### HasDidResetFlaggedCount

`func (o *SetCommentApprovedResponse) HasDidResetFlaggedCount() bool`

HasDidResetFlaggedCount returns a boolean if a field has been set.

### GetStatus

`func (o *SetCommentApprovedResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SetCommentApprovedResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SetCommentApprovedResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


