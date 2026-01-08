# APIDomainConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Domain** | **string** |  | 
**EmailFromName** | Pointer to **NullableString** |  | [optional] 
**EmailFromEmail** | Pointer to **NullableString** |  | [optional] 
**EmailHeaders** | Pointer to **map[string]string** | Construct a type with a set of properties K of type T | [optional] 
**WpSyncToken** | Pointer to **NullableString** |  | [optional] 
**WpSynced** | Pointer to **bool** |  | [optional] 
**WpURL** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**AutoAddedDate** | Pointer to **time.Time** |  | [optional] 
**SiteType** | Pointer to [**ImportedSiteType**](ImportedSiteType.md) |  | [optional] 
**LogoSrc** | Pointer to **NullableString** |  | [optional] 
**LogoSrc100px** | Pointer to **NullableString** |  | [optional] 
**FooterUnsubscribeURL** | Pointer to **string** |  | [optional] 
**DisableUnsubscribeLinks** | Pointer to **bool** |  | [optional] 

## Methods

### NewAPIDomainConfiguration

`func NewAPIDomainConfiguration(id string, domain string, createdAt time.Time, ) *APIDomainConfiguration`

NewAPIDomainConfiguration instantiates a new APIDomainConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIDomainConfigurationWithDefaults

`func NewAPIDomainConfigurationWithDefaults() *APIDomainConfiguration`

NewAPIDomainConfigurationWithDefaults instantiates a new APIDomainConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APIDomainConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIDomainConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIDomainConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetDomain

