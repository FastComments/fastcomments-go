# AggregationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Data** | [**[]AggregationItem**](AggregationItem.md) |  | 
**Stats** | Pointer to [**AggregationResponseStats**](AggregationResponseStats.md) |  | [optional] 

## Methods

### NewAggregationResponse

`func NewAggregationResponse(status APIStatus, data []AggregationItem, ) *AggregationResponse`

NewAggregationResponse instantiates a new AggregationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationResponseWithDefaults

`func NewAggregationResponseWithDefaults() *AggregationResponse`

NewAggregationResponseWithDefaults instantiates a new AggregationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AggregationResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AggregationResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AggregationResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *AggregationResponse) GetData() []AggregationItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AggregationResponse) GetDataOk() (*[]AggregationItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AggregationResponse) SetData(v []AggregationItem)`

SetData sets Data field to given value.


### GetStats

`func (o *AggregationResponse) GetStats() AggregationResponseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *AggregationResponse) GetStatsOk() (*AggregationResponseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *AggregationResponse) SetStats(v AggregationResponseStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *AggregationResponse) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


