# GetV2PageReactsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReactedIds** | Pointer to **[]string** |  | [optional] 
**Counts** | Pointer to **map[string]float64** | Construct a type with a set of properties K of type T | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetV2PageReactsResponse

`func NewGetV2PageReactsResponse(status APIStatus, reason string, code string, ) *GetV2PageReactsResponse`

NewGetV2PageReactsResponse instantiates a new GetV2PageReactsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetV2PageReactsResponseWithDefaults

`func NewGetV2PageReactsResponseWithDefaults() *GetV2PageReactsResponse`

NewGetV2PageReactsResponseWithDefaults instantiates a new GetV2PageReactsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReactedIds

`func (o *GetV2PageReactsResponse) GetReactedIds() []string`

GetReactedIds returns the ReactedIds field if non-nil, zero value otherwise.

### GetReactedIdsOk

`func (o *GetV2PageReactsResponse) GetReactedIdsOk() (*[]string, bool)`

GetReactedIdsOk returns a tuple with the ReactedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactedIds

`func (o *GetV2PageReactsResponse) SetReactedIds(v []string)`

SetReactedIds sets ReactedIds field to given value.

### HasReactedIds

`func (o *GetV2PageReactsResponse) HasReactedIds() bool`

HasReactedIds returns a boolean if a field has been set.

### GetCounts

`func (o *GetV2PageReactsResponse) GetCounts() map[string]float64`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *GetV2PageReactsResponse) GetCountsOk() (*map[string]float64, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *GetV2PageReactsResponse) SetCounts(v map[string]float64)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *GetV2PageReactsResponse) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetStatus

`func (o *GetV2PageReactsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetV2PageReactsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetV2PageReactsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetReason

`func (o *GetV2PageReactsResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetV2PageReactsResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetV2PageReactsResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetV2PageReactsResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetV2PageReactsResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetV2PageReactsResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetV2PageReactsResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetV2PageReactsResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetV2PageReactsResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetV2PageReactsResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetV2PageReactsResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetV2PageReactsResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetV2PageReactsResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetV2PageReactsResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetV2PageReactsResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetV2PageReactsResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetV2PageReactsResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetV2PageReactsResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetV2PageReactsResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetV2PageReactsResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetV2PageReactsResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetV2PageReactsResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetV2PageReactsResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetV2PageReactsResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetV2PageReactsResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetV2PageReactsResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


