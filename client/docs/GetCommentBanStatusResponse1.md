# GetCommentBanStatusResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**EmailDomain** | **NullableString** |  | 
**CanIPBan** | **NullableBool** |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetCommentBanStatusResponse1

`func NewGetCommentBanStatusResponse1(status APIStatus, emailDomain NullableString, canIPBan NullableBool, reason string, code string, ) *GetCommentBanStatusResponse1`

NewGetCommentBanStatusResponse1 instantiates a new GetCommentBanStatusResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommentBanStatusResponse1WithDefaults

`func NewGetCommentBanStatusResponse1WithDefaults() *GetCommentBanStatusResponse1`

NewGetCommentBanStatusResponse1WithDefaults instantiates a new GetCommentBanStatusResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetCommentBanStatusResponse1) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCommentBanStatusResponse1) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCommentBanStatusResponse1) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetEmailDomain

`func (o *GetCommentBanStatusResponse1) GetEmailDomain() string`

GetEmailDomain returns the EmailDomain field if non-nil, zero value otherwise.

### GetEmailDomainOk

`func (o *GetCommentBanStatusResponse1) GetEmailDomainOk() (*string, bool)`

GetEmailDomainOk returns a tuple with the EmailDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomain

`func (o *GetCommentBanStatusResponse1) SetEmailDomain(v string)`

SetEmailDomain sets EmailDomain field to given value.


### SetEmailDomainNil

`func (o *GetCommentBanStatusResponse1) SetEmailDomainNil(b bool)`

 SetEmailDomainNil sets the value for EmailDomain to be an explicit nil

### UnsetEmailDomain
`func (o *GetCommentBanStatusResponse1) UnsetEmailDomain()`

UnsetEmailDomain ensures that no value is present for EmailDomain, not even an explicit nil
### GetCanIPBan

`func (o *GetCommentBanStatusResponse1) GetCanIPBan() bool`

GetCanIPBan returns the CanIPBan field if non-nil, zero value otherwise.

### GetCanIPBanOk

`func (o *GetCommentBanStatusResponse1) GetCanIPBanOk() (*bool, bool)`

GetCanIPBanOk returns a tuple with the CanIPBan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanIPBan

`func (o *GetCommentBanStatusResponse1) SetCanIPBan(v bool)`

SetCanIPBan sets CanIPBan field to given value.


### SetCanIPBanNil

`func (o *GetCommentBanStatusResponse1) SetCanIPBanNil(b bool)`

 SetCanIPBanNil sets the value for CanIPBan to be an explicit nil

### UnsetCanIPBan
`func (o *GetCommentBanStatusResponse1) UnsetCanIPBan()`

UnsetCanIPBan ensures that no value is present for CanIPBan, not even an explicit nil
### GetReason

`func (o *GetCommentBanStatusResponse1) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetCommentBanStatusResponse1) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetCommentBanStatusResponse1) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetCommentBanStatusResponse1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetCommentBanStatusResponse1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetCommentBanStatusResponse1) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetCommentBanStatusResponse1) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetCommentBanStatusResponse1) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetCommentBanStatusResponse1) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetCommentBanStatusResponse1) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetCommentBanStatusResponse1) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetCommentBanStatusResponse1) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetCommentBanStatusResponse1) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetCommentBanStatusResponse1) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetCommentBanStatusResponse1) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetCommentBanStatusResponse1) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetCommentBanStatusResponse1) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetCommentBanStatusResponse1) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetCommentBanStatusResponse1) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetCommentBanStatusResponse1) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetCommentBanStatusResponse1) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetCommentBanStatusResponse1) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetCommentBanStatusResponse1) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetCommentBanStatusResponse1) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetCommentBanStatusResponse1) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetCommentBanStatusResponse1) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


