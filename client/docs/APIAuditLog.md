# APIAuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**ResourceName** | **string** |  | 
**CrudType** | **string** |  | 
**From** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Ip** | Pointer to **NullableString** |  | [optional] 
**When** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ServerStartDate** | Pointer to **time.Time** |  | [optional] 
**ObjectDetails** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewAPIAuditLog

`func NewAPIAuditLog(id string, resourceName string, crudType string, ) *APIAuditLog`

NewAPIAuditLog instantiates a new APIAuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIAuditLogWithDefaults

`func NewAPIAuditLogWithDefaults() *APIAuditLog`

NewAPIAuditLogWithDefaults instantiates a new APIAuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APIAuditLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIAuditLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIAuditLog) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *APIAuditLog) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *APIAuditLog) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *APIAuditLog) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *APIAuditLog) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *APIAuditLog) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *APIAuditLog) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *APIAuditLog) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *APIAuditLog) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetResourceName

`func (o *APIAuditLog) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *APIAuditLog) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *APIAuditLog) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetCrudType

`func (o *APIAuditLog) GetCrudType() string`

GetCrudType returns the CrudType field if non-nil, zero value otherwise.

### GetCrudTypeOk

`func (o *APIAuditLog) GetCrudTypeOk() (*string, bool)`

GetCrudTypeOk returns a tuple with the CrudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrudType

`func (o *APIAuditLog) SetCrudType(v string)`

SetCrudType sets CrudType field to given value.


### GetFrom

`func (o *APIAuditLog) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *APIAuditLog) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *APIAuditLog) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *APIAuditLog) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetUrl

`func (o *APIAuditLog) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *APIAuditLog) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *APIAuditLog) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *APIAuditLog) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *APIAuditLog) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *APIAuditLog) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetIp

`func (o *APIAuditLog) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *APIAuditLog) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *APIAuditLog) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *APIAuditLog) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *APIAuditLog) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *APIAuditLog) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetWhen

`func (o *APIAuditLog) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *APIAuditLog) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *APIAuditLog) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *APIAuditLog) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetDescription

`func (o *APIAuditLog) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *APIAuditLog) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *APIAuditLog) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *APIAuditLog) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetServerStartDate

`func (o *APIAuditLog) GetServerStartDate() time.Time`

GetServerStartDate returns the ServerStartDate field if non-nil, zero value otherwise.

### GetServerStartDateOk

`func (o *APIAuditLog) GetServerStartDateOk() (*time.Time, bool)`

GetServerStartDateOk returns a tuple with the ServerStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStartDate

`func (o *APIAuditLog) SetServerStartDate(v time.Time)`

SetServerStartDate sets ServerStartDate field to given value.

### HasServerStartDate

`func (o *APIAuditLog) HasServerStartDate() bool`

HasServerStartDate returns a boolean if a field has been set.

### GetObjectDetails

`func (o *APIAuditLog) GetObjectDetails() map[string]interface{}`

GetObjectDetails returns the ObjectDetails field if non-nil, zero value otherwise.

### GetObjectDetailsOk

`func (o *APIAuditLog) GetObjectDetailsOk() (*map[string]interface{}, bool)`

GetObjectDetailsOk returns a tuple with the ObjectDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDetails

`func (o *APIAuditLog) SetObjectDetails(v map[string]interface{})`

SetObjectDetails sets ObjectDetails field to given value.

### HasObjectDetails

`func (o *APIAuditLog) HasObjectDetails() bool`

HasObjectDetails returns a boolean if a field has been set.

### SetObjectDetailsNil

`func (o *APIAuditLog) SetObjectDetailsNil(b bool)`

 SetObjectDetailsNil sets the value for ObjectDetails to be an explicit nil

### UnsetObjectDetails
`func (o *APIAuditLog) UnsetObjectDetails()`

UnsetObjectDetails ensures that no value is present for ObjectDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


