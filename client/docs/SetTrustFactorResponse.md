# SetTrustFactorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviousManualTrustFactor** | Pointer to **float64** |  | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewSetTrustFactorResponse

`func NewSetTrustFactorResponse(status APIStatus, reason string, code string, ) *SetTrustFactorResponse`

NewSetTrustFactorResponse instantiates a new SetTrustFactorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetTrustFactorResponseWithDefaults

`func NewSetTrustFactorResponseWithDefaults() *SetTrustFactorResponse`

NewSetTrustFactorResponseWithDefaults instantiates a new SetTrustFactorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviousManualTrustFactor

`func (o *SetTrustFactorResponse) GetPreviousManualTrustFactor() float64`

GetPreviousManualTrustFactor returns the PreviousManualTrustFactor field if non-nil, zero value otherwise.

### GetPreviousManualTrustFactorOk

`func (o *SetTrustFactorResponse) GetPreviousManualTrustFactorOk() (*float64, bool)`

GetPreviousManualTrustFactorOk returns a tuple with the PreviousManualTrustFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousManualTrustFactor

`func (o *SetTrustFactorResponse) SetPreviousManualTrustFactor(v float64)`

SetPreviousManualTrustFactor sets PreviousManualTrustFactor field to given value.

### HasPreviousManualTrustFactor

`func (o *SetTrustFactorResponse) HasPreviousManualTrustFactor() bool`

HasPreviousManualTrustFactor returns a boolean if a field has been set.

### GetStatus

`func (o *SetTrustFactorResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SetTrustFactorResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SetTrustFactorResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetReason

`func (o *SetTrustFactorResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SetTrustFactorResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SetTrustFactorResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *SetTrustFactorResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SetTrustFactorResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SetTrustFactorResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *SetTrustFactorResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *SetTrustFactorResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *SetTrustFactorResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *SetTrustFactorResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *SetTrustFactorResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *SetTrustFactorResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *SetTrustFactorResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *SetTrustFactorResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *SetTrustFactorResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *SetTrustFactorResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *SetTrustFactorResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *SetTrustFactorResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *SetTrustFactorResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *SetTrustFactorResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *SetTrustFactorResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *SetTrustFactorResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *SetTrustFactorResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *SetTrustFactorResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *SetTrustFactorResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *SetTrustFactorResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


