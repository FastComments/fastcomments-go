# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Username** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**WebsiteUrl** | Pointer to **NullableString** |  | [optional] 
**Email** | **NullableString** |  | 
**PendingEmail** | Pointer to **string** |  | [optional] 
**SignUpDate** | **int64** |  | 
**CreatedFromUrlId** | Pointer to **NullableString** |  | [optional] 
**CreatedFromTenantId** | **NullableString** |  | 
**CreatedFromIpHashed** | **string** |  | 
**Verified** | **bool** |  | 
**LoginId** | **string** |  | 
**LoginIdDate** | **int64** |  | 
**LoginCount** | Pointer to **int32** |  | [optional] 
**OptedInNotifications** | Pointer to **bool** |  | [optional] 
**OptedInTenantNotifications** | Pointer to **bool** |  | [optional] 
**HideAccountCode** | Pointer to **bool** |  | [optional] 
**AvatarSrc** | Pointer to **NullableString** |  | [optional] 
**IsFastCommentsHelpRequestAdmin** | Pointer to **bool** |  | [optional] 
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
**IsImpersonator** | Pointer to **bool** |  | [optional] 
**IsCouponManager** | Pointer to **bool** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**DigestEmailFrequency** | Pointer to [**DigestEmailFrequency**](DigestEmailFrequency.md) |  | [optional] 
**IgnoredAddToMySiteMessages** | Pointer to **bool** |  | [optional] 
**LastLoginDate** | Pointer to **time.Time** |  | [optional] 
**DisplayLabel** | Pointer to **string** |  | [optional] 
**IsProfileActivityPrivate** | Pointer to **bool** |  | [optional] 
**IsProfileCommentsPrivate** | Pointer to **bool** |  | [optional] 
**IsProfileDMDisabled** | Pointer to **bool** |  | [optional] 
**ProfileCommentApprovalMode** | Pointer to **float64** |  | [optional] 
**Karma** | Pointer to **float64** |  | [optional] 
**PasswordHash** | Pointer to **string** |  | [optional] 
**AverageTicketAckTimeMS** | Pointer to **NullableFloat64** |  | [optional] 
**HasBlockedUsers** | Pointer to **bool** |  | [optional] 
**Bio** | Pointer to **string** |  | [optional] 
**HeaderBackgroundSrc** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**CountryFlag** | Pointer to **string** |  | [optional] 
**SocialLinks** | Pointer to **[]string** |  | [optional] 
**HasTwoFactor** | Pointer to **bool** |  | [optional] 

## Methods

### NewUser

