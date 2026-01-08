# GetTenantsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Tenants** | [**[]APITenant**](APITenant.md) |  | 

## Methods

### NewGetTenantsResponse

`func NewGetTenantsResponse(status APIStatus, tenants []APITenant, ) *GetTenantsResponse`

NewGetTenantsResponse instantiates a new GetTenantsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantsResponseWithDefaults

`func NewGetTenantsResponseWithDefaults() *GetTenantsResponse`

NewGetTenantsResponseWithDefaults instantiates a new GetTenantsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetTenantsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTenantsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTenantsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTenants

`func (o *GetTenantsResponse) GetTenants() []APITenant`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetTenantsResponse) GetTenantsOk() (*[]APITenant, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetTenantsResponse) SetTenants(v []APITenant)`

SetTenants sets Tenants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


