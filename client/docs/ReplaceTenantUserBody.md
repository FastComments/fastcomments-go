# ReplaceTenantUserBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Email** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**WebsiteUrl** | Pointer to **string** |  | [optional] 
**SignUpDate** | Pointer to **float64** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**LoginCount** | Pointer to **float64** |  | [optional] 
**OptedInNotifications** | Pointer to **bool** |  | [optional] 
**OptedInTenantNotifications** | Pointer to **bool** |  | [optional] 
**HideAccountCode** | Pointer to **bool** |  | [optional] 
**AvatarSrc** | Pointer to **string** |  | [optional] 
**IsHelpRequestAdmin** | Pointer to **bool** |  | [optional] 
**IsAccountOwner** | Pointer to **bool** |  | [optional] 
**IsAdminAdmin** | Pointer to **bool** |  | [optional] 
**IsBillingAdmin** | Pointer to **bool** |  | [optional] 
**IsAnalyticsAdmin** | Pointer to **bool** |  | [optional] 
**IsCustomizationAdmin** | Pointer to **bool** |  | [optional] 
**IsManageDataAdmin** | Pointer to **bool** |  | [optional] 
**IsCommentModeratorAdmin** | Pointer to **bool** |  | [optional] 
**IsAPIAdmin** | Pointer to **bool** |  | [optional] 
**ModeratorIds** | Pointer to **[]string** |  | [optional] 
**DigestEmailFrequency** | Pointer to **float64** |  | [optional] 
**DisplayLabel** | Pointer to **string** |  | [optional] 
**CreatedFromUrlId** | Pointer to **string** |  | [optional] 
**CreatedFromTenantId** | Pointer to **string** |  | [optional] 
**LastLoginDate** | Pointer to **float64** |  | [optional] 
**Karma** | Pointer to **float64** |  | [optional] 

## Methods

### NewReplaceTenantUserBody

`func NewReplaceTenantUserBody(username string, email string, ) *ReplaceTenantUserBody`

NewReplaceTenantUserBody instantiates a new ReplaceTenantUserBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceTenantUserBodyWithDefaults

`func NewReplaceTenantUserBodyWithDefaults() *ReplaceTenantUserBody`

NewReplaceTenantUserBodyWithDefaults instantiates a new ReplaceTenantUserBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ReplaceTenantUserBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ReplaceTenantUserBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ReplaceTenantUserBody) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *ReplaceTenantUserBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ReplaceTenantUserBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ReplaceTenantUserBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDisplayName

`func (o *ReplaceTenantUserBody) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ReplaceTenantUserBody) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ReplaceTenantUserBody) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ReplaceTenantUserBody) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *ReplaceTenantUserBody) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *ReplaceTenantUserBody) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *ReplaceTenantUserBody) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *ReplaceTenantUserBody) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetSignUpDate

`func (o *ReplaceTenantUserBody) GetSignUpDate() float64`

GetSignUpDate returns the SignUpDate field if non-nil, zero value otherwise.

### GetSignUpDateOk

`func (o *ReplaceTenantUserBody) GetSignUpDateOk() (*float64, bool)`

GetSignUpDateOk returns a tuple with the SignUpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUpDate

`func (o *ReplaceTenantUserBody) SetSignUpDate(v float64)`

SetSignUpDate sets SignUpDate field to given value.

### HasSignUpDate

`func (o *ReplaceTenantUserBody) HasSignUpDate() bool`

HasSignUpDate returns a boolean if a field has been set.

### GetLocale

`func (o *ReplaceTenantUserBody) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ReplaceTenantUserBody) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ReplaceTenantUserBody) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *ReplaceTenantUserBody) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetVerified

