# BulkAggregateQuestionResultsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | [**[]BulkAggregateQuestionItem**](BulkAggregateQuestionItem.md) |  | 

## Methods

### NewBulkAggregateQuestionResultsRequest

`func NewBulkAggregateQuestionResultsRequest(aggregations []BulkAggregateQuestionItem, ) *BulkAggregateQuestionResultsRequest`

NewBulkAggregateQuestionResultsRequest instantiates a new BulkAggregateQuestionResultsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAggregateQuestionResultsRequestWithDefaults

`func NewBulkAggregateQuestionResultsRequestWithDefaults() *BulkAggregateQuestionResultsRequest`

NewBulkAggregateQuestionResultsRequestWithDefaults instantiates a new BulkAggregateQuestionResultsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *BulkAggregateQuestionResultsRequest) GetAggregations() []BulkAggregateQuestionItem`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *BulkAggregateQuestionResultsRequest) GetAggregationsOk() (*[]BulkAggregateQuestionItem, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *BulkAggregateQuestionResultsRequest) SetAggregations(v []BulkAggregateQuestionItem)`

SetAggregations sets Aggregations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


