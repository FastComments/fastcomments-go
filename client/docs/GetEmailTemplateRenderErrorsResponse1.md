# GetEmailTemplateRenderErrorsResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**RenderErrors** | [**[]EmailTemplateRenderErrorResponse**](EmailTemplateRenderErrorResponse.md) |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetEmailTemplateRenderErrorsResponse1

`func NewGetEmailTemplateRenderErrorsResponse1(status APIStatus, renderErrors []EmailTemplateRenderErrorResponse, reason string, code string, ) *GetEmailTemplateRenderErrorsResponse1`

NewGetEmailTemplateRenderErrorsResponse1 instantiates a new GetEmailTemplateRenderErrorsResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEmailTemplateRenderErrorsResponse1WithDefaults

`func NewGetEmailTemplateRenderErrorsResponse1WithDefaults() *GetEmailTemplateRenderErrorsResponse1`

NewGetEmailTemplateRenderErrorsResponse1WithDefaults instantiates a new GetEmailTemplateRenderErrorsResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetEmailTemplateRenderErrorsResponse1) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEmailTemplateRenderErrorsResponse1) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEmailTemplateRenderErrorsResponse1) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetRenderErrors

`func (o *GetEmailTemplateRenderErrorsResponse1) GetRenderErrors() []EmailTemplateRenderErrorResponse`

GetRenderErrors returns the RenderErrors field if non-nil, zero value otherwise.

### GetRenderErrorsOk

`func (o *GetEmailTemplateRenderErrorsResponse1) GetRenderErrorsOk() (*[]EmailTemplateRenderErrorResponse, bool)`

GetRenderErrorsOk returns a tuple with the RenderErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderErrors

`func (o *GetEmailTemplateRenderErrorsResponse1) SetRenderErrors(v []EmailTemplateRenderErrorResponse)`

SetRenderErrors sets RenderErrors field to given value.


### GetReason

`func (o *GetEmailTemplateRenderErrorsResponse1) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetEmailTemplateRenderErrorsResponse1) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetEmailTemplateRenderErrorsResponse1) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetEmailTemplateRenderErrorsResponse1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetEmailTemplateRenderErrorsResponse1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetEmailTemplateRenderErrorsResponse1) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetEmailTemplateRenderErrorsResponse1) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetEmailTemplateRenderErrorsResponse1) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetEmailTemplateRenderErrorsResponse1) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetEmailTemplateRenderErrorsResponse1) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetEmailTemplateRenderErrorsResponse1) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetEmailTemplateRenderErrorsResponse1) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetEmailTemplateRenderErrorsResponse1) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetEmailTemplateRenderErrorsResponse1) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetEmailTemplateRenderErrorsResponse1) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetEmailTemplateRenderErrorsResponse1) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetEmailTemplateRenderErrorsResponse1) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetEmailTemplateRenderErrorsResponse1) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetEmailTemplateRenderErrorsResponse1) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetEmailTemplateRenderErrorsResponse1) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetEmailTemplateRenderErrorsResponse1) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetEmailTemplateRenderErrorsResponse1) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetEmailTemplateRenderErrorsResponse1) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetEmailTemplateRenderErrorsResponse1) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetEmailTemplateRenderErrorsResponse1) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetEmailTemplateRenderErrorsResponse1) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