`func (o *ReplaceTenantUserBody) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ReplaceTenantUserBody) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ReplaceTenantUserBody) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ReplaceTenantUserBody) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetLoginCount

`func (o *ReplaceTenantUserBody) GetLoginCount() float64`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *ReplaceTenantUserBody) GetLoginCountOk() (*float64, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *ReplaceTenantUserBody) SetLoginCount(v float64)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *ReplaceTenantUserBody) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### GetOptedInNotifications

`func (o *ReplaceTenantUserBody) GetOptedInNotifications() bool`

GetOptedInNotifications returns the OptedInNotifications field if non-nil, zero value otherwise.

### GetOptedInNotificationsOk

`func (o *ReplaceTenantUserBody) GetOptedInNotificationsOk() (*bool, bool)`

GetOptedInNotificationsOk returns a tuple with the OptedInNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedInNotifications

`func (o *ReplaceTenantUserBody) SetOptedInNotifications(v bool)`

SetOptedInNotifications sets OptedInNotifications field to given value.

### HasOptedInNotifications

`func (o *ReplaceTenantUserBody) HasOptedInNotifications() bool`

HasOptedInNotifications returns a boolean if a field has been set.

### GetOptedInTenantNotifications

`func (o *ReplaceTenantUserBody) GetOptedInTenantNotifications() bool`

GetOptedInTenantNotifications returns the OptedInTenantNotifications field if non-nil, zero value otherwise.

### GetOptedInTenantNotificationsOk

`func (o *ReplaceTenantUserBody) GetOptedInTenantNotificationsOk() (*bool, bool)`

GetOptedInTenantNotificationsOk returns a tuple with the OptedInTenantNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedInTenantNotifications

`func (o *ReplaceTenantUserBody) SetOptedInTenantNotifications(v bool)`

SetOptedInTenantNotifications sets OptedInTenantNotifications field to given value.

### HasOptedInTenantNotifications

`func (o *ReplaceTenantUserBody) HasOptedInTenantNotifications() bool`

HasOptedInTenantNotifications returns a boolean if a field has been set.

### GetHideAccountCode

`func (o *ReplaceTenantUserBody) GetHideAccountCode() bool`

GetHideAccountCode returns the HideAccountCode field if non-nil, zero value otherwise.

### GetHideAccountCodeOk

`func (o *ReplaceTenantUserBody) GetHideAccountCodeOk() (*bool, bool)`

GetHideAccountCodeOk returns a tuple with the HideAccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAccountCode

`func (o *ReplaceTenantUserBody) SetHideAccountCode(v bool)`

SetHideAccountCode sets HideAccountCode field to given value.

### HasHideAccountCode

`func (o *ReplaceTenantUserBody) HasHideAccountCode() bool`

HasHideAccountCode returns a boolean if a field has been set.

### GetAvatarSrc

`func (o *ReplaceTenantUserBody) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *ReplaceTenantUserBody) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *ReplaceTenantUserBody) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *ReplaceTenantUserBody) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### GetIsHelpRequestAdmin

`func (o *ReplaceTenantUserBody) GetIsHelpRequestAdmin() bool`

GetIsHelpRequestAdmin returns the IsHelpRequestAdmin field if non-nil, zero value otherwise.

### GetIsHelpRequestAdminOk

`func (o *ReplaceTenantUserBody) GetIsHelpRequestAdminOk() (*bool, bool)`

GetIsHelpRequestAdminOk returns a tuple with the IsHelpRequestAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHelpRequestAdmin

`func (o *ReplaceTenantUserBody) SetIsHelpRequestAdmin(v bool)`

SetIsHelpRequestAdmin sets IsHelpRequestAdmin field to given value.

### HasIsHelpRequestAdmin

`func (o *ReplaceTenantUserBody) HasIsHelpRequestAdmin() bool`

HasIsHelpRequestAdmin returns a boolean if a field has been set.

### GetIsAccountOwner

`func (o *ReplaceTenantUserBody) GetIsAccountOwner() bool`

GetIsAccountOwner returns the IsAccountOwner field if non-nil, zero value otherwise.

### GetIsAccountOwnerOk

`func (o *ReplaceTenantUserBody) GetIsAccountOwnerOk() (*bool, bool)`

GetIsAccountOwnerOk returns a tuple with the IsAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountOwner

`func (o *ReplaceTenantUserBody) SetIsAccountOwner(v bool)`

SetIsAccountOwner sets IsAccountOwner field to given value.

### HasIsAccountOwner

`func (o *ReplaceTenantUserBody) HasIsAccountOwner() bool`

HasIsAccountOwner returns a boolean if a field has been set.

### GetIsAdminAdmin

`func (o *ReplaceTenantUserBody) GetIsAdminAdmin() bool`

GetIsAdminAdmin returns the IsAdminAdmin field if non-nil, zero value otherwise.

### GetIsAdminAdminOk

`func (o *ReplaceTenantUserBody) GetIsAdminAdminOk() (*bool, bool)`

GetIsAdminAdminOk returns a tuple with the IsAdminAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdminAdmin

`func (o *ReplaceTenantUserBody) SetIsAdminAdmin(v bool)`

SetIsAdminAdmin sets IsAdminAdmin field to given value.

### HasIsAdminAdmin

`func (o *ReplaceTenantUserBody) HasIsAdminAdmin() bool`

HasIsAdminAdmin returns a boolean if a field has been set.

### GetIsBillingAdmin

`func (o *ReplaceTenantUserBody) GetIsBillingAdmin() bool`

