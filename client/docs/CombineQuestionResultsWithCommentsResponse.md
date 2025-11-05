# CombineQuestionResultsWithCommentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Data** | [**FindCommentsByRangeResponse**](FindCommentsByRangeResponse.md) |  | 

## Methods

### NewCombineQuestionResultsWithCommentsResponse

`func NewCombineQuestionResultsWithCommentsResponse(status APIStatus, data FindCommentsByRangeResponse, ) *CombineQuestionResultsWithCommentsResponse`

NewCombineQuestionResultsWithCommentsResponse instantiates a new CombineQuestionResultsWithCommentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombineQuestionResultsWithCommentsResponseWithDefaults

`func NewCombineQuestionResultsWithCommentsResponseWithDefaults() *CombineQuestionResultsWithCommentsResponse`

NewCombineQuestionResultsWithCommentsResponseWithDefaults instantiates a new CombineQuestionResultsWithCommentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CombineQuestionResultsWithCommentsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CombineQuestionResultsWithCommentsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CombineQuestionResultsWithCommentsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *CombineQuestionResultsWithCommentsResponse) GetData() FindCommentsByRangeResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CombineQuestionResultsWithCommentsResponse) GetDataOk() (*FindCommentsByRangeResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CombineQuestionResultsWithCommentsResponse) SetData(v FindCommentsByRangeResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


