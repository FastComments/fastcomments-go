# GetOfflineUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextAfterUserId** | **NullableString** |  | 
**NextAfterName** | **NullableString** |  | 
**Users** | [**[]PageUserEntry**](PageUserEntry.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetOfflineUsersResponse

`func NewGetOfflineUsersResponse(nextAfterUserId NullableString, nextAfterName NullableString, users []PageUserEntry, status APIStatus, reason string, code string, ) *GetOfflineUsersResponse`

NewGetOfflineUsersResponse instantiates a new GetOfflineUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOfflineUsersResponseWithDefaults

`func NewGetOfflineUsersResponseWithDefaults() *GetOfflineUsersResponse`

NewGetOfflineUsersResponseWithDefaults instantiates a new GetOfflineUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextAfterUserId

`func (o *GetOfflineUsersResponse) GetNextAfterUserId() string`

GetNextAfterUserId returns the NextAfterUserId field if non-nil, zero value otherwise.

### GetNextAfterUserIdOk

`func (o *GetOfflineUsersResponse) GetNextAfterUserIdOk() (*string, bool)`

GetNextAfterUserIdOk returns a tuple with the NextAfterUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAfterUserId

`func (o *GetOfflineUsersResponse) SetNextAfterUserId(v string)`

SetNextAfterUserId sets NextAfterUserId field to given value.


### SetNextAfterUserIdNil

`func (o *GetOfflineUsersResponse) SetNextAfterUserIdNil(b bool)`

 SetNextAfterUserIdNil sets the value for NextAfterUserId to be an explicit nil

### UnsetNextAfterUserId
`func (o *GetOfflineUsersResponse) UnsetNextAfterUserId()`

UnsetNextAfterUserId ensures that no value is present for NextAfterUserId, not even an explicit nil
### GetNextAfterName

`func (o *GetOfflineUsersResponse) GetNextAfterName() string`

GetNextAfterName returns the NextAfterName field if non-nil, zero value otherwise.

### GetNextAfterNameOk

`func (o *GetOfflineUsersResponse) GetNextAfterNameOk() (*string, bool)`

GetNextAfterNameOk returns a tuple with the NextAfterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAfterName

`func (o *GetOfflineUsersResponse) SetNextAfterName(v string)`

SetNextAfterName sets NextAfterName field to given value.


### SetNextAfterNameNil

`func (o *GetOfflineUsersResponse) SetNextAfterNameNil(b bool)`

 SetNextAfterNameNil sets the value for NextAfterName to be an explicit nil

### UnsetNextAfterName
`func (o *GetOfflineUsersResponse) UnsetNextAfterName()`

UnsetNextAfterName ensures that no value is present for NextAfterName, not even an explicit nil
### GetUsers

`func (o *GetOfflineUsersResponse) GetUsers() []PageUserEntry`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetOfflineUsersResponse) GetUsersOk() (*[]PageUserEntry, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetOfflineUsersResponse) SetUsers(v []PageUserEntry)`

SetUsers sets Users field to given value.


### GetStatus

`func (o *GetOfflineUsersResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOfflineUsersResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOfflineUsersResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetReason

`func (o *GetOfflineUsersResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetOfflineUsersResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetOfflineUsersResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetOfflineUsersResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetOfflineUsersResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetOfflineUsersResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetOfflineUsersResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetOfflineUsersResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetOfflineUsersResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetOfflineUsersResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetOfflineUsersResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetOfflineUsersResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetOfflineUsersResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetOfflineUsersResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetOfflineUsersResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetOfflineUsersResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetOfflineUsersResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetOfflineUsersResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetOfflineUsersResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetOfflineUsersResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetOfflineUsersResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetOfflineUsersResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetOfflineUsersResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetOfflineUsersResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetOfflineUsersResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetOfflineUsersResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


