# UserNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 
**AnonUserId** | Pointer to **NullableString** |  | [optional] 
**UrlId** | **string** |  | 
**Url** | **string** |  | 
**PageTitle** | Pointer to **NullableString** |  | [optional] 
**RelatedObjectType** | [**NotificationObjectType**](NotificationObjectType.md) |  | 
**RelatedObjectId** | **string** |  | 
**Viewed** | **bool** |  | 
**IsUnreadMessage** | **bool** |  | 
**Sent** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**Type** | [**NotificationType**](NotificationType.md) |  | 
**FromCommentId** | Pointer to **NullableString** |  | [optional] 
**FromVoteId** | Pointer to **NullableString** |  | [optional] 
**FromUserName** | Pointer to **NullableString** |  | [optional] 
**FromUserId** | Pointer to **NullableString** |  | [optional] 
**FromUserAvatarSrc** | Pointer to **NullableString** |  | [optional] 
**OptedOut** | **bool** |  | 
**Count** | Pointer to **int64** |  | [optional] 
**RelatedIds** | Pointer to **[]string** |  | [optional] 
**FromUserIds** | Pointer to **[]string** |  | [optional] 
**FromUserNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUserNotification

`func NewUserNotification(id string, tenantId string, urlId string, url string, relatedObjectType NotificationObjectType, relatedObjectId string, viewed bool, isUnreadMessage bool, sent bool, createdAt time.Time, type_ NotificationType, optedOut bool, ) *UserNotification`

NewUserNotification instantiates a new UserNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNotificationWithDefaults

`func NewUserNotificationWithDefaults() *UserNotification`

NewUserNotificationWithDefaults instantiates a new UserNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserNotification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserNotification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserNotification) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *UserNotification) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserNotification) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserNotification) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUserId

`func (o *UserNotification) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserNotification) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserNotification) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserNotification) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *UserNotification) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UserNotification) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetAnonUserId

`func (o *UserNotification) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *UserNotification) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *UserNotification) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.

### HasAnonUserId

`func (o *UserNotification) HasAnonUserId() bool`

HasAnonUserId returns a boolean if a field has been set.

### SetAnonUserIdNil

`func (o *UserNotification) SetAnonUserIdNil(b bool)`

 SetAnonUserIdNil sets the value for AnonUserId to be an explicit nil

### UnsetAnonUserId
`func (o *UserNotification) UnsetAnonUserId()`

UnsetAnonUserId ensures that no value is present for AnonUserId, not even an explicit nil
### GetUrlId

`func (o *UserNotification) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *UserNotification) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *UserNotification) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetUrl

`func (o *UserNotification) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserNotification) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserNotification) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPageTitle

`func (o *UserNotification) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *UserNotification) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *UserNotification) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *UserNotification) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### SetPageTitleNil

`func (o *UserNotification) SetPageTitleNil(b bool)`

 SetPageTitleNil sets the value for PageTitle to be an explicit nil

### UnsetPageTitle
`func (o *UserNotification) UnsetPageTitle()`

UnsetPageTitle ensures that no value is present for PageTitle, not even an explicit nil
### GetRelatedObjectType

`func (o *UserNotification) GetRelatedObjectType() NotificationObjectType`

GetRelatedObjectType returns the RelatedObjectType field if non-nil, zero value otherwise.

### GetRelatedObjectTypeOk

`func (o *UserNotification) GetRelatedObjectTypeOk() (*NotificationObjectType, bool)`

GetRelatedObjectTypeOk returns a tuple with the RelatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectType

`func (o *UserNotification) SetRelatedObjectType(v NotificationObjectType)`

SetRelatedObjectType sets RelatedObjectType field to given value.


### GetRelatedObjectId

`func (o *UserNotification) GetRelatedObjectId() string`

GetRelatedObjectId returns the RelatedObjectId field if non-nil, zero value otherwise.

### GetRelatedObjectIdOk

`func (o *UserNotification) GetRelatedObjectIdOk() (*string, bool)`

GetRelatedObjectIdOk returns a tuple with the RelatedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectId

`func (o *UserNotification) SetRelatedObjectId(v string)`

SetRelatedObjectId sets RelatedObjectId field to given value.


### GetViewed

`func (o *UserNotification) GetViewed() bool`

GetViewed returns the Viewed field if non-nil, zero value otherwise.

### GetViewedOk

`func (o *UserNotification) GetViewedOk() (*bool, bool)`

GetViewedOk returns a tuple with the Viewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewed

`func (o *UserNotification) SetViewed(v bool)`

SetViewed sets Viewed field to given value.


### GetIsUnreadMessage

`func (o *UserNotification) GetIsUnreadMessage() bool`

GetIsUnreadMessage returns the IsUnreadMessage field if non-nil, zero value otherwise.

### GetIsUnreadMessageOk

`func (o *UserNotification) GetIsUnreadMessageOk() (*bool, bool)`

GetIsUnreadMessageOk returns a tuple with the IsUnreadMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnreadMessage

`func (o *UserNotification) SetIsUnreadMessage(v bool)`

SetIsUnreadMessage sets IsUnreadMessage field to given value.


### GetSent

`func (o *UserNotification) GetSent() bool`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *UserNotification) GetSentOk() (*bool, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *UserNotification) SetSent(v bool)`