`func NewUser(id string, username string, email NullableString, signUpDate int64, createdFromTenantId NullableString, createdFromIpHashed string, verified bool, loginId string, loginIdDate int64, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *User) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *User) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *User) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *User) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *User) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *User) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *User) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *User) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *User) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *User) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *User) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *User) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *User) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *User) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *User) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPendingEmail

`func (o *User) GetPendingEmail() string`

GetPendingEmail returns the PendingEmail field if non-nil, zero value otherwise.

### GetPendingEmailOk

`func (o *User) GetPendingEmailOk() (*string, bool)`

GetPendingEmailOk returns a tuple with the PendingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingEmail

`func (o *User) SetPendingEmail(v string)`

SetPendingEmail sets PendingEmail field to given value.

### HasPendingEmail

`func (o *User) HasPendingEmail() bool`

HasPendingEmail returns a boolean if a field has been set.

### GetSignUpDate

`func (o *User) GetSignUpDate() int64`

GetSignUpDate returns the SignUpDate field if non-nil, zero value otherwise.

### GetSignUpDateOk

`func (o *User) GetSignUpDateOk() (*int64, bool)`

GetSignUpDateOk returns a tuple with the SignUpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUpDate

`func (o *User) SetSignUpDate(v int64)`

SetSignUpDate sets SignUpDate field to given value.


### GetCreatedFromUrlId

`func (o *User) GetCreatedFromUrlId() string`

GetCreatedFromUrlId returns the CreatedFromUrlId field if non-nil, zero value otherwise.

### GetCreatedFromUrlIdOk

`func (o *User) GetCreatedFromUrlIdOk() (*string, bool)`

GetCreatedFromUrlIdOk returns a tuple with the CreatedFromUrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFromUrlId

`func (o *User) SetCreatedFromUrlId(v string)`

SetCreatedFromUrlId sets CreatedFromUrlId field to given value.

### HasCreatedFromUrlId

`func (o *User) HasCreatedFromUrlId() bool`

HasCreatedFromUrlId returns a boolean if a field has been set.

### SetCreatedFromUrlIdNil

`func (o *User) SetCreatedFromUrlIdNil(b bool)`

 SetCreatedFromUrlIdNil sets the value for CreatedFromUrlId to be an explicit nil

### UnsetCreatedFromUrlId
`func (o *User) UnsetCreatedFromUrlId()`

UnsetCreatedFromUrlId ensures that no value is present for CreatedFromUrlId, not even an explicit nil
### GetCreatedFromTenantId

`func (o *User) GetCreatedFromTenantId() string`

GetCreatedFromTenantId returns the CreatedFromTenantId field if non-nil, zero value otherwise.

### GetCreatedFromTenantIdOk

`func (o *User) GetCreatedFromTenantIdOk() (*string, bool)`

GetCreatedFromTenantIdOk returns a tuple with the CreatedFromTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFromTenantId

`func (o *User) SetCreatedFromTenantId(v string)`

SetCreatedFromTenantId sets CreatedFromTenantId field to given value.


### SetCreatedFromTenantIdNil

`func (o *User) SetCreatedFromTenantIdNil(b bool)`

 SetCreatedFromTenantIdNil sets the value for CreatedFromTenantId to be an explicit nil

### UnsetCreatedFromTenantId
`func (o *User) UnsetCreatedFromTenantId()`

UnsetCreatedFromTenantId ensures that no value is present for CreatedFromTenantId, not even an explicit nil
### GetCreatedFromIpHashed

`func (o *User) GetCreatedFromIpHashed() string`

GetCreatedFromIpHashed returns the CreatedFromIpHashed field if non-nil, zero value otherwise.

### GetCreatedFromIpHashedOk

`func (o *User) GetCreatedFromIpHashedOk() (*string, bool)`

GetCreatedFromIpHashedOk returns a tuple with the CreatedFromIpHashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFromIpHashed

`func (o *User) SetCreatedFromIpHashed(v string)`

SetCreatedFromIpHashed sets CreatedFromIpHashed field to given value.


### GetVerified

`func (o *User) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *User) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *User) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetLoginId

`func (o *User) GetLoginId() string`

GetLoginId returns the LoginId field if non-nil, zero value otherwise.

### GetLoginIdOk

`func (o *User) GetLoginIdOk() (*string, bool)`

GetLoginIdOk returns a tuple with the LoginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginId

`func (o *User) SetLoginId(v string)`

SetLoginId sets LoginId field to given value.


### GetLoginIdDate

`func (o *User) GetLoginIdDate() int64`

GetLoginIdDate returns the LoginIdDate field if non-nil, zero value otherwise.

### GetLoginIdDateOk

`func (o *User) GetLoginIdDateOk() (*int64, bool)`

GetLoginIdDateOk returns a tuple with the LoginIdDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginIdDate

`func (o *User) SetLoginIdDate(v int64)`

SetLoginIdDate sets LoginIdDate field to given value.


### GetLoginCount

`func (o *User) GetLoginCount() int32`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *User) GetLoginCountOk() (*int32, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *User) SetLoginCount(v int32)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *User) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### GetOptedInNotifications

`func (o *User) GetOptedInNotifications() bool`

GetOptedInNotifications returns the OptedInNotifications field if non-nil, zero value otherwise.

### GetOptedInNotificationsOk

`func (o *User) GetOptedInNotificationsOk() (*bool, bool)`

GetOptedInNotificationsOk returns a tuple with the OptedInNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedInNotifications

`func (o *User) SetOptedInNotifications(v bool)`

SetOptedInNotifications sets OptedInNotifications field to given value.

### HasOptedInNotifications

`func (o *User) HasOptedInNotifications() bool`

HasOptedInNotifications returns a boolean if a field has been set.

### GetOptedInTenantNotifications

`func (o *User) GetOptedInTenantNotifications() bool`

GetOptedInTenantNotifications returns the OptedInTenantNotifications field if non-nil, zero value otherwise.

