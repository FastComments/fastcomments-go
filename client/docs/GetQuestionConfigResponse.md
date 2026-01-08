# GetQuestionConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**QuestionConfig** | [**QuestionConfig**](QuestionConfig.md) |  | 

## Methods

### NewGetQuestionConfigResponse

`func NewGetQuestionConfigResponse(status APIStatus, questionConfig QuestionConfig, ) *GetQuestionConfigResponse`

NewGetQuestionConfigResponse instantiates a new GetQuestionConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQuestionConfigResponseWithDefaults

`func NewGetQuestionConfigResponseWithDefaults() *GetQuestionConfigResponse`

NewGetQuestionConfigResponseWithDefaults instantiates a new GetQuestionConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetQuestionConfigResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetQuestionConfigResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetQuestionConfigResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetQuestionConfig

`func (o *GetQuestionConfigResponse) GetQuestionConfig() QuestionConfig`

GetQuestionConfig returns the QuestionConfig field if non-nil, zero value otherwise.

### GetQuestionConfigOk

`func (o *GetQuestionConfigResponse) GetQuestionConfigOk() (*QuestionConfig, bool)`

GetQuestionConfigOk returns a tuple with the QuestionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionConfig

`func (o *GetQuestionConfigResponse) SetQuestionConfig(v QuestionConfig)`

SetQuestionConfig sets QuestionConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


