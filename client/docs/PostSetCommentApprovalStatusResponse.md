# PostSetCommentApprovalStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DidResetFlaggedCount** | Pointer to **bool** |  | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewPostSetCommentApprovalStatusResponse

`func NewPostSetCommentApprovalStatusResponse(status APIStatus, reason string, code string, ) *PostSetCommentApprovalStatusResponse`

NewPostSetCommentApprovalStatusResponse instantiates a new PostSetCommentApprovalStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSetCommentApprovalStatusResponseWithDefaults

`func NewPostSetCommentApprovalStatusResponseWithDefaults() *PostSetCommentApprovalStatusResponse`

NewPostSetCommentApprovalStatusResponseWithDefaults instantiates a new PostSetCommentApprovalStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDidResetFlaggedCount

`func (o *PostSetCommentApprovalStatusResponse) GetDidResetFlaggedCount() bool`

GetDidResetFlaggedCount returns the DidResetFlaggedCount field if non-nil, zero value otherwise.

### GetDidResetFlaggedCountOk

`func (o *PostSetCommentApprovalStatusResponse) GetDidResetFlaggedCountOk() (*bool, bool)`

GetDidResetFlaggedCountOk returns a tuple with the DidResetFlaggedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDidResetFlaggedCount

`func (o *PostSetCommentApprovalStatusResponse) SetDidResetFlaggedCount(v bool)`

SetDidResetFlaggedCount sets DidResetFlaggedCount field to given value.

### HasDidResetFlaggedCount

`func (o *PostSetCommentApprovalStatusResponse) HasDidResetFlaggedCount() bool`

HasDidResetFlaggedCount returns a boolean if a field has been set.

### GetStatus

`func (o *PostSetCommentApprovalStatusResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostSetCommentApprovalStatusResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostSetCommentApprovalStatusResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetReason

`func (o *PostSetCommentApprovalStatusResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PostSetCommentApprovalStatusResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PostSetCommentApprovalStatusResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *PostSetCommentApprovalStatusResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PostSetCommentApprovalStatusResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PostSetCommentApprovalStatusResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *PostSetCommentApprovalStatusResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *PostSetCommentApprovalStatusResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *PostSetCommentApprovalStatusResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *PostSetCommentApprovalStatusResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *PostSetCommentApprovalStatusResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *PostSetCommentApprovalStatusResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *PostSetCommentApprovalStatusResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *PostSetCommentApprovalStatusResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *PostSetCommentApprovalStatusResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *PostSetCommentApprovalStatusResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *PostSetCommentApprovalStatusResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *PostSetCommentApprovalStatusResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *PostSetCommentApprovalStatusResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *PostSetCommentApprovalStatusResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *PostSetCommentApprovalStatusResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *PostSetCommentApprovalStatusResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *PostSetCommentApprovalStatusResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *PostSetCommentApprovalStatusResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *PostSetCommentApprovalStatusResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *PostSetCommentApprovalStatusResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


