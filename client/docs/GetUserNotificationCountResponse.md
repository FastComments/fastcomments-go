# GetUserNotificationCountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Count** | **int64** |  | 

## Methods

### NewGetUserNotificationCountResponse

`func NewGetUserNotificationCountResponse(status APIStatus, count int64, ) *GetUserNotificationCountResponse`

NewGetUserNotificationCountResponse instantiates a new GetUserNotificationCountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserNotificationCountResponseWithDefaults

`func NewGetUserNotificationCountResponseWithDefaults() *GetUserNotificationCountResponse`

NewGetUserNotificationCountResponseWithDefaults instantiates a new GetUserNotificationCountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetUserNotificationCountResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUserNotificationCountResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUserNotificationCountResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetCount

`func (o *GetUserNotificationCountResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetUserNotificationCountResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetUserNotificationCountResponse) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


