# GetTicket200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Ticket** | [**APITicketDetail**](APITicketDetail.md) |  | 
**AvailableStates** | **[]float64** |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetTicket200Response

`func NewGetTicket200Response(status APIStatus, ticket APITicketDetail, availableStates []float64, reason string, code string, ) *GetTicket200Response`

NewGetTicket200Response instantiates a new GetTicket200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTicket200ResponseWithDefaults

`func NewGetTicket200ResponseWithDefaults() *GetTicket200Response`

NewGetTicket200ResponseWithDefaults instantiates a new GetTicket200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetTicket200Response) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTicket200Response) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTicket200Response) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTicket

`func (o *GetTicket200Response) GetTicket() APITicketDetail`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *GetTicket200Response) GetTicketOk() (*APITicketDetail, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *GetTicket200Response) SetTicket(v APITicketDetail)`

SetTicket sets Ticket field to given value.


### GetAvailableStates

`func (o *GetTicket200Response) GetAvailableStates() []float64`

GetAvailableStates returns the AvailableStates field if non-nil, zero value otherwise.

### GetAvailableStatesOk

`func (o *GetTicket200Response) GetAvailableStatesOk() (*[]float64, bool)`

GetAvailableStatesOk returns a tuple with the AvailableStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStates

`func (o *GetTicket200Response) SetAvailableStates(v []float64)`

SetAvailableStates sets AvailableStates field to given value.


### GetReason

`func (o *GetTicket200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetTicket200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetTicket200Response) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetTicket200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetTicket200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetTicket200Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetTicket200Response) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetTicket200Response) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetTicket200Response) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetTicket200Response) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetTicket200Response) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetTicket200Response) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetTicket200Response) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetTicket200Response) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetTicket200Response) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetTicket200Response) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetTicket200Response) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetTicket200Response) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetTicket200Response) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetTicket200Response) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetTicket200Response) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetTicket200Response) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetTicket200Response) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetTicket200Response) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetTicket200Response) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetTicket200Response) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