SetSent sets Sent field to given value.


### GetCreatedAt

`func (o *UserNotification) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserNotification) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserNotification) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetType

`func (o *UserNotification) GetType() NotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserNotification) GetTypeOk() (*NotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserNotification) SetType(v NotificationType)`

SetType sets Type field to given value.


### GetFromCommentId

`func (o *UserNotification) GetFromCommentId() string`

GetFromCommentId returns the FromCommentId field if non-nil, zero value otherwise.

### GetFromCommentIdOk

`func (o *UserNotification) GetFromCommentIdOk() (*string, bool)`

GetFromCommentIdOk returns a tuple with the FromCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCommentId

`func (o *UserNotification) SetFromCommentId(v string)`

SetFromCommentId sets FromCommentId field to given value.

### HasFromCommentId

`func (o *UserNotification) HasFromCommentId() bool`

HasFromCommentId returns a boolean if a field has been set.

### SetFromCommentIdNil

`func (o *UserNotification) SetFromCommentIdNil(b bool)`

 SetFromCommentIdNil sets the value for FromCommentId to be an explicit nil

### UnsetFromCommentId
`func (o *UserNotification) UnsetFromCommentId()`

UnsetFromCommentId ensures that no value is present for FromCommentId, not even an explicit nil
### GetFromVoteId

`func (o *UserNotification) GetFromVoteId() string`

GetFromVoteId returns the FromVoteId field if non-nil, zero value otherwise.

### GetFromVoteIdOk

`func (o *UserNotification) GetFromVoteIdOk() (*string, bool)`

GetFromVoteIdOk returns a tuple with the FromVoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromVoteId

`func (o *UserNotification) SetFromVoteId(v string)`

SetFromVoteId sets FromVoteId field to given value.

### HasFromVoteId

`func (o *UserNotification) HasFromVoteId() bool`

HasFromVoteId returns a boolean if a field has been set.

### SetFromVoteIdNil

`func (o *UserNotification) SetFromVoteIdNil(b bool)`

 SetFromVoteIdNil sets the value for FromVoteId to be an explicit nil

### UnsetFromVoteId
`func (o *UserNotification) UnsetFromVoteId()`

UnsetFromVoteId ensures that no value is present for FromVoteId, not even an explicit nil
### GetFromUserName

`func (o *UserNotification) GetFromUserName() string`

GetFromUserName returns the FromUserName field if non-nil, zero value otherwise.

### GetFromUserNameOk

`func (o *UserNotification) GetFromUserNameOk() (*string, bool)`

GetFromUserNameOk returns a tuple with the FromUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserName

`func (o *UserNotification) SetFromUserName(v string)`

SetFromUserName sets FromUserName field to given value.

### HasFromUserName

`func (o *UserNotification) HasFromUserName() bool`

HasFromUserName returns a boolean if a field has been set.

### SetFromUserNameNil

`func (o *UserNotification) SetFromUserNameNil(b bool)`

 SetFromUserNameNil sets the value for FromUserName to be an explicit nil

### UnsetFromUserName
`func (o *UserNotification) UnsetFromUserName()`

UnsetFromUserName ensures that no value is present for FromUserName, not even an explicit nil
### GetFromUserId

`func (o *UserNotification) GetFromUserId() string`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *UserNotification) GetFromUserIdOk() (*string, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *UserNotification) SetFromUserId(v string)`

