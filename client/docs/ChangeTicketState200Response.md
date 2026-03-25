# ChangeTicketState200Response

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

### NewChangeTicketState200Response

`func NewChangeTicketState200Response(status APIStatus, ticket APITicket, reason string, code string, ) *ChangeTicketState200Response`

NewChangeTicketState200Response instantiates a new ChangeTicketState200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeTicketState200ResponseWithDefaults

`func NewChangeTicketState200ResponseWithDefaults() *ChangeTicketState200Response`

NewChangeTicketState200ResponseWithDefaults instantiates a new ChangeTicketState200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChangeTicketState200Response) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChangeTicketState200Response) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChangeTicketState200Response) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTicket

`func (o *ChangeTicketState200Response) GetTicket() APITicket`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *ChangeTicketState200Response) GetTicketOk() (*APITicket, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *ChangeTicketState200Response) SetTicket(v APITicket)`

SetTicket sets Ticket field to given value.


### GetReason

`func (o *ChangeTicketState200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ChangeTicketState200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ChangeTicketState200Response) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *ChangeTicketState200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ChangeTicketState200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ChangeTicketState200Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *ChangeTicketState200Response) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *ChangeTicketState200Response) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *ChangeTicketState200Response) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *ChangeTicketState200Response) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *ChangeTicketState200Response) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *ChangeTicketState200Response) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *ChangeTicketState200Response) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *ChangeTicketState200Response) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *ChangeTicketState200Response) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *ChangeTicketState200Response) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *ChangeTicketState200Response) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *ChangeTicketState200Response) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *ChangeTicketState200Response) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *ChangeTicketState200Response) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *ChangeTicketState200Response) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *ChangeTicketState200Response) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *ChangeTicketState200Response) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *ChangeTicketState200Response) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *ChangeTicketState200Response) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *ChangeTicketState200Response) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


