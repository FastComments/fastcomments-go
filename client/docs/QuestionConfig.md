# QuestionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Name** | **string** |  | 
**Question** | **string** |  | 
**SummaryLabel** | Pointer to **string** |  | [optional] 
**HelpText** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**UsedCount** | **float64** |  | 
**LastUsed** | **time.Time** |  | 
**Type** | **string** |  | 
**NumStars** | **float64** |  | 
**Min** | **float64** |  | 
**Max** | **float64** |  | 
**DefaultValue** | **float64** |  | 
**LabelNegative** | **string** |  | 
**LabelPositive** | **string** |  | 
**CustomOptions** | [**[]QuestionConfigCustomOptionsInner**](QuestionConfigCustomOptionsInner.md) |  | 
**SubQuestionIds** | **[]string** |  | 
**AlwaysShowSubQuestions** | **bool** |  | 
**ReportingOrder** | **float64** |  | 

## Methods

### NewQuestionConfig

`func NewQuestionConfig(id string, tenantId string, name string, question string, helpText string, createdAt time.Time, createdBy string, usedCount float64, lastUsed time.Time, type_ string, numStars float64, min float64, max float64, defaultValue float64, labelNegative string, labelPositive string, customOptions []QuestionConfigCustomOptionsInner, subQuestionIds []string, alwaysShowSubQuestions bool, reportingOrder float64, ) *QuestionConfig`

NewQuestionConfig instantiates a new QuestionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionConfigWithDefaults

`func NewQuestionConfigWithDefaults() *QuestionConfig`

NewQuestionConfigWithDefaults instantiates a new QuestionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionConfig) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *QuestionConfig) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *QuestionConfig) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *QuestionConfig) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetName

`func (o *QuestionConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuestionConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuestionConfig) SetName(v string)`

SetName sets Name field to given value.


### GetQuestion

`func (o *QuestionConfig) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *QuestionConfig) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *QuestionConfig) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetSummaryLabel

`func (o *QuestionConfig) GetSummaryLabel() string`

GetSummaryLabel returns the SummaryLabel field if non-nil, zero value otherwise.

### GetSummaryLabelOk

`func (o *QuestionConfig) GetSummaryLabelOk() (*string, bool)`

GetSummaryLabelOk returns a tuple with the SummaryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryLabel

`func (o *QuestionConfig) SetSummaryLabel(v string)`

SetSummaryLabel sets SummaryLabel field to given value.

### HasSummaryLabel

`func (o *QuestionConfig) HasSummaryLabel() bool`

HasSummaryLabel returns a boolean if a field has been set.

### GetHelpText

`func (o *QuestionConfig) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *QuestionConfig) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *QuestionConfig) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.


### GetCreatedAt

`func (o *QuestionConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QuestionConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QuestionConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *QuestionConfig) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *QuestionConfig) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *QuestionConfig) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetUsedCount

`func (o *QuestionConfig) GetUsedCount() float64`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *QuestionConfig) GetUsedCountOk() (*float64, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *QuestionConfig) SetUsedCount(v float64)`

SetUsedCount sets UsedCount field to given value.


### GetLastUsed

`func (o *QuestionConfig) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *QuestionConfig) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *QuestionConfig) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.


### GetType

`func (o *QuestionConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuestionConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuestionConfig) SetType(v string)`

SetType sets Type field to given value.


### GetNumStars

`func (o *QuestionConfig) GetNumStars() float64`

GetNumStars returns the NumStars field if non-nil, zero value otherwise.

### GetNumStarsOk

`func (o *QuestionConfig) GetNumStarsOk() (*float64, bool)`

GetNumStarsOk returns a tuple with the NumStars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStars

`func (o *QuestionConfig) SetNumStars(v float64)`

SetNumStars sets NumStars field to given value.


### GetMin

`func (o *QuestionConfig) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *QuestionConfig) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *QuestionConfig) SetMin(v float64)`

SetMin sets Min field to given value.


### GetMax

`func (o *QuestionConfig) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *QuestionConfig) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *QuestionConfig) SetMax(v float64)`

SetMax sets Max field to given value.


### GetDefaultValue

`func (o *QuestionConfig) GetDefaultValue() float64`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *QuestionConfig) GetDefaultValueOk() (*float64, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *QuestionConfig) SetDefaultValue(v float64)`

SetDefaultValue sets DefaultValue field to given value.


### GetLabelNegative

`func (o *QuestionConfig) GetLabelNegative() string`

GetLabelNegative returns the LabelNegative field if non-nil, zero value otherwise.

### GetLabelNegativeOk

`func (o *QuestionConfig) GetLabelNegativeOk() (*string, bool)`

GetLabelNegativeOk returns a tuple with the LabelNegative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelNegative

`func (o *QuestionConfig) SetLabelNegative(v string)`

SetLabelNegative sets LabelNegative field to given value.


### GetLabelPositive

`func (o *QuestionConfig) GetLabelPositive() string`

GetLabelPositive returns the LabelPositive field if non-nil, zero value otherwise.

### GetLabelPositiveOk

`func (o *QuestionConfig) GetLabelPositiveOk() (*string, bool)`

GetLabelPositiveOk returns a tuple with the LabelPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPositive

`func (o *QuestionConfig) SetLabelPositive(v string)`

SetLabelPositive sets LabelPositive field to given value.


### GetCustomOptions

`func (o *QuestionConfig) GetCustomOptions() []QuestionConfigCustomOptionsInner`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *QuestionConfig) GetCustomOptionsOk() (*[]QuestionConfigCustomOptionsInner, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *QuestionConfig) SetCustomOptions(v []QuestionConfigCustomOptionsInner)`

SetCustomOptions sets CustomOptions field to given value.


### GetSubQuestionIds

`func (o *QuestionConfig) GetSubQuestionIds() []string`

GetSubQuestionIds returns the SubQuestionIds field if non-nil, zero value otherwise.

### GetSubQuestionIdsOk

`func (o *QuestionConfig) GetSubQuestionIdsOk() (*[]string, bool)`

GetSubQuestionIdsOk returns a tuple with the SubQuestionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubQuestionIds

`func (o *QuestionConfig) SetSubQuestionIds(v []string)`

SetSubQuestionIds sets SubQuestionIds field to given value.


### GetAlwaysShowSubQuestions

`func (o *QuestionConfig) GetAlwaysShowSubQuestions() bool`

GetAlwaysShowSubQuestions returns the AlwaysShowSubQuestions field if non-nil, zero value otherwise.

### GetAlwaysShowSubQuestionsOk

`func (o *QuestionConfig) GetAlwaysShowSubQuestionsOk() (*bool, bool)`

GetAlwaysShowSubQuestionsOk returns a tuple with the AlwaysShowSubQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysShowSubQuestions

`func (o *QuestionConfig) SetAlwaysShowSubQuestions(v bool)`

SetAlwaysShowSubQuestions sets AlwaysShowSubQuestions field to given value.


### GetReportingOrder

`func (o *QuestionConfig) GetReportingOrder() float64`

GetReportingOrder returns the ReportingOrder field if non-nil, zero value otherwise.

### GetReportingOrderOk

`func (o *QuestionConfig) GetReportingOrderOk() (*float64, bool)`

GetReportingOrderOk returns a tuple with the ReportingOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingOrder

`func (o *QuestionConfig) SetReportingOrder(v float64)`

SetReportingOrder sets ReportingOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