### GetOptedInTenantNotificationsOk

`func (o *User) GetOptedInTenantNotificationsOk() (*bool, bool)`

GetOptedInTenantNotificationsOk returns a tuple with the OptedInTenantNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedInTenantNotifications

`func (o *User) SetOptedInTenantNotifications(v bool)`

SetOptedInTenantNotifications sets OptedInTenantNotifications field to given value.

### HasOptedInTenantNotifications

`func (o *User) HasOptedInTenantNotifications() bool`

HasOptedInTenantNotifications returns a boolean if a field has been set.

### GetHideAccountCode

`func (o *User) GetHideAccountCode() bool`

GetHideAccountCode returns the HideAccountCode field if non-nil, zero value otherwise.

### GetHideAccountCodeOk

`func (o *User) GetHideAccountCodeOk() (*bool, bool)`

GetHideAccountCodeOk returns a tuple with the HideAccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAccountCode

`func (o *User) SetHideAccountCode(v bool)`

SetHideAccountCode sets HideAccountCode field to given value.

### HasHideAccountCode

`func (o *User) HasHideAccountCode() bool`

HasHideAccountCode returns a boolean if a field has been set.

### GetAvatarSrc

`func (o *User) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *User) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *User) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *User) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### SetAvatarSrcNil

`func (o *User) SetAvatarSrcNil(b bool)`

 SetAvatarSrcNil sets the value for AvatarSrc to be an explicit nil

### UnsetAvatarSrc
`func (o *User) UnsetAvatarSrc()`

UnsetAvatarSrc ensures that no value is present for AvatarSrc, not even an explicit nil
### GetIsFastCommentsHelpRequestAdmin

`func (o *User) GetIsFastCommentsHelpRequestAdmin() bool`

GetIsFastCommentsHelpRequestAdmin returns the IsFastCommentsHelpRequestAdmin field if non-nil, zero value otherwise.

### GetIsFastCommentsHelpRequestAdminOk

`func (o *User) GetIsFastCommentsHelpRequestAdminOk() (*bool, bool)`

GetIsFastCommentsHelpRequestAdminOk returns a tuple with the IsFastCommentsHelpRequestAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFastCommentsHelpRequestAdmin

`func (o *User) SetIsFastCommentsHelpRequestAdmin(v bool)`

SetIsFastCommentsHelpRequestAdmin sets IsFastCommentsHelpRequestAdmin field to given value.

### HasIsFastCommentsHelpRequestAdmin

`func (o *User) HasIsFastCommentsHelpRequestAdmin() bool`

HasIsFastCommentsHelpRequestAdmin returns a boolean if a field has been set.

### GetIsHelpRequestAdmin

`func (o *User) GetIsHelpRequestAdmin() bool`

GetIsHelpRequestAdmin returns the IsHelpRequestAdmin field if non-nil, zero value otherwise.

### GetIsHelpRequestAdminOk

`func (o *User) GetIsHelpRequestAdminOk() (*bool, bool)`

GetIsHelpRequestAdminOk returns a tuple with the IsHelpRequestAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHelpRequestAdmin

`func (o *User) SetIsHelpRequestAdmin(v bool)`

SetIsHelpRequestAdmin sets IsHelpRequestAdmin field to given value.

### HasIsHelpRequestAdmin

`func (o *User) HasIsHelpRequestAdmin() bool`

HasIsHelpRequestAdmin returns a boolean if a field has been set.

### GetIsAccountOwner

`func (o *User) GetIsAccountOwner() bool`

GetIsAccountOwner returns the IsAccountOwner field if non-nil, zero value otherwise.

### GetIsAccountOwnerOk

`func (o *User) GetIsAccountOwnerOk() (*bool, bool)`

GetIsAccountOwnerOk returns a tuple with the IsAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountOwner

`func (o *User) SetIsAccountOwner(v bool)`

SetIsAccountOwner sets IsAccountOwner field to given value.

### HasIsAccountOwner

`func (o *User) HasIsAccountOwner() bool`

HasIsAccountOwner returns a boolean if a field has been set.

### GetIsAdminAdmin

`func (o *User) GetIsAdminAdmin() bool`

GetIsAdminAdmin returns the IsAdminAdmin field if non-nil, zero value otherwise.

### GetIsAdminAdminOk

`func (o *User) GetIsAdminAdminOk() (*bool, bool)`

GetIsAdminAdminOk returns a tuple with the IsAdminAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdminAdmin

`func (o *User) SetIsAdminAdmin(v bool)`

SetIsAdminAdmin sets IsAdminAdmin field to given value.

### HasIsAdminAdmin

`func (o *User) HasIsAdminAdmin() bool`

HasIsAdminAdmin returns a boolean if a field has been set.

### GetIsBillingAdmin

`func (o *User) GetIsBillingAdmin() bool`

GetIsBillingAdmin returns the IsBillingAdmin field if non-nil, zero value otherwise.

### GetIsBillingAdminOk

`func (o *User) GetIsBillingAdminOk() (*bool, bool)`

GetIsBillingAdminOk returns a tuple with the IsBillingAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillingAdmin

`func (o *User) SetIsBillingAdmin(v bool)`

SetIsBillingAdmin sets IsBillingAdmin field to given value.

### HasIsBillingAdmin

`func (o *User) HasIsBillingAdmin() bool`

HasIsBillingAdmin returns a boolean if a field has been set.

### GetIsAnalyticsAdmin

`func (o *User) GetIsAnalyticsAdmin() bool`

GetIsAnalyticsAdmin returns the IsAnalyticsAdmin field if non-nil, zero value otherwise.

### GetIsAnalyticsAdminOk

`func (o *User) GetIsAnalyticsAdminOk() (*bool, bool)`

GetIsAnalyticsAdminOk returns a tuple with the IsAnalyticsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnalyticsAdmin

`func (o *User) SetIsAnalyticsAdmin(v bool)`

SetIsAnalyticsAdmin sets IsAnalyticsAdmin field to given value.

### HasIsAnalyticsAdmin

`func (o *User) HasIsAnalyticsAdmin() bool`

HasIsAnalyticsAdmin returns a boolean if a field has been set.

### GetIsCustomizationAdmin

`func (o *User) GetIsCustomizationAdmin() bool`

GetIsCustomizationAdmin returns the IsCustomizationAdmin field if non-nil, zero value otherwise.

### GetIsCustomizationAdminOk

`func (o *User) GetIsCustomizationAdminOk() (*bool, bool)`

GetIsCustomizationAdminOk returns a tuple with the IsCustomizationAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomizationAdmin

`func (o *User) SetIsCustomizationAdmin(v bool)`

SetIsCustomizationAdmin sets IsCustomizationAdmin field to given value.

### HasIsCustomizationAdmin

`func (o *User) HasIsCustomizationAdmin() bool`

HasIsCustomizationAdmin returns a boolean if a field has been set.

### GetIsManageDataAdmin

`func (o *User) GetIsManageDataAdmin() bool`

GetIsManageDataAdmin returns the IsManageDataAdmin field if non-nil, zero value otherwise.

### GetIsManageDataAdminOk

`func (o *User) GetIsManageDataAdminOk() (*bool, bool)`

GetIsManageDataAdminOk returns a tuple with the IsManageDataAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManageDataAdmin

`func (o *User) SetIsManageDataAdmin(v bool)`

SetIsManageDataAdmin sets IsManageDataAdmin field to given value.

### HasIsManageDataAdmin

`func (o *User) HasIsManageDataAdmin() bool`

HasIsManageDataAdmin returns a boolean if a field has been set.

### GetIsCommentModeratorAdmin

`func (o *User) GetIsCommentModeratorAdmin() bool`

GetIsCommentModeratorAdmin returns the IsCommentModeratorAdmin field if non-nil, zero value otherwise.

### GetIsCommentModeratorAdminOk

`func (o *User) GetIsCommentModeratorAdminOk() (*bool, bool)`

GetIsCommentModeratorAdminOk returns a tuple with the IsCommentModeratorAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCommentModeratorAdmin

`func (o *User) SetIsCommentModeratorAdmin(v bool)`

SetIsCommentModeratorAdmin sets IsCommentModeratorAdmin field to given value.

### HasIsCommentModeratorAdmin

`func (o *User) HasIsCommentModeratorAdmin() bool`

HasIsCommentModeratorAdmin returns a boolean if a field has been set.

### GetIsAPIAdmin

`func (o *User) GetIsAPIAdmin() bool`

GetIsAPIAdmin returns the IsAPIAdmin field if non-nil, zero value otherwise.

### GetIsAPIAdminOk

`func (o *User) GetIsAPIAdminOk() (*bool, bool)`

GetIsAPIAdminOk returns a tuple with the IsAPIAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAPIAdmin

`func (o *User) SetIsAPIAdmin(v bool)`

SetIsAPIAdmin sets IsAPIAdmin field to given value.

### HasIsAPIAdmin

`func (o *User) HasIsAPIAdmin() bool`

HasIsAPIAdmin returns a boolean if a field has been set.

### GetModeratorIds

`func (o *User) GetModeratorIds() []string`

GetModeratorIds returns the ModeratorIds field if non-nil, zero value otherwise.

### GetModeratorIdsOk

`func (o *User) GetModeratorIdsOk() (*[]string, bool)`

GetModeratorIdsOk returns a tuple with the ModeratorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeratorIds

`func (o *User) SetModeratorIds(v []string)`

SetModeratorIds sets ModeratorIds field to given value.

### HasModeratorIds

`func (o *User) HasModeratorIds() bool`

HasModeratorIds returns a boolean if a field has been set.

### GetIsImpersonator

`func (o *User) GetIsImpersonator() bool`

GetIsImpersonator returns the IsImpersonator field if non-nil, zero value otherwise.

### GetIsImpersonatorOk

`func (o *User) GetIsImpersonatorOk() (*bool, bool)`

GetIsImpersonatorOk returns a tuple with the IsImpersonator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImpersonator

`func (o *User) SetIsImpersonator(v bool)`

SetIsImpersonator sets IsImpersonator field to given value.

### HasIsImpersonator

`func (o *User) HasIsImpersonator() bool`

HasIsImpersonator returns a boolean if a field has been set.

### GetIsCouponManager

`func (o *User) GetIsCouponManager() bool`

GetIsCouponManager returns the IsCouponManager field if non-nil, zero value otherwise.

### GetIsCouponManagerOk

`func (o *User) GetIsCouponManagerOk() (*bool, bool)`

GetIsCouponManagerOk returns a tuple with the IsCouponManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCouponManager

`func (o *User) SetIsCouponManager(v bool)`

SetIsCouponManager sets IsCouponManager field to given value.

### HasIsCouponManager

`func (o *User) HasIsCouponManager() bool`

HasIsCouponManager returns a boolean if a field has been set.

### GetLocale

`func (o *User) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *User) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *User) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *User) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetDigestEmailFrequency

