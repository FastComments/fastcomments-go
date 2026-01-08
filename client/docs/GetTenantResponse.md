# GetTenantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Tenant** | [**APITenant**](APITenant.md) |  | 

## Methods

### NewGetTenantResponse

`func NewGetTenantResponse(status APIStatus, tenant APITenant, ) *GetTenantResponse`

NewGetTenantResponse instantiates a new GetTenantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantResponseWithDefaults

`func NewGetTenantResponseWithDefaults() *GetTenantResponse`

NewGetTenantResponseWithDefaults instantiates a new GetTenantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetTenantResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTenantResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTenantResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *GetTenantResponse) GetTenant() APITenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *GetTenantResponse) GetTenantOk() (*APITenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *GetTenantResponse) SetTenant(v APITenant)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


