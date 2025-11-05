# AggregationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to [**[]QueryPredicate**](QueryPredicate.md) |  | [optional] 
**ResourceName** | **string** |  | 
**GroupBy** | Pointer to **[]string** |  | [optional] 
**Operations** | [**[]AggregationOperation**](AggregationOperation.md) |  | 
**Sort** | Pointer to [**AggregationRequestSort**](AggregationRequestSort.md) |  | [optional] 

## Methods

### NewAggregationRequest

`func NewAggregationRequest(resourceName string, operations []AggregationOperation, ) *AggregationRequest`

NewAggregationRequest instantiates a new AggregationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationRequestWithDefaults

`func NewAggregationRequestWithDefaults() *AggregationRequest`

NewAggregationRequestWithDefaults instantiates a new AggregationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *AggregationRequest) GetQuery() []QueryPredicate`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AggregationRequest) GetQueryOk() (*[]QueryPredicate, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AggregationRequest) SetQuery(v []QueryPredicate)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *AggregationRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetResourceName

`func (o *AggregationRequest) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AggregationRequest) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AggregationRequest) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetGroupBy

`func (o *AggregationRequest) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *AggregationRequest) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *AggregationRequest) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *AggregationRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetOperations

`func (o *AggregationRequest) GetOperations() []AggregationOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *AggregationRequest) GetOperationsOk() (*[]AggregationOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *AggregationRequest) SetOperations(v []AggregationOperation)`

SetOperations sets Operations field to given value.


### GetSort

`func (o *AggregationRequest) GetSort() AggregationRequestSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *AggregationRequest) GetSortOk() (*AggregationRequestSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *AggregationRequest) SetSort(v AggregationRequestSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *AggregationRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


