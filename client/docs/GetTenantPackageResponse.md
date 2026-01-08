# GetTenantPackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**TenantPackage** | [**TenantPackage**](TenantPackage.md) |  | 

## Methods

### NewGetTenantPackageResponse

`func NewGetTenantPackageResponse(status APIStatus, tenantPackage TenantPackage, ) *GetTenantPackageResponse`

NewGetTenantPackageResponse instantiates a new GetTenantPackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantPackageResponseWithDefaults

`func NewGetTenantPackageResponseWithDefaults() *GetTenantPackageResponse`

NewGetTenantPackageResponseWithDefaults instantiates a new GetTenantPackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetTenantPackageResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTenantPackageResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTenantPackageResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTenantPackage

`func (o *GetTenantPackageResponse) GetTenantPackage() TenantPackage`

GetTenantPackage returns the TenantPackage field if non-nil, zero value otherwise.

### GetTenantPackageOk

`func (o *GetTenantPackageResponse) GetTenantPackageOk() (*TenantPackage, bool)`

GetTenantPackageOk returns a tuple with the TenantPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPackage

`func (o *GetTenantPackageResponse) SetTenantPackage(v TenantPackage)`

SetTenantPackage sets TenantPackage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