`func (o *User) GetDigestEmailFrequency() DigestEmailFrequency`

GetDigestEmailFrequency returns the DigestEmailFrequency field if non-nil, zero value otherwise.

### GetDigestEmailFrequencyOk

`func (o *User) GetDigestEmailFrequencyOk() (*DigestEmailFrequency, bool)`

GetDigestEmailFrequencyOk returns a tuple with the DigestEmailFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestEmailFrequency

`func (o *User) SetDigestEmailFrequency(v DigestEmailFrequency)`

SetDigestEmailFrequency sets DigestEmailFrequency field to given value.

### HasDigestEmailFrequency

`func (o *User) HasDigestEmailFrequency() bool`

HasDigestEmailFrequency returns a boolean if a field has been set.

### GetIgnoredAddToMySiteMessages

`func (o *User) GetIgnoredAddToMySiteMessages() bool`

GetIgnoredAddToMySiteMessages returns the IgnoredAddToMySiteMessages field if non-nil, zero value otherwise.

### GetIgnoredAddToMySiteMessagesOk

`func (o *User) GetIgnoredAddToMySiteMessagesOk() (*bool, bool)`

GetIgnoredAddToMySiteMessagesOk returns a tuple with the IgnoredAddToMySiteMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredAddToMySiteMessages

