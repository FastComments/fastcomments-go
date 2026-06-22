# PostVoteResponse

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

### NewPostVoteResponse

`func NewPostVoteResponse(status APIStatus, reason string, code string, ) *PostVoteResponse`

NewPostVoteResponse instantiates a new PostVoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostVoteResponseWithDefaults

`func NewPostVoteResponseWithDefaults() *PostVoteResponse`

NewPostVoteResponseWithDefaults instantiates a new PostVoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PostVoteResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostVoteResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostVoteResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetVoteId

`func (o *PostVoteResponse) GetVoteId() string`

GetVoteId returns the VoteId field if non-nil, zero value otherwise.

### GetVoteIdOk

`func (o *PostVoteResponse) GetVoteIdOk() (*string, bool)`

GetVoteIdOk returns a tuple with the VoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteId

`func (o *PostVoteResponse) SetVoteId(v string)`

SetVoteId sets VoteId field to given value.

### HasVoteId

`func (o *PostVoteResponse) HasVoteId() bool`

HasVoteId returns a boolean if a field has been set.

### GetIsVerified

`func (o *PostVoteResponse) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *PostVoteResponse) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *PostVoteResponse) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *PostVoteResponse) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetUser

`func (o *PostVoteResponse) GetUser() VoteResponseUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PostVoteResponse) GetUserOk() (*VoteResponseUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PostVoteResponse) SetUser(v VoteResponseUser)`

SetUser sets User field to given value.

### HasUser

`func (o *PostVoteResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetEditKey

`func (o *PostVoteResponse) GetEditKey() string`

GetEditKey returns the EditKey field if non-nil, zero value otherwise.

### GetEditKeyOk

`func (o *PostVoteResponse) GetEditKeyOk() (*string, bool)`

GetEditKeyOk returns a tuple with the EditKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditKey

`func (o *PostVoteResponse) SetEditKey(v string)`

SetEditKey sets EditKey field to given value.

### HasEditKey

`func (o *PostVoteResponse) HasEditKey() bool`

HasEditKey returns a boolean if a field has been set.

### GetReason

`func (o *PostVoteResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PostVoteResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PostVoteResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *PostVoteResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PostVoteResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PostVoteResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *PostVoteResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *PostVoteResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *PostVoteResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *PostVoteResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *PostVoteResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *PostVoteResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *PostVoteResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *PostVoteResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *PostVoteResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *PostVoteResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *PostVoteResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *PostVoteResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *PostVoteResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *PostVoteResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *PostVoteResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *PostVoteResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *PostVoteResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *PostVoteResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *PostVoteResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *PostVoteResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