`func (o *APIDomainConfiguration) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *APIDomainConfiguration) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *APIDomainConfiguration) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEmailFromName

`func (o *APIDomainConfiguration) GetEmailFromName() string`

GetEmailFromName returns the EmailFromName field if non-nil, zero value otherwise.

### GetEmailFromNameOk

`func (o *APIDomainConfiguration) GetEmailFromNameOk() (*string, bool)`

GetEmailFromNameOk returns a tuple with the EmailFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFromName

`func (o *APIDomainConfiguration) SetEmailFromName(v string)`

SetEmailFromName sets EmailFromName field to given value.

### HasEmailFromName

`func (o *APIDomainConfiguration) HasEmailFromName() bool`

HasEmailFromName returns a boolean if a field has been set.

### SetEmailFromNameNil

`func (o *APIDomainConfiguration) SetEmailFromNameNil(b bool)`

 SetEmailFromNameNil sets the value for EmailFromName to be an explicit nil

### UnsetEmailFromName
`func (o *APIDomainConfiguration) UnsetEmailFromName()`

UnsetEmailFromName ensures that no value is present for EmailFromName, not even an explicit nil
### GetEmailFromEmail

`func (o *APIDomainConfiguration) GetEmailFromEmail() string`

GetEmailFromEmail returns the EmailFromEmail field if non-nil, zero value otherwise.

### GetEmailFromEmailOk

`func (o *APIDomainConfiguration) GetEmailFromEmailOk() (*string, bool)`

GetEmailFromEmailOk returns a tuple with the EmailFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFromEmail

`func (o *APIDomainConfiguration) SetEmailFromEmail(v string)`

SetEmailFromEmail sets EmailFromEmail field to given value.

### HasEmailFromEmail

`func (o *APIDomainConfiguration) HasEmailFromEmail() bool`

HasEmailFromEmail returns a boolean if a field has been set.

### SetEmailFromEmailNil

`func (o *APIDomainConfiguration) SetEmailFromEmailNil(b bool)`

 SetEmailFromEmailNil sets the value for EmailFromEmail to be an explicit nil

### UnsetEmailFromEmail
`func (o *APIDomainConfiguration) UnsetEmailFromEmail()`

UnsetEmailFromEmail ensures that no value is present for EmailFromEmail, not even an explicit nil
### GetEmailHeaders

`func (o *APIDomainConfiguration) GetEmailHeaders() map[string]string`

GetEmailHeaders returns the EmailHeaders field if non-nil, zero value otherwise.

### GetEmailHeadersOk

`func (o *APIDomainConfiguration) GetEmailHeadersOk() (*map[string]string, bool)`

GetEmailHeadersOk returns a tuple with the EmailHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailHeaders

`func (o *APIDomainConfiguration) SetEmailHeaders(v map[string]string)`

SetEmailHeaders sets EmailHeaders field to given value.

### HasEmailHeaders

`func (o *APIDomainConfiguration) HasEmailHeaders() bool`

HasEmailHeaders returns a boolean if a field has been set.

### GetWpSyncToken

`func (o *APIDomainConfiguration) GetWpSyncToken() string`

GetWpSyncToken returns the WpSyncToken field if non-nil, zero value otherwise.

### GetWpSyncTokenOk

`func (o *APIDomainConfiguration) GetWpSyncTokenOk() (*string, bool)`

GetWpSyncTokenOk returns a tuple with the WpSyncToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpSyncToken

`func (o *APIDomainConfiguration) SetWpSyncToken(v string)`

SetWpSyncToken sets WpSyncToken field to given value.

### HasWpSyncToken

`func (o *APIDomainConfiguration) HasWpSyncToken() bool`

HasWpSyncToken returns a boolean if a field has been set.

### SetWpSyncTokenNil

`func (o *APIDomainConfiguration) SetWpSyncTokenNil(b bool)`

 SetWpSyncTokenNil sets the value for WpSyncToken to be an explicit nil

### UnsetWpSyncToken
`func (o *APIDomainConfiguration) UnsetWpSyncToken()`

UnsetWpSyncToken ensures that no value is present for WpSyncToken, not even an explicit nil
### GetWpSynced

`func (o *APIDomainConfiguration) GetWpSynced() bool`

GetWpSynced returns the WpSynced field if non-nil, zero value otherwise.

### GetWpSyncedOk

`func (o *APIDomainConfiguration) GetWpSyncedOk() (*bool, bool)`

GetWpSyncedOk returns a tuple with the WpSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpSynced

`func (o *APIDomainConfiguration) SetWpSynced(v bool)`

SetWpSynced sets WpSynced field to given value.

### HasWpSynced

`func (o *APIDomainConfiguration) HasWpSynced() bool`

HasWpSynced returns a boolean if a field has been set.

### GetWpURL

`func (o *APIDomainConfiguration) GetWpURL() string`

GetWpURL returns the WpURL field if non-nil, zero value otherwise.

### GetWpURLOk

`func (o *APIDomainConfiguration) GetWpURLOk() (*string, bool)`

GetWpURLOk returns a tuple with the WpURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpURL

`func (o *APIDomainConfiguration) SetWpURL(v string)`

SetWpURL sets WpURL field to given value.

### HasWpURL

`func (o *APIDomainConfiguration) HasWpURL() bool`

HasWpURL returns a boolean if a field has been set.

### SetWpURLNil

`func (o *APIDomainConfiguration) SetWpURLNil(b bool)`

 SetWpURLNil sets the value for WpURL to be an explicit nil

### UnsetWpURL
`func (o *APIDomainConfiguration) UnsetWpURL()`

UnsetWpURL ensures that no value is present for WpURL, not even an explicit nil
### GetCreatedAt

`func (o *APIDomainConfiguration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APIDomainConfiguration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APIDomainConfiguration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetAutoAddedDate

`func (o *APIDomainConfiguration) GetAutoAddedDate() time.Time`

GetAutoAddedDate returns the AutoAddedDate field if non-nil, zero value otherwise.

### GetAutoAddedDateOk

`func (o *APIDomainConfiguration) GetAutoAddedDateOk() (*time.Time, bool)`

GetAutoAddedDateOk returns a tuple with the AutoAddedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAddedDate

`func (o *APIDomainConfiguration) SetAutoAddedDate(v time.Time)`

SetAutoAddedDate sets AutoAddedDate field to given value.

### HasAutoAddedDate

`func (o *APIDomainConfiguration) HasAutoAddedDate() bool`

HasAutoAddedDate returns a boolean if a field has been set.

### GetSiteType

`func (o *APIDomainConfiguration) GetSiteType() ImportedSiteType`

GetSiteType returns the SiteType field if non-nil, zero value otherwise.

### GetSiteTypeOk

`func (o *APIDomainConfiguration) GetSiteTypeOk() (*ImportedSiteType, bool)`

GetSiteTypeOk returns a tuple with the SiteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteType

`func (o *APIDomainConfiguration) SetSiteType(v ImportedSiteType)`

SetSiteType sets SiteType field to given value.

### HasSiteType

`func (o *APIDomainConfiguration) HasSiteType() bool`

HasSiteType returns a boolean if a field has been set.

### GetLogoSrc

`func (o *APIDomainConfiguration) GetLogoSrc() string`

GetLogoSrc returns the LogoSrc field if non-nil, zero value otherwise.

### GetLogoSrcOk

`func (o *APIDomainConfiguration) GetLogoSrcOk() (*string, bool)`

GetLogoSrcOk returns a tuple with the LogoSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoSrc

`func (o *APIDomainConfiguration) SetLogoSrc(v string)`

SetLogoSrc sets LogoSrc field to given value.

### HasLogoSrc

`func (o *APIDomainConfiguration) HasLogoSrc() bool`

HasLogoSrc returns a boolean if a field has been set.

### SetLogoSrcNil

`func (o *APIDomainConfiguration) SetLogoSrcNil(b bool)`

 SetLogoSrcNil sets the value for LogoSrc to be an explicit nil

### UnsetLogoSrc
`func (o *APIDomainConfiguration) UnsetLogoSrc()`

UnsetLogoSrc ensures that no value is present for LogoSrc, not even an explicit nil
### GetLogoSrc100px

`func (o *APIDomainConfiguration) GetLogoSrc100px() string`

GetLogoSrc100px returns the LogoSrc100px field if non-nil, zero value otherwise.

### GetLogoSrc100pxOk

`func (o *APIDomainConfiguration) GetLogoSrc100pxOk() (*string, bool)`

GetLogoSrc100pxOk returns a tuple with the LogoSrc100px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoSrc100px

`func (o *APIDomainConfiguration) SetLogoSrc100px(v string)`

SetLogoSrc100px sets LogoSrc100px field to given value.

### HasLogoSrc100px

`func (o *APIDomainConfiguration) HasLogoSrc100px() bool`

HasLogoSrc100px returns a boolean if a field has been set.

### SetLogoSrc100pxNil

`func (o *APIDomainConfiguration) SetLogoSrc100pxNil(b bool)`

 SetLogoSrc100pxNil sets the value for LogoSrc100px to be an explicit nil

### UnsetLogoSrc100px
`func (o *APIDomainConfiguration) UnsetLogoSrc100px()`

UnsetLogoSrc100px ensures that no value is present for LogoSrc100px, not even an explicit nil
### GetFooterUnsubscribeURL

`func (o *APIDomainConfiguration) GetFooterUnsubscribeURL() string`

GetFooterUnsubscribeURL returns the FooterUnsubscribeURL field if non-nil, zero value otherwise.

### GetFooterUnsubscribeURLOk

`func (o *APIDomainConfiguration) GetFooterUnsubscribeURLOk() (*string, bool)`

GetFooterUnsubscribeURLOk returns a tuple with the FooterUnsubscribeURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterUnsubscribeURL

`func (o *APIDomainConfiguration) SetFooterUnsubscribeURL(v string)`

SetFooterUnsubscribeURL sets FooterUnsubscribeURL field to given value.

### HasFooterUnsubscribeURL

`func (o *APIDomainConfiguration) HasFooterUnsubscribeURL() bool`

HasFooterUnsubscribeURL returns a boolean if a field has been set.

### GetDisableUnsubscribeLinks

`func (o *APIDomainConfiguration) GetDisableUnsubscribeLinks() bool`

GetDisableUnsubscribeLinks returns the DisableUnsubscribeLinks field if non-nil, zero value otherwise.

### GetDisableUnsubscribeLinksOk

`func (o *APIDomainConfiguration) GetDisableUnsubscribeLinksOk() (*bool, bool)`

GetDisableUnsubscribeLinksOk returns a tuple with the DisableUnsubscribeLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUnsubscribeLinks

`func (o *APIDomainConfiguration) SetDisableUnsubscribeLinks(v bool)`

SetDisableUnsubscribeLinks sets DisableUnsubscribeLinks field to given value.

### HasDisableUnsubscribeLinks

`func (o *APIDomainConfiguration) HasDisableUnsubscribeLinks() bool`

HasDisableUnsubscribeLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


