# CreateTenantUserBody

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

## Methods

### NewCreateTenantUserBody

`func NewCreateTenantUserBody(username string, email string, ) *CreateTenantUserBody`

NewCreateTenantUserBody instantiates a new CreateTenantUserBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenantUserBodyWithDefaults

`func NewCreateTenantUserBodyWithDefaults() *CreateTenantUserBody`

NewCreateTenantUserBodyWithDefaults instantiates a new CreateTenantUserBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateTenantUserBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateTenantUserBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateTenantUserBody) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *CreateTenantUserBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateTenantUserBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateTenantUserBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDisplayName

`func (o *CreateTenantUserBody) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateTenantUserBody) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateTenantUserBody) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateTenantUserBody) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *CreateTenantUserBody) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *CreateTenantUserBody) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *CreateTenantUserBody) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *CreateTenantUserBody) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetSignUpDate

`func (o *CreateTenantUserBody) GetSignUpDate() float64`

GetSignUpDate returns the SignUpDate field if non-nil, zero value otherwise.

### GetSignUpDateOk

`func (o *CreateTenantUserBody) GetSignUpDateOk() (*float64, bool)`

GetSignUpDateOk returns a tuple with the SignUpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUpDate

`func (o *CreateTenantUserBody) SetSignUpDate(v float64)`

SetSignUpDate sets SignUpDate field to given value.

### HasSignUpDate

`func (o *CreateTenantUserBody) HasSignUpDate() bool`

HasSignUpDate returns a boolean if a field has been set.

### GetLocale

`func (o *CreateTenantUserBody) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CreateTenantUserBody) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CreateTenantUserBody) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CreateTenantUserBody) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetVerified

`func (o *CreateTenantUserBody) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *CreateTenantUserBody) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *CreateTenantUserBody) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *CreateTenantUserBody) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetLoginCount

`func (o *CreateTenantUserBody) GetLoginCount() float64`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *CreateTenantUserBody) GetLoginCountOk() (*float64, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *CreateTenantUserBody) SetLoginCount(v float64)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *CreateTenantUserBody) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### GetOptedInNotifications

`func (o *CreateTenantUserBody) GetOptedInNotifications() bool`

GetOptedInNotifications returns the OptedInNotifications field if non-nil, zero value otherwise.

### GetOptedInNotificationsOk

`func (o *CreateTenantUserBody) GetOptedInNotificationsOk() (*bool, bool)`

GetOptedInNotificationsOk returns a tuple with the OptedInNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedInNotifications

`func (o *CreateTenantUserBody) SetOptedInNotifications(v bool)`

SetOptedInNotifications sets OptedInNotifications field to given value.

### HasOptedInNotifications

`func (o *CreateTenantUserBody) HasOptedInNotifications() bool`

HasOptedInNotifications returns a boolean if a field has been set.

### GetOptedInTenantNotifications

`func (o *CreateTenantUserBody) GetOptedInTenantNotifications() bool`

GetOptedInTenantNotifications returns the OptedInTenantNotifications field if non-nil, zero value otherwise.

### GetOptedInTenantNotificationsOk

`func (o *CreateTenantUserBody) GetOptedInTenantNotificationsOk() (*bool, bool)`

GetOptedInTenantNotificationsOk returns a tuple with the OptedInTenantNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedInTenantNotifications

`func (o *CreateTenantUserBody) SetOptedInTenantNotifications(v bool)`

SetOptedInTenantNotifications sets OptedInTenantNotifications field to given value.

### HasOptedInTenantNotifications

`func (o *CreateTenantUserBody) HasOptedInTenantNotifications() bool`

HasOptedInTenantNotifications returns a boolean if a field has been set.

### GetHideAccountCode

`func (o *CreateTenantUserBody) GetHideAccountCode() bool`

GetHideAccountCode returns the HideAccountCode field if non-nil, zero value otherwise.

### GetHideAccountCodeOk

`func (o *CreateTenantUserBody) GetHideAccountCodeOk() (*bool, bool)`

GetHideAccountCodeOk returns a tuple with the HideAccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAccountCode

`func (o *CreateTenantUserBody) SetHideAccountCode(v bool)`

SetHideAccountCode sets HideAccountCode field to given value.

### HasHideAccountCode

`func (o *CreateTenantUserBody) HasHideAccountCode() bool`

HasHideAccountCode returns a boolean if a field has been set.

### GetAvatarSrc

`func (o *CreateTenantUserBody) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *CreateTenantUserBody) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *CreateTenantUserBody) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *CreateTenantUserBody) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### GetIsHelpRequestAdmin

