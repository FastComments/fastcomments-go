# UpdateEmailTemplateBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailTemplateId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Ejs** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**TranslationOverridesByLocale** | Pointer to **map[string]map[string]string** | Construct a type with a set of properties K of type T | [optional] 
**TestData** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewUpdateEmailTemplateBody

`func NewUpdateEmailTemplateBody() *UpdateEmailTemplateBody`

NewUpdateEmailTemplateBody instantiates a new UpdateEmailTemplateBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEmailTemplateBodyWithDefaults

`func NewUpdateEmailTemplateBodyWithDefaults() *UpdateEmailTemplateBody`

NewUpdateEmailTemplateBodyWithDefaults instantiates a new UpdateEmailTemplateBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailTemplateId

`func (o *UpdateEmailTemplateBody) GetEmailTemplateId() string`

GetEmailTemplateId returns the EmailTemplateId field if non-nil, zero value otherwise.

### GetEmailTemplateIdOk

`func (o *UpdateEmailTemplateBody) GetEmailTemplateIdOk() (*string, bool)`

GetEmailTemplateIdOk returns a tuple with the EmailTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateId

`func (o *UpdateEmailTemplateBody) SetEmailTemplateId(v string)`

SetEmailTemplateId sets EmailTemplateId field to given value.

### HasEmailTemplateId

`func (o *UpdateEmailTemplateBody) HasEmailTemplateId() bool`

HasEmailTemplateId returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateEmailTemplateBody) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateEmailTemplateBody) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateEmailTemplateBody) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateEmailTemplateBody) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEjs

`func (o *UpdateEmailTemplateBody) GetEjs() string`

GetEjs returns the Ejs field if non-nil, zero value otherwise.

### GetEjsOk

`func (o *UpdateEmailTemplateBody) GetEjsOk() (*string, bool)`

GetEjsOk returns a tuple with the Ejs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEjs

`func (o *UpdateEmailTemplateBody) SetEjs(v string)`

SetEjs sets Ejs field to given value.

### HasEjs

`func (o *UpdateEmailTemplateBody) HasEjs() bool`

HasEjs returns a boolean if a field has been set.

### GetDomain

`func (o *UpdateEmailTemplateBody) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UpdateEmailTemplateBody) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UpdateEmailTemplateBody) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UpdateEmailTemplateBody) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTranslationOverridesByLocale

`func (o *UpdateEmailTemplateBody) GetTranslationOverridesByLocale() map[string]map[string]string`

GetTranslationOverridesByLocale returns the TranslationOverridesByLocale field if non-nil, zero value otherwise.

### GetTranslationOverridesByLocaleOk

`func (o *UpdateEmailTemplateBody) GetTranslationOverridesByLocaleOk() (*map[string]map[string]string, bool)`

GetTranslationOverridesByLocaleOk returns a tuple with the TranslationOverridesByLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationOverridesByLocale

`func (o *UpdateEmailTemplateBody) SetTranslationOverridesByLocale(v map[string]map[string]string)`

SetTranslationOverridesByLocale sets TranslationOverridesByLocale field to given value.

### HasTranslationOverridesByLocale

`func (o *UpdateEmailTemplateBody) HasTranslationOverridesByLocale() bool`

HasTranslationOverridesByLocale returns a boolean if a field has been set.

### GetTestData

`func (o *UpdateEmailTemplateBody) GetTestData() map[string]interface{}`

GetTestData returns the TestData field if non-nil, zero value otherwise.

### GetTestDataOk

`func (o *UpdateEmailTemplateBody) GetTestDataOk() (*map[string]interface{}, bool)`

GetTestDataOk returns a tuple with the TestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestData

`func (o *UpdateEmailTemplateBody) SetTestData(v map[string]interface{})`

SetTestData sets TestData field to given value.

### HasTestData

`func (o *UpdateEmailTemplateBody) HasTestData() bool`

HasTestData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