SetFromUserId sets FromUserId field to given value.

### HasFromUserId

`func (o *UserNotification) HasFromUserId() bool`

HasFromUserId returns a boolean if a field has been set.

### SetFromUserIdNil

`func (o *UserNotification) SetFromUserIdNil(b bool)`

 SetFromUserIdNil sets the value for FromUserId to be an explicit nil

### UnsetFromUserId
`func (o *UserNotification) UnsetFromUserId()`

UnsetFromUserId ensures that no value is present for FromUserId, not even an explicit nil
### GetFromUserAvatarSrc

`func (o *UserNotification) GetFromUserAvatarSrc() string`

GetFromUserAvatarSrc returns the FromUserAvatarSrc field if non-nil, zero value otherwise.

### GetFromUserAvatarSrcOk

`func (o *UserNotification) GetFromUserAvatarSrcOk() (*string, bool)`

GetFromUserAvatarSrcOk returns a tuple with the FromUserAvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserAvatarSrc

`func (o *UserNotification) SetFromUserAvatarSrc(v string)`

SetFromUserAvatarSrc sets FromUserAvatarSrc field to given value.

### HasFromUserAvatarSrc

`func (o *UserNotification) HasFromUserAvatarSrc() bool`

HasFromUserAvatarSrc returns a boolean if a field has been set.

### SetFromUserAvatarSrcNil

`func (o *UserNotification) SetFromUserAvatarSrcNil(b bool)`

 SetFromUserAvatarSrcNil sets the value for FromUserAvatarSrc to be an explicit nil

### UnsetFromUserAvatarSrc
`func (o *UserNotification) UnsetFromUserAvatarSrc()`

UnsetFromUserAvatarSrc ensures that no value is present for FromUserAvatarSrc, not even an explicit nil
### GetOptedOut

`func (o *UserNotification) GetOptedOut() bool`

GetOptedOut returns the OptedOut field if non-nil, zero value otherwise.

### GetOptedOutOk

`func (o *UserNotification) GetOptedOutOk() (*bool, bool)`

GetOptedOutOk returns a tuple with the OptedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedOut

`func (o *UserNotification) SetOptedOut(v bool)`

SetOptedOut sets OptedOut field to given value.


### GetCount

`func (o *UserNotification) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UserNotification) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UserNotification) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *UserNotification) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetRelatedIds

`func (o *UserNotification) GetRelatedIds() []string`

GetRelatedIds returns the RelatedIds field if non-nil, zero value otherwise.

### GetRelatedIdsOk

`func (o *UserNotification) GetRelatedIdsOk() (*[]string, bool)`

GetRelatedIdsOk returns a tuple with the RelatedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIds

`func (o *UserNotification) SetRelatedIds(v []string)`

SetRelatedIds sets RelatedIds field to given value.

### HasRelatedIds

`func (o *UserNotification) HasRelatedIds() bool`

HasRelatedIds returns a boolean if a field has been set.

### GetFromUserIds

`func (o *UserNotification) GetFromUserIds() []string`

GetFromUserIds returns the FromUserIds field if non-nil, zero value otherwise.

### GetFromUserIdsOk

`func (o *UserNotification) GetFromUserIdsOk() (*[]string, bool)`

GetFromUserIdsOk returns a tuple with the FromUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserIds

`func (o *UserNotification) SetFromUserIds(v []string)`

SetFromUserIds sets FromUserIds field to given value.

### HasFromUserIds

`func (o *UserNotification) HasFromUserIds() bool`

HasFromUserIds returns a boolean if a field has been set.

### GetFromUserNames

`func (o *UserNotification) GetFromUserNames() []string`

GetFromUserNames returns the FromUserNames field if non-nil, zero value otherwise.

### GetFromUserNamesOk

`func (o *UserNotification) GetFromUserNamesOk() (*[]string, bool)`

GetFromUserNamesOk returns a tuple with the FromUserNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserNames

`func (o *UserNotification) SetFromUserNames(v []string)`

SetFromUserNames sets FromUserNames field to given value.

### HasFromUserNames

`func (o *UserNotification) HasFromUserNames() bool`

HasFromUserNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