GetIsBillingAdmin returns the IsBillingAdmin field if non-nil, zero value otherwise.

### GetIsBillingAdminOk

`func (o *ReplaceTenantUserBody) GetIsBillingAdminOk() (*bool, bool)`

GetIsBillingAdminOk returns a tuple with the IsBillingAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillingAdmin

`func (o *ReplaceTenantUserBody) SetIsBillingAdmin(v bool)`

SetIsBillingAdmin sets IsBillingAdmin field to given value.

### HasIsBillingAdmin

`func (o *ReplaceTenantUserBody) HasIsBillingAdmin() bool`

HasIsBillingAdmin returns a boolean if a field has been set.

### GetIsAnalyticsAdmin

`func (o *ReplaceTenantUserBody) GetIsAnalyticsAdmin() bool`

GetIsAnalyticsAdmin returns the IsAnalyticsAdmin field if non-nil, zero value otherwise.

### GetIsAnalyticsAdminOk

`func (o *ReplaceTenantUserBody) GetIsAnalyticsAdminOk() (*bool, bool)`

GetIsAnalyticsAdminOk returns a tuple with the IsAnalyticsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnalyticsAdmin

`func (o *ReplaceTenantUserBody) SetIsAnalyticsAdmin(v bool)`

SetIsAnalyticsAdmin sets IsAnalyticsAdmin field to given value.

### HasIsAnalyticsAdmin

`func (o *ReplaceTenantUserBody) HasIsAnalyticsAdmin() bool`

HasIsAnalyticsAdmin returns a boolean if a field has been set.

### GetIsCustomizationAdmin

`func (o *ReplaceTenantUserBody) GetIsCustomizationAdmin() bool`

GetIsCustomizationAdmin returns the IsCustomizationAdmin field if non-nil, zero value otherwise.

### GetIsCustomizationAdminOk

`func (o *ReplaceTenantUserBody) GetIsCustomizationAdminOk() (*bool, bool)`

GetIsCustomizationAdminOk returns a tuple with the IsCustomizationAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomizationAdmin

`func (o *ReplaceTenantUserBody) SetIsCustomizationAdmin(v bool)`

SetIsCustomizationAdmin sets IsCustomizationAdmin field to given value.

### HasIsCustomizationAdmin

`func (o *ReplaceTenantUserBody) HasIsCustomizationAdmin() bool`

HasIsCustomizationAdmin returns a boolean if a field has been set.

### GetIsManageDataAdmin

`func (o *ReplaceTenantUserBody) GetIsManageDataAdmin() bool`

GetIsManageDataAdmin returns the IsManageDataAdmin field if non-nil, zero value otherwise.

### GetIsManageDataAdminOk

`func (o *ReplaceTenantUserBody) GetIsManageDataAdminOk() (*bool, bool)`

GetIsManageDataAdminOk returns a tuple with the IsManageDataAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManageDataAdmin

`func (o *ReplaceTenantUserBody) SetIsManageDataAdmin(v bool)`

SetIsManageDataAdmin sets IsManageDataAdmin field to given value.

### HasIsManageDataAdmin

`func (o *ReplaceTenantUserBody) HasIsManageDataAdmin() bool`

HasIsManageDataAdmin returns a boolean if a field has been set.

### GetIsCommentModeratorAdmin

`func (o *ReplaceTenantUserBody) GetIsCommentModeratorAdmin() bool`

GetIsCommentModeratorAdmin returns the IsCommentModeratorAdmin field if non-nil, zero value otherwise.

### GetIsCommentModeratorAdminOk

`func (o *ReplaceTenantUserBody) GetIsCommentModeratorAdminOk() (*bool, bool)`

GetIsCommentModeratorAdminOk returns a tuple with the IsCommentModeratorAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCommentModeratorAdmin

`func (o *ReplaceTenantUserBody) SetIsCommentModeratorAdmin(v bool)`

SetIsCommentModeratorAdmin sets IsCommentModeratorAdmin field to given value.

### HasIsCommentModeratorAdmin

`func (o *ReplaceTenantUserBody) HasIsCommentModeratorAdmin() bool`

HasIsCommentModeratorAdmin returns a boolean if a field has been set.

### GetIsAPIAdmin

`func (o *ReplaceTenantUserBody) GetIsAPIAdmin() bool`

GetIsAPIAdmin returns the IsAPIAdmin field if non-nil, zero value otherwise.

### GetIsAPIAdminOk

`func (o *ReplaceTenantUserBody) GetIsAPIAdminOk() (*bool, bool)`

