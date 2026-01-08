# UpdateQuestionConfigBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**HelpText** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**NumStars** | Pointer to **float64** |  | [optional] 
**Min** | Pointer to **float64** |  | [optional] 
**Max** | Pointer to **float64** |  | [optional] 
**DefaultValue** | Pointer to **float64** |  | [optional] 
**LabelNegative** | Pointer to **string** |  | [optional] 
**LabelPositive** | Pointer to **string** |  | [optional] 
**CustomOptions** | Pointer to [**[]QuestionConfigCustomOptionsInner**](QuestionConfigCustomOptionsInner.md) |  | [optional] 
**SubQuestionIds** | Pointer to **[]string** |  | [optional] 
**AlwaysShowSubQuestions** | Pointer to **bool** |  | [optional] 
**ReportingOrder** | Pointer to **float64** |  | [optional] 

## Methods

### NewUpdateQuestionConfigBody

`func NewUpdateQuestionConfigBody() *UpdateQuestionConfigBody`

NewUpdateQuestionConfigBody instantiates a new UpdateQuestionConfigBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateQuestionConfigBodyWithDefaults

`func NewUpdateQuestionConfigBodyWithDefaults() *UpdateQuestionConfigBody`

NewUpdateQuestionConfigBodyWithDefaults instantiates a new UpdateQuestionConfigBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateQuestionConfigBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateQuestionConfigBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateQuestionConfigBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateQuestionConfigBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuestion

`func (o *UpdateQuestionConfigBody) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *UpdateQuestionConfigBody) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *UpdateQuestionConfigBody) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *UpdateQuestionConfigBody) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetHelpText

`func (o *UpdateQuestionConfigBody) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *UpdateQuestionConfigBody) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *UpdateQuestionConfigBody) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *UpdateQuestionConfigBody) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### GetType

`func (o *UpdateQuestionConfigBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateQuestionConfigBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateQuestionConfigBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateQuestionConfigBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNumStars

`func (o *UpdateQuestionConfigBody) GetNumStars() float64`

GetNumStars returns the NumStars field if non-nil, zero value otherwise.

### GetNumStarsOk

`func (o *UpdateQuestionConfigBody) GetNumStarsOk() (*float64, bool)`

GetNumStarsOk returns a tuple with the NumStars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStars

`func (o *UpdateQuestionConfigBody) SetNumStars(v float64)`

SetNumStars sets NumStars field to given value.

### HasNumStars

`func (o *UpdateQuestionConfigBody) HasNumStars() bool`

HasNumStars returns a boolean if a field has been set.

### GetMin

`func (o *UpdateQuestionConfigBody) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *UpdateQuestionConfigBody) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *UpdateQuestionConfigBody) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *UpdateQuestionConfigBody) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *UpdateQuestionConfigBody) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *UpdateQuestionConfigBody) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *UpdateQuestionConfigBody) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *UpdateQuestionConfigBody) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetDefaultValue

`func (o *UpdateQuestionConfigBody) GetDefaultValue() float64`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *UpdateQuestionConfigBody) GetDefaultValueOk() (*float64, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *UpdateQuestionConfigBody) SetDefaultValue(v float64)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *UpdateQuestionConfigBody) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetLabelNegative

`func (o *UpdateQuestionConfigBody) GetLabelNegative() string`

GetLabelNegative returns the LabelNegative field if non-nil, zero value otherwise.

### GetLabelNegativeOk

`func (o *UpdateQuestionConfigBody) GetLabelNegativeOk() (*string, bool)`

GetLabelNegativeOk returns a tuple with the LabelNegative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelNegative

`func (o *UpdateQuestionConfigBody) SetLabelNegative(v string)`

SetLabelNegative sets LabelNegative field to given value.

### HasLabelNegative

`func (o *UpdateQuestionConfigBody) HasLabelNegative() bool`

HasLabelNegative returns a boolean if a field has been set.

### GetLabelPositive

`func (o *UpdateQuestionConfigBody) GetLabelPositive() string`

GetLabelPositive returns the LabelPositive field if non-nil, zero value otherwise.

### GetLabelPositiveOk

`func (o *UpdateQuestionConfigBody) GetLabelPositiveOk() (*string, bool)`

GetLabelPositiveOk returns a tuple with the LabelPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPositive

`func (o *UpdateQuestionConfigBody) SetLabelPositive(v string)`

SetLabelPositive sets LabelPositive field to given value.

### HasLabelPositive

`func (o *UpdateQuestionConfigBody) HasLabelPositive() bool`

HasLabelPositive returns a boolean if a field has been set.

### GetCustomOptions

`func (o *UpdateQuestionConfigBody) GetCustomOptions() []QuestionConfigCustomOptionsInner`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *UpdateQuestionConfigBody) GetCustomOptionsOk() (*[]QuestionConfigCustomOptionsInner, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *UpdateQuestionConfigBody) SetCustomOptions(v []QuestionConfigCustomOptionsInner)`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *UpdateQuestionConfigBody) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetSubQuestionIds

`func (o *UpdateQuestionConfigBody) GetSubQuestionIds() []string`

GetSubQuestionIds returns the SubQuestionIds field if non-nil, zero value otherwise.

### GetSubQuestionIdsOk

`func (o *UpdateQuestionConfigBody) GetSubQuestionIdsOk() (*[]string, bool)`

GetSubQuestionIdsOk returns a tuple with the SubQuestionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubQuestionIds

`func (o *UpdateQuestionConfigBody) SetSubQuestionIds(v []string)`

SetSubQuestionIds sets SubQuestionIds field to given value.

### HasSubQuestionIds

`func (o *UpdateQuestionConfigBody) HasSubQuestionIds() bool`

HasSubQuestionIds returns a boolean if a field has been set.

### GetAlwaysShowSubQuestions

`func (o *UpdateQuestionConfigBody) GetAlwaysShowSubQuestions() bool`

GetAlwaysShowSubQuestions returns the AlwaysShowSubQuestions field if non-nil, zero value otherwise.

### GetAlwaysShowSubQuestionsOk

`func (o *UpdateQuestionConfigBody) GetAlwaysShowSubQuestionsOk() (*bool, bool)`

GetAlwaysShowSubQuestionsOk returns a tuple with the AlwaysShowSubQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysShowSubQuestions

`func (o *UpdateQuestionConfigBody) SetAlwaysShowSubQuestions(v bool)`

SetAlwaysShowSubQuestions sets AlwaysShowSubQuestions field to given value.

### HasAlwaysShowSubQuestions

`func (o *UpdateQuestionConfigBody) HasAlwaysShowSubQuestions() bool`

HasAlwaysShowSubQuestions returns a boolean if a field has been set.

### GetReportingOrder

`func (o *UpdateQuestionConfigBody) GetReportingOrder() float64`

GetReportingOrder returns the ReportingOrder field if non-nil, zero value otherwise.

### GetReportingOrderOk

`func (o *UpdateQuestionConfigBody) GetReportingOrderOk() (*float64, bool)`

GetReportingOrderOk returns a tuple with the ReportingOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingOrder

`func (o *UpdateQuestionConfigBody) SetReportingOrder(v float64)`

SetReportingOrder sets ReportingOrder field to given value.

### HasReportingOrder

`func (o *UpdateQuestionConfigBody) HasReportingOrder() bool`

HasReportingOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


