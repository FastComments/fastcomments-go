# GetMyNotificationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Translations** | Pointer to **map[string]string** | Construct a type with a set of properties K of type T | [optional] 
**IsSubscribed** | **bool** |  | 
**HasMore** | **bool** |  | 
**Notifications** | [**[]RenderableUserNotification**](RenderableUserNotification.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewGetMyNotificationsResponse

`func NewGetMyNotificationsResponse(isSubscribed bool, hasMore bool, notifications []RenderableUserNotification, status APIStatus, ) *GetMyNotificationsResponse`

NewGetMyNotificationsResponse instantiates a new GetMyNotificationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMyNotificationsResponseWithDefaults

`func NewGetMyNotificationsResponseWithDefaults() *GetMyNotificationsResponse`

NewGetMyNotificationsResponseWithDefaults instantiates a new GetMyNotificationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTranslations

`func (o *GetMyNotificationsResponse) GetTranslations() map[string]string`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *GetMyNotificationsResponse) GetTranslationsOk() (*map[string]string, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *GetMyNotificationsResponse) SetTranslations(v map[string]string)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *GetMyNotificationsResponse) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *GetMyNotificationsResponse) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *GetMyNotificationsResponse) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *GetMyNotificationsResponse) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.


### GetHasMore

`func (o *GetMyNotificationsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetMyNotificationsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetMyNotificationsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNotifications

`func (o *GetMyNotificationsResponse) GetNotifications() []RenderableUserNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *GetMyNotificationsResponse) GetNotificationsOk() (*[]RenderableUserNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *GetMyNotificationsResponse) SetNotifications(v []RenderableUserNotification)`

SetNotifications sets Notifications field to given value.


### GetStatus

`func (o *GetMyNotificationsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMyNotificationsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMyNotificationsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


