# APIGetUserBadgeProgressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**UserBadgeProgress** | [**UserBadgeProgress**](UserBadgeProgress.md) |  | 

## Methods

### NewAPIGetUserBadgeProgressResponse

`func NewAPIGetUserBadgeProgressResponse(status APIStatus, userBadgeProgress UserBadgeProgress, ) *APIGetUserBadgeProgressResponse`

NewAPIGetUserBadgeProgressResponse instantiates a new APIGetUserBadgeProgressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIGetUserBadgeProgressResponseWithDefaults

`func NewAPIGetUserBadgeProgressResponseWithDefaults() *APIGetUserBadgeProgressResponse`

NewAPIGetUserBadgeProgressResponseWithDefaults instantiates a new APIGetUserBadgeProgressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *APIGetUserBadgeProgressResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *APIGetUserBadgeProgressResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *APIGetUserBadgeProgressResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetUserBadgeProgress

`func (o *APIGetUserBadgeProgressResponse) GetUserBadgeProgress() UserBadgeProgress`

GetUserBadgeProgress returns the UserBadgeProgress field if non-nil, zero value otherwise.

### GetUserBadgeProgressOk

`func (o *APIGetUserBadgeProgressResponse) GetUserBadgeProgressOk() (*UserBadgeProgress, bool)`

GetUserBadgeProgressOk returns a tuple with the UserBadgeProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBadgeProgress

`func (o *APIGetUserBadgeProgressResponse) SetUserBadgeProgress(v UserBadgeProgress)`

SetUserBadgeProgress sets UserBadgeProgress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


