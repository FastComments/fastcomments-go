# PatchDomainConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | **interface{}** |  | 
**Status** | **interface{}** |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 

## Methods

### NewPatchDomainConfigResponse

`func NewPatchDomainConfigResponse(configuration interface{}, status interface{}, reason string, code string, ) *PatchDomainConfigResponse`

NewPatchDomainConfigResponse instantiates a new PatchDomainConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchDomainConfigResponseWithDefaults

`func NewPatchDomainConfigResponseWithDefaults() *PatchDomainConfigResponse`

NewPatchDomainConfigResponseWithDefaults instantiates a new PatchDomainConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *PatchDomainConfigResponse) GetConfiguration() interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PatchDomainConfigResponse) GetConfigurationOk() (*interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PatchDomainConfigResponse) SetConfiguration(v interface{})`

SetConfiguration sets Configuration field to given value.


### SetConfigurationNil

`func (o *PatchDomainConfigResponse) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *PatchDomainConfigResponse) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetStatus

`func (o *PatchDomainConfigResponse) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchDomainConfigResponse) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchDomainConfigResponse) SetStatus(v interface{})`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *PatchDomainConfigResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PatchDomainConfigResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetReason

`func (o *PatchDomainConfigResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PatchDomainConfigResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PatchDomainConfigResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *PatchDomainConfigResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PatchDomainConfigResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PatchDomainConfigResponse) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


