# GetTicketResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Ticket** | [**APITicketDetail**](APITicketDetail.md) |  | 
**AvailableStates** | **[]float64** |  | 

## Methods

### NewGetTicketResponse

`func NewGetTicketResponse(status APIStatus, ticket APITicketDetail, availableStates []float64, ) *GetTicketResponse`

NewGetTicketResponse instantiates a new GetTicketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTicketResponseWithDefaults

`func NewGetTicketResponseWithDefaults() *GetTicketResponse`

NewGetTicketResponseWithDefaults instantiates a new GetTicketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetTicketResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTicketResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTicketResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTicket

`func (o *GetTicketResponse) GetTicket() APITicketDetail`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *GetTicketResponse) GetTicketOk() (*APITicketDetail, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *GetTicketResponse) SetTicket(v APITicketDetail)`

SetTicket sets Ticket field to given value.


### GetAvailableStates

`func (o *GetTicketResponse) GetAvailableStates() []float64`

GetAvailableStates returns the AvailableStates field if non-nil, zero value otherwise.

### GetAvailableStatesOk

`func (o *GetTicketResponse) GetAvailableStatesOk() (*[]float64, bool)`

GetAvailableStatesOk returns a tuple with the AvailableStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStates

`func (o *GetTicketResponse) SetAvailableStates(v []float64)`

SetAvailableStates sets AvailableStates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


