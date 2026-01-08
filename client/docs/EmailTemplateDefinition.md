# EmailTemplateDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailTemplateId** | **string** |  | 
**DefaultTestData** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 
**DefaultTranslationsByLocale** | **map[string]map[string]string** | Construct a type with a set of properties K of type T | 
**DefaultEJS** | **string** |  | 

## Methods

### NewEmailTemplateDefinition

`func NewEmailTemplateDefinition(emailTemplateId string, defaultTestData map[string]interface{}, defaultTranslationsByLocale map[string]map[string]string, defaultEJS string, ) *EmailTemplateDefinition`

NewEmailTemplateDefinition instantiates a new EmailTemplateDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTemplateDefinitionWithDefaults

`func NewEmailTemplateDefinitionWithDefaults() *EmailTemplateDefinition`

NewEmailTemplateDefinitionWithDefaults instantiates a new EmailTemplateDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailTemplateId

`func (o *EmailTemplateDefinition) GetEmailTemplateId() string`

GetEmailTemplateId returns the EmailTemplateId field if non-nil, zero value otherwise.

### GetEmailTemplateIdOk

`func (o *EmailTemplateDefinition) GetEmailTemplateIdOk() (*string, bool)`

GetEmailTemplateIdOk returns a tuple with the EmailTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateId

`func (o *EmailTemplateDefinition) SetEmailTemplateId(v string)`

SetEmailTemplateId sets EmailTemplateId field to given value.


### GetDefaultTestData

`func (o *EmailTemplateDefinition) GetDefaultTestData() map[string]interface{}`

GetDefaultTestData returns the DefaultTestData field if non-nil, zero value otherwise.

### GetDefaultTestDataOk

`func (o *EmailTemplateDefinition) GetDefaultTestDataOk() (*map[string]interface{}, bool)`

GetDefaultTestDataOk returns a tuple with the DefaultTestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTestData

`func (o *EmailTemplateDefinition) SetDefaultTestData(v map[string]interface{})`

SetDefaultTestData sets DefaultTestData field to given value.


### GetDefaultTranslationsByLocale

`func (o *EmailTemplateDefinition) GetDefaultTranslationsByLocale() map[string]map[string]string`

GetDefaultTranslationsByLocale returns the DefaultTranslationsByLocale field if non-nil, zero value otherwise.

### GetDefaultTranslationsByLocaleOk

`func (o *EmailTemplateDefinition) GetDefaultTranslationsByLocaleOk() (*map[string]map[string]string, bool)`

GetDefaultTranslationsByLocaleOk returns a tuple with the DefaultTranslationsByLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTranslationsByLocale

`func (o *EmailTemplateDefinition) SetDefaultTranslationsByLocale(v map[string]map[string]string)`

SetDefaultTranslationsByLocale sets DefaultTranslationsByLocale field to given value.


### GetDefaultEJS

`func (o *EmailTemplateDefinition) GetDefaultEJS() string`

GetDefaultEJS returns the DefaultEJS field if non-nil, zero value otherwise.

### GetDefaultEJSOk

`func (o *EmailTemplateDefinition) GetDefaultEJSOk() (*string, bool)`

GetDefaultEJSOk returns a tuple with the DefaultEJS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEJS

`func (o *EmailTemplateDefinition) SetDefaultEJS(v string)`

SetDefaultEJS sets DefaultEJS field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


