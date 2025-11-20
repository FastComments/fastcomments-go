# HeaderState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**NotificationType** | **map[string]interface{}** |  | 
**UserId** | **string** |  | 
**UserIdWS** | **string** |  | 
**NotificationCounts** | [**[]NotificationAndCount**](NotificationAndCount.md) |  | 

## Methods

### NewHeaderState

`func NewHeaderState(status APIStatus, notificationType map[string]interface{}, userId string, userIdWS string, notificationCounts []NotificationAndCount, ) *HeaderState`

NewHeaderState instantiates a new HeaderState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderStateWithDefaults

`func NewHeaderStateWithDefaults() *HeaderState`

NewHeaderStateWithDefaults instantiates a new HeaderState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HeaderState) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HeaderState) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HeaderState) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetNotificationType

`func (o *HeaderState) GetNotificationType() map[string]interface{}`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *HeaderState) GetNotificationTypeOk() (*map[string]interface{}, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *HeaderState) SetNotificationType(v map[string]interface{})`

SetNotificationType sets NotificationType field to given value.


### GetUserId

`func (o *HeaderState) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HeaderState) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HeaderState) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserIdWS

`func (o *HeaderState) GetUserIdWS() string`

GetUserIdWS returns the UserIdWS field if non-nil, zero value otherwise.

### GetUserIdWSOk

`func (o *HeaderState) GetUserIdWSOk() (*string, bool)`

GetUserIdWSOk returns a tuple with the UserIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdWS

`func (o *HeaderState) SetUserIdWS(v string)`

SetUserIdWS sets UserIdWS field to given value.


### GetNotificationCounts

`func (o *HeaderState) GetNotificationCounts() []NotificationAndCount`

GetNotificationCounts returns the NotificationCounts field if non-nil, zero value otherwise.

### GetNotificationCountsOk

`func (o *HeaderState) GetNotificationCountsOk() (*[]NotificationAndCount, bool)`

GetNotificationCountsOk returns a tuple with the NotificationCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCounts

`func (o *HeaderState) SetNotificationCounts(v []NotificationAndCount)`

SetNotificationCounts sets NotificationCounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


