# CustomEmailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**EmailTemplateId** | **string** |  | 
**DisplayName** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **NullableTime** |  | 
**UpdatedByUserId** | **NullableString** |  | 
**Domain** | Pointer to **NullableString** |  | [optional] 
**Ejs** | **string** |  | 
**TranslationOverridesByLocale** | **map[string]map[string]string** | Construct a type with a set of properties K of type T | 
**TestData** | **interface{}** |  | 

## Methods

### NewCustomEmailTemplate

`func NewCustomEmailTemplate(id string, tenantId string, emailTemplateId string, displayName string, createdAt time.Time, updatedAt NullableTime, updatedByUserId NullableString, ejs string, translationOverridesByLocale map[string]map[string]string, testData interface{}, ) *CustomEmailTemplate`

NewCustomEmailTemplate instantiates a new CustomEmailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEmailTemplateWithDefaults

`func NewCustomEmailTemplateWithDefaults() *CustomEmailTemplate`

NewCustomEmailTemplateWithDefaults instantiates a new CustomEmailTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomEmailTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomEmailTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomEmailTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *CustomEmailTemplate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CustomEmailTemplate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CustomEmailTemplate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetEmailTemplateId

`func (o *CustomEmailTemplate) GetEmailTemplateId() string`

GetEmailTemplateId returns the EmailTemplateId field if non-nil, zero value otherwise.

### GetEmailTemplateIdOk

`func (o *CustomEmailTemplate) GetEmailTemplateIdOk() (*string, bool)`

GetEmailTemplateIdOk returns a tuple with the EmailTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateId

`func (o *CustomEmailTemplate) SetEmailTemplateId(v string)`

SetEmailTemplateId sets EmailTemplateId field to given value.


### GetDisplayName

`func (o *CustomEmailTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CustomEmailTemplate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CustomEmailTemplate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetCreatedAt

`func (o *CustomEmailTemplate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomEmailTemplate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomEmailTemplate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CustomEmailTemplate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomEmailTemplate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomEmailTemplate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *CustomEmailTemplate) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CustomEmailTemplate) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUpdatedByUserId

`func (o *CustomEmailTemplate) GetUpdatedByUserId() string`

GetUpdatedByUserId returns the UpdatedByUserId field if non-nil, zero value otherwise.

### GetUpdatedByUserIdOk

`func (o *CustomEmailTemplate) GetUpdatedByUserIdOk() (*string, bool)`

GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserId

`func (o *CustomEmailTemplate) SetUpdatedByUserId(v string)`

SetUpdatedByUserId sets UpdatedByUserId field to given value.


### SetUpdatedByUserIdNil

`func (o *CustomEmailTemplate) SetUpdatedByUserIdNil(b bool)`

 SetUpdatedByUserIdNil sets the value for UpdatedByUserId to be an explicit nil

### UnsetUpdatedByUserId
`func (o *CustomEmailTemplate) UnsetUpdatedByUserId()`

UnsetUpdatedByUserId ensures that no value is present for UpdatedByUserId, not even an explicit nil
### GetDomain

`func (o *CustomEmailTemplate) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CustomEmailTemplate) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CustomEmailTemplate) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CustomEmailTemplate) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *CustomEmailTemplate) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *CustomEmailTemplate) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetEjs

`func (o *CustomEmailTemplate) GetEjs() string`

GetEjs returns the Ejs field if non-nil, zero value otherwise.

### GetEjsOk

`func (o *CustomEmailTemplate) GetEjsOk() (*string, bool)`

GetEjsOk returns a tuple with the Ejs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEjs

`func (o *CustomEmailTemplate) SetEjs(v string)`

SetEjs sets Ejs field to given value.


### GetTranslationOverridesByLocale

`func (o *CustomEmailTemplate) GetTranslationOverridesByLocale() map[string]map[string]string`

GetTranslationOverridesByLocale returns the TranslationOverridesByLocale field if non-nil, zero value otherwise.

### GetTranslationOverridesByLocaleOk

`func (o *CustomEmailTemplate) GetTranslationOverridesByLocaleOk() (*map[string]map[string]string, bool)`

GetTranslationOverridesByLocaleOk returns a tuple with the TranslationOverridesByLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationOverridesByLocale

`func (o *CustomEmailTemplate) SetTranslationOverridesByLocale(v map[string]map[string]string)`

SetTranslationOverridesByLocale sets TranslationOverridesByLocale field to given value.


### GetTestData

`func (o *CustomEmailTemplate) GetTestData() interface{}`

GetTestData returns the TestData field if non-nil, zero value otherwise.

### GetTestDataOk

`func (o *CustomEmailTemplate) GetTestDataOk() (*interface{}, bool)`

GetTestDataOk returns a tuple with the TestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestData

`func (o *CustomEmailTemplate) SetTestData(v interface{})`

SetTestData sets TestData field to given value.


### SetTestDataNil

`func (o *CustomEmailTemplate) SetTestDataNil(b bool)`

 SetTestDataNil sets the value for TestData to be an explicit nil

### UnsetTestData
`func (o *CustomEmailTemplate) UnsetTestData()`

UnsetTestData ensures that no value is present for TestData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


