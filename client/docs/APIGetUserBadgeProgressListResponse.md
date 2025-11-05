# APIGetUserBadgeProgressListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**UserBadgeProgresses** | [**[]UserBadgeProgress**](UserBadgeProgress.md) |  | 

## Methods

### NewAPIGetUserBadgeProgressListResponse

`func NewAPIGetUserBadgeProgressListResponse(status APIStatus, userBadgeProgresses []UserBadgeProgress, ) *APIGetUserBadgeProgressListResponse`

NewAPIGetUserBadgeProgressListResponse instantiates a new APIGetUserBadgeProgressListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIGetUserBadgeProgressListResponseWithDefaults

`func NewAPIGetUserBadgeProgressListResponseWithDefaults() *APIGetUserBadgeProgressListResponse`

NewAPIGetUserBadgeProgressListResponseWithDefaults instantiates a new APIGetUserBadgeProgressListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *APIGetUserBadgeProgressListResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *APIGetUserBadgeProgressListResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *APIGetUserBadgeProgressListResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetUserBadgeProgresses

`func (o *APIGetUserBadgeProgressListResponse) GetUserBadgeProgresses() []UserBadgeProgress`

GetUserBadgeProgresses returns the UserBadgeProgresses field if non-nil, zero value otherwise.

### GetUserBadgeProgressesOk

`func (o *APIGetUserBadgeProgressListResponse) GetUserBadgeProgressesOk() (*[]UserBadgeProgress, bool)`

GetUserBadgeProgressesOk returns a tuple with the UserBadgeProgresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBadgeProgresses

`func (o *APIGetUserBadgeProgressListResponse) SetUserBadgeProgresses(v []UserBadgeProgress)`

SetUserBadgeProgresses sets UserBadgeProgresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


