# GetOnlineUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextAfterUserId** | **NullableString** |  | 
**NextAfterName** | **NullableString** |  | 
**TotalCount** | **float64** |  | 
**AnonCount** | **float64** |  | 
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

### NewGetOnlineUsersResponse

`func NewGetOnlineUsersResponse(nextAfterUserId NullableString, nextAfterName NullableString, totalCount float64, anonCount float64, users []PageUserEntry, status APIStatus, reason string, code string, ) *GetOnlineUsersResponse`

NewGetOnlineUsersResponse instantiates a new GetOnlineUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOnlineUsersResponseWithDefaults

`func NewGetOnlineUsersResponseWithDefaults() *GetOnlineUsersResponse`

NewGetOnlineUsersResponseWithDefaults instantiates a new GetOnlineUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextAfterUserId

`func (o *GetOnlineUsersResponse) GetNextAfterUserId() string`

GetNextAfterUserId returns the NextAfterUserId field if non-nil, zero value otherwise.

### GetNextAfterUserIdOk

`func (o *GetOnlineUsersResponse) GetNextAfterUserIdOk() (*string, bool)`

GetNextAfterUserIdOk returns a tuple with the NextAfterUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAfterUserId

`func (o *GetOnlineUsersResponse) SetNextAfterUserId(v string)`

SetNextAfterUserId sets NextAfterUserId field to given value.


### SetNextAfterUserIdNil

`func (o *GetOnlineUsersResponse) SetNextAfterUserIdNil(b bool)`

 SetNextAfterUserIdNil sets the value for NextAfterUserId to be an explicit nil

### UnsetNextAfterUserId
`func (o *GetOnlineUsersResponse) UnsetNextAfterUserId()`

UnsetNextAfterUserId ensures that no value is present for NextAfterUserId, not even an explicit nil
### GetNextAfterName

`func (o *GetOnlineUsersResponse) GetNextAfterName() string`

GetNextAfterName returns the NextAfterName field if non-nil, zero value otherwise.

### GetNextAfterNameOk

`func (o *GetOnlineUsersResponse) GetNextAfterNameOk() (*string, bool)`

GetNextAfterNameOk returns a tuple with the NextAfterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAfterName

`func (o *GetOnlineUsersResponse) SetNextAfterName(v string)`

SetNextAfterName sets NextAfterName field to given value.


### SetNextAfterNameNil

`func (o *GetOnlineUsersResponse) SetNextAfterNameNil(b bool)`

 SetNextAfterNameNil sets the value for NextAfterName to be an explicit nil

### UnsetNextAfterName
`func (o *GetOnlineUsersResponse) UnsetNextAfterName()`

UnsetNextAfterName ensures that no value is present for NextAfterName, not even an explicit nil
### GetTotalCount

`func (o *GetOnlineUsersResponse) GetTotalCount() float64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetOnlineUsersResponse) GetTotalCountOk() (*float64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetOnlineUsersResponse) SetTotalCount(v float64)`

SetTotalCount sets TotalCount field to given value.


### GetAnonCount

`func (o *GetOnlineUsersResponse) GetAnonCount() float64`

GetAnonCount returns the AnonCount field if non-nil, zero value otherwise.

### GetAnonCountOk

`func (o *GetOnlineUsersResponse) GetAnonCountOk() (*float64, bool)`

GetAnonCountOk returns a tuple with the AnonCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonCount

`func (o *GetOnlineUsersResponse) SetAnonCount(v float64)`

SetAnonCount sets AnonCount field to given value.


### GetUsers

`func (o *GetOnlineUsersResponse) GetUsers() []PageUserEntry`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetOnlineUsersResponse) GetUsersOk() (*[]PageUserEntry, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetOnlineUsersResponse) SetUsers(v []PageUserEntry)`

SetUsers sets Users field to given value.


### GetStatus

`func (o *GetOnlineUsersResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOnlineUsersResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOnlineUsersResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetReason

`func (o *GetOnlineUsersResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetOnlineUsersResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetOnlineUsersResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetOnlineUsersResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetOnlineUsersResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetOnlineUsersResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetOnlineUsersResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetOnlineUsersResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetOnlineUsersResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetOnlineUsersResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetOnlineUsersResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetOnlineUsersResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetOnlineUsersResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetOnlineUsersResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetOnlineUsersResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetOnlineUsersResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetOnlineUsersResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetOnlineUsersResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetOnlineUsersResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetOnlineUsersResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetOnlineUsersResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetOnlineUsersResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetOnlineUsersResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetOnlineUsersResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetOnlineUsersResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetOnlineUsersResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


