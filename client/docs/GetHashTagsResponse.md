# GetHashTagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**HashTags** | [**[]TenantHashTag**](TenantHashTag.md) |  | 

## Methods

### NewGetHashTagsResponse

`func NewGetHashTagsResponse(status APIStatus, hashTags []TenantHashTag, ) *GetHashTagsResponse`

NewGetHashTagsResponse instantiates a new GetHashTagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHashTagsResponseWithDefaults

`func NewGetHashTagsResponseWithDefaults() *GetHashTagsResponse`

NewGetHashTagsResponseWithDefaults instantiates a new GetHashTagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetHashTagsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetHashTagsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetHashTagsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetHashTags

`func (o *GetHashTagsResponse) GetHashTags() []TenantHashTag`

GetHashTags returns the HashTags field if non-nil, zero value otherwise.

### GetHashTagsOk

`func (o *GetHashTagsResponse) GetHashTagsOk() (*[]TenantHashTag, bool)`

GetHashTagsOk returns a tuple with the HashTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashTags

`func (o *GetHashTagsResponse) SetHashTags(v []TenantHashTag)`

SetHashTags sets HashTags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


