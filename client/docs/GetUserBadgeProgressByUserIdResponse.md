# GetUserBadgeProgressByUserIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**UserBadgeProgress** | [**UserBadgeProgress**](UserBadgeProgress.md) |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetUserBadgeProgressByUserIdResponse

`func NewGetUserBadgeProgressByUserIdResponse(status APIStatus, userBadgeProgress UserBadgeProgress, reason string, code string, ) *GetUserBadgeProgressByUserIdResponse`

NewGetUserBadgeProgressByUserIdResponse instantiates a new GetUserBadgeProgressByUserIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserBadgeProgressByUserIdResponseWithDefaults

`func NewGetUserBadgeProgressByUserIdResponseWithDefaults() *GetUserBadgeProgressByUserIdResponse`

NewGetUserBadgeProgressByUserIdResponseWithDefaults instantiates a new GetUserBadgeProgressByUserIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetUserBadgeProgressByUserIdResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUserBadgeProgressByUserIdResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUserBadgeProgressByUserIdResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetUserBadgeProgress

`func (o *GetUserBadgeProgressByUserIdResponse) GetUserBadgeProgress() UserBadgeProgress`

GetUserBadgeProgress returns the UserBadgeProgress field if non-nil, zero value otherwise.

### GetUserBadgeProgressOk

`func (o *GetUserBadgeProgressByUserIdResponse) GetUserBadgeProgressOk() (*UserBadgeProgress, bool)`

GetUserBadgeProgressOk returns a tuple with the UserBadgeProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBadgeProgress

`func (o *GetUserBadgeProgressByUserIdResponse) SetUserBadgeProgress(v UserBadgeProgress)`

SetUserBadgeProgress sets UserBadgeProgress field to given value.


### GetReason

`func (o *GetUserBadgeProgressByUserIdResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetUserBadgeProgressByUserIdResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetUserBadgeProgressByUserIdResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetUserBadgeProgressByUserIdResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetUserBadgeProgressByUserIdResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetUserBadgeProgressByUserIdResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetUserBadgeProgressByUserIdResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetUserBadgeProgressByUserIdResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetUserBadgeProgressByUserIdResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetUserBadgeProgressByUserIdResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetUserBadgeProgressByUserIdResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetUserBadgeProgressByUserIdResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetUserBadgeProgressByUserIdResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetUserBadgeProgressByUserIdResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetUserBadgeProgressByUserIdResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetUserBadgeProgressByUserIdResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetUserBadgeProgressByUserIdResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetUserBadgeProgressByUserIdResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetUserBadgeProgressByUserIdResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetUserBadgeProgressByUserIdResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetUserBadgeProgressByUserIdResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetUserBadgeProgressByUserIdResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetUserBadgeProgressByUserIdResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetUserBadgeProgressByUserIdResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetUserBadgeProgressByUserIdResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetUserBadgeProgressByUserIdResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


