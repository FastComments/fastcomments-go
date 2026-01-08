# UpdateTenantUserBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**WebsiteUrl** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**SignUpDate** | Pointer to **float64** |  | [optional] 
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
**Locale** | Pointer to **string** |  | [optional] 
**DigestEmailFrequency** | Pointer to **float64** |  | [optional] 
**DisplayLabel** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateTenantUserBody

`func NewUpdateTenantUserBody() *UpdateTenantUserBody`

NewUpdateTenantUserBody instantiates a new UpdateTenantUserBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTenantUserBodyWithDefaults

`func NewUpdateTenantUserBodyWithDefaults() *UpdateTenantUserBody`

NewUpdateTenantUserBodyWithDefaults instantiates a new UpdateTenantUserBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UpdateTenantUserBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateTenantUserBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateTenantUserBody) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateTenantUserBody) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateTenantUserBody) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateTenantUserBody) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateTenantUserBody) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateTenantUserBody) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *UpdateTenantUserBody) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *UpdateTenantUserBody) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *UpdateTenantUserBody) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *UpdateTenantUserBody) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateTenantUserBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateTenantUserBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateTenantUserBody) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateTenantUserBody) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSignUpDate

`func (o *UpdateTenantUserBody) GetSignUpDate() float64`

GetSignUpDate returns the SignUpDate field if non-nil, zero value otherwise.

### GetSignUpDateOk

`func (o *UpdateTenantUserBody) GetSignUpDateOk() (*float64, bool)`

GetSignUpDateOk returns a tuple with the SignUpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUpDate

`func (o *UpdateTenantUserBody) SetSignUpDate(v float64)`

SetSignUpDate sets SignUpDate field to given value.

### HasSignUpDate

`func (o *UpdateTenantUserBody) HasSignUpDate() bool`

HasSignUpDate returns a boolean if a field has been set.

### GetVerified

`func (o *UpdateTenantUserBody) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *UpdateTenantUserBody) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *UpdateTenantUserBody) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *UpdateTenantUserBody) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetLoginCount

`func (o *UpdateTenantUserBody) GetLoginCount() float64`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *UpdateTenantUserBody) GetLoginCountOk() (*float64, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *UpdateTenantUserBody) SetLoginCount(v float64)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *UpdateTenantUserBody) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### GetOptedInNotifications

`func (o *UpdateTenantUserBody) GetOptedInNotifications() bool`

GetOptedInNotifications returns the OptedInNotifications field if non-nil, zero value otherwise.

### GetOptedInNotificationsOk

`func (o *UpdateTenantUserBody) GetOptedInNotificationsOk() (*bool, bool)`

GetOptedInNotificationsOk returns a tuple with the OptedInNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedInNotifications

`func (o *UpdateTenantUserBody) SetOptedInNotifications(v bool)`

SetOptedInNotifications sets OptedInNotifications field to given value.

### HasOptedInNotifications

`func (o *UpdateTenantUserBody) HasOptedInNotifications() bool`

HasOptedInNotifications returns a boolean if a field has been set.

### GetOptedInTenantNotifications

`func (o *UpdateTenantUserBody) GetOptedInTenantNotifications() bool`

GetOptedInTenantNotifications returns the OptedInTenantNotifications field if non-nil, zero value otherwise.

### GetOptedInTenantNotificationsOk

`func (o *UpdateTenantUserBody) GetOptedInTenantNotificationsOk() (*bool, bool)`

GetOptedInTenantNotificationsOk returns a tuple with the OptedInTenantNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedInTenantNotifications

`func (o *UpdateTenantUserBody) SetOptedInTenantNotifications(v bool)`

SetOptedInTenantNotifications sets OptedInTenantNotifications field to given value.

### HasOptedInTenantNotifications

`func (o *UpdateTenantUserBody) HasOptedInTenantNotifications() bool`

HasOptedInTenantNotifications returns a boolean if a field has been set.

### GetHideAccountCode

`func (o *UpdateTenantUserBody) GetHideAccountCode() bool`

GetHideAccountCode returns the HideAccountCode field if non-nil, zero value otherwise.

### GetHideAccountCodeOk

`func (o *UpdateTenantUserBody) GetHideAccountCodeOk() (*bool, bool)`

GetHideAccountCodeOk returns a tuple with the HideAccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAccountCode

`func (o *UpdateTenantUserBody) SetHideAccountCode(v bool)`

SetHideAccountCode sets HideAccountCode field to given value.

### HasHideAccountCode

`func (o *UpdateTenantUserBody) HasHideAccountCode() bool`

HasHideAccountCode returns a boolean if a field has been set.

### GetAvatarSrc

`func (o *UpdateTenantUserBody) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *UpdateTenantUserBody) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *UpdateTenantUserBody) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *UpdateTenantUserBody) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### GetIsHelpRequestAdmin

