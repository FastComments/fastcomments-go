# AggregationAPIError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**ValidResourceNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAggregationAPIError

`func NewAggregationAPIError(status APIStatus, reason string, code string, ) *AggregationAPIError`

NewAggregationAPIError instantiates a new AggregationAPIError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationAPIErrorWithDefaults

`func NewAggregationAPIErrorWithDefaults() *AggregationAPIError`

NewAggregationAPIErrorWithDefaults instantiates a new AggregationAPIError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AggregationAPIError) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AggregationAPIError) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AggregationAPIError) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetReason

`func (o *AggregationAPIError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AggregationAPIError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AggregationAPIError) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *AggregationAPIError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AggregationAPIError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AggregationAPIError) SetCode(v string)`

SetCode sets Code field to given value.


### GetValidResourceNames

`func (o *AggregationAPIError) GetValidResourceNames() []string`

GetValidResourceNames returns the ValidResourceNames field if non-nil, zero value otherwise.

### GetValidResourceNamesOk

`func (o *AggregationAPIError) GetValidResourceNamesOk() (*[]string, bool)`

GetValidResourceNamesOk returns a tuple with the ValidResourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidResourceNames

`func (o *AggregationAPIError) SetValidResourceNames(v []string)`

SetValidResourceNames sets ValidResourceNames field to given value.

### HasValidResourceNames

`func (o *AggregationAPIError) HasValidResourceNames() bool`

HasValidResourceNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


