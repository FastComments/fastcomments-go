# SetUserTrustFactorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviousManualTrustFactor** | Pointer to **float64** |  | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewSetUserTrustFactorResponse

`func NewSetUserTrustFactorResponse(status APIStatus, ) *SetUserTrustFactorResponse`

NewSetUserTrustFactorResponse instantiates a new SetUserTrustFactorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetUserTrustFactorResponseWithDefaults

`func NewSetUserTrustFactorResponseWithDefaults() *SetUserTrustFactorResponse`

NewSetUserTrustFactorResponseWithDefaults instantiates a new SetUserTrustFactorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviousManualTrustFactor

`func (o *SetUserTrustFactorResponse) GetPreviousManualTrustFactor() float64`

GetPreviousManualTrustFactor returns the PreviousManualTrustFactor field if non-nil, zero value otherwise.

### GetPreviousManualTrustFactorOk

`func (o *SetUserTrustFactorResponse) GetPreviousManualTrustFactorOk() (*float64, bool)`

GetPreviousManualTrustFactorOk returns a tuple with the PreviousManualTrustFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousManualTrustFactor

`func (o *SetUserTrustFactorResponse) SetPreviousManualTrustFactor(v float64)`

SetPreviousManualTrustFactor sets PreviousManualTrustFactor field to given value.

### HasPreviousManualTrustFactor

`func (o *SetUserTrustFactorResponse) HasPreviousManualTrustFactor() bool`

HasPreviousManualTrustFactor returns a boolean if a field has been set.

### GetStatus

`func (o *SetUserTrustFactorResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SetUserTrustFactorResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SetUserTrustFactorResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


