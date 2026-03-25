# TOSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**TextByLocale** | Pointer to **map[string]string** | Construct a type with a set of properties K of type T | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTOSConfig

`func NewTOSConfig() *TOSConfig`

NewTOSConfig instantiates a new TOSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTOSConfigWithDefaults

`func NewTOSConfigWithDefaults() *TOSConfig`

NewTOSConfigWithDefaults instantiates a new TOSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *TOSConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TOSConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TOSConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TOSConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTextByLocale

`func (o *TOSConfig) GetTextByLocale() map[string]string`

GetTextByLocale returns the TextByLocale field if non-nil, zero value otherwise.

### GetTextByLocaleOk

`func (o *TOSConfig) GetTextByLocaleOk() (*map[string]string, bool)`

GetTextByLocaleOk returns a tuple with the TextByLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextByLocale

`func (o *TOSConfig) SetTextByLocale(v map[string]string)`

SetTextByLocale sets TextByLocale field to given value.

### HasTextByLocale

`func (o *TOSConfig) HasTextByLocale() bool`

HasTextByLocale returns a boolean if a field has been set.

### GetLastUpdated

`func (o *TOSConfig) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *TOSConfig) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *TOSConfig) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *TOSConfig) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


