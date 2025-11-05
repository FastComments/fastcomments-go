# UpdateDomainConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** |  | 
**EmailFromName** | Pointer to **string** |  | [optional] 
**EmailFromEmail** | Pointer to **string** |  | [optional] 
**LogoSrc** | Pointer to **string** |  | [optional] 
**LogoSrc100px** | Pointer to **string** |  | [optional] 
**FooterUnsubscribeURL** | Pointer to **string** |  | [optional] 
**EmailHeaders** | Pointer to **map[string]string** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewUpdateDomainConfigParams

`func NewUpdateDomainConfigParams(domain string, ) *UpdateDomainConfigParams`

NewUpdateDomainConfigParams instantiates a new UpdateDomainConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainConfigParamsWithDefaults

`func NewUpdateDomainConfigParamsWithDefaults() *UpdateDomainConfigParams`

NewUpdateDomainConfigParamsWithDefaults instantiates a new UpdateDomainConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *UpdateDomainConfigParams) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UpdateDomainConfigParams) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UpdateDomainConfigParams) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEmailFromName

`func (o *UpdateDomainConfigParams) GetEmailFromName() string`

GetEmailFromName returns the EmailFromName field if non-nil, zero value otherwise.

### GetEmailFromNameOk

`func (o *UpdateDomainConfigParams) GetEmailFromNameOk() (*string, bool)`

GetEmailFromNameOk returns a tuple with the EmailFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFromName

`func (o *UpdateDomainConfigParams) SetEmailFromName(v string)`

SetEmailFromName sets EmailFromName field to given value.

### HasEmailFromName

`func (o *UpdateDomainConfigParams) HasEmailFromName() bool`

HasEmailFromName returns a boolean if a field has been set.

### GetEmailFromEmail

`func (o *UpdateDomainConfigParams) GetEmailFromEmail() string`

GetEmailFromEmail returns the EmailFromEmail field if non-nil, zero value otherwise.

### GetEmailFromEmailOk

`func (o *UpdateDomainConfigParams) GetEmailFromEmailOk() (*string, bool)`

GetEmailFromEmailOk returns a tuple with the EmailFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFromEmail

`func (o *UpdateDomainConfigParams) SetEmailFromEmail(v string)`

SetEmailFromEmail sets EmailFromEmail field to given value.

### HasEmailFromEmail

`func (o *UpdateDomainConfigParams) HasEmailFromEmail() bool`

HasEmailFromEmail returns a boolean if a field has been set.

### GetLogoSrc

`func (o *UpdateDomainConfigParams) GetLogoSrc() string`

GetLogoSrc returns the LogoSrc field if non-nil, zero value otherwise.

### GetLogoSrcOk

`func (o *UpdateDomainConfigParams) GetLogoSrcOk() (*string, bool)`

GetLogoSrcOk returns a tuple with the LogoSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoSrc

`func (o *UpdateDomainConfigParams) SetLogoSrc(v string)`

SetLogoSrc sets LogoSrc field to given value.

### HasLogoSrc

`func (o *UpdateDomainConfigParams) HasLogoSrc() bool`

HasLogoSrc returns a boolean if a field has been set.

### GetLogoSrc100px

`func (o *UpdateDomainConfigParams) GetLogoSrc100px() string`

GetLogoSrc100px returns the LogoSrc100px field if non-nil, zero value otherwise.

### GetLogoSrc100pxOk

`func (o *UpdateDomainConfigParams) GetLogoSrc100pxOk() (*string, bool)`

GetLogoSrc100pxOk returns a tuple with the LogoSrc100px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoSrc100px

`func (o *UpdateDomainConfigParams) SetLogoSrc100px(v string)`

SetLogoSrc100px sets LogoSrc100px field to given value.

### HasLogoSrc100px

`func (o *UpdateDomainConfigParams) HasLogoSrc100px() bool`

HasLogoSrc100px returns a boolean if a field has been set.

### GetFooterUnsubscribeURL

`func (o *UpdateDomainConfigParams) GetFooterUnsubscribeURL() string`

GetFooterUnsubscribeURL returns the FooterUnsubscribeURL field if non-nil, zero value otherwise.

### GetFooterUnsubscribeURLOk

`func (o *UpdateDomainConfigParams) GetFooterUnsubscribeURLOk() (*string, bool)`

GetFooterUnsubscribeURLOk returns a tuple with the FooterUnsubscribeURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterUnsubscribeURL

`func (o *UpdateDomainConfigParams) SetFooterUnsubscribeURL(v string)`

SetFooterUnsubscribeURL sets FooterUnsubscribeURL field to given value.

### HasFooterUnsubscribeURL

`func (o *UpdateDomainConfigParams) HasFooterUnsubscribeURL() bool`

HasFooterUnsubscribeURL returns a boolean if a field has been set.

### GetEmailHeaders

`func (o *UpdateDomainConfigParams) GetEmailHeaders() map[string]string`

GetEmailHeaders returns the EmailHeaders field if non-nil, zero value otherwise.

### GetEmailHeadersOk

`func (o *UpdateDomainConfigParams) GetEmailHeadersOk() (*map[string]string, bool)`

GetEmailHeadersOk returns a tuple with the EmailHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailHeaders

`func (o *UpdateDomainConfigParams) SetEmailHeaders(v map[string]string)`

SetEmailHeaders sets EmailHeaders field to given value.

### HasEmailHeaders

`func (o *UpdateDomainConfigParams) HasEmailHeaders() bool`

HasEmailHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