`func (o *User) SetIgnoredAddToMySiteMessages(v bool)`

SetIgnoredAddToMySiteMessages sets IgnoredAddToMySiteMessages field to given value.

### HasIgnoredAddToMySiteMessages

`func (o *User) HasIgnoredAddToMySiteMessages() bool`

HasIgnoredAddToMySiteMessages returns a boolean if a field has been set.

### GetLastLoginDate

`func (o *User) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *User) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *User) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *User) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *User) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *User) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *User) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *User) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetIsProfileActivityPrivate

`func (o *User) GetIsProfileActivityPrivate() bool`

GetIsProfileActivityPrivate returns the IsProfileActivityPrivate field if non-nil, zero value otherwise.

### GetIsProfileActivityPrivateOk

`func (o *User) GetIsProfileActivityPrivateOk() (*bool, bool)`

GetIsProfileActivityPrivateOk returns a tuple with the IsProfileActivityPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProfileActivityPrivate

`func (o *User) SetIsProfileActivityPrivate(v bool)`

SetIsProfileActivityPrivate sets IsProfileActivityPrivate field to given value.

### HasIsProfileActivityPrivate

`func (o *User) HasIsProfileActivityPrivate() bool`

HasIsProfileActivityPrivate returns a boolean if a field has been set.

### GetIsProfileCommentsPrivate

`func (o *User) GetIsProfileCommentsPrivate() bool`

GetIsProfileCommentsPrivate returns the IsProfileCommentsPrivate field if non-nil, zero value otherwise.

### GetIsProfileCommentsPrivateOk

`func (o *User) GetIsProfileCommentsPrivateOk() (*bool, bool)`

GetIsProfileCommentsPrivateOk returns a tuple with the IsProfileCommentsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProfileCommentsPrivate

`func (o *User) SetIsProfileCommentsPrivate(v bool)`

SetIsProfileCommentsPrivate sets IsProfileCommentsPrivate field to given value.

### HasIsProfileCommentsPrivate

`func (o *User) HasIsProfileCommentsPrivate() bool`

HasIsProfileCommentsPrivate returns a boolean if a field has been set.

### GetIsProfileDMDisabled

`func (o *User) GetIsProfileDMDisabled() bool`

GetIsProfileDMDisabled returns the IsProfileDMDisabled field if non-nil, zero value otherwise.

### GetIsProfileDMDisabledOk

`func (o *User) GetIsProfileDMDisabledOk() (*bool, bool)`

GetIsProfileDMDisabledOk returns a tuple with the IsProfileDMDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProfileDMDisabled

`func (o *User) SetIsProfileDMDisabled(v bool)`

SetIsProfileDMDisabled sets IsProfileDMDisabled field to given value.

### HasIsProfileDMDisabled

`func (o *User) HasIsProfileDMDisabled() bool`

HasIsProfileDMDisabled returns a boolean if a field has been set.

### GetProfileCommentApprovalMode

`func (o *User) GetProfileCommentApprovalMode() float64`

GetProfileCommentApprovalMode returns the ProfileCommentApprovalMode field if non-nil, zero value otherwise.

### GetProfileCommentApprovalModeOk

`func (o *User) GetProfileCommentApprovalModeOk() (*float64, bool)`

GetProfileCommentApprovalModeOk returns a tuple with the ProfileCommentApprovalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileCommentApprovalMode

`func (o *User) SetProfileCommentApprovalMode(v float64)`

SetProfileCommentApprovalMode sets ProfileCommentApprovalMode field to given value.

### HasProfileCommentApprovalMode

`func (o *User) HasProfileCommentApprovalMode() bool`

HasProfileCommentApprovalMode returns a boolean if a field has been set.

### GetKarma

`func (o *User) GetKarma() float64`

GetKarma returns the Karma field if non-nil, zero value otherwise.

### GetKarmaOk

`func (o *User) GetKarmaOk() (*float64, bool)`

GetKarmaOk returns a tuple with the Karma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKarma

`func (o *User) SetKarma(v float64)`

SetKarma sets Karma field to given value.

### HasKarma

`func (o *User) HasKarma() bool`

HasKarma returns a boolean if a field has been set.

### GetPasswordHash

`func (o *User) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *User) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *User) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *User) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetAverageTicketAckTimeMS

