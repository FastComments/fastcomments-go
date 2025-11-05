# CombineCommentsWithQuestionResults200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ImportedAPIStatusFAILED**](ImportedAPIStatusFAILED.md) |  | 
**Data** | [**FindCommentsByRangeResponse**](FindCommentsByRangeResponse.md) |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewCombineCommentsWithQuestionResults200Response

`func NewCombineCommentsWithQuestionResults200Response(status ImportedAPIStatusFAILED, data FindCommentsByRangeResponse, reason string, code string, ) *CombineCommentsWithQuestionResults200Response`

NewCombineCommentsWithQuestionResults200Response instantiates a new CombineCommentsWithQuestionResults200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombineCommentsWithQuestionResults200ResponseWithDefaults

`func NewCombineCommentsWithQuestionResults200ResponseWithDefaults() *CombineCommentsWithQuestionResults200Response`

NewCombineCommentsWithQuestionResults200ResponseWithDefaults instantiates a new CombineCommentsWithQuestionResults200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CombineCommentsWithQuestionResults200Response) GetStatus() ImportedAPIStatusFAILED`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CombineCommentsWithQuestionResults200Response) GetStatusOk() (*ImportedAPIStatusFAILED, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CombineCommentsWithQuestionResults200Response) SetStatus(v ImportedAPIStatusFAILED)`

SetStatus sets Status field to given value.


### GetData

`func (o *CombineCommentsWithQuestionResults200Response) GetData() FindCommentsByRangeResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CombineCommentsWithQuestionResults200Response) GetDataOk() (*FindCommentsByRangeResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CombineCommentsWithQuestionResults200Response) SetData(v FindCommentsByRangeResponse)`

SetData sets Data field to given value.


### GetReason

`func (o *CombineCommentsWithQuestionResults200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CombineCommentsWithQuestionResults200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CombineCommentsWithQuestionResults200Response) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *CombineCommentsWithQuestionResults200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CombineCommentsWithQuestionResults200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CombineCommentsWithQuestionResults200Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *CombineCommentsWithQuestionResults200Response) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *CombineCommentsWithQuestionResults200Response) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *CombineCommentsWithQuestionResults200Response) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *CombineCommentsWithQuestionResults200Response) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *CombineCommentsWithQuestionResults200Response) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *CombineCommentsWithQuestionResults200Response) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *CombineCommentsWithQuestionResults200Response) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *CombineCommentsWithQuestionResults200Response) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *CombineCommentsWithQuestionResults200Response) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *CombineCommentsWithQuestionResults200Response) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *CombineCommentsWithQuestionResults200Response) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *CombineCommentsWithQuestionResults200Response) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *CombineCommentsWithQuestionResults200Response) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *CombineCommentsWithQuestionResults200Response) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *CombineCommentsWithQuestionResults200Response) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *CombineCommentsWithQuestionResults200Response) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *CombineCommentsWithQuestionResults200Response) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *CombineCommentsWithQuestionResults200Response) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *CombineCommentsWithQuestionResults200Response) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *CombineCommentsWithQuestionResults200Response) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


