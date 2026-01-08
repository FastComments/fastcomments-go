# GetNotificationCountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Count** | **float64** |  | 

## Methods

### NewGetNotificationCountResponse

`func NewGetNotificationCountResponse(status APIStatus, count float64, ) *GetNotificationCountResponse`

NewGetNotificationCountResponse instantiates a new GetNotificationCountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationCountResponseWithDefaults

`func NewGetNotificationCountResponseWithDefaults() *GetNotificationCountResponse`

NewGetNotificationCountResponseWithDefaults instantiates a new GetNotificationCountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetNotificationCountResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNotificationCountResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNotificationCountResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetCount

`func (o *GetNotificationCountResponse) GetCount() float64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetNotificationCountResponse) GetCountOk() (*float64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetNotificationCountResponse) SetCount(v float64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