`func (o *UpdateTenantUserBody) GetIsHelpRequestAdmin() bool`

GetIsHelpRequestAdmin returns the IsHelpRequestAdmin field if non-nil, zero value otherwise.

### GetIsHelpRequestAdminOk

`func (o *UpdateTenantUserBody) GetIsHelpRequestAdminOk() (*bool, bool)`

GetIsHelpRequestAdminOk returns a tuple with the IsHelpRequestAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHelpRequestAdmin

`func (o *UpdateTenantUserBody) SetIsHelpRequestAdmin(v bool)`

SetIsHelpRequestAdmin sets IsHelpRequestAdmin field to given value.

### HasIsHelpRequestAdmin

`func (o *UpdateTenantUserBody) HasIsHelpRequestAdmin() bool`

HasIsHelpRequestAdmin returns a boolean if a field has been set.

### GetIsAccountOwner

`func (o *UpdateTenantUserBody) GetIsAccountOwner() bool`

GetIsAccountOwner returns the IsAccountOwner field if non-nil, zero value otherwise.

### GetIsAccountOwnerOk

`func (o *UpdateTenantUserBody) GetIsAccountOwnerOk() (*bool, bool)`

GetIsAccountOwnerOk returns a tuple with the IsAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountOwner

`func (o *UpdateTenantUserBody) SetIsAccountOwner(v bool)`

SetIsAccountOwner sets IsAccountOwner field to given value.

### HasIsAccountOwner

`func (o *UpdateTenantUserBody) HasIsAccountOwner() bool`

HasIsAccountOwner returns a boolean if a field has been set.

### GetIsAdminAdmin

`func (o *UpdateTenantUserBody) GetIsAdminAdmin() bool`

GetIsAdminAdmin returns the IsAdminAdmin field if non-nil, zero value otherwise.

### GetIsAdminAdminOk

`func (o *UpdateTenantUserBody) GetIsAdminAdminOk() (*bool, bool)`

GetIsAdminAdminOk returns a tuple with the IsAdminAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdminAdmin

`func (o *UpdateTenantUserBody) SetIsAdminAdmin(v bool)`

SetIsAdminAdmin sets IsAdminAdmin field to given value.

### HasIsAdminAdmin

`func (o *UpdateTenantUserBody) HasIsAdminAdmin() bool`

HasIsAdminAdmin returns a boolean if a field has been set.

### GetIsBillingAdmin

`func (o *UpdateTenantUserBody) GetIsBillingAdmin() bool`

GetIsBillingAdmin returns the IsBillingAdmin field if non-nil, zero value otherwise.

### GetIsBillingAdminOk

`func (o *UpdateTenantUserBody) GetIsBillingAdminOk() (*bool, bool)`

GetIsBillingAdminOk returns a tuple with the IsBillingAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillingAdmin

`func (o *UpdateTenantUserBody) SetIsBillingAdmin(v bool)`

SetIsBillingAdmin sets IsBillingAdmin field to given value.

### HasIsBillingAdmin

`func (o *UpdateTenantUserBody) HasIsBillingAdmin() bool`

HasIsBillingAdmin returns a boolean if a field has been set.

### GetIsAnalyticsAdmin

`func (o *UpdateTenantUserBody) GetIsAnalyticsAdmin() bool`

GetIsAnalyticsAdmin returns the IsAnalyticsAdmin field if non-nil, zero value otherwise.

### GetIsAnalyticsAdminOk

`func (o *UpdateTenantUserBody) GetIsAnalyticsAdminOk() (*bool, bool)`

GetIsAnalyticsAdminOk returns a tuple with the IsAnalyticsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnalyticsAdmin

`func (o *UpdateTenantUserBody) SetIsAnalyticsAdmin(v bool)`

SetIsAnalyticsAdmin sets IsAnalyticsAdmin field to given value.

### HasIsAnalyticsAdmin

`func (o *UpdateTenantUserBody) HasIsAnalyticsAdmin() bool`

HasIsAnalyticsAdmin returns a boolean if a field has been set.

### GetIsCustomizationAdmin

`func (o *UpdateTenantUserBody) GetIsCustomizationAdmin() bool`

GetIsCustomizationAdmin returns the IsCustomizationAdmin field if non-nil, zero value otherwise.

### GetIsCustomizationAdminOk

`func (o *UpdateTenantUserBody) GetIsCustomizationAdminOk() (*bool, bool)`

GetIsCustomizationAdminOk returns a tuple with the IsCustomizationAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomizationAdmin

`func (o *UpdateTenantUserBody) SetIsCustomizationAdmin(v bool)`

SetIsCustomizationAdmin sets IsCustomizationAdmin field to given value.

### HasIsCustomizationAdmin

`func (o *UpdateTenantUserBody) HasIsCustomizationAdmin() bool`

HasIsCustomizationAdmin returns a boolean if a field has been set.

### GetIsManageDataAdmin

`func (o *UpdateTenantUserBody) GetIsManageDataAdmin() bool`

GetIsManageDataAdmin returns the IsManageDataAdmin field if non-nil, zero value otherwise.

### GetIsManageDataAdminOk

`func (o *UpdateTenantUserBody) GetIsManageDataAdminOk() (*bool, bool)`

GetIsManageDataAdminOk returns a tuple with the IsManageDataAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManageDataAdmin

`func (o *UpdateTenantUserBody) SetIsManageDataAdmin(v bool)`

SetIsManageDataAdmin sets IsManageDataAdmin field to given value.

### HasIsManageDataAdmin

`func (o *UpdateTenantUserBody) HasIsManageDataAdmin() bool`

HasIsManageDataAdmin returns a boolean if a field has been set.

### GetIsCommentModeratorAdmin

`func (o *UpdateTenantUserBody) GetIsCommentModeratorAdmin() bool`

GetIsCommentModeratorAdmin returns the IsCommentModeratorAdmin field if non-nil, zero value otherwise.

### GetIsCommentModeratorAdminOk

`func (o *UpdateTenantUserBody) GetIsCommentModeratorAdminOk() (*bool, bool)`

GetIsCommentModeratorAdminOk returns a tuple with the IsCommentModeratorAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCommentModeratorAdmin

`func (o *UpdateTenantUserBody) SetIsCommentModeratorAdmin(v bool)`

SetIsCommentModeratorAdmin sets IsCommentModeratorAdmin field to given value.

### HasIsCommentModeratorAdmin

`func (o *UpdateTenantUserBody) HasIsCommentModeratorAdmin() bool`

HasIsCommentModeratorAdmin returns a boolean if a field has been set.

### GetIsAPIAdmin

`func (o *UpdateTenantUserBody) GetIsAPIAdmin() bool`

GetIsAPIAdmin returns the IsAPIAdmin field if non-nil, zero value otherwise.

### GetIsAPIAdminOk

`func (o *UpdateTenantUserBody) GetIsAPIAdminOk() (*bool, bool)`

GetIsAPIAdminOk returns a tuple with the IsAPIAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAPIAdmin

`func (o *UpdateTenantUserBody) SetIsAPIAdmin(v bool)`

SetIsAPIAdmin sets IsAPIAdmin field to given value.

### HasIsAPIAdmin

`func (o *UpdateTenantUserBody) HasIsAPIAdmin() bool`

HasIsAPIAdmin returns a boolean if a field has been set.

### GetModeratorIds

`func (o *UpdateTenantUserBody) GetModeratorIds() []string`

GetModeratorIds returns the ModeratorIds field if non-nil, zero value otherwise.

### GetModeratorIdsOk

`func (o *UpdateTenantUserBody) GetModeratorIdsOk() (*[]string, bool)`

GetModeratorIdsOk returns a tuple with the ModeratorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeratorIds

`func (o *UpdateTenantUserBody) SetModeratorIds(v []string)`

SetModeratorIds sets ModeratorIds field to given value.

### HasModeratorIds

`func (o *UpdateTenantUserBody) HasModeratorIds() bool`

HasModeratorIds returns a boolean if a field has been set.

### GetLocale

`func (o *UpdateTenantUserBody) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UpdateTenantUserBody) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UpdateTenantUserBody) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UpdateTenantUserBody) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetDigestEmailFrequency

`func (o *UpdateTenantUserBody) GetDigestEmailFrequency() float64`

GetDigestEmailFrequency returns the DigestEmailFrequency field if non-nil, zero value otherwise.

### GetDigestEmailFrequencyOk

`func (o *UpdateTenantUserBody) GetDigestEmailFrequencyOk() (*float64, bool)`

GetDigestEmailFrequencyOk returns a tuple with the DigestEmailFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestEmailFrequency

`func (o *UpdateTenantUserBody) SetDigestEmailFrequency(v float64)`

SetDigestEmailFrequency sets DigestEmailFrequency field to given value.

### HasDigestEmailFrequency

`func (o *UpdateTenantUserBody) HasDigestEmailFrequency() bool`

HasDigestEmailFrequency returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *UpdateTenantUserBody) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *UpdateTenantUserBody) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *UpdateTenantUserBody) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *UpdateTenantUserBody) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