`func (o *CreateTenantUserBody) GetIsHelpRequestAdmin() bool`

GetIsHelpRequestAdmin returns the IsHelpRequestAdmin field if non-nil, zero value otherwise.

### GetIsHelpRequestAdminOk

`func (o *CreateTenantUserBody) GetIsHelpRequestAdminOk() (*bool, bool)`

GetIsHelpRequestAdminOk returns a tuple with the IsHelpRequestAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHelpRequestAdmin

`func (o *CreateTenantUserBody) SetIsHelpRequestAdmin(v bool)`

SetIsHelpRequestAdmin sets IsHelpRequestAdmin field to given value.

### HasIsHelpRequestAdmin

`func (o *CreateTenantUserBody) HasIsHelpRequestAdmin() bool`

HasIsHelpRequestAdmin returns a boolean if a field has been set.

### GetIsAccountOwner

`func (o *CreateTenantUserBody) GetIsAccountOwner() bool`

GetIsAccountOwner returns the IsAccountOwner field if non-nil, zero value otherwise.

### GetIsAccountOwnerOk

`func (o *CreateTenantUserBody) GetIsAccountOwnerOk() (*bool, bool)`

GetIsAccountOwnerOk returns a tuple with the IsAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountOwner

`func (o *CreateTenantUserBody) SetIsAccountOwner(v bool)`

SetIsAccountOwner sets IsAccountOwner field to given value.

### HasIsAccountOwner

`func (o *CreateTenantUserBody) HasIsAccountOwner() bool`

HasIsAccountOwner returns a boolean if a field has been set.

### GetIsAdminAdmin

`func (o *CreateTenantUserBody) GetIsAdminAdmin() bool`

GetIsAdminAdmin returns the IsAdminAdmin field if non-nil, zero value otherwise.

### GetIsAdminAdminOk

`func (o *CreateTenantUserBody) GetIsAdminAdminOk() (*bool, bool)`

GetIsAdminAdminOk returns a tuple with the IsAdminAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdminAdmin

`func (o *CreateTenantUserBody) SetIsAdminAdmin(v bool)`

SetIsAdminAdmin sets IsAdminAdmin field to given value.

### HasIsAdminAdmin

`func (o *CreateTenantUserBody) HasIsAdminAdmin() bool`

HasIsAdminAdmin returns a boolean if a field has been set.

### GetIsBillingAdmin

`func (o *CreateTenantUserBody) GetIsBillingAdmin() bool`

GetIsBillingAdmin returns the IsBillingAdmin field if non-nil, zero value otherwise.

### GetIsBillingAdminOk

`func (o *CreateTenantUserBody) GetIsBillingAdminOk() (*bool, bool)`

GetIsBillingAdminOk returns a tuple with the IsBillingAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillingAdmin

`func (o *CreateTenantUserBody) SetIsBillingAdmin(v bool)`

SetIsBillingAdmin sets IsBillingAdmin field to given value.

### HasIsBillingAdmin

`func (o *CreateTenantUserBody) HasIsBillingAdmin() bool`

HasIsBillingAdmin returns a boolean if a field has been set.

### GetIsAnalyticsAdmin

`func (o *CreateTenantUserBody) GetIsAnalyticsAdmin() bool`

GetIsAnalyticsAdmin returns the IsAnalyticsAdmin field if non-nil, zero value otherwise.

### GetIsAnalyticsAdminOk

`func (o *CreateTenantUserBody) GetIsAnalyticsAdminOk() (*bool, bool)`

GetIsAnalyticsAdminOk returns a tuple with the IsAnalyticsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnalyticsAdmin

`func (o *CreateTenantUserBody) SetIsAnalyticsAdmin(v bool)`

SetIsAnalyticsAdmin sets IsAnalyticsAdmin field to given value.

### HasIsAnalyticsAdmin

`func (o *CreateTenantUserBody) HasIsAnalyticsAdmin() bool`

HasIsAnalyticsAdmin returns a boolean if a field has been set.

### GetIsCustomizationAdmin

`func (o *CreateTenantUserBody) GetIsCustomizationAdmin() bool`

GetIsCustomizationAdmin returns the IsCustomizationAdmin field if non-nil, zero value otherwise.

### GetIsCustomizationAdminOk

`func (o *CreateTenantUserBody) GetIsCustomizationAdminOk() (*bool, bool)`

GetIsCustomizationAdminOk returns a tuple with the IsCustomizationAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomizationAdmin

