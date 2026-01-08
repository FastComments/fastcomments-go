# TenantHashTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**TenantId** | **string** |  | 
**Tag** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewTenantHashTag

`func NewTenantHashTag(id string, createdAt time.Time, tenantId string, tag string, ) *TenantHashTag`

NewTenantHashTag instantiates a new TenantHashTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantHashTagWithDefaults

`func NewTenantHashTagWithDefaults() *TenantHashTag`

NewTenantHashTagWithDefaults instantiates a new TenantHashTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantHashTag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantHashTag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantHashTag) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *TenantHashTag) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TenantHashTag) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TenantHashTag) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTenantId

`func (o *TenantHashTag) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantHashTag) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantHashTag) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetTag

`func (o *TenantHashTag) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TenantHashTag) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TenantHashTag) SetTag(v string)`

SetTag sets Tag field to given value.


### GetUrl

`func (o *TenantHashTag) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TenantHashTag) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TenantHashTag) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TenantHashTag) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


