# GetSearchSuggestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Pages** | Pointer to [**[]ModerationPageSearchProjected**](ModerationPageSearchProjected.md) |  | [optional] 
**Users** | Pointer to [**[]ModerationUserSearchProjected**](ModerationUserSearchProjected.md) |  | [optional] 
**Code** | **string** |  | 
**Reason** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetSearchSuggestResponse

`func NewGetSearchSuggestResponse(status APIStatus, code string, reason string, ) *GetSearchSuggestResponse`

NewGetSearchSuggestResponse instantiates a new GetSearchSuggestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSearchSuggestResponseWithDefaults

`func NewGetSearchSuggestResponseWithDefaults() *GetSearchSuggestResponse`

NewGetSearchSuggestResponseWithDefaults instantiates a new GetSearchSuggestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetSearchSuggestResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSearchSuggestResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSearchSuggestResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetPages

`func (o *GetSearchSuggestResponse) GetPages() []ModerationPageSearchProjected`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *GetSearchSuggestResponse) GetPagesOk() (*[]ModerationPageSearchProjected, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *GetSearchSuggestResponse) SetPages(v []ModerationPageSearchProjected)`

SetPages sets Pages field to given value.

### HasPages

`func (o *GetSearchSuggestResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetUsers

`func (o *GetSearchSuggestResponse) GetUsers() []ModerationUserSearchProjected`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetSearchSuggestResponse) GetUsersOk() (*[]ModerationUserSearchProjected, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetSearchSuggestResponse) SetUsers(v []ModerationUserSearchProjected)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GetSearchSuggestResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetCode

`func (o *GetSearchSuggestResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetSearchSuggestResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetSearchSuggestResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetReason

`func (o *GetSearchSuggestResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetSearchSuggestResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetSearchSuggestResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSecondaryCode

`func (o *GetSearchSuggestResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetSearchSuggestResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetSearchSuggestResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetSearchSuggestResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetSearchSuggestResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetSearchSuggestResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetSearchSuggestResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetSearchSuggestResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetSearchSuggestResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetSearchSuggestResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetSearchSuggestResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetSearchSuggestResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetSearchSuggestResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetSearchSuggestResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetSearchSuggestResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetSearchSuggestResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetSearchSuggestResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetSearchSuggestResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetSearchSuggestResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetSearchSuggestResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


