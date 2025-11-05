# AggregationResponseStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeMS** | **int64** |  | 
**Scanned** | **int64** |  | 

## Methods

### NewAggregationResponseStats

`func NewAggregationResponseStats(timeMS int64, scanned int64, ) *AggregationResponseStats`

NewAggregationResponseStats instantiates a new AggregationResponseStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationResponseStatsWithDefaults

`func NewAggregationResponseStatsWithDefaults() *AggregationResponseStats`

NewAggregationResponseStatsWithDefaults instantiates a new AggregationResponseStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeMS

`func (o *AggregationResponseStats) GetTimeMS() int64`

GetTimeMS returns the TimeMS field if non-nil, zero value otherwise.

### GetTimeMSOk

`func (o *AggregationResponseStats) GetTimeMSOk() (*int64, bool)`

GetTimeMSOk returns a tuple with the TimeMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMS

`func (o *AggregationResponseStats) SetTimeMS(v int64)`

SetTimeMS sets TimeMS field to given value.


### GetScanned

`func (o *AggregationResponseStats) GetScanned() int64`

GetScanned returns the Scanned field if non-nil, zero value otherwise.

### GetScannedOk

`func (o *AggregationResponseStats) GetScannedOk() (*int64, bool)`

GetScannedOk returns a tuple with the Scanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanned

`func (o *AggregationResponseStats) SetScanned(v int64)`

SetScanned sets Scanned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


