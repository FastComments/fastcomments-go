# CreateTicketResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Ticket** | [**APITicket**](APITicket.md) |  | 

## Methods

### NewCreateTicketResponse

`func NewCreateTicketResponse(status APIStatus, ticket APITicket, ) *CreateTicketResponse`

NewCreateTicketResponse instantiates a new CreateTicketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTicketResponseWithDefaults

`func NewCreateTicketResponseWithDefaults() *CreateTicketResponse`

NewCreateTicketResponseWithDefaults instantiates a new CreateTicketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateTicketResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTicketResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTicketResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTicket

`func (o *CreateTicketResponse) GetTicket() APITicket`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *CreateTicketResponse) GetTicketOk() (*APITicket, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *CreateTicketResponse) SetTicket(v APITicket)`

SetTicket sets Ticket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


