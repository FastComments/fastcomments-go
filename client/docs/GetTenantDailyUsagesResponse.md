# GetTenantDailyUsagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**TenantDailyUsages** | [**[]APITenantDailyUsage**](APITenantDailyUsage.md) |  | 

## Methods

### NewGetTenantDailyUsagesResponse

`func NewGetTenantDailyUsagesResponse(status APIStatus, tenantDailyUsages []APITenantDailyUsage, ) *GetTenantDailyUsagesResponse`

NewGetTenantDailyUsagesResponse instantiates a new GetTenantDailyUsagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantDailyUsagesResponseWithDefaults

`func NewGetTenantDailyUsagesResponseWithDefaults() *GetTenantDailyUsagesResponse`

NewGetTenantDailyUsagesResponseWithDefaults instantiates a new GetTenantDailyUsagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetTenantDailyUsagesResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTenantDailyUsagesResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTenantDailyUsagesResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTenantDailyUsages

`func (o *GetTenantDailyUsagesResponse) GetTenantDailyUsages() []APITenantDailyUsage`

GetTenantDailyUsages returns the TenantDailyUsages field if non-nil, zero value otherwise.

### GetTenantDailyUsagesOk

`func (o *GetTenantDailyUsagesResponse) GetTenantDailyUsagesOk() (*[]APITenantDailyUsage, bool)`

GetTenantDailyUsagesOk returns a tuple with the TenantDailyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantDailyUsages

`func (o *GetTenantDailyUsagesResponse) SetTenantDailyUsages(v []APITenantDailyUsage)`

SetTenantDailyUsages sets TenantDailyUsages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


