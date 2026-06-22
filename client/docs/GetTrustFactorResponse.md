# GetTrustFactorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManualTrustFactor** | Pointer to **float64** |  | [optional] 
**AutoTrustFactor** | Pointer to **float64** |  | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetTrustFactorResponse

`func NewGetTrustFactorResponse(status APIStatus, reason string, code string, ) *GetTrustFactorResponse`

NewGetTrustFactorResponse instantiates a new GetTrustFactorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTrustFactorResponseWithDefaults

`func NewGetTrustFactorResponseWithDefaults() *GetTrustFactorResponse`

NewGetTrustFactorResponseWithDefaults instantiates a new GetTrustFactorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManualTrustFactor

`func (o *GetTrustFactorResponse) GetManualTrustFactor() float64`

GetManualTrustFactor returns the ManualTrustFactor field if non-nil, zero value otherwise.

### GetManualTrustFactorOk

`func (o *GetTrustFactorResponse) GetManualTrustFactorOk() (*float64, bool)`

GetManualTrustFactorOk returns a tuple with the ManualTrustFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualTrustFactor

`func (o *GetTrustFactorResponse) SetManualTrustFactor(v float64)`

SetManualTrustFactor sets ManualTrustFactor field to given value.

### HasManualTrustFactor

`func (o *GetTrustFactorResponse) HasManualTrustFactor() bool`

HasManualTrustFactor returns a boolean if a field has been set.

### GetAutoTrustFactor

`func (o *GetTrustFactorResponse) GetAutoTrustFactor() float64`

GetAutoTrustFactor returns the AutoTrustFactor field if non-nil, zero value otherwise.

### GetAutoTrustFactorOk

`func (o *GetTrustFactorResponse) GetAutoTrustFactorOk() (*float64, bool)`

GetAutoTrustFactorOk returns a tuple with the AutoTrustFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTrustFactor

`func (o *GetTrustFactorResponse) SetAutoTrustFactor(v float64)`

SetAutoTrustFactor sets AutoTrustFactor field to given value.

### HasAutoTrustFactor

`func (o *GetTrustFactorResponse) HasAutoTrustFactor() bool`

HasAutoTrustFactor returns a boolean if a field has been set.

### GetStatus

`func (o *GetTrustFactorResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTrustFactorResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTrustFactorResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetReason

`func (o *GetTrustFactorResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetTrustFactorResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetTrustFactorResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetTrustFactorResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetTrustFactorResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetTrustFactorResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetTrustFactorResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetTrustFactorResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetTrustFactorResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetTrustFactorResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetTrustFactorResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetTrustFactorResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetTrustFactorResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetTrustFactorResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetTrustFactorResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetTrustFactorResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetTrustFactorResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetTrustFactorResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetTrustFactorResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetTrustFactorResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetTrustFactorResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetTrustFactorResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetTrustFactorResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetTrustFactorResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetTrustFactorResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetTrustFactorResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


