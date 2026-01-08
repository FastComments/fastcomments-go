# GetQuestionResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**QuestionResult** | [**QuestionResult**](QuestionResult.md) |  | 

## Methods

### NewGetQuestionResultResponse

`func NewGetQuestionResultResponse(status APIStatus, questionResult QuestionResult, ) *GetQuestionResultResponse`

NewGetQuestionResultResponse instantiates a new GetQuestionResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQuestionResultResponseWithDefaults

`func NewGetQuestionResultResponseWithDefaults() *GetQuestionResultResponse`

NewGetQuestionResultResponseWithDefaults instantiates a new GetQuestionResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetQuestionResultResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetQuestionResultResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetQuestionResultResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetQuestionResult

`func (o *GetQuestionResultResponse) GetQuestionResult() QuestionResult`

GetQuestionResult returns the QuestionResult field if non-nil, zero value otherwise.

### GetQuestionResultOk

`func (o *GetQuestionResultResponse) GetQuestionResultOk() (*QuestionResult, bool)`

GetQuestionResultOk returns a tuple with the QuestionResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionResult

`func (o *GetQuestionResultResponse) SetQuestionResult(v QuestionResult)`

SetQuestionResult sets QuestionResult field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


