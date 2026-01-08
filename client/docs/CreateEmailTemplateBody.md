# CreateEmailTemplateBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailTemplateId** | **string** |  | 
**DisplayName** | **string** |  | 
**Ejs** | **string** |  | 
**Domain** | Pointer to **string** |  | [optional] 
**TranslationOverridesByLocale** | Pointer to **map[string]map[string]string** | Construct a type with a set of properties K of type T | [optional] 
**TestData** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewCreateEmailTemplateBody

`func NewCreateEmailTemplateBody(emailTemplateId string, displayName string, ejs string, ) *CreateEmailTemplateBody`

NewCreateEmailTemplateBody instantiates a new CreateEmailTemplateBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmailTemplateBodyWithDefaults

`func NewCreateEmailTemplateBodyWithDefaults() *CreateEmailTemplateBody`

NewCreateEmailTemplateBodyWithDefaults instantiates a new CreateEmailTemplateBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailTemplateId

`func (o *CreateEmailTemplateBody) GetEmailTemplateId() string`

GetEmailTemplateId returns the EmailTemplateId field if non-nil, zero value otherwise.

### GetEmailTemplateIdOk

`func (o *CreateEmailTemplateBody) GetEmailTemplateIdOk() (*string, bool)`

GetEmailTemplateIdOk returns a tuple with the EmailTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateId

`func (o *CreateEmailTemplateBody) SetEmailTemplateId(v string)`

SetEmailTemplateId sets EmailTemplateId field to given value.


### GetDisplayName

`func (o *CreateEmailTemplateBody) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateEmailTemplateBody) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateEmailTemplateBody) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEjs

`func (o *CreateEmailTemplateBody) GetEjs() string`

GetEjs returns the Ejs field if non-nil, zero value otherwise.

### GetEjsOk

`func (o *CreateEmailTemplateBody) GetEjsOk() (*string, bool)`

GetEjsOk returns a tuple with the Ejs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEjs

`func (o *CreateEmailTemplateBody) SetEjs(v string)`

SetEjs sets Ejs field to given value.


### GetDomain

`func (o *CreateEmailTemplateBody) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateEmailTemplateBody) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateEmailTemplateBody) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateEmailTemplateBody) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTranslationOverridesByLocale

`func (o *CreateEmailTemplateBody) GetTranslationOverridesByLocale() map[string]map[string]string`

GetTranslationOverridesByLocale returns the TranslationOverridesByLocale field if non-nil, zero value otherwise.

### GetTranslationOverridesByLocaleOk

`func (o *CreateEmailTemplateBody) GetTranslationOverridesByLocaleOk() (*map[string]map[string]string, bool)`

GetTranslationOverridesByLocaleOk returns a tuple with the TranslationOverridesByLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationOverridesByLocale

`func (o *CreateEmailTemplateBody) SetTranslationOverridesByLocale(v map[string]map[string]string)`

SetTranslationOverridesByLocale sets TranslationOverridesByLocale field to given value.

### HasTranslationOverridesByLocale

`func (o *CreateEmailTemplateBody) HasTranslationOverridesByLocale() bool`

HasTranslationOverridesByLocale returns a boolean if a field has been set.

### GetTestData

`func (o *CreateEmailTemplateBody) GetTestData() map[string]interface{}`

GetTestData returns the TestData field if non-nil, zero value otherwise.

### GetTestDataOk

`func (o *CreateEmailTemplateBody) GetTestDataOk() (*map[string]interface{}, bool)`

GetTestDataOk returns a tuple with the TestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestData

`func (o *CreateEmailTemplateBody) SetTestData(v map[string]interface{})`

SetTestData sets TestData field to given value.

### HasTestData

`func (o *CreateEmailTemplateBody) HasTestData() bool`

HasTestData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