`func (o *CreateTenantUserBody) SetIsCustomizationAdmin(v bool)`

SetIsCustomizationAdmin sets IsCustomizationAdmin field to given value.

### HasIsCustomizationAdmin

`func (o *CreateTenantUserBody) HasIsCustomizationAdmin() bool`

HasIsCustomizationAdmin returns a boolean if a field has been set.

### GetIsManageDataAdmin

`func (o *CreateTenantUserBody) GetIsManageDataAdmin() bool`

GetIsManageDataAdmin returns the IsManageDataAdmin field if non-nil, zero value otherwise.

### GetIsManageDataAdminOk

`func (o *CreateTenantUserBody) GetIsManageDataAdminOk() (*bool, bool)`

GetIsManageDataAdminOk returns a tuple with the IsManageDataAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManageDataAdmin

`func (o *CreateTenantUserBody) SetIsManageDataAdmin(v bool)`

SetIsManageDataAdmin sets IsManageDataAdmin field to given value.

### HasIsManageDataAdmin

`func (o *CreateTenantUserBody) HasIsManageDataAdmin() bool`

HasIsManageDataAdmin returns a boolean if a field has been set.

### GetIsCommentModeratorAdmin

`func (o *CreateTenantUserBody) GetIsCommentModeratorAdmin() bool`

GetIsCommentModeratorAdmin returns the IsCommentModeratorAdmin field if non-nil, zero value otherwise.

### GetIsCommentModeratorAdminOk

`func (o *CreateTenantUserBody) GetIsCommentModeratorAdminOk() (*bool, bool)`

GetIsCommentModeratorAdminOk returns a tuple with the IsCommentModeratorAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCommentModeratorAdmin

`func (o *CreateTenantUserBody) SetIsCommentModeratorAdmin(v bool)`

SetIsCommentModeratorAdmin sets IsCommentModeratorAdmin field to given value.

### HasIsCommentModeratorAdmin

`func (o *CreateTenantUserBody) HasIsCommentModeratorAdmin() bool`

HasIsCommentModeratorAdmin returns a boolean if a field has been set.

### GetIsAPIAdmin

`func (o *CreateTenantUserBody) GetIsAPIAdmin() bool`

GetIsAPIAdmin returns the IsAPIAdmin field if non-nil, zero value otherwise.

### GetIsAPIAdminOk

`func (o *CreateTenantUserBody) GetIsAPIAdminOk() (*bool, bool)`

GetIsAPIAdminOk returns a tuple with the IsAPIAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAPIAdmin

`func (o *CreateTenantUserBody) SetIsAPIAdmin(v bool)`

SetIsAPIAdmin sets IsAPIAdmin field to given value.

### HasIsAPIAdmin

`func (o *CreateTenantUserBody) HasIsAPIAdmin() bool`

HasIsAPIAdmin returns a boolean if a field has been set.

### GetModeratorIds

`func (o *CreateTenantUserBody) GetModeratorIds() []string`

GetModeratorIds returns the ModeratorIds field if non-nil, zero value otherwise.

### GetModeratorIdsOk

`func (o *CreateTenantUserBody) GetModeratorIdsOk() (*[]string, bool)`

GetModeratorIdsOk returns a tuple with the ModeratorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeratorIds

`func (o *CreateTenantUserBody) SetModeratorIds(v []string)`

SetModeratorIds sets ModeratorIds field to given value.

### HasModeratorIds

`func (o *CreateTenantUserBody) HasModeratorIds() bool`

HasModeratorIds returns a boolean if a field has been set.

### GetDigestEmailFrequency

`func (o *CreateTenantUserBody) GetDigestEmailFrequency() float64`

GetDigestEmailFrequency returns the DigestEmailFrequency field if non-nil, zero value otherwise.

### GetDigestEmailFrequencyOk

`func (o *CreateTenantUserBody) GetDigestEmailFrequencyOk() (*float64, bool)`

GetDigestEmailFrequencyOk returns a tuple with the DigestEmailFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestEmailFrequency

`func (o *CreateTenantUserBody) SetDigestEmailFrequency(v float64)`

SetDigestEmailFrequency sets DigestEmailFrequency field to given value.

### HasDigestEmailFrequency

`func (o *CreateTenantUserBody) HasDigestEmailFrequency() bool`

HasDigestEmailFrequency returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *CreateTenantUserBody) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *CreateTenantUserBody) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *CreateTenantUserBody) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *CreateTenantUserBody) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


