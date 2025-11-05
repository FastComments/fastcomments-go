# BulkAggregateQuestionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggId** | **string** |  | 
**QuestionId** | Pointer to **string** |  | [optional] 
**QuestionIds** | Pointer to **[]string** |  | [optional] 
**UrlId** | Pointer to **string** |  | [optional] 
**TimeBucket** | Pointer to [**AggregateTimeBucket**](AggregateTimeBucket.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBulkAggregateQuestionItem

`func NewBulkAggregateQuestionItem(aggId string, ) *BulkAggregateQuestionItem`

NewBulkAggregateQuestionItem instantiates a new BulkAggregateQuestionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAggregateQuestionItemWithDefaults

`func NewBulkAggregateQuestionItemWithDefaults() *BulkAggregateQuestionItem`

NewBulkAggregateQuestionItemWithDefaults instantiates a new BulkAggregateQuestionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggId

`func (o *BulkAggregateQuestionItem) GetAggId() string`

GetAggId returns the AggId field if non-nil, zero value otherwise.

### GetAggIdOk

`func (o *BulkAggregateQuestionItem) GetAggIdOk() (*string, bool)`

GetAggIdOk returns a tuple with the AggId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggId

`func (o *BulkAggregateQuestionItem) SetAggId(v string)`

SetAggId sets AggId field to given value.


### GetQuestionId

`func (o *BulkAggregateQuestionItem) GetQuestionId() string`

GetQuestionId returns the QuestionId field if non-nil, zero value otherwise.

### GetQuestionIdOk

`func (o *BulkAggregateQuestionItem) GetQuestionIdOk() (*string, bool)`

GetQuestionIdOk returns a tuple with the QuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionId

`func (o *BulkAggregateQuestionItem) SetQuestionId(v string)`

SetQuestionId sets QuestionId field to given value.

### HasQuestionId

`func (o *BulkAggregateQuestionItem) HasQuestionId() bool`

HasQuestionId returns a boolean if a field has been set.

### GetQuestionIds

`func (o *BulkAggregateQuestionItem) GetQuestionIds() []string`

GetQuestionIds returns the QuestionIds field if non-nil, zero value otherwise.

### GetQuestionIdsOk

`func (o *BulkAggregateQuestionItem) GetQuestionIdsOk() (*[]string, bool)`

GetQuestionIdsOk returns a tuple with the QuestionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionIds

`func (o *BulkAggregateQuestionItem) SetQuestionIds(v []string)`

SetQuestionIds sets QuestionIds field to given value.

### HasQuestionIds

`func (o *BulkAggregateQuestionItem) HasQuestionIds() bool`

HasQuestionIds returns a boolean if a field has been set.

### GetUrlId

`func (o *BulkAggregateQuestionItem) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *BulkAggregateQuestionItem) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *BulkAggregateQuestionItem) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.

### HasUrlId

`func (o *BulkAggregateQuestionItem) HasUrlId() bool`

HasUrlId returns a boolean if a field has been set.

### GetTimeBucket

`func (o *BulkAggregateQuestionItem) GetTimeBucket() AggregateTimeBucket`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *BulkAggregateQuestionItem) GetTimeBucketOk() (*AggregateTimeBucket, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *BulkAggregateQuestionItem) SetTimeBucket(v AggregateTimeBucket)`

SetTimeBucket sets TimeBucket field to given value.

### HasTimeBucket

`func (o *BulkAggregateQuestionItem) HasTimeBucket() bool`

HasTimeBucket returns a boolean if a field has been set.

### GetStartDate

`func (o *BulkAggregateQuestionItem) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BulkAggregateQuestionItem) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BulkAggregateQuestionItem) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BulkAggregateQuestionItem) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


