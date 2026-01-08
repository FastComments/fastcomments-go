# BulkCreateHashTagsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**Tags** | [**[]BulkCreateHashTagsBodyTagsInner**](BulkCreateHashTagsBodyTagsInner.md) |  | 

## Methods

### NewBulkCreateHashTagsBody

`func NewBulkCreateHashTagsBody(tags []BulkCreateHashTagsBodyTagsInner, ) *BulkCreateHashTagsBody`

NewBulkCreateHashTagsBody instantiates a new BulkCreateHashTagsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCreateHashTagsBodyWithDefaults

`func NewBulkCreateHashTagsBodyWithDefaults() *BulkCreateHashTagsBody`

NewBulkCreateHashTagsBodyWithDefaults instantiates a new BulkCreateHashTagsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *BulkCreateHashTagsBody) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BulkCreateHashTagsBody) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BulkCreateHashTagsBody) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BulkCreateHashTagsBody) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTags

`func (o *BulkCreateHashTagsBody) GetTags() []BulkCreateHashTagsBodyTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkCreateHashTagsBody) GetTagsOk() (*[]BulkCreateHashTagsBodyTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkCreateHashTagsBody) SetTags(v []BulkCreateHashTagsBodyTagsInner)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


