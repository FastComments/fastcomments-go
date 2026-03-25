# GetTickets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Tickets** | [**[]APITicket**](APITicket.md) |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetTickets200Response

`func NewGetTickets200Response(status APIStatus, tickets []APITicket, reason string, code string, ) *GetTickets200Response`

NewGetTickets200Response instantiates a new GetTickets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTickets200ResponseWithDefaults

`func NewGetTickets200ResponseWithDefaults() *GetTickets200Response`

NewGetTickets200ResponseWithDefaults instantiates a new GetTickets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetTickets200Response) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTickets200Response) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTickets200Response) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTickets

`func (o *GetTickets200Response) GetTickets() []APITicket`

GetTickets returns the Tickets field if non-nil, zero value otherwise.

### GetTicketsOk

`func (o *GetTickets200Response) GetTicketsOk() (*[]APITicket, bool)`

GetTicketsOk returns a tuple with the Tickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickets

`func (o *GetTickets200Response) SetTickets(v []APITicket)`

SetTickets sets Tickets field to given value.


### GetReason

`func (o *GetTickets200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetTickets200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetTickets200Response) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetTickets200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetTickets200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetTickets200Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetTickets200Response) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetTickets200Response) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetTickets200Response) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetTickets200Response) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetTickets200Response) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetTickets200Response) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetTickets200Response) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetTickets200Response) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetTickets200Response) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetTickets200Response) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetTickets200Response) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetTickets200Response) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetTickets200Response) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetTickets200Response) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetTickets200Response) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetTickets200Response) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetTickets200Response) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetTickets200Response) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetTickets200Response) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetTickets200Response) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


