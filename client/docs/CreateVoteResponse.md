# CreateVoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**VoteId** | Pointer to **string** |  | [optional] 
**IsVerified** | Pointer to **bool** |  | [optional] 
**User** | Pointer to [**VoteResponseUser**](VoteResponseUser.md) |  | [optional] 
**EditKey** | Pointer to **string** |  | [optional] 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewCreateVoteResponse

`func NewCreateVoteResponse(status APIStatus, reason string, code string, ) *CreateVoteResponse`

NewCreateVoteResponse instantiates a new CreateVoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVoteResponseWithDefaults

`func NewCreateVoteResponseWithDefaults() *CreateVoteResponse`

NewCreateVoteResponseWithDefaults instantiates a new CreateVoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateVoteResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateVoteResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateVoteResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetVoteId

`func (o *CreateVoteResponse) GetVoteId() string`

GetVoteId returns the VoteId field if non-nil, zero value otherwise.

### GetVoteIdOk

`func (o *CreateVoteResponse) GetVoteIdOk() (*string, bool)`

GetVoteIdOk returns a tuple with the VoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteId

`func (o *CreateVoteResponse) SetVoteId(v string)`

SetVoteId sets VoteId field to given value.

### HasVoteId

`func (o *CreateVoteResponse) HasVoteId() bool`

HasVoteId returns a boolean if a field has been set.

### GetIsVerified

`func (o *CreateVoteResponse) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *CreateVoteResponse) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *CreateVoteResponse) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *CreateVoteResponse) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetUser

`func (o *CreateVoteResponse) GetUser() VoteResponseUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateVoteResponse) GetUserOk() (*VoteResponseUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateVoteResponse) SetUser(v VoteResponseUser)`

SetUser sets User field to given value.

### HasUser

`func (o *CreateVoteResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetEditKey

`func (o *CreateVoteResponse) GetEditKey() string`

GetEditKey returns the EditKey field if non-nil, zero value otherwise.

### GetEditKeyOk

`func (o *CreateVoteResponse) GetEditKeyOk() (*string, bool)`

GetEditKeyOk returns a tuple with the EditKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditKey

`func (o *CreateVoteResponse) SetEditKey(v string)`

SetEditKey sets EditKey field to given value.

### HasEditKey

`func (o *CreateVoteResponse) HasEditKey() bool`

HasEditKey returns a boolean if a field has been set.

### GetReason

`func (o *CreateVoteResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateVoteResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateVoteResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *CreateVoteResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateVoteResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateVoteResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *CreateVoteResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *CreateVoteResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *CreateVoteResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *CreateVoteResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *CreateVoteResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *CreateVoteResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *CreateVoteResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *CreateVoteResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *CreateVoteResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *CreateVoteResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *CreateVoteResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *CreateVoteResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *CreateVoteResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *CreateVoteResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *CreateVoteResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *CreateVoteResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *CreateVoteResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *CreateVoteResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *CreateVoteResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *CreateVoteResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


