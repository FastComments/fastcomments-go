# ChangeTicketStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Ticket** | [**APITicket**](APITicket.md) |  | 

## Methods

### NewChangeTicketStateResponse

`func NewChangeTicketStateResponse(status APIStatus, ticket APITicket, ) *ChangeTicketStateResponse`

NewChangeTicketStateResponse instantiates a new ChangeTicketStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeTicketStateResponseWithDefaults

`func NewChangeTicketStateResponseWithDefaults() *ChangeTicketStateResponse`

NewChangeTicketStateResponseWithDefaults instantiates a new ChangeTicketStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChangeTicketStateResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChangeTicketStateResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChangeTicketStateResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTicket

`func (o *ChangeTicketStateResponse) GetTicket() APITicket`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *ChangeTicketStateResponse) GetTicketOk() (*APITicket, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *ChangeTicketStateResponse) SetTicket(v APITicket)`

SetTicket sets Ticket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


