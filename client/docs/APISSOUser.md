# APISSOUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Username** | **string** |  | 
**WebsiteUrl** | **string** |  | 
**Email** | **string** |  | 
**SignUpDate** | **int64** |  | 
**CreatedFromUrlId** | **string** |  | 
**LoginCount** | **int32** |  | 
**AvatarSrc** | **string** |  | 
**OptedInNotifications** | **bool** |  | 
**OptedInSubscriptionNotifications** | **bool** |  | 
**DisplayLabel** | **string** |  | 
**DisplayName** | **string** |  | 
**IsAccountOwner** | Pointer to **bool** |  | [optional] 
**IsAdminAdmin** | Pointer to **bool** |  | [optional] 
**IsCommentModeratorAdmin** | Pointer to **bool** |  | [optional] 
**IsProfileActivityPrivate** | Pointer to **bool** |  | [optional] 
**IsProfileCommentsPrivate** | Pointer to **bool** |  | [optional] 
**IsProfileDMDisabled** | Pointer to **bool** |  | [optional] 
**HasBlockedUsers** | Pointer to **bool** |  | [optional] 
**GroupIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAPISSOUser

`func NewAPISSOUser(id string, username string, websiteUrl string, email string, signUpDate int64, createdFromUrlId string, loginCount int32, avatarSrc string, optedInNotifications bool, optedInSubscriptionNotifications bool, displayLabel string, displayName string, ) *APISSOUser`

NewAPISSOUser instantiates a new APISSOUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPISSOUserWithDefaults

`func NewAPISSOUserWithDefaults() *APISSOUser`

NewAPISSOUserWithDefaults instantiates a new APISSOUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APISSOUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APISSOUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APISSOUser) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *APISSOUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *APISSOUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *APISSOUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetWebsiteUrl

`func (o *APISSOUser) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *APISSOUser) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *APISSOUser) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.


### GetEmail

`func (o *APISSOUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *APISSOUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *APISSOUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetSignUpDate

`func (o *APISSOUser) GetSignUpDate() int64`

GetSignUpDate returns the SignUpDate field if non-nil, zero value otherwise.

### GetSignUpDateOk

`func (o *APISSOUser) GetSignUpDateOk() (*int64, bool)`

GetSignUpDateOk returns a tuple with the SignUpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUpDate

`func (o *APISSOUser) SetSignUpDate(v int64)`

SetSignUpDate sets SignUpDate field to given value.


### GetCreatedFromUrlId

`func (o *APISSOUser) GetCreatedFromUrlId() string`

GetCreatedFromUrlId returns the CreatedFromUrlId field if non-nil, zero value otherwise.

### GetCreatedFromUrlIdOk

`func (o *APISSOUser) GetCreatedFromUrlIdOk() (*string, bool)`

GetCreatedFromUrlIdOk returns a tuple with the CreatedFromUrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFromUrlId

`func (o *APISSOUser) SetCreatedFromUrlId(v string)`

SetCreatedFromUrlId sets CreatedFromUrlId field to given value.


### GetLoginCount

`func (o *APISSOUser) GetLoginCount() int32`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *APISSOUser) GetLoginCountOk() (*int32, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *APISSOUser) SetLoginCount(v int32)`

SetLoginCount sets LoginCount field to given value.


### GetAvatarSrc

`func (o *APISSOUser) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *APISSOUser) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *APISSOUser) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.


### GetOptedInNotifications

`func (o *APISSOUser) GetOptedInNotifications() bool`

GetOptedInNotifications returns the OptedInNotifications field if non-nil, zero value otherwise.

### GetOptedInNotificationsOk

`func (o *APISSOUser) GetOptedInNotificationsOk() (*bool, bool)`

GetOptedInNotificationsOk returns a tuple with the OptedInNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedInNotifications

`func (o *APISSOUser) SetOptedInNotifications(v bool)`

SetOptedInNotifications sets OptedInNotifications field to given value.


### GetOptedInSubscriptionNotifications

`func (o *APISSOUser) GetOptedInSubscriptionNotifications() bool`

GetOptedInSubscriptionNotifications returns the OptedInSubscriptionNotifications field if non-nil, zero value otherwise.

### GetOptedInSubscriptionNotificationsOk

`func (o *APISSOUser) GetOptedInSubscriptionNotificationsOk() (*bool, bool)`

GetOptedInSubscriptionNotificationsOk returns a tuple with the OptedInSubscriptionNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedInSubscriptionNotifications

`func (o *APISSOUser) SetOptedInSubscriptionNotifications(v bool)`

SetOptedInSubscriptionNotifications sets OptedInSubscriptionNotifications field to given value.


### GetDisplayLabel

