# UpdateUserNotificationCommentSubscriptionStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**MatchedCount** | **int64** |  | 
**ModifiedCount** | **int64** |  | 
**Note** | **string** |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewUpdateUserNotificationCommentSubscriptionStatusResponse

`func NewUpdateUserNotificationCommentSubscriptionStatusResponse(status APIStatus, matchedCount int64, modifiedCount int64, note string, reason string, code string, ) *UpdateUserNotificationCommentSubscriptionStatusResponse`

NewUpdateUserNotificationCommentSubscriptionStatusResponse instantiates a new UpdateUserNotificationCommentSubscriptionStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserNotificationCommentSubscriptionStatusResponseWithDefaults

`func NewUpdateUserNotificationCommentSubscriptionStatusResponseWithDefaults() *UpdateUserNotificationCommentSubscriptionStatusResponse`

NewUpdateUserNotificationCommentSubscriptionStatusResponseWithDefaults instantiates a new UpdateUserNotificationCommentSubscriptionStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetMatchedCount

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetMatchedCount() int64`

GetMatchedCount returns the MatchedCount field if non-nil, zero value otherwise.

### GetMatchedCountOk

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetMatchedCountOk() (*int64, bool)`

GetMatchedCountOk returns a tuple with the MatchedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedCount

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) SetMatchedCount(v int64)`

SetMatchedCount sets MatchedCount field to given value.


### GetModifiedCount

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetModifiedCount() int64`

GetModifiedCount returns the ModifiedCount field if non-nil, zero value otherwise.

### GetModifiedCountOk

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetModifiedCountOk() (*int64, bool)`

GetModifiedCountOk returns a tuple with the ModifiedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedCount

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) SetModifiedCount(v int64)`

SetModifiedCount sets ModifiedCount field to given value.


### GetNote

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) SetNote(v string)`

SetNote sets Note field to given value.


### GetReason

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *UpdateUserNotificationCommentSubscriptionStatusResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


