# GetTicketsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Tickets** | [**[]APITicket**](APITicket.md) |  | 

## Methods

### NewGetTicketsResponse

`func NewGetTicketsResponse(status APIStatus, tickets []APITicket, ) *GetTicketsResponse`

NewGetTicketsResponse instantiates a new GetTicketsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTicketsResponseWithDefaults

`func NewGetTicketsResponseWithDefaults() *GetTicketsResponse`

NewGetTicketsResponseWithDefaults instantiates a new GetTicketsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetTicketsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTicketsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTicketsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTickets

`func (o *GetTicketsResponse) GetTickets() []APITicket`

GetTickets returns the Tickets field if non-nil, zero value otherwise.

### GetTicketsOk

`func (o *GetTicketsResponse) GetTicketsOk() (*[]APITicket, bool)`

GetTicketsOk returns a tuple with the Tickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickets

`func (o *GetTicketsResponse) SetTickets(v []APITicket)`

SetTickets sets Tickets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


