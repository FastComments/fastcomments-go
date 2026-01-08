# CreateHashTagBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**Tag** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateHashTagBody

`func NewCreateHashTagBody(tag string, ) *CreateHashTagBody`

NewCreateHashTagBody instantiates a new CreateHashTagBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHashTagBodyWithDefaults

`func NewCreateHashTagBodyWithDefaults() *CreateHashTagBody`

NewCreateHashTagBodyWithDefaults instantiates a new CreateHashTagBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateHashTagBody) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateHashTagBody) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateHashTagBody) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateHashTagBody) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTag

`func (o *CreateHashTagBody) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateHashTagBody) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateHashTagBody) SetTag(v string)`

SetTag sets Tag field to given value.


### GetUrl

`func (o *CreateHashTagBody) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateHashTagBody) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateHashTagBody) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateHashTagBody) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


