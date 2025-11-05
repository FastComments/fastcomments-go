# APIGetUserBadgeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**UserBadge** | [**UserBadge**](UserBadge.md) |  | 

## Methods

### NewAPIGetUserBadgeResponse

`func NewAPIGetUserBadgeResponse(status APIStatus, userBadge UserBadge, ) *APIGetUserBadgeResponse`

NewAPIGetUserBadgeResponse instantiates a new APIGetUserBadgeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIGetUserBadgeResponseWithDefaults

`func NewAPIGetUserBadgeResponseWithDefaults() *APIGetUserBadgeResponse`

NewAPIGetUserBadgeResponseWithDefaults instantiates a new APIGetUserBadgeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *APIGetUserBadgeResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *APIGetUserBadgeResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *APIGetUserBadgeResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetUserBadge

`func (o *APIGetUserBadgeResponse) GetUserBadge() UserBadge`

GetUserBadge returns the UserBadge field if non-nil, zero value otherwise.

### GetUserBadgeOk

`func (o *APIGetUserBadgeResponse) GetUserBadgeOk() (*UserBadge, bool)`

GetUserBadgeOk returns a tuple with the UserBadge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBadge

`func (o *APIGetUserBadgeResponse) SetUserBadge(v UserBadge)`

SetUserBadge sets UserBadge field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


