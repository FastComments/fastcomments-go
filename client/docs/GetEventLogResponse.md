# GetEventLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]EventLogEntry**](EventLogEntry.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewGetEventLogResponse

`func NewGetEventLogResponse(events []EventLogEntry, status APIStatus, ) *GetEventLogResponse`

NewGetEventLogResponse instantiates a new GetEventLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventLogResponseWithDefaults

`func NewGetEventLogResponseWithDefaults() *GetEventLogResponse`

NewGetEventLogResponseWithDefaults instantiates a new GetEventLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *GetEventLogResponse) GetEvents() []EventLogEntry`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetEventLogResponse) GetEventsOk() (*[]EventLogEntry, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetEventLogResponse) SetEvents(v []EventLogEntry)`

SetEvents sets Events field to given value.


### GetStatus

`func (o *GetEventLogResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEventLogResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEventLogResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


