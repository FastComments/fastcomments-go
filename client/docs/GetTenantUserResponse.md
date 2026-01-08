# GetTenantUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**TenantUser** | [**User**](User.md) |  | 

## Methods

### NewGetTenantUserResponse

`func NewGetTenantUserResponse(status APIStatus, tenantUser User, ) *GetTenantUserResponse`

NewGetTenantUserResponse instantiates a new GetTenantUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantUserResponseWithDefaults

`func NewGetTenantUserResponseWithDefaults() *GetTenantUserResponse`

NewGetTenantUserResponseWithDefaults instantiates a new GetTenantUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetTenantUserResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTenantUserResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTenantUserResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTenantUser

`func (o *GetTenantUserResponse) GetTenantUser() User`

GetTenantUser returns the TenantUser field if non-nil, zero value otherwise.

### GetTenantUserOk

`func (o *GetTenantUserResponse) GetTenantUserOk() (*User, bool)`

GetTenantUserOk returns a tuple with the TenantUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUser

`func (o *GetTenantUserResponse) SetTenantUser(v User)`

SetTenantUser sets TenantUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


