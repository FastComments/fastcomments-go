# AddDomainConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** |  | 
**Code** | **string** |  | 
**Status** | **interface{}** |  | 
**Configuration** | **interface{}** |  | 

## Methods

### NewAddDomainConfigResponse

`func NewAddDomainConfigResponse(reason string, code string, status interface{}, configuration interface{}, ) *AddDomainConfigResponse`

NewAddDomainConfigResponse instantiates a new AddDomainConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDomainConfigResponseWithDefaults

`func NewAddDomainConfigResponseWithDefaults() *AddDomainConfigResponse`

NewAddDomainConfigResponseWithDefaults instantiates a new AddDomainConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *AddDomainConfigResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AddDomainConfigResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AddDomainConfigResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *AddDomainConfigResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddDomainConfigResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddDomainConfigResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetStatus

`func (o *AddDomainConfigResponse) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddDomainConfigResponse) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddDomainConfigResponse) SetStatus(v interface{})`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *AddDomainConfigResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AddDomainConfigResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetConfiguration

`func (o *AddDomainConfigResponse) GetConfiguration() interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AddDomainConfigResponse) GetConfigurationOk() (*interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AddDomainConfigResponse) SetConfiguration(v interface{})`

SetConfiguration sets Configuration field to given value.


### SetConfigurationNil

`func (o *AddDomainConfigResponse) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *AddDomainConfigResponse) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


