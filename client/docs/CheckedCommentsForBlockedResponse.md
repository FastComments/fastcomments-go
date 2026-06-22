# CheckedCommentsForBlockedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommentStatuses** | **map[string]bool** | Construct a type with a set of properties K of type T | 
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewCheckedCommentsForBlockedResponse

`func NewCheckedCommentsForBlockedResponse(commentStatuses map[string]bool, status APIStatus, reason string, code string, ) *CheckedCommentsForBlockedResponse`

NewCheckedCommentsForBlockedResponse instantiates a new CheckedCommentsForBlockedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckedCommentsForBlockedResponseWithDefaults

`func NewCheckedCommentsForBlockedResponseWithDefaults() *CheckedCommentsForBlockedResponse`

NewCheckedCommentsForBlockedResponseWithDefaults instantiates a new CheckedCommentsForBlockedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentStatuses

`func (o *CheckedCommentsForBlockedResponse) GetCommentStatuses() map[string]bool`

GetCommentStatuses returns the CommentStatuses field if non-nil, zero value otherwise.

### GetCommentStatusesOk

`func (o *CheckedCommentsForBlockedResponse) GetCommentStatusesOk() (*map[string]bool, bool)`

GetCommentStatusesOk returns a tuple with the CommentStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentStatuses

`func (o *CheckedCommentsForBlockedResponse) SetCommentStatuses(v map[string]bool)`

SetCommentStatuses sets CommentStatuses field to given value.


### GetStatus

`func (o *CheckedCommentsForBlockedResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckedCommentsForBlockedResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckedCommentsForBlockedResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetReason

`func (o *CheckedCommentsForBlockedResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CheckedCommentsForBlockedResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CheckedCommentsForBlockedResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *CheckedCommentsForBlockedResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CheckedCommentsForBlockedResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CheckedCommentsForBlockedResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *CheckedCommentsForBlockedResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *CheckedCommentsForBlockedResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *CheckedCommentsForBlockedResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *CheckedCommentsForBlockedResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *CheckedCommentsForBlockedResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *CheckedCommentsForBlockedResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *CheckedCommentsForBlockedResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *CheckedCommentsForBlockedResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *CheckedCommentsForBlockedResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *CheckedCommentsForBlockedResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *CheckedCommentsForBlockedResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *CheckedCommentsForBlockedResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *CheckedCommentsForBlockedResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *CheckedCommentsForBlockedResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *CheckedCommentsForBlockedResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *CheckedCommentsForBlockedResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *CheckedCommentsForBlockedResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *CheckedCommentsForBlockedResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *CheckedCommentsForBlockedResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *CheckedCommentsForBlockedResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


