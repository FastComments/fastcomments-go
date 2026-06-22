# UpdateUserNotificationPageSubscriptionStatusResponse

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

### NewUpdateUserNotificationPageSubscriptionStatusResponse

`func NewUpdateUserNotificationPageSubscriptionStatusResponse(status APIStatus, matchedCount int64, modifiedCount int64, note string, reason string, code string, ) *UpdateUserNotificationPageSubscriptionStatusResponse`

NewUpdateUserNotificationPageSubscriptionStatusResponse instantiates a new UpdateUserNotificationPageSubscriptionStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserNotificationPageSubscriptionStatusResponseWithDefaults

`func NewUpdateUserNotificationPageSubscriptionStatusResponseWithDefaults() *UpdateUserNotificationPageSubscriptionStatusResponse`

NewUpdateUserNotificationPageSubscriptionStatusResponseWithDefaults instantiates a new UpdateUserNotificationPageSubscriptionStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetMatchedCount

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetMatchedCount() int64`

GetMatchedCount returns the MatchedCount field if non-nil, zero value otherwise.

### GetMatchedCountOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetMatchedCountOk() (*int64, bool)`

GetMatchedCountOk returns a tuple with the MatchedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedCount

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetMatchedCount(v int64)`

SetMatchedCount sets MatchedCount field to given value.


### GetModifiedCount

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetModifiedCount() int64`

GetModifiedCount returns the ModifiedCount field if non-nil, zero value otherwise.

### GetModifiedCountOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetModifiedCountOk() (*int64, bool)`

GetModifiedCountOk returns a tuple with the ModifiedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedCount

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetModifiedCount(v int64)`

SetModifiedCount sets ModifiedCount field to given value.


### GetNote

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetNote(v string)`

SetNote sets Note field to given value.


### GetReason

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


