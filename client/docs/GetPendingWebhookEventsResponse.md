# GetPendingWebhookEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**PendingWebhookEvents** | [**[]PendingCommentToSyncOutbound**](PendingCommentToSyncOutbound.md) |  | 

## Methods

### NewGetPendingWebhookEventsResponse

`func NewGetPendingWebhookEventsResponse(status APIStatus, pendingWebhookEvents []PendingCommentToSyncOutbound, ) *GetPendingWebhookEventsResponse`

NewGetPendingWebhookEventsResponse instantiates a new GetPendingWebhookEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPendingWebhookEventsResponseWithDefaults

`func NewGetPendingWebhookEventsResponseWithDefaults() *GetPendingWebhookEventsResponse`

NewGetPendingWebhookEventsResponseWithDefaults instantiates a new GetPendingWebhookEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetPendingWebhookEventsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPendingWebhookEventsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPendingWebhookEventsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetPendingWebhookEvents

`func (o *GetPendingWebhookEventsResponse) GetPendingWebhookEvents() []PendingCommentToSyncOutbound`

GetPendingWebhookEvents returns the PendingWebhookEvents field if non-nil, zero value otherwise.

### GetPendingWebhookEventsOk

`func (o *GetPendingWebhookEventsResponse) GetPendingWebhookEventsOk() (*[]PendingCommentToSyncOutbound, bool)`

GetPendingWebhookEventsOk returns a tuple with the PendingWebhookEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingWebhookEvents

`func (o *GetPendingWebhookEventsResponse) SetPendingWebhookEvents(v []PendingCommentToSyncOutbound)`

SetPendingWebhookEvents sets PendingWebhookEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


