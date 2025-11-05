# AggregateQuestionResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Data** | [**QuestionResultAggregationOverall**](QuestionResultAggregationOverall.md) |  | 

## Methods

### NewAggregateQuestionResultsResponse

`func NewAggregateQuestionResultsResponse(status APIStatus, data QuestionResultAggregationOverall, ) *AggregateQuestionResultsResponse`

NewAggregateQuestionResultsResponse instantiates a new AggregateQuestionResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateQuestionResultsResponseWithDefaults

`func NewAggregateQuestionResultsResponseWithDefaults() *AggregateQuestionResultsResponse`

NewAggregateQuestionResultsResponseWithDefaults instantiates a new AggregateQuestionResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AggregateQuestionResultsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AggregateQuestionResultsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AggregateQuestionResultsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *AggregateQuestionResultsResponse) GetData() QuestionResultAggregationOverall`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AggregateQuestionResultsResponse) GetDataOk() (*QuestionResultAggregationOverall, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AggregateQuestionResultsResponse) SetData(v QuestionResultAggregationOverall)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


