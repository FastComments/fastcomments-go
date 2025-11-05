# QueryPredicate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | [**QueryPredicateValue**](QueryPredicateValue.md) |  | 
**Operator** | **string** |  | 

## Methods

### NewQueryPredicate

`func NewQueryPredicate(key string, value QueryPredicateValue, operator string, ) *QueryPredicate`

NewQueryPredicate instantiates a new QueryPredicate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPredicateWithDefaults

`func NewQueryPredicateWithDefaults() *QueryPredicate`

NewQueryPredicateWithDefaults instantiates a new QueryPredicate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *QueryPredicate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *QueryPredicate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *QueryPredicate) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *QueryPredicate) GetValue() QueryPredicateValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QueryPredicate) GetValueOk() (*QueryPredicateValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QueryPredicate) SetValue(v QueryPredicateValue)`

SetValue sets Value field to given value.


### GetOperator

`func (o *QueryPredicate) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *QueryPredicate) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *QueryPredicate) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


