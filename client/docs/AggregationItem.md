# AggregationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to **map[string]string** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewAggregationItem

`func NewAggregationItem() *AggregationItem`

NewAggregationItem instantiates a new AggregationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationItemWithDefaults

`func NewAggregationItemWithDefaults() *AggregationItem`

NewAggregationItemWithDefaults instantiates a new AggregationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *AggregationItem) GetGroups() map[string]string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AggregationItem) GetGroupsOk() (*map[string]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AggregationItem) SetGroups(v map[string]string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AggregationItem) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