`func (o *User) GetAverageTicketAckTimeMS() float64`

GetAverageTicketAckTimeMS returns the AverageTicketAckTimeMS field if non-nil, zero value otherwise.

### GetAverageTicketAckTimeMSOk

`func (o *User) GetAverageTicketAckTimeMSOk() (*float64, bool)`

GetAverageTicketAckTimeMSOk returns a tuple with the AverageTicketAckTimeMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTicketAckTimeMS

`func (o *User) SetAverageTicketAckTimeMS(v float64)`

SetAverageTicketAckTimeMS sets AverageTicketAckTimeMS field to given value.

### HasAverageTicketAckTimeMS

`func (o *User) HasAverageTicketAckTimeMS() bool`

HasAverageTicketAckTimeMS returns a boolean if a field has been set.

### SetAverageTicketAckTimeMSNil

`func (o *User) SetAverageTicketAckTimeMSNil(b bool)`

 SetAverageTicketAckTimeMSNil sets the value for AverageTicketAckTimeMS to be an explicit nil

### UnsetAverageTicketAckTimeMS
`func (o *User) UnsetAverageTicketAckTimeMS()`

UnsetAverageTicketAckTimeMS ensures that no value is present for AverageTicketAckTimeMS, not even an explicit nil
### GetHasBlockedUsers

