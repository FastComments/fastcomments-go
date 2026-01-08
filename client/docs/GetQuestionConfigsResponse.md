# GetQuestionConfigsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**QuestionConfigs** | [**[]QuestionConfig**](QuestionConfig.md) |  | 

## Methods

### NewGetQuestionConfigsResponse

`func NewGetQuestionConfigsResponse(status APIStatus, questionConfigs []QuestionConfig, ) *GetQuestionConfigsResponse`

NewGetQuestionConfigsResponse instantiates a new GetQuestionConfigsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQuestionConfigsResponseWithDefaults

`func NewGetQuestionConfigsResponseWithDefaults() *GetQuestionConfigsResponse`

NewGetQuestionConfigsResponseWithDefaults instantiates a new GetQuestionConfigsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetQuestionConfigsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetQuestionConfigsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetQuestionConfigsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetQuestionConfigs

`func (o *GetQuestionConfigsResponse) GetQuestionConfigs() []QuestionConfig`

GetQuestionConfigs returns the QuestionConfigs field if non-nil, zero value otherwise.

### GetQuestionConfigsOk

`func (o *GetQuestionConfigsResponse) GetQuestionConfigsOk() (*[]QuestionConfig, bool)`

GetQuestionConfigsOk returns a tuple with the QuestionConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionConfigs

`func (o *GetQuestionConfigsResponse) SetQuestionConfigs(v []QuestionConfig)`

SetQuestionConfigs sets QuestionConfigs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


