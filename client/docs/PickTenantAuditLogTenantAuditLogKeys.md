# PickTenantAuditLogTenantAuditLogKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**ResourceName** | **string** |  | 
**CrudType** | **string** |  | 
**From** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**When** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ServerStartDate** | Pointer to **time.Time** |  | [optional] 
**ObjectDetails** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewPickTenantAuditLogTenantAuditLogKeys

`func NewPickTenantAuditLogTenantAuditLogKeys(id string, resourceName string, crudType string, ) *PickTenantAuditLogTenantAuditLogKeys`

NewPickTenantAuditLogTenantAuditLogKeys instantiates a new PickTenantAuditLogTenantAuditLogKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickTenantAuditLogTenantAuditLogKeysWithDefaults

`func NewPickTenantAuditLogTenantAuditLogKeysWithDefaults() *PickTenantAuditLogTenantAuditLogKeys`

NewPickTenantAuditLogTenantAuditLogKeysWithDefaults instantiates a new PickTenantAuditLogTenantAuditLogKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PickTenantAuditLogTenantAuditLogKeys) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PickTenantAuditLogTenantAuditLogKeys) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PickTenantAuditLogTenantAuditLogKeys) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserId

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PickTenantAuditLogTenantAuditLogKeys) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PickTenantAuditLogTenantAuditLogKeys) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PickTenantAuditLogTenantAuditLogKeys) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PickTenantAuditLogTenantAuditLogKeys) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetResourceName

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *PickTenantAuditLogTenantAuditLogKeys) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetCrudType

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetCrudType() string`

GetCrudType returns the CrudType field if non-nil, zero value otherwise.

### GetCrudTypeOk

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetCrudTypeOk() (*string, bool)`

GetCrudTypeOk returns a tuple with the CrudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrudType

`func (o *PickTenantAuditLogTenantAuditLogKeys) SetCrudType(v string)`

SetCrudType sets CrudType field to given value.


### GetFrom

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PickTenantAuditLogTenantAuditLogKeys) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PickTenantAuditLogTenantAuditLogKeys) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetIp

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *PickTenantAuditLogTenantAuditLogKeys) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *PickTenantAuditLogTenantAuditLogKeys) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetWhen

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *PickTenantAuditLogTenantAuditLogKeys) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *PickTenantAuditLogTenantAuditLogKeys) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetDescription

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PickTenantAuditLogTenantAuditLogKeys) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PickTenantAuditLogTenantAuditLogKeys) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetServerStartDate

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetServerStartDate() time.Time`

GetServerStartDate returns the ServerStartDate field if non-nil, zero value otherwise.

### GetServerStartDateOk

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetServerStartDateOk() (*time.Time, bool)`

GetServerStartDateOk returns a tuple with the ServerStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStartDate

`func (o *PickTenantAuditLogTenantAuditLogKeys) SetServerStartDate(v time.Time)`

SetServerStartDate sets ServerStartDate field to given value.

### HasServerStartDate

`func (o *PickTenantAuditLogTenantAuditLogKeys) HasServerStartDate() bool`

HasServerStartDate returns a boolean if a field has been set.

### GetObjectDetails

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetObjectDetails() map[string]interface{}`

GetObjectDetails returns the ObjectDetails field if non-nil, zero value otherwise.

### GetObjectDetailsOk

`func (o *PickTenantAuditLogTenantAuditLogKeys) GetObjectDetailsOk() (*map[string]interface{}, bool)`

GetObjectDetailsOk returns a tuple with the ObjectDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDetails

`func (o *PickTenantAuditLogTenantAuditLogKeys) SetObjectDetails(v map[string]interface{})`

SetObjectDetails sets ObjectDetails field to given value.

### HasObjectDetails

`func (o *PickTenantAuditLogTenantAuditLogKeys) HasObjectDetails() bool`

HasObjectDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


