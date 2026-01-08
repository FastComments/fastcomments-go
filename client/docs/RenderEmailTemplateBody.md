# RenderEmailTemplateBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailTemplateId** | **string** |  | 
**Ejs** | **string** |  | 
**TestData** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 
**TranslationOverridesByLocale** | Pointer to **map[string]map[string]string** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewRenderEmailTemplateBody

`func NewRenderEmailTemplateBody(emailTemplateId string, ejs string, ) *RenderEmailTemplateBody`

NewRenderEmailTemplateBody instantiates a new RenderEmailTemplateBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderEmailTemplateBodyWithDefaults

`func NewRenderEmailTemplateBodyWithDefaults() *RenderEmailTemplateBody`

NewRenderEmailTemplateBodyWithDefaults instantiates a new RenderEmailTemplateBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailTemplateId

`func (o *RenderEmailTemplateBody) GetEmailTemplateId() string`

GetEmailTemplateId returns the EmailTemplateId field if non-nil, zero value otherwise.

### GetEmailTemplateIdOk

`func (o *RenderEmailTemplateBody) GetEmailTemplateIdOk() (*string, bool)`

GetEmailTemplateIdOk returns a tuple with the EmailTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateId

`func (o *RenderEmailTemplateBody) SetEmailTemplateId(v string)`

SetEmailTemplateId sets EmailTemplateId field to given value.


### GetEjs

`func (o *RenderEmailTemplateBody) GetEjs() string`

GetEjs returns the Ejs field if non-nil, zero value otherwise.

### GetEjsOk

`func (o *RenderEmailTemplateBody) GetEjsOk() (*string, bool)`

GetEjsOk returns a tuple with the Ejs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEjs

`func (o *RenderEmailTemplateBody) SetEjs(v string)`

SetEjs sets Ejs field to given value.


### GetTestData

`func (o *RenderEmailTemplateBody) GetTestData() map[string]interface{}`

GetTestData returns the TestData field if non-nil, zero value otherwise.

### GetTestDataOk

`func (o *RenderEmailTemplateBody) GetTestDataOk() (*map[string]interface{}, bool)`

GetTestDataOk returns a tuple with the TestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestData

`func (o *RenderEmailTemplateBody) SetTestData(v map[string]interface{})`

SetTestData sets TestData field to given value.

### HasTestData

`func (o *RenderEmailTemplateBody) HasTestData() bool`

HasTestData returns a boolean if a field has been set.

### GetTranslationOverridesByLocale

`func (o *RenderEmailTemplateBody) GetTranslationOverridesByLocale() map[string]map[string]string`

GetTranslationOverridesByLocale returns the TranslationOverridesByLocale field if non-nil, zero value otherwise.

### GetTranslationOverridesByLocaleOk

`func (o *RenderEmailTemplateBody) GetTranslationOverridesByLocaleOk() (*map[string]map[string]string, bool)`

GetTranslationOverridesByLocaleOk returns a tuple with the TranslationOverridesByLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationOverridesByLocale

`func (o *RenderEmailTemplateBody) SetTranslationOverridesByLocale(v map[string]map[string]string)`

SetTranslationOverridesByLocale sets TranslationOverridesByLocale field to given value.

### HasTranslationOverridesByLocale

`func (o *RenderEmailTemplateBody) HasTranslationOverridesByLocale() bool`

HasTranslationOverridesByLocale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


