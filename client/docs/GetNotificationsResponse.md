# GetNotificationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Notifications** | [**[]UserNotification**](UserNotification.md) |  | 

## Methods

### NewGetNotificationsResponse

`func NewGetNotificationsResponse(status APIStatus, notifications []UserNotification, ) *GetNotificationsResponse`

NewGetNotificationsResponse instantiates a new GetNotificationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationsResponseWithDefaults

`func NewGetNotificationsResponseWithDefaults() *GetNotificationsResponse`

NewGetNotificationsResponseWithDefaults instantiates a new GetNotificationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetNotificationsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNotificationsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNotificationsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetNotifications

`func (o *GetNotificationsResponse) GetNotifications() []UserNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *GetNotificationsResponse) GetNotificationsOk() (*[]UserNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *GetNotificationsResponse) SetNotifications(v []UserNotification)`

SetNotifications sets Notifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


