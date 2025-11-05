# FlagCommentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** |  | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Code** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**WasUnapproved** | Pointer to **bool** |  | [optional] 

## Methods

### NewFlagCommentResponse

`func NewFlagCommentResponse(status APIStatus, ) *FlagCommentResponse`

NewFlagCommentResponse instantiates a new FlagCommentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagCommentResponseWithDefaults

`func NewFlagCommentResponseWithDefaults() *FlagCommentResponse`

NewFlagCommentResponseWithDefaults instantiates a new FlagCommentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *FlagCommentResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *FlagCommentResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *FlagCommentResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *FlagCommentResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetStatus

`func (o *FlagCommentResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlagCommentResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlagCommentResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetCode

`func (o *FlagCommentResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FlagCommentResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FlagCommentResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FlagCommentResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *FlagCommentResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *FlagCommentResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *FlagCommentResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *FlagCommentResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetWasUnapproved

`func (o *FlagCommentResponse) GetWasUnapproved() bool`

GetWasUnapproved returns the WasUnapproved field if non-nil, zero value otherwise.

### GetWasUnapprovedOk

`func (o *FlagCommentResponse) GetWasUnapprovedOk() (*bool, bool)`

GetWasUnapprovedOk returns a tuple with the WasUnapproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasUnapproved

`func (o *FlagCommentResponse) SetWasUnapproved(v bool)`

SetWasUnapproved sets WasUnapproved field to given value.

### HasWasUnapproved

`func (o *FlagCommentResponse) HasWasUnapproved() bool`

HasWasUnapproved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


