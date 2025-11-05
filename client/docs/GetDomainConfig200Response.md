# GetDomainConfig200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | **interface{}** |  | 
**Status** | **interface{}** |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 

## Methods

### NewGetDomainConfig200Response

`func NewGetDomainConfig200Response(configuration interface{}, status interface{}, reason string, code string, ) *GetDomainConfig200Response`

NewGetDomainConfig200Response instantiates a new GetDomainConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainConfig200ResponseWithDefaults

`func NewGetDomainConfig200ResponseWithDefaults() *GetDomainConfig200Response`

NewGetDomainConfig200ResponseWithDefaults instantiates a new GetDomainConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *GetDomainConfig200Response) GetConfiguration() interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetDomainConfig200Response) GetConfigurationOk() (*interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetDomainConfig200Response) SetConfiguration(v interface{})`

SetConfiguration sets Configuration field to given value.


### SetConfigurationNil

`func (o *GetDomainConfig200Response) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *GetDomainConfig200Response) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetStatus

`func (o *GetDomainConfig200Response) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDomainConfig200Response) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDomainConfig200Response) SetStatus(v interface{})`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *GetDomainConfig200Response) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetDomainConfig200Response) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetReason

`func (o *GetDomainConfig200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetDomainConfig200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetDomainConfig200Response) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetDomainConfig200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetDomainConfig200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetDomainConfig200Response) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


