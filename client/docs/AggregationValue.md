# AggregationValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to **map[string]string** | Construct a type with a set of properties K of type T | [optional] 
**StringValue** | Pointer to **string** |  | [optional] 
**NumericValue** | Pointer to **float64** |  | [optional] 
**DistinctCount** | Pointer to **int64** |  | [optional] 
**DistinctCounts** | Pointer to **map[string]float64** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewAggregationValue

`func NewAggregationValue() *AggregationValue`

NewAggregationValue instantiates a new AggregationValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationValueWithDefaults

`func NewAggregationValueWithDefaults() *AggregationValue`

NewAggregationValueWithDefaults instantiates a new AggregationValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *AggregationValue) GetGroups() map[string]string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AggregationValue) GetGroupsOk() (*map[string]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AggregationValue) SetGroups(v map[string]string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AggregationValue) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetStringValue

`func (o *AggregationValue) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *AggregationValue) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *AggregationValue) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *AggregationValue) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetNumericValue

`func (o *AggregationValue) GetNumericValue() float64`

GetNumericValue returns the NumericValue field if non-nil, zero value otherwise.

### GetNumericValueOk

`func (o *AggregationValue) GetNumericValueOk() (*float64, bool)`

GetNumericValueOk returns a tuple with the NumericValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericValue

`func (o *AggregationValue) SetNumericValue(v float64)`

SetNumericValue sets NumericValue field to given value.

### HasNumericValue

`func (o *AggregationValue) HasNumericValue() bool`

HasNumericValue returns a boolean if a field has been set.

### GetDistinctCount

`func (o *AggregationValue) GetDistinctCount() int64`

GetDistinctCount returns the DistinctCount field if non-nil, zero value otherwise.

### GetDistinctCountOk

`func (o *AggregationValue) GetDistinctCountOk() (*int64, bool)`

GetDistinctCountOk returns a tuple with the DistinctCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctCount

`func (o *AggregationValue) SetDistinctCount(v int64)`

SetDistinctCount sets DistinctCount field to given value.

### HasDistinctCount

`func (o *AggregationValue) HasDistinctCount() bool`

HasDistinctCount returns a boolean if a field has been set.

### GetDistinctCounts

`func (o *AggregationValue) GetDistinctCounts() map[string]float64`

GetDistinctCounts returns the DistinctCounts field if non-nil, zero value otherwise.

### GetDistinctCountsOk

`func (o *AggregationValue) GetDistinctCountsOk() (*map[string]float64, bool)`

GetDistinctCountsOk returns a tuple with the DistinctCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctCounts

`func (o *AggregationValue) SetDistinctCounts(v map[string]float64)`

SetDistinctCounts sets DistinctCounts field to given value.

### HasDistinctCounts

`func (o *AggregationValue) HasDistinctCounts() bool`

HasDistinctCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


