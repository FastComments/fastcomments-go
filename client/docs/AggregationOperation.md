# AggregationOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The field to operate on | 
**Op** | [**AggregationOpType**](AggregationOpType.md) |  | 
**Alias** | Pointer to **string** | Optional alias for the output; if not provided, a default alias is computed | [optional] 
**ExpandArray** | Pointer to **bool** |  | [optional] 

## Methods

### NewAggregationOperation

`func NewAggregationOperation(field string, op AggregationOpType, ) *AggregationOperation`

NewAggregationOperation instantiates a new AggregationOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationOperationWithDefaults

`func NewAggregationOperationWithDefaults() *AggregationOperation`

NewAggregationOperationWithDefaults instantiates a new AggregationOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AggregationOperation) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggregationOperation) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggregationOperation) SetField(v string)`

SetField sets Field field to given value.


### GetOp

`func (o *AggregationOperation) GetOp() AggregationOpType`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AggregationOperation) GetOpOk() (*AggregationOpType, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AggregationOperation) SetOp(v AggregationOpType)`

SetOp sets Op field to given value.


### GetAlias

`func (o *AggregationOperation) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AggregationOperation) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AggregationOperation) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *AggregationOperation) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetExpandArray

`func (o *AggregationOperation) GetExpandArray() bool`

GetExpandArray returns the ExpandArray field if non-nil, zero value otherwise.

### GetExpandArrayOk

`func (o *AggregationOperation) GetExpandArrayOk() (*bool, bool)`

GetExpandArrayOk returns a tuple with the ExpandArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandArray

`func (o *AggregationOperation) SetExpandArray(v bool)`

SetExpandArray sets ExpandArray field to given value.

### HasExpandArray

`func (o *AggregationOperation) HasExpandArray() bool`

HasExpandArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


