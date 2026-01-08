# UpdateHashTagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**HashTag** | [**TenantHashTag**](TenantHashTag.md) |  | 

## Methods

### NewUpdateHashTagResponse

`func NewUpdateHashTagResponse(status APIStatus, hashTag TenantHashTag, ) *UpdateHashTagResponse`

NewUpdateHashTagResponse instantiates a new UpdateHashTagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHashTagResponseWithDefaults

`func NewUpdateHashTagResponseWithDefaults() *UpdateHashTagResponse`

NewUpdateHashTagResponseWithDefaults instantiates a new UpdateHashTagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateHashTagResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateHashTagResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateHashTagResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetHashTag

`func (o *UpdateHashTagResponse) GetHashTag() TenantHashTag`

GetHashTag returns the HashTag field if non-nil, zero value otherwise.

### GetHashTagOk

`func (o *UpdateHashTagResponse) GetHashTagOk() (*TenantHashTag, bool)`

GetHashTagOk returns a tuple with the HashTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashTag

`func (o *UpdateHashTagResponse) SetHashTag(v TenantHashTag)`

SetHashTag sets HashTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


