# GetUserTrustFactorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManualTrustFactor** | Pointer to **float64** |  | [optional] 
**AutoTrustFactor** | Pointer to **float64** |  | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewGetUserTrustFactorResponse

`func NewGetUserTrustFactorResponse(status APIStatus, ) *GetUserTrustFactorResponse`

NewGetUserTrustFactorResponse instantiates a new GetUserTrustFactorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserTrustFactorResponseWithDefaults

`func NewGetUserTrustFactorResponseWithDefaults() *GetUserTrustFactorResponse`

NewGetUserTrustFactorResponseWithDefaults instantiates a new GetUserTrustFactorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManualTrustFactor

`func (o *GetUserTrustFactorResponse) GetManualTrustFactor() float64`

GetManualTrustFactor returns the ManualTrustFactor field if non-nil, zero value otherwise.

### GetManualTrustFactorOk

`func (o *GetUserTrustFactorResponse) GetManualTrustFactorOk() (*float64, bool)`

GetManualTrustFactorOk returns a tuple with the ManualTrustFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualTrustFactor

`func (o *GetUserTrustFactorResponse) SetManualTrustFactor(v float64)`

SetManualTrustFactor sets ManualTrustFactor field to given value.

### HasManualTrustFactor

`func (o *GetUserTrustFactorResponse) HasManualTrustFactor() bool`

HasManualTrustFactor returns a boolean if a field has been set.

### GetAutoTrustFactor

`func (o *GetUserTrustFactorResponse) GetAutoTrustFactor() float64`

GetAutoTrustFactor returns the AutoTrustFactor field if non-nil, zero value otherwise.

### GetAutoTrustFactorOk

`func (o *GetUserTrustFactorResponse) GetAutoTrustFactorOk() (*float64, bool)`

GetAutoTrustFactorOk returns a tuple with the AutoTrustFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTrustFactor

`func (o *GetUserTrustFactorResponse) SetAutoTrustFactor(v float64)`

SetAutoTrustFactor sets AutoTrustFactor field to given value.

### HasAutoTrustFactor

`func (o *GetUserTrustFactorResponse) HasAutoTrustFactor() bool`

HasAutoTrustFactor returns a boolean if a field has been set.

### GetStatus

`func (o *GetUserTrustFactorResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUserTrustFactorResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUserTrustFactorResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


