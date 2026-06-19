# GetUserInternalProfileResponseProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommenterName** | Pointer to **string** |  | [optional] 
**FirstCommentDate** | Pointer to **NullableTime** |  | [optional] 
**IpHash** | Pointer to **string** |  | [optional] 
**CountryFlag** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**WebsiteUrl** | Pointer to **NullableString** |  | [optional] 
**Bio** | Pointer to **string** |  | [optional] 
**Karma** | Pointer to **float64** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**AvatarSrc** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**CommenterEmail** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**AnonUserId** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetUserInternalProfileResponseProfile

`func NewGetUserInternalProfileResponseProfile() *GetUserInternalProfileResponseProfile`

NewGetUserInternalProfileResponseProfile instantiates a new GetUserInternalProfileResponseProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserInternalProfileResponseProfileWithDefaults

`func NewGetUserInternalProfileResponseProfileWithDefaults() *GetUserInternalProfileResponseProfile`

NewGetUserInternalProfileResponseProfileWithDefaults instantiates a new GetUserInternalProfileResponseProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommenterName

`func (o *GetUserInternalProfileResponseProfile) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *GetUserInternalProfileResponseProfile) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *GetUserInternalProfileResponseProfile) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.

### HasCommenterName

`func (o *GetUserInternalProfileResponseProfile) HasCommenterName() bool`

HasCommenterName returns a boolean if a field has been set.

### GetFirstCommentDate

`func (o *GetUserInternalProfileResponseProfile) GetFirstCommentDate() time.Time`

GetFirstCommentDate returns the FirstCommentDate field if non-nil, zero value otherwise.

### GetFirstCommentDateOk

`func (o *GetUserInternalProfileResponseProfile) GetFirstCommentDateOk() (*time.Time, bool)`

GetFirstCommentDateOk returns a tuple with the FirstCommentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCommentDate

`func (o *GetUserInternalProfileResponseProfile) SetFirstCommentDate(v time.Time)`

SetFirstCommentDate sets FirstCommentDate field to given value.

### HasFirstCommentDate

`func (o *GetUserInternalProfileResponseProfile) HasFirstCommentDate() bool`

HasFirstCommentDate returns a boolean if a field has been set.

### SetFirstCommentDateNil

`func (o *GetUserInternalProfileResponseProfile) SetFirstCommentDateNil(b bool)`

 SetFirstCommentDateNil sets the value for FirstCommentDate to be an explicit nil

### UnsetFirstCommentDate
`func (o *GetUserInternalProfileResponseProfile) UnsetFirstCommentDate()`

UnsetFirstCommentDate ensures that no value is present for FirstCommentDate, not even an explicit nil
### GetIpHash

`func (o *GetUserInternalProfileResponseProfile) GetIpHash() string`

GetIpHash returns the IpHash field if non-nil, zero value otherwise.

### GetIpHashOk

`func (o *GetUserInternalProfileResponseProfile) GetIpHashOk() (*string, bool)`

GetIpHashOk returns a tuple with the IpHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpHash

`func (o *GetUserInternalProfileResponseProfile) SetIpHash(v string)`

SetIpHash sets IpHash field to given value.

### HasIpHash

`func (o *GetUserInternalProfileResponseProfile) HasIpHash() bool`

HasIpHash returns a boolean if a field has been set.

### GetCountryFlag

`func (o *GetUserInternalProfileResponseProfile) GetCountryFlag() string`

GetCountryFlag returns the CountryFlag field if non-nil, zero value otherwise.

### GetCountryFlagOk

`func (o *GetUserInternalProfileResponseProfile) GetCountryFlagOk() (*string, bool)`

GetCountryFlagOk returns a tuple with the CountryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryFlag

`func (o *GetUserInternalProfileResponseProfile) SetCountryFlag(v string)`

SetCountryFlag sets CountryFlag field to given value.

### HasCountryFlag

`func (o *GetUserInternalProfileResponseProfile) HasCountryFlag() bool`

HasCountryFlag returns a boolean if a field has been set.

### GetCountryCode

`func (o *GetUserInternalProfileResponseProfile) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GetUserInternalProfileResponseProfile) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GetUserInternalProfileResponseProfile) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *GetUserInternalProfileResponseProfile) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *GetUserInternalProfileResponseProfile) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *GetUserInternalProfileResponseProfile) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *GetUserInternalProfileResponseProfile) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *GetUserInternalProfileResponseProfile) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *GetUserInternalProfileResponseProfile) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *GetUserInternalProfileResponseProfile) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetBio

