# GetTenantManualBadgesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Badges** | [**[]TenantBadge**](TenantBadge.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewGetTenantManualBadgesResponse

`func NewGetTenantManualBadgesResponse(badges []TenantBadge, status APIStatus, ) *GetTenantManualBadgesResponse`

NewGetTenantManualBadgesResponse instantiates a new GetTenantManualBadgesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantManualBadgesResponseWithDefaults

`func NewGetTenantManualBadgesResponseWithDefaults() *GetTenantManualBadgesResponse`

NewGetTenantManualBadgesResponseWithDefaults instantiates a new GetTenantManualBadgesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadges

`func (o *GetTenantManualBadgesResponse) GetBadges() []TenantBadge`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *GetTenantManualBadgesResponse) GetBadgesOk() (*[]TenantBadge, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *GetTenantManualBadgesResponse) SetBadges(v []TenantBadge)`

SetBadges sets Badges field to given value.


### GetStatus

`func (o *GetTenantManualBadgesResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTenantManualBadgesResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTenantManualBadgesResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


