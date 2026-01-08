# CreateTenantPackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**TenantPackage** | [**TenantPackage**](TenantPackage.md) |  | 

## Methods

### NewCreateTenantPackageResponse

`func NewCreateTenantPackageResponse(status APIStatus, tenantPackage TenantPackage, ) *CreateTenantPackageResponse`

NewCreateTenantPackageResponse instantiates a new CreateTenantPackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenantPackageResponseWithDefaults

`func NewCreateTenantPackageResponseWithDefaults() *CreateTenantPackageResponse`

NewCreateTenantPackageResponseWithDefaults instantiates a new CreateTenantPackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateTenantPackageResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTenantPackageResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTenantPackageResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTenantPackage

`func (o *CreateTenantPackageResponse) GetTenantPackage() TenantPackage`

GetTenantPackage returns the TenantPackage field if non-nil, zero value otherwise.

### GetTenantPackageOk

`func (o *CreateTenantPackageResponse) GetTenantPackageOk() (*TenantPackage, bool)`

GetTenantPackageOk returns a tuple with the TenantPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPackage

`func (o *CreateTenantPackageResponse) SetTenantPackage(v TenantPackage)`

SetTenantPackage sets TenantPackage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