`func (o *User) GetHasBlockedUsers() bool`

GetHasBlockedUsers returns the HasBlockedUsers field if non-nil, zero value otherwise.

### GetHasBlockedUsersOk

`func (o *User) GetHasBlockedUsersOk() (*bool, bool)`

GetHasBlockedUsersOk returns a tuple with the HasBlockedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBlockedUsers

`func (o *User) SetHasBlockedUsers(v bool)`

SetHasBlockedUsers sets HasBlockedUsers field to given value.

### HasHasBlockedUsers

`func (o *User) HasHasBlockedUsers() bool`

HasHasBlockedUsers returns a boolean if a field has been set.

### GetBio

`func (o *User) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *User) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *User) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *User) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetHeaderBackgroundSrc

`func (o *User) GetHeaderBackgroundSrc() string`

GetHeaderBackgroundSrc returns the HeaderBackgroundSrc field if non-nil, zero value otherwise.

### GetHeaderBackgroundSrcOk

`func (o *User) GetHeaderBackgroundSrcOk() (*string, bool)`

GetHeaderBackgroundSrcOk returns a tuple with the HeaderBackgroundSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderBackgroundSrc

`func (o *User) SetHeaderBackgroundSrc(v string)`

SetHeaderBackgroundSrc sets HeaderBackgroundSrc field to given value.

### HasHeaderBackgroundSrc

`func (o *User) HasHeaderBackgroundSrc() bool`

HasHeaderBackgroundSrc returns a boolean if a field has been set.

### GetCountryCode

`func (o *User) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *User) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *User) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *User) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryFlag

`func (o *User) GetCountryFlag() string`

GetCountryFlag returns the CountryFlag field if non-nil, zero value otherwise.

### GetCountryFlagOk

`func (o *User) GetCountryFlagOk() (*string, bool)`

GetCountryFlagOk returns a tuple with the CountryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryFlag

`func (o *User) SetCountryFlag(v string)`

SetCountryFlag sets CountryFlag field to given value.

### HasCountryFlag

`func (o *User) HasCountryFlag() bool`

HasCountryFlag returns a boolean if a field has been set.

### GetSocialLinks

`func (o *User) GetSocialLinks() []string`

GetSocialLinks returns the SocialLinks field if non-nil, zero value otherwise.

### GetSocialLinksOk

`func (o *User) GetSocialLinksOk() (*[]string, bool)`

GetSocialLinksOk returns a tuple with the SocialLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialLinks

`func (o *User) SetSocialLinks(v []string)`

SetSocialLinks sets SocialLinks field to given value.

### HasSocialLinks

`func (o *User) HasSocialLinks() bool`

HasSocialLinks returns a boolean if a field has been set.

### GetHasTwoFactor

`func (o *User) GetHasTwoFactor() bool`

GetHasTwoFactor returns the HasTwoFactor field if non-nil, zero value otherwise.

### GetHasTwoFactorOk

`func (o *User) GetHasTwoFactorOk() (*bool, bool)`

GetHasTwoFactorOk returns a tuple with the HasTwoFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTwoFactor

`func (o *User) SetHasTwoFactor(v bool)`

SetHasTwoFactor sets HasTwoFactor field to given value.

### HasHasTwoFactor

`func (o *User) HasHasTwoFactor() bool`

HasHasTwoFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


