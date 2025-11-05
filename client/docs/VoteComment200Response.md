# VoteComment200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ImportedAPIStatusFAILED**](ImportedAPIStatusFAILED.md) |  | 
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

### NewVoteComment200Response

`func NewVoteComment200Response(status ImportedAPIStatusFAILED, reason string, code string, ) *VoteComment200Response`

NewVoteComment200Response instantiates a new VoteComment200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoteComment200ResponseWithDefaults

`func NewVoteComment200ResponseWithDefaults() *VoteComment200Response`

NewVoteComment200ResponseWithDefaults instantiates a new VoteComment200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VoteComment200Response) GetStatus() ImportedAPIStatusFAILED`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VoteComment200Response) GetStatusOk() (*ImportedAPIStatusFAILED, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VoteComment200Response) SetStatus(v ImportedAPIStatusFAILED)`

SetStatus sets Status field to given value.


### GetVoteId

`func (o *VoteComment200Response) GetVoteId() string`

GetVoteId returns the VoteId field if non-nil, zero value otherwise.

### GetVoteIdOk

`func (o *VoteComment200Response) GetVoteIdOk() (*string, bool)`

GetVoteIdOk returns a tuple with the VoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteId

`func (o *VoteComment200Response) SetVoteId(v string)`

SetVoteId sets VoteId field to given value.

### HasVoteId

`func (o *VoteComment200Response) HasVoteId() bool`

HasVoteId returns a boolean if a field has been set.

### GetIsVerified

`func (o *VoteComment200Response) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *VoteComment200Response) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *VoteComment200Response) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *VoteComment200Response) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetUser

`func (o *VoteComment200Response) GetUser() VoteResponseUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *VoteComment200Response) GetUserOk() (*VoteResponseUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *VoteComment200Response) SetUser(v VoteResponseUser)`

SetUser sets User field to given value.

### HasUser

`func (o *VoteComment200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetEditKey

`func (o *VoteComment200Response) GetEditKey() string`

GetEditKey returns the EditKey field if non-nil, zero value otherwise.

### GetEditKeyOk

`func (o *VoteComment200Response) GetEditKeyOk() (*string, bool)`

GetEditKeyOk returns a tuple with the EditKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditKey

`func (o *VoteComment200Response) SetEditKey(v string)`

SetEditKey sets EditKey field to given value.

### HasEditKey

`func (o *VoteComment200Response) HasEditKey() bool`

HasEditKey returns a boolean if a field has been set.

### GetReason

`func (o *VoteComment200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *VoteComment200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *VoteComment200Response) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *VoteComment200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VoteComment200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VoteComment200Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *VoteComment200Response) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *VoteComment200Response) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *VoteComment200Response) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *VoteComment200Response) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *VoteComment200Response) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *VoteComment200Response) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *VoteComment200Response) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *VoteComment200Response) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *VoteComment200Response) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *VoteComment200Response) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *VoteComment200Response) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *VoteComment200Response) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *VoteComment200Response) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *VoteComment200Response) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *VoteComment200Response) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *VoteComment200Response) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *VoteComment200Response) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *VoteComment200Response) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *VoteComment200Response) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *VoteComment200Response) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


