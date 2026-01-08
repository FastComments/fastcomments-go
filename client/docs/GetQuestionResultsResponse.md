# GetQuestionResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**QuestionResults** | [**[]QuestionResult**](QuestionResult.md) |  | 

## Methods

### NewGetQuestionResultsResponse

`func NewGetQuestionResultsResponse(status APIStatus, questionResults []QuestionResult, ) *GetQuestionResultsResponse`

NewGetQuestionResultsResponse instantiates a new GetQuestionResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQuestionResultsResponseWithDefaults

`func NewGetQuestionResultsResponseWithDefaults() *GetQuestionResultsResponse`

NewGetQuestionResultsResponseWithDefaults instantiates a new GetQuestionResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetQuestionResultsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetQuestionResultsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetQuestionResultsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetQuestionResults

`func (o *GetQuestionResultsResponse) GetQuestionResults() []QuestionResult`

GetQuestionResults returns the QuestionResults field if non-nil, zero value otherwise.

### GetQuestionResultsOk

`func (o *GetQuestionResultsResponse) GetQuestionResultsOk() (*[]QuestionResult, bool)`

GetQuestionResultsOk returns a tuple with the QuestionResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionResults

`func (o *GetQuestionResultsResponse) SetQuestionResults(v []QuestionResult)`

SetQuestionResults sets QuestionResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


