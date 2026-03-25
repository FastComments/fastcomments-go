# CreateTicket200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Ticket** | [**APITicket**](APITicket.md) |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewCreateTicket200Response

`func NewCreateTicket200Response(status APIStatus, ticket APITicket, reason string, code string, ) *CreateTicket200Response`

NewCreateTicket200Response instantiates a new CreateTicket200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTicket200ResponseWithDefaults

`func NewCreateTicket200ResponseWithDefaults() *CreateTicket200Response`

NewCreateTicket200ResponseWithDefaults instantiates a new CreateTicket200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateTicket200Response) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTicket200Response) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTicket200Response) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTicket

`func (o *CreateTicket200Response) GetTicket() APITicket`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *CreateTicket200Response) GetTicketOk() (*APITicket, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *CreateTicket200Response) SetTicket(v APITicket)`

SetTicket sets Ticket field to given value.


### GetReason

`func (o *CreateTicket200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateTicket200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateTicket200Response) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *CreateTicket200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateTicket200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateTicket200Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *CreateTicket200Response) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *CreateTicket200Response) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *CreateTicket200Response) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *CreateTicket200Response) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *CreateTicket200Response) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *CreateTicket200Response) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *CreateTicket200Response) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *CreateTicket200Response) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *CreateTicket200Response) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *CreateTicket200Response) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *CreateTicket200Response) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *CreateTicket200Response) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *CreateTicket200Response) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *CreateTicket200Response) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *CreateTicket200Response) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *CreateTicket200Response) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *CreateTicket200Response) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *CreateTicket200Response) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *CreateTicket200Response) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *CreateTicket200Response) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


