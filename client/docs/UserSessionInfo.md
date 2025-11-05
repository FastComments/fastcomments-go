# UserSessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Authorized** | Pointer to **bool** |  | [optional] 
**AvatarSrc** | Pointer to **NullableString** |  | [optional] 
**Badges** | Pointer to [**[]CommentUserBadgeInfo**](CommentUserBadgeInfo.md) |  | [optional] 
**DisplayLabel** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**GroupIds** | Pointer to **[]string** |  | [optional] 
**HasBlockedUsers** | Pointer to **bool** |  | [optional] 
**IsAnonSession** | Pointer to **bool** |  | [optional] 
**SessionId** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**WebsiteUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewUserSessionInfo

`func NewUserSessionInfo() *UserSessionInfo`

NewUserSessionInfo instantiates a new UserSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionInfoWithDefaults

`func NewUserSessionInfoWithDefaults() *UserSessionInfo`

NewUserSessionInfoWithDefaults instantiates a new UserSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSessionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSessionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSessionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserSessionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthorized

`func (o *UserSessionInfo) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *UserSessionInfo) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *UserSessionInfo) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *UserSessionInfo) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### GetAvatarSrc

`func (o *UserSessionInfo) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *UserSessionInfo) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *UserSessionInfo) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *UserSessionInfo) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### SetAvatarSrcNil

`func (o *UserSessionInfo) SetAvatarSrcNil(b bool)`

 SetAvatarSrcNil sets the value for AvatarSrc to be an explicit nil

### UnsetAvatarSrc
`func (o *UserSessionInfo) UnsetAvatarSrc()`

UnsetAvatarSrc ensures that no value is present for AvatarSrc, not even an explicit nil
### GetBadges

`func (o *UserSessionInfo) GetBadges() []CommentUserBadgeInfo`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *UserSessionInfo) GetBadgesOk() (*[]CommentUserBadgeInfo, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *UserSessionInfo) SetBadges(v []CommentUserBadgeInfo)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *UserSessionInfo) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *UserSessionInfo) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *UserSessionInfo) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *UserSessionInfo) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *UserSessionInfo) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserSessionInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserSessionInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserSessionInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserSessionInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *UserSessionInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSessionInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSessionInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserSessionInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserSessionInfo) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserSessionInfo) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetGroupIds

`func (o *UserSessionInfo) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *UserSessionInfo) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *UserSessionInfo) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *UserSessionInfo) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetHasBlockedUsers

`func (o *UserSessionInfo) GetHasBlockedUsers() bool`

GetHasBlockedUsers returns the HasBlockedUsers field if non-nil, zero value otherwise.

### GetHasBlockedUsersOk

`func (o *UserSessionInfo) GetHasBlockedUsersOk() (*bool, bool)`

GetHasBlockedUsersOk returns a tuple with the HasBlockedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBlockedUsers

`func (o *UserSessionInfo) SetHasBlockedUsers(v bool)`

SetHasBlockedUsers sets HasBlockedUsers field to given value.

### HasHasBlockedUsers

`func (o *UserSessionInfo) HasHasBlockedUsers() bool`

HasHasBlockedUsers returns a boolean if a field has been set.

### GetIsAnonSession

`func (o *UserSessionInfo) GetIsAnonSession() bool`

GetIsAnonSession returns the IsAnonSession field if non-nil, zero value otherwise.

### GetIsAnonSessionOk

`func (o *UserSessionInfo) GetIsAnonSessionOk() (*bool, bool)`

GetIsAnonSessionOk returns a tuple with the IsAnonSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonSession

`func (o *UserSessionInfo) SetIsAnonSession(v bool)`

SetIsAnonSession sets IsAnonSession field to given value.

### HasIsAnonSession

`func (o *UserSessionInfo) HasIsAnonSession() bool`

HasIsAnonSession returns a boolean if a field has been set.

### GetSessionId

`func (o *UserSessionInfo) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *UserSessionInfo) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *UserSessionInfo) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *UserSessionInfo) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *UserSessionInfo) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *UserSessionInfo) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetUsername

`func (o *UserSessionInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserSessionInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserSessionInfo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserSessionInfo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *UserSessionInfo) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *UserSessionInfo) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *UserSessionInfo) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *UserSessionInfo) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


