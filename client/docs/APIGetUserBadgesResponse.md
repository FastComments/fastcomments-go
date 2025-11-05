# APIGetUserBadgesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**UserBadges** | [**[]UserBadge**](UserBadge.md) |  | 

## Methods

### NewAPIGetUserBadgesResponse

`func NewAPIGetUserBadgesResponse(status APIStatus, userBadges []UserBadge, ) *APIGetUserBadgesResponse`

NewAPIGetUserBadgesResponse instantiates a new APIGetUserBadgesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIGetUserBadgesResponseWithDefaults

`func NewAPIGetUserBadgesResponseWithDefaults() *APIGetUserBadgesResponse`

NewAPIGetUserBadgesResponseWithDefaults instantiates a new APIGetUserBadgesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *APIGetUserBadgesResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *APIGetUserBadgesResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *APIGetUserBadgesResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetUserBadges

`func (o *APIGetUserBadgesResponse) GetUserBadges() []UserBadge`

GetUserBadges returns the UserBadges field if non-nil, zero value otherwise.

### GetUserBadgesOk

`func (o *APIGetUserBadgesResponse) GetUserBadgesOk() (*[]UserBadge, bool)`

GetUserBadgesOk returns a tuple with the UserBadges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBadges

`func (o *APIGetUserBadgesResponse) SetUserBadges(v []UserBadge)`

SetUserBadges sets UserBadges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