`func (o *GetUserInternalProfileResponseProfile) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *GetUserInternalProfileResponseProfile) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *GetUserInternalProfileResponseProfile) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *GetUserInternalProfileResponseProfile) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetKarma

`func (o *GetUserInternalProfileResponseProfile) GetKarma() float64`

GetKarma returns the Karma field if non-nil, zero value otherwise.

### GetKarmaOk

`func (o *GetUserInternalProfileResponseProfile) GetKarmaOk() (*float64, bool)`

GetKarmaOk returns a tuple with the Karma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKarma

`func (o *GetUserInternalProfileResponseProfile) SetKarma(v float64)`

SetKarma sets Karma field to given value.

### HasKarma

`func (o *GetUserInternalProfileResponseProfile) HasKarma() bool`

HasKarma returns a boolean if a field has been set.

### GetLocale

`func (o *GetUserInternalProfileResponseProfile) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *GetUserInternalProfileResponseProfile) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *GetUserInternalProfileResponseProfile) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *GetUserInternalProfileResponseProfile) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetVerified

`func (o *GetUserInternalProfileResponseProfile) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *GetUserInternalProfileResponseProfile) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *GetUserInternalProfileResponseProfile) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *GetUserInternalProfileResponseProfile) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetAvatarSrc

`func (o *GetUserInternalProfileResponseProfile) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *GetUserInternalProfileResponseProfile) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *GetUserInternalProfileResponseProfile) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *GetUserInternalProfileResponseProfile) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### SetAvatarSrcNil

`func (o *GetUserInternalProfileResponseProfile) SetAvatarSrcNil(b bool)`

 SetAvatarSrcNil sets the value for AvatarSrc to be an explicit nil

### UnsetAvatarSrc
`func (o *GetUserInternalProfileResponseProfile) UnsetAvatarSrc()`

UnsetAvatarSrc ensures that no value is present for AvatarSrc, not even an explicit nil
### GetDisplayName

`func (o *GetUserInternalProfileResponseProfile) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetUserInternalProfileResponseProfile) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetUserInternalProfileResponseProfile) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetUserInternalProfileResponseProfile) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUsername

`func (o *GetUserInternalProfileResponseProfile) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetUserInternalProfileResponseProfile) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetUserInternalProfileResponseProfile) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetUserInternalProfileResponseProfile) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetCommenterEmail

`func (o *GetUserInternalProfileResponseProfile) GetCommenterEmail() string`

GetCommenterEmail returns the CommenterEmail field if non-nil, zero value otherwise.

### GetCommenterEmailOk

`func (o *GetUserInternalProfileResponseProfile) GetCommenterEmailOk() (*string, bool)`

GetCommenterEmailOk returns a tuple with the CommenterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterEmail

`func (o *GetUserInternalProfileResponseProfile) SetCommenterEmail(v string)`

SetCommenterEmail sets CommenterEmail field to given value.

### HasCommenterEmail

`func (o *GetUserInternalProfileResponseProfile) HasCommenterEmail() bool`

HasCommenterEmail returns a boolean if a field has been set.

### SetCommenterEmailNil

`func (o *GetUserInternalProfileResponseProfile) SetCommenterEmailNil(b bool)`

 SetCommenterEmailNil sets the value for CommenterEmail to be an explicit nil

### UnsetCommenterEmail
`func (o *GetUserInternalProfileResponseProfile) UnsetCommenterEmail()`

UnsetCommenterEmail ensures that no value is present for CommenterEmail, not even an explicit nil
### GetEmail

`func (o *GetUserInternalProfileResponseProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetUserInternalProfileResponseProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetUserInternalProfileResponseProfile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetUserInternalProfileResponseProfile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GetUserInternalProfileResponseProfile) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GetUserInternalProfileResponseProfile) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAnonUserId

`func (o *GetUserInternalProfileResponseProfile) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *GetUserInternalProfileResponseProfile) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *GetUserInternalProfileResponseProfile) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.

### HasAnonUserId

`func (o *GetUserInternalProfileResponseProfile) HasAnonUserId() bool`

HasAnonUserId returns a boolean if a field has been set.

### SetAnonUserIdNil

`func (o *GetUserInternalProfileResponseProfile) SetAnonUserIdNil(b bool)`

 SetAnonUserIdNil sets the value for AnonUserId to be an explicit nil

### UnsetAnonUserId
`func (o *GetUserInternalProfileResponseProfile) UnsetAnonUserId()`

UnsetAnonUserId ensures that no value is present for AnonUserId, not even an explicit nil
### GetUserId

`func (o *GetUserInternalProfileResponseProfile) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetUserInternalProfileResponseProfile) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetUserInternalProfileResponseProfile) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetUserInternalProfileResponseProfile) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *GetUserInternalProfileResponseProfile) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *GetUserInternalProfileResponseProfile) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


