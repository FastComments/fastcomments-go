# GetTranslationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Translations** | **map[string]string** | Construct a type with a set of properties K of type T | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewGetTranslationsResponse

`func NewGetTranslationsResponse(translations map[string]string, status APIStatus, ) *GetTranslationsResponse`

NewGetTranslationsResponse instantiates a new GetTranslationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTranslationsResponseWithDefaults

`func NewGetTranslationsResponseWithDefaults() *GetTranslationsResponse`

NewGetTranslationsResponseWithDefaults instantiates a new GetTranslationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTranslations

`func (o *GetTranslationsResponse) GetTranslations() map[string]string`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *GetTranslationsResponse) GetTranslationsOk() (*map[string]string, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *GetTranslationsResponse) SetTranslations(v map[string]string)`

SetTranslations sets Translations field to given value.


### GetStatus

`func (o *GetTranslationsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTranslationsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTranslationsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


