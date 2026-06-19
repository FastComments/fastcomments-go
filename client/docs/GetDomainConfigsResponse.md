# GetDomainConfigsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configurations** | **interface{}** |  | 
**Status** | **interface{}** |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 

## Methods

### NewGetDomainConfigsResponse

`func NewGetDomainConfigsResponse(configurations interface{}, status interface{}, reason string, code string, ) *GetDomainConfigsResponse`

NewGetDomainConfigsResponse instantiates a new GetDomainConfigsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainConfigsResponseWithDefaults

`func NewGetDomainConfigsResponseWithDefaults() *GetDomainConfigsResponse`

NewGetDomainConfigsResponseWithDefaults instantiates a new GetDomainConfigsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurations

`func (o *GetDomainConfigsResponse) GetConfigurations() interface{}`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *GetDomainConfigsResponse) GetConfigurationsOk() (*interface{}, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *GetDomainConfigsResponse) SetConfigurations(v interface{})`

SetConfigurations sets Configurations field to given value.


### SetConfigurationsNil

`func (o *GetDomainConfigsResponse) SetConfigurationsNil(b bool)`

 SetConfigurationsNil sets the value for Configurations to be an explicit nil

### UnsetConfigurations
`func (o *GetDomainConfigsResponse) UnsetConfigurations()`

UnsetConfigurations ensures that no value is present for Configurations, not even an explicit nil
### GetStatus

`func (o *GetDomainConfigsResponse) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDomainConfigsResponse) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDomainConfigsResponse) SetStatus(v interface{})`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *GetDomainConfigsResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetDomainConfigsResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetReason

`func (o *GetDomainConfigsResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetDomainConfigsResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetDomainConfigsResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetDomainConfigsResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetDomainConfigsResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetDomainConfigsResponse) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


