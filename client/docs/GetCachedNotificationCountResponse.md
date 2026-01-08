# GetCachedNotificationCountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Data** | [**UserNotificationCount**](UserNotificationCount.md) |  | 

## Methods

### NewGetCachedNotificationCountResponse

`func NewGetCachedNotificationCountResponse(status APIStatus, data UserNotificationCount, ) *GetCachedNotificationCountResponse`

NewGetCachedNotificationCountResponse instantiates a new GetCachedNotificationCountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCachedNotificationCountResponseWithDefaults

`func NewGetCachedNotificationCountResponseWithDefaults() *GetCachedNotificationCountResponse`

NewGetCachedNotificationCountResponseWithDefaults instantiates a new GetCachedNotificationCountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetCachedNotificationCountResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCachedNotificationCountResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCachedNotificationCountResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *GetCachedNotificationCountResponse) GetData() UserNotificationCount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCachedNotificationCountResponse) GetDataOk() (*UserNotificationCount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCachedNotificationCountResponse) SetData(v UserNotificationCount)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


