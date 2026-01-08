# CreateQuestionConfigBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Question** | **string** |  | 
**HelpText** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**NumStars** | Pointer to **float64** |  | [optional] 
**Min** | Pointer to **float64** |  | [optional] 
**Max** | Pointer to **float64** |  | [optional] 
**DefaultValue** | Pointer to **float64** |  | [optional] 
**LabelNegative** | Pointer to **string** |  | [optional] 
**LabelPositive** | Pointer to **string** |  | [optional] 
**CustomOptions** | Pointer to [**[]QuestionConfigCustomOptionsInner**](QuestionConfigCustomOptionsInner.md) |  | [optional] 
**SubQuestionIds** | Pointer to **[]string** |  | [optional] 
**AlwaysShowSubQuestions** | Pointer to **bool** |  | [optional] 
**ReportingOrder** | **float64** |  | 

## Methods

### NewCreateQuestionConfigBody

`func NewCreateQuestionConfigBody(name string, question string, type_ string, reportingOrder float64, ) *CreateQuestionConfigBody`

NewCreateQuestionConfigBody instantiates a new CreateQuestionConfigBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuestionConfigBodyWithDefaults

`func NewCreateQuestionConfigBodyWithDefaults() *CreateQuestionConfigBody`

NewCreateQuestionConfigBodyWithDefaults instantiates a new CreateQuestionConfigBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateQuestionConfigBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateQuestionConfigBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateQuestionConfigBody) SetName(v string)`

SetName sets Name field to given value.


### GetQuestion

`func (o *CreateQuestionConfigBody) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *CreateQuestionConfigBody) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *CreateQuestionConfigBody) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetHelpText

`func (o *CreateQuestionConfigBody) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *CreateQuestionConfigBody) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *CreateQuestionConfigBody) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *CreateQuestionConfigBody) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### GetType

`func (o *CreateQuestionConfigBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateQuestionConfigBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateQuestionConfigBody) SetType(v string)`

SetType sets Type field to given value.


### GetNumStars

`func (o *CreateQuestionConfigBody) GetNumStars() float64`

GetNumStars returns the NumStars field if non-nil, zero value otherwise.

### GetNumStarsOk

`func (o *CreateQuestionConfigBody) GetNumStarsOk() (*float64, bool)`

GetNumStarsOk returns a tuple with the NumStars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStars

`func (o *CreateQuestionConfigBody) SetNumStars(v float64)`

SetNumStars sets NumStars field to given value.

### HasNumStars

`func (o *CreateQuestionConfigBody) HasNumStars() bool`

HasNumStars returns a boolean if a field has been set.

### GetMin

`func (o *CreateQuestionConfigBody) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *CreateQuestionConfigBody) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *CreateQuestionConfigBody) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *CreateQuestionConfigBody) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *CreateQuestionConfigBody) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *CreateQuestionConfigBody) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *CreateQuestionConfigBody) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *CreateQuestionConfigBody) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetDefaultValue

`func (o *CreateQuestionConfigBody) GetDefaultValue() float64`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CreateQuestionConfigBody) GetDefaultValueOk() (*float64, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CreateQuestionConfigBody) SetDefaultValue(v float64)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CreateQuestionConfigBody) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetLabelNegative

`func (o *CreateQuestionConfigBody) GetLabelNegative() string`

GetLabelNegative returns the LabelNegative field if non-nil, zero value otherwise.

### GetLabelNegativeOk

`func (o *CreateQuestionConfigBody) GetLabelNegativeOk() (*string, bool)`

GetLabelNegativeOk returns a tuple with the LabelNegative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelNegative

`func (o *CreateQuestionConfigBody) SetLabelNegative(v string)`

SetLabelNegative sets LabelNegative field to given value.

### HasLabelNegative

`func (o *CreateQuestionConfigBody) HasLabelNegative() bool`

HasLabelNegative returns a boolean if a field has been set.

### GetLabelPositive

`func (o *CreateQuestionConfigBody) GetLabelPositive() string`

GetLabelPositive returns the LabelPositive field if non-nil, zero value otherwise.

### GetLabelPositiveOk

`func (o *CreateQuestionConfigBody) GetLabelPositiveOk() (*string, bool)`

GetLabelPositiveOk returns a tuple with the LabelPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPositive

`func (o *CreateQuestionConfigBody) SetLabelPositive(v string)`

SetLabelPositive sets LabelPositive field to given value.

### HasLabelPositive

`func (o *CreateQuestionConfigBody) HasLabelPositive() bool`

HasLabelPositive returns a boolean if a field has been set.

### GetCustomOptions

`func (o *CreateQuestionConfigBody) GetCustomOptions() []QuestionConfigCustomOptionsInner`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *CreateQuestionConfigBody) GetCustomOptionsOk() (*[]QuestionConfigCustomOptionsInner, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *CreateQuestionConfigBody) SetCustomOptions(v []QuestionConfigCustomOptionsInner)`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *CreateQuestionConfigBody) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetSubQuestionIds

`func (o *CreateQuestionConfigBody) GetSubQuestionIds() []string`

GetSubQuestionIds returns the SubQuestionIds field if non-nil, zero value otherwise.

### GetSubQuestionIdsOk

`func (o *CreateQuestionConfigBody) GetSubQuestionIdsOk() (*[]string, bool)`

GetSubQuestionIdsOk returns a tuple with the SubQuestionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubQuestionIds

`func (o *CreateQuestionConfigBody) SetSubQuestionIds(v []string)`

SetSubQuestionIds sets SubQuestionIds field to given value.

### HasSubQuestionIds

`func (o *CreateQuestionConfigBody) HasSubQuestionIds() bool`

HasSubQuestionIds returns a boolean if a field has been set.

### GetAlwaysShowSubQuestions

`func (o *CreateQuestionConfigBody) GetAlwaysShowSubQuestions() bool`

GetAlwaysShowSubQuestions returns the AlwaysShowSubQuestions field if non-nil, zero value otherwise.

### GetAlwaysShowSubQuestionsOk

`func (o *CreateQuestionConfigBody) GetAlwaysShowSubQuestionsOk() (*bool, bool)`

GetAlwaysShowSubQuestionsOk returns a tuple with the AlwaysShowSubQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysShowSubQuestions

`func (o *CreateQuestionConfigBody) SetAlwaysShowSubQuestions(v bool)`

SetAlwaysShowSubQuestions sets AlwaysShowSubQuestions field to given value.

### HasAlwaysShowSubQuestions

`func (o *CreateQuestionConfigBody) HasAlwaysShowSubQuestions() bool`

HasAlwaysShowSubQuestions returns a boolean if a field has been set.

### GetReportingOrder

`func (o *CreateQuestionConfigBody) GetReportingOrder() float64`

GetReportingOrder returns the ReportingOrder field if non-nil, zero value otherwise.

### GetReportingOrderOk

`func (o *CreateQuestionConfigBody) GetReportingOrderOk() (*float64, bool)`

GetReportingOrderOk returns a tuple with the ReportingOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingOrder

`func (o *CreateQuestionConfigBody) SetReportingOrder(v float64)`

SetReportingOrder sets ReportingOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


