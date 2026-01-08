# GetTenantUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**TenantUsers** | [**[]User**](User.md) |  | 

## Methods

### NewGetTenantUsersResponse

`func NewGetTenantUsersResponse(status APIStatus, tenantUsers []User, ) *GetTenantUsersResponse`

NewGetTenantUsersResponse instantiates a new GetTenantUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantUsersResponseWithDefaults

`func NewGetTenantUsersResponseWithDefaults() *GetTenantUsersResponse`

NewGetTenantUsersResponseWithDefaults instantiates a new GetTenantUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetTenantUsersResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTenantUsersResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTenantUsersResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTenantUsers

`func (o *GetTenantUsersResponse) GetTenantUsers() []User`

GetTenantUsers returns the TenantUsers field if non-nil, zero value otherwise.

### GetTenantUsersOk

`func (o *GetTenantUsersResponse) GetTenantUsersOk() (*[]User, bool)`

GetTenantUsersOk returns a tuple with the TenantUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUsers

`func (o *GetTenantUsersResponse) SetTenantUsers(v []User)`

SetTenantUsers sets TenantUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