GetIsAPIAdminOk returns a tuple with the IsAPIAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAPIAdmin

`func (o *ReplaceTenantUserBody) SetIsAPIAdmin(v bool)`

SetIsAPIAdmin sets IsAPIAdmin field to given value.

### HasIsAPIAdmin

`func (o *ReplaceTenantUserBody) HasIsAPIAdmin() bool`

HasIsAPIAdmin returns a boolean if a field has been set.

### GetModeratorIds

`func (o *ReplaceTenantUserBody) GetModeratorIds() []string`

GetModeratorIds returns the ModeratorIds field if non-nil, zero value otherwise.

### GetModeratorIdsOk

`func (o *ReplaceTenantUserBody) GetModeratorIdsOk() (*[]string, bool)`

GetModeratorIdsOk returns a tuple with the ModeratorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeratorIds

`func (o *ReplaceTenantUserBody) SetModeratorIds(v []string)`

SetModeratorIds sets ModeratorIds field to given value.

### HasModeratorIds

`func (o *ReplaceTenantUserBody) HasModeratorIds() bool`

HasModeratorIds returns a boolean if a field has been set.

### GetDigestEmailFrequency

`func (o *ReplaceTenantUserBody) GetDigestEmailFrequency() float64`

GetDigestEmailFrequency returns the DigestEmailFrequency field if non-nil, zero value otherwise.

### GetDigestEmailFrequencyOk

`func (o *ReplaceTenantUserBody) GetDigestEmailFrequencyOk() (*float64, bool)`

GetDigestEmailFrequencyOk returns a tuple with the DigestEmailFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestEmailFrequency

`func (o *ReplaceTenantUserBody) SetDigestEmailFrequency(v float64)`

SetDigestEmailFrequency sets DigestEmailFrequency field to given value.

### HasDigestEmailFrequency

`func (o *ReplaceTenantUserBody) HasDigestEmailFrequency() bool`

HasDigestEmailFrequency returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *ReplaceTenantUserBody) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *ReplaceTenantUserBody) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *ReplaceTenantUserBody) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *ReplaceTenantUserBody) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetCreatedFromUrlId

`func (o *ReplaceTenantUserBody) GetCreatedFromUrlId() string`

GetCreatedFromUrlId returns the CreatedFromUrlId field if non-nil, zero value otherwise.

### GetCreatedFromUrlIdOk

`func (o *ReplaceTenantUserBody) GetCreatedFromUrlIdOk() (*string, bool)`

GetCreatedFromUrlIdOk returns a tuple with the CreatedFromUrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFromUrlId

`func (o *ReplaceTenantUserBody) SetCreatedFromUrlId(v string)`

SetCreatedFromUrlId sets CreatedFromUrlId field to given value.

### HasCreatedFromUrlId

`func (o *ReplaceTenantUserBody) HasCreatedFromUrlId() bool`

HasCreatedFromUrlId returns a boolean if a field has been set.

### GetCreatedFromTenantId

`func (o *ReplaceTenantUserBody) GetCreatedFromTenantId() string`

GetCreatedFromTenantId returns the CreatedFromTenantId field if non-nil, zero value otherwise.

### GetCreatedFromTenantIdOk

`func (o *ReplaceTenantUserBody) GetCreatedFromTenantIdOk() (*string, bool)`

GetCreatedFromTenantIdOk returns a tuple with the CreatedFromTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFromTenantId

`func (o *ReplaceTenantUserBody) SetCreatedFromTenantId(v string)`

SetCreatedFromTenantId sets CreatedFromTenantId field to given value.

### HasCreatedFromTenantId

`func (o *ReplaceTenantUserBody) HasCreatedFromTenantId() bool`

HasCreatedFromTenantId returns a boolean if a field has been set.

### GetLastLoginDate

`func (o *ReplaceTenantUserBody) GetLastLoginDate() float64`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *ReplaceTenantUserBody) GetLastLoginDateOk() (*float64, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *ReplaceTenantUserBody) SetLastLoginDate(v float64)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *ReplaceTenantUserBody) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### GetKarma

`func (o *ReplaceTenantUserBody) GetKarma() float64`

GetKarma returns the Karma field if non-nil, zero value otherwise.

### GetKarmaOk

`func (o *ReplaceTenantUserBody) GetKarmaOk() (*float64, bool)`

GetKarmaOk returns a tuple with the Karma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKarma

`func (o *ReplaceTenantUserBody) SetKarma(v float64)`

SetKarma sets Karma field to given value.

### HasKarma

`func (o *ReplaceTenantUserBody) HasKarma() bool`

HasKarma returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


