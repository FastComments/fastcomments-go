# CreateUserBadgeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**UserBadge** | [**UserBadge**](UserBadge.md) |  | 
**Notes** | Pointer to **[]string** |  | [optional] 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewCreateUserBadgeResponse

`func NewCreateUserBadgeResponse(status APIStatus, userBadge UserBadge, reason string, code string, ) *CreateUserBadgeResponse`

NewCreateUserBadgeResponse instantiates a new CreateUserBadgeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserBadgeResponseWithDefaults

`func NewCreateUserBadgeResponseWithDefaults() *CreateUserBadgeResponse`

NewCreateUserBadgeResponseWithDefaults instantiates a new CreateUserBadgeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateUserBadgeResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateUserBadgeResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateUserBadgeResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetUserBadge

`func (o *CreateUserBadgeResponse) GetUserBadge() UserBadge`

GetUserBadge returns the UserBadge field if non-nil, zero value otherwise.

### GetUserBadgeOk

`func (o *CreateUserBadgeResponse) GetUserBadgeOk() (*UserBadge, bool)`

GetUserBadgeOk returns a tuple with the UserBadge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBadge

`func (o *CreateUserBadgeResponse) SetUserBadge(v UserBadge)`

SetUserBadge sets UserBadge field to given value.


### GetNotes

`func (o *CreateUserBadgeResponse) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateUserBadgeResponse) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateUserBadgeResponse) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateUserBadgeResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetReason

`func (o *CreateUserBadgeResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateUserBadgeResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateUserBadgeResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *CreateUserBadgeResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateUserBadgeResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateUserBadgeResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *CreateUserBadgeResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *CreateUserBadgeResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *CreateUserBadgeResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *CreateUserBadgeResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *CreateUserBadgeResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *CreateUserBadgeResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *CreateUserBadgeResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *CreateUserBadgeResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *CreateUserBadgeResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *CreateUserBadgeResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *CreateUserBadgeResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *CreateUserBadgeResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *CreateUserBadgeResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *CreateUserBadgeResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *CreateUserBadgeResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *CreateUserBadgeResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *CreateUserBadgeResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *CreateUserBadgeResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *CreateUserBadgeResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *CreateUserBadgeResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