`func (o *APISSOUser) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *APISSOUser) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *APISSOUser) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.


### GetDisplayName

`func (o *APISSOUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *APISSOUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *APISSOUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIsAccountOwner

`func (o *APISSOUser) GetIsAccountOwner() bool`

GetIsAccountOwner returns the IsAccountOwner field if non-nil, zero value otherwise.

### GetIsAccountOwnerOk

`func (o *APISSOUser) GetIsAccountOwnerOk() (*bool, bool)`

GetIsAccountOwnerOk returns a tuple with the IsAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountOwner

`func (o *APISSOUser) SetIsAccountOwner(v bool)`

SetIsAccountOwner sets IsAccountOwner field to given value.

### HasIsAccountOwner

`func (o *APISSOUser) HasIsAccountOwner() bool`

HasIsAccountOwner returns a boolean if a field has been set.

### GetIsAdminAdmin

`func (o *APISSOUser) GetIsAdminAdmin() bool`

GetIsAdminAdmin returns the IsAdminAdmin field if non-nil, zero value otherwise.

### GetIsAdminAdminOk

`func (o *APISSOUser) GetIsAdminAdminOk() (*bool, bool)`

GetIsAdminAdminOk returns a tuple with the IsAdminAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdminAdmin

`func (o *APISSOUser) SetIsAdminAdmin(v bool)`

SetIsAdminAdmin sets IsAdminAdmin field to given value.

### HasIsAdminAdmin

`func (o *APISSOUser) HasIsAdminAdmin() bool`

HasIsAdminAdmin returns a boolean if a field has been set.

### GetIsCommentModeratorAdmin

`func (o *APISSOUser) GetIsCommentModeratorAdmin() bool`

GetIsCommentModeratorAdmin returns the IsCommentModeratorAdmin field if non-nil, zero value otherwise.

### GetIsCommentModeratorAdminOk

`func (o *APISSOUser) GetIsCommentModeratorAdminOk() (*bool, bool)`

GetIsCommentModeratorAdminOk returns a tuple with the IsCommentModeratorAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCommentModeratorAdmin

`func (o *APISSOUser) SetIsCommentModeratorAdmin(v bool)`

SetIsCommentModeratorAdmin sets IsCommentModeratorAdmin field to given value.

### HasIsCommentModeratorAdmin

`func (o *APISSOUser) HasIsCommentModeratorAdmin() bool`

HasIsCommentModeratorAdmin returns a boolean if a field has been set.

### GetIsProfileActivityPrivate

`func (o *APISSOUser) GetIsProfileActivityPrivate() bool`

GetIsProfileActivityPrivate returns the IsProfileActivityPrivate field if non-nil, zero value otherwise.

### GetIsProfileActivityPrivateOk

`func (o *APISSOUser) GetIsProfileActivityPrivateOk() (*bool, bool)`

GetIsProfileActivityPrivateOk returns a tuple with the IsProfileActivityPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProfileActivityPrivate

`func (o *APISSOUser) SetIsProfileActivityPrivate(v bool)`

SetIsProfileActivityPrivate sets IsProfileActivityPrivate field to given value.

### HasIsProfileActivityPrivate

`func (o *APISSOUser) HasIsProfileActivityPrivate() bool`

HasIsProfileActivityPrivate returns a boolean if a field has been set.

### GetIsProfileCommentsPrivate

`func (o *APISSOUser) GetIsProfileCommentsPrivate() bool`

GetIsProfileCommentsPrivate returns the IsProfileCommentsPrivate field if non-nil, zero value otherwise.

### GetIsProfileCommentsPrivateOk

`func (o *APISSOUser) GetIsProfileCommentsPrivateOk() (*bool, bool)`

GetIsProfileCommentsPrivateOk returns a tuple with the IsProfileCommentsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProfileCommentsPrivate

`func (o *APISSOUser) SetIsProfileCommentsPrivate(v bool)`

SetIsProfileCommentsPrivate sets IsProfileCommentsPrivate field to given value.

### HasIsProfileCommentsPrivate

`func (o *APISSOUser) HasIsProfileCommentsPrivate() bool`

HasIsProfileCommentsPrivate returns a boolean if a field has been set.

### GetIsProfileDMDisabled

`func (o *APISSOUser) GetIsProfileDMDisabled() bool`

GetIsProfileDMDisabled returns the IsProfileDMDisabled field if non-nil, zero value otherwise.

### GetIsProfileDMDisabledOk

`func (o *APISSOUser) GetIsProfileDMDisabledOk() (*bool, bool)`

GetIsProfileDMDisabledOk returns a tuple with the IsProfileDMDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProfileDMDisabled

`func (o *APISSOUser) SetIsProfileDMDisabled(v bool)`

SetIsProfileDMDisabled sets IsProfileDMDisabled field to given value.

### HasIsProfileDMDisabled

`func (o *APISSOUser) HasIsProfileDMDisabled() bool`

HasIsProfileDMDisabled returns a boolean if a field has been set.

### GetHasBlockedUsers

`func (o *APISSOUser) GetHasBlockedUsers() bool`

GetHasBlockedUsers returns the HasBlockedUsers field if non-nil, zero value otherwise.

### GetHasBlockedUsersOk

`func (o *APISSOUser) GetHasBlockedUsersOk() (*bool, bool)`

GetHasBlockedUsersOk returns a tuple with the HasBlockedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBlockedUsers

`func (o *APISSOUser) SetHasBlockedUsers(v bool)`

SetHasBlockedUsers sets HasBlockedUsers field to given value.

### HasHasBlockedUsers

`func (o *APISSOUser) HasHasBlockedUsers() bool`

HasHasBlockedUsers returns a boolean if a field has been set.

### GetGroupIds

`func (o *APISSOUser) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *APISSOUser) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *APISSOUser) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *APISSOUser) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


