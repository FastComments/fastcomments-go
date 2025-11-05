# GetDomainConfigs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configurations** | **interface{}** |  | 
**Status** | **interface{}** |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 

## Methods

### NewGetDomainConfigs200Response

`func NewGetDomainConfigs200Response(configurations interface{}, status interface{}, reason string, code string, ) *GetDomainConfigs200Response`

NewGetDomainConfigs200Response instantiates a new GetDomainConfigs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainConfigs200ResponseWithDefaults

`func NewGetDomainConfigs200ResponseWithDefaults() *GetDomainConfigs200Response`

NewGetDomainConfigs200ResponseWithDefaults instantiates a new GetDomainConfigs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurations

`func (o *GetDomainConfigs200Response) GetConfigurations() interface{}`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *GetDomainConfigs200Response) GetConfigurationsOk() (*interface{}, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *GetDomainConfigs200Response) SetConfigurations(v interface{})`

SetConfigurations sets Configurations field to given value.


### SetConfigurationsNil

`func (o *GetDomainConfigs200Response) SetConfigurationsNil(b bool)`

 SetConfigurationsNil sets the value for Configurations to be an explicit nil

### UnsetConfigurations
`func (o *GetDomainConfigs200Response) UnsetConfigurations()`

UnsetConfigurations ensures that no value is present for Configurations, not even an explicit nil
### GetStatus

`func (o *GetDomainConfigs200Response) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDomainConfigs200Response) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDomainConfigs200Response) SetStatus(v interface{})`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *GetDomainConfigs200Response) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetDomainConfigs200Response) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetReason

`func (o *GetDomainConfigs200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetDomainConfigs200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetDomainConfigs200Response) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetDomainConfigs200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetDomainConfigs200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetDomainConfigs200Response) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


