# BulkAggregateQuestionResultsResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Data** | [**map[string]QuestionResultAggregationOverall**](QuestionResultAggregationOverall.md) | Construct a type with a set of properties K of type T | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewBulkAggregateQuestionResultsResponse1

`func NewBulkAggregateQuestionResultsResponse1(status APIStatus, data map[string]QuestionResultAggregationOverall, reason string, code string, ) *BulkAggregateQuestionResultsResponse1`

NewBulkAggregateQuestionResultsResponse1 instantiates a new BulkAggregateQuestionResultsResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAggregateQuestionResultsResponse1WithDefaults

`func NewBulkAggregateQuestionResultsResponse1WithDefaults() *BulkAggregateQuestionResultsResponse1`

NewBulkAggregateQuestionResultsResponse1WithDefaults instantiates a new BulkAggregateQuestionResultsResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BulkAggregateQuestionResultsResponse1) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkAggregateQuestionResultsResponse1) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkAggregateQuestionResultsResponse1) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *BulkAggregateQuestionResultsResponse1) GetData() map[string]QuestionResultAggregationOverall`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BulkAggregateQuestionResultsResponse1) GetDataOk() (*map[string]QuestionResultAggregationOverall, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BulkAggregateQuestionResultsResponse1) SetData(v map[string]QuestionResultAggregationOverall)`

SetData sets Data field to given value.


### GetReason

`func (o *BulkAggregateQuestionResultsResponse1) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BulkAggregateQuestionResultsResponse1) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BulkAggregateQuestionResultsResponse1) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *BulkAggregateQuestionResultsResponse1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BulkAggregateQuestionResultsResponse1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BulkAggregateQuestionResultsResponse1) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *BulkAggregateQuestionResultsResponse1) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *BulkAggregateQuestionResultsResponse1) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *BulkAggregateQuestionResultsResponse1) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *BulkAggregateQuestionResultsResponse1) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *BulkAggregateQuestionResultsResponse1) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *BulkAggregateQuestionResultsResponse1) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *BulkAggregateQuestionResultsResponse1) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *BulkAggregateQuestionResultsResponse1) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *BulkAggregateQuestionResultsResponse1) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *BulkAggregateQuestionResultsResponse1) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *BulkAggregateQuestionResultsResponse1) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *BulkAggregateQuestionResultsResponse1) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *BulkAggregateQuestionResultsResponse1) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *BulkAggregateQuestionResultsResponse1) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *BulkAggregateQuestionResultsResponse1) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *BulkAggregateQuestionResultsResponse1) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *BulkAggregateQuestionResultsResponse1) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *BulkAggregateQuestionResultsResponse1) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *BulkAggregateQuestionResultsResponse1) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *BulkAggregateQuestionResultsResponse1) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


