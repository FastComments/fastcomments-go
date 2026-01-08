# GetTenantPackagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**TenantPackages** | [**[]TenantPackage**](TenantPackage.md) |  | 

## Methods

### NewGetTenantPackagesResponse

`func NewGetTenantPackagesResponse(status APIStatus, tenantPackages []TenantPackage, ) *GetTenantPackagesResponse`

NewGetTenantPackagesResponse instantiates a new GetTenantPackagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantPackagesResponseWithDefaults

`func NewGetTenantPackagesResponseWithDefaults() *GetTenantPackagesResponse`

NewGetTenantPackagesResponseWithDefaults instantiates a new GetTenantPackagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetTenantPackagesResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTenantPackagesResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTenantPackagesResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTenantPackages

`func (o *GetTenantPackagesResponse) GetTenantPackages() []TenantPackage`

GetTenantPackages returns the TenantPackages field if non-nil, zero value otherwise.

### GetTenantPackagesOk

`func (o *GetTenantPackagesResponse) GetTenantPackagesOk() (*[]TenantPackage, bool)`

GetTenantPackagesOk returns a tuple with the TenantPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPackages

`func (o *GetTenantPackagesResponse) SetTenantPackages(v []TenantPackage)`

SetTenantPackages sets TenantPackages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


