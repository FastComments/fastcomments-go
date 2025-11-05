# QuestionResultAggregationOverall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataByDateBucket** | Pointer to [**map[string]QuestionDatum**](QuestionDatum.md) | Construct a type with a set of properties K of type T | [optional] 
**DataByUrlId** | Pointer to [**map[string]QuestionDatum**](QuestionDatum.md) | Construct a type with a set of properties K of type T | [optional] 
**CountsByValue** | Pointer to **map[string]int32** |  | [optional] 
**Total** | **int64** |  | 
**Average** | Pointer to **float64** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewQuestionResultAggregationOverall

`func NewQuestionResultAggregationOverall(total int64, createdAt time.Time, ) *QuestionResultAggregationOverall`

NewQuestionResultAggregationOverall instantiates a new QuestionResultAggregationOverall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionResultAggregationOverallWithDefaults

`func NewQuestionResultAggregationOverallWithDefaults() *QuestionResultAggregationOverall`

NewQuestionResultAggregationOverallWithDefaults instantiates a new QuestionResultAggregationOverall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataByDateBucket

`func (o *QuestionResultAggregationOverall) GetDataByDateBucket() map[string]QuestionDatum`

GetDataByDateBucket returns the DataByDateBucket field if non-nil, zero value otherwise.

### GetDataByDateBucketOk

`func (o *QuestionResultAggregationOverall) GetDataByDateBucketOk() (*map[string]QuestionDatum, bool)`

GetDataByDateBucketOk returns a tuple with the DataByDateBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataByDateBucket

`func (o *QuestionResultAggregationOverall) SetDataByDateBucket(v map[string]QuestionDatum)`

SetDataByDateBucket sets DataByDateBucket field to given value.

### HasDataByDateBucket

`func (o *QuestionResultAggregationOverall) HasDataByDateBucket() bool`

HasDataByDateBucket returns a boolean if a field has been set.

### GetDataByUrlId

`func (o *QuestionResultAggregationOverall) GetDataByUrlId() map[string]QuestionDatum`

GetDataByUrlId returns the DataByUrlId field if non-nil, zero value otherwise.

### GetDataByUrlIdOk

`func (o *QuestionResultAggregationOverall) GetDataByUrlIdOk() (*map[string]QuestionDatum, bool)`

GetDataByUrlIdOk returns a tuple with the DataByUrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataByUrlId

`func (o *QuestionResultAggregationOverall) SetDataByUrlId(v map[string]QuestionDatum)`

SetDataByUrlId sets DataByUrlId field to given value.

### HasDataByUrlId

`func (o *QuestionResultAggregationOverall) HasDataByUrlId() bool`

HasDataByUrlId returns a boolean if a field has been set.

### GetCountsByValue

`func (o *QuestionResultAggregationOverall) GetCountsByValue() map[string]int32`

GetCountsByValue returns the CountsByValue field if non-nil, zero value otherwise.

### GetCountsByValueOk

`func (o *QuestionResultAggregationOverall) GetCountsByValueOk() (*map[string]int32, bool)`

GetCountsByValueOk returns a tuple with the CountsByValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountsByValue

`func (o *QuestionResultAggregationOverall) SetCountsByValue(v map[string]int32)`

SetCountsByValue sets CountsByValue field to given value.

### HasCountsByValue

`func (o *QuestionResultAggregationOverall) HasCountsByValue() bool`

HasCountsByValue returns a boolean if a field has been set.

### GetTotal

`func (o *QuestionResultAggregationOverall) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QuestionResultAggregationOverall) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QuestionResultAggregationOverall) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetAverage

`func (o *QuestionResultAggregationOverall) GetAverage() float64`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *QuestionResultAggregationOverall) GetAverageOk() (*float64, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *QuestionResultAggregationOverall) SetAverage(v float64)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *QuestionResultAggregationOverall) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *QuestionResultAggregationOverall) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QuestionResultAggregationOverall) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QuestionResultAggregationOverall) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


