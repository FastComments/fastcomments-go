# RenderableUserNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | Pointer to **string** |  | [optional] 
**ContextHTML** | Pointer to **string** |  | [optional] 
**FromUserNames** | Pointer to **[]string** |  | [optional] 
**FromUserIds** | Pointer to **[]string** |  | [optional] 
**RelatedIds** | Pointer to **[]string** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**OptedOut** | **bool** |  | 
**FromUserAvatarSrc** | Pointer to **NullableString** |  | [optional] 
**FromUserId** | Pointer to **NullableString** |  | [optional] 
**FromUserName** | Pointer to **NullableString** |  | [optional] 
**FromCommentId** | Pointer to **NullableString** |  | [optional] 
**Type** | [**NotificationType**](NotificationType.md) |  | 
**CreatedAt** | **string** |  | 
**Sent** | **string** |  | 
**Viewed** | **string** |  | 
**RelatedObjectId** | **string** |  | 
**RelatedObjectType** | [**NotificationObjectType**](NotificationObjectType.md) |  | 
**PageTitle** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** |  | 
**UrlId** | **string** |  | 
**Id** | **string** |  | 

## Methods

### NewRenderableUserNotification

`func NewRenderableUserNotification(optedOut bool, type_ NotificationType, createdAt string, sent string, viewed string, relatedObjectId string, relatedObjectType NotificationObjectType, url string, urlId string, id string, ) *RenderableUserNotification`

NewRenderableUserNotification instantiates a new RenderableUserNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderableUserNotificationWithDefaults

`func NewRenderableUserNotificationWithDefaults() *RenderableUserNotification`

NewRenderableUserNotificationWithDefaults instantiates a new RenderableUserNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *RenderableUserNotification) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *RenderableUserNotification) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *RenderableUserNotification) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *RenderableUserNotification) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetContextHTML

`func (o *RenderableUserNotification) GetContextHTML() string`

GetContextHTML returns the ContextHTML field if non-nil, zero value otherwise.

### GetContextHTMLOk

`func (o *RenderableUserNotification) GetContextHTMLOk() (*string, bool)`

GetContextHTMLOk returns a tuple with the ContextHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextHTML

`func (o *RenderableUserNotification) SetContextHTML(v string)`

SetContextHTML sets ContextHTML field to given value.

### HasContextHTML

`func (o *RenderableUserNotification) HasContextHTML() bool`

HasContextHTML returns a boolean if a field has been set.

### GetFromUserNames

`func (o *RenderableUserNotification) GetFromUserNames() []string`

GetFromUserNames returns the FromUserNames field if non-nil, zero value otherwise.

### GetFromUserNamesOk

`func (o *RenderableUserNotification) GetFromUserNamesOk() (*[]string, bool)`

GetFromUserNamesOk returns a tuple with the FromUserNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserNames

`func (o *RenderableUserNotification) SetFromUserNames(v []string)`

SetFromUserNames sets FromUserNames field to given value.

### HasFromUserNames

`func (o *RenderableUserNotification) HasFromUserNames() bool`

HasFromUserNames returns a boolean if a field has been set.

### GetFromUserIds

`func (o *RenderableUserNotification) GetFromUserIds() []string`

GetFromUserIds returns the FromUserIds field if non-nil, zero value otherwise.

### GetFromUserIdsOk

`func (o *RenderableUserNotification) GetFromUserIdsOk() (*[]string, bool)`

GetFromUserIdsOk returns a tuple with the FromUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserIds

`func (o *RenderableUserNotification) SetFromUserIds(v []string)`

SetFromUserIds sets FromUserIds field to given value.

### HasFromUserIds

`func (o *RenderableUserNotification) HasFromUserIds() bool`

HasFromUserIds returns a boolean if a field has been set.

### GetRelatedIds

`func (o *RenderableUserNotification) GetRelatedIds() []string`

GetRelatedIds returns the RelatedIds field if non-nil, zero value otherwise.

### GetRelatedIdsOk

`func (o *RenderableUserNotification) GetRelatedIdsOk() (*[]string, bool)`

GetRelatedIdsOk returns a tuple with the RelatedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIds

`func (o *RenderableUserNotification) SetRelatedIds(v []string)`

SetRelatedIds sets RelatedIds field to given value.

### HasRelatedIds

`func (o *RenderableUserNotification) HasRelatedIds() bool`

HasRelatedIds returns a boolean if a field has been set.

### GetCount

`func (o *RenderableUserNotification) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RenderableUserNotification) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RenderableUserNotification) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *RenderableUserNotification) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOptedOut

`func (o *RenderableUserNotification) GetOptedOut() bool`

GetOptedOut returns the OptedOut field if non-nil, zero value otherwise.

### GetOptedOutOk

`func (o *RenderableUserNotification) GetOptedOutOk() (*bool, bool)`

GetOptedOutOk returns a tuple with the OptedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedOut

`func (o *RenderableUserNotification) SetOptedOut(v bool)`

SetOptedOut sets OptedOut field to given value.


### GetFromUserAvatarSrc

`func (o *RenderableUserNotification) GetFromUserAvatarSrc() string`

GetFromUserAvatarSrc returns the FromUserAvatarSrc field if non-nil, zero value otherwise.

### GetFromUserAvatarSrcOk

`func (o *RenderableUserNotification) GetFromUserAvatarSrcOk() (*string, bool)`

GetFromUserAvatarSrcOk returns a tuple with the FromUserAvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserAvatarSrc

`func (o *RenderableUserNotification) SetFromUserAvatarSrc(v string)`

SetFromUserAvatarSrc sets FromUserAvatarSrc field to given value.

### HasFromUserAvatarSrc

`func (o *RenderableUserNotification) HasFromUserAvatarSrc() bool`

HasFromUserAvatarSrc returns a boolean if a field has been set.

### SetFromUserAvatarSrcNil

`func (o *RenderableUserNotification) SetFromUserAvatarSrcNil(b bool)`

 SetFromUserAvatarSrcNil sets the value for FromUserAvatarSrc to be an explicit nil

### UnsetFromUserAvatarSrc
`func (o *RenderableUserNotification) UnsetFromUserAvatarSrc()`

UnsetFromUserAvatarSrc ensures that no value is present for FromUserAvatarSrc, not even an explicit nil
### GetFromUserId

`func (o *RenderableUserNotification) GetFromUserId() string`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *RenderableUserNotification) GetFromUserIdOk() (*string, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *RenderableUserNotification) SetFromUserId(v string)`

SetFromUserId sets FromUserId field to given value.

### HasFromUserId

`func (o *RenderableUserNotification) HasFromUserId() bool`

HasFromUserId returns a boolean if a field has been set.

### SetFromUserIdNil

`func (o *RenderableUserNotification) SetFromUserIdNil(b bool)`

 SetFromUserIdNil sets the value for FromUserId to be an explicit nil

### UnsetFromUserId
`func (o *RenderableUserNotification) UnsetFromUserId()`

UnsetFromUserId ensures that no value is present for FromUserId, not even an explicit nil
### GetFromUserName

`func (o *RenderableUserNotification) GetFromUserName() string`

GetFromUserName returns the FromUserName field if non-nil, zero value otherwise.

### GetFromUserNameOk

`func (o *RenderableUserNotification) GetFromUserNameOk() (*string, bool)`

GetFromUserNameOk returns a tuple with the FromUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserName

`func (o *RenderableUserNotification) SetFromUserName(v string)`

SetFromUserName sets FromUserName field to given value.

### HasFromUserName

`func (o *RenderableUserNotification) HasFromUserName() bool`

HasFromUserName returns a boolean if a field has been set.

### SetFromUserNameNil

`func (o *RenderableUserNotification) SetFromUserNameNil(b bool)`

 SetFromUserNameNil sets the value for FromUserName to be an explicit nil

### UnsetFromUserName
`func (o *RenderableUserNotification) UnsetFromUserName()`

UnsetFromUserName ensures that no value is present for FromUserName, not even an explicit nil
### GetFromCommentId

`func (o *RenderableUserNotification) GetFromCommentId() string`

GetFromCommentId returns the FromCommentId field if non-nil, zero value otherwise.

### GetFromCommentIdOk

`func (o *RenderableUserNotification) GetFromCommentIdOk() (*string, bool)`

GetFromCommentIdOk returns a tuple with the FromCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCommentId

`func (o *RenderableUserNotification) SetFromCommentId(v string)`

SetFromCommentId sets FromCommentId field to given value.

### HasFromCommentId

`func (o *RenderableUserNotification) HasFromCommentId() bool`

HasFromCommentId returns a boolean if a field has been set.

### SetFromCommentIdNil

`func (o *RenderableUserNotification) SetFromCommentIdNil(b bool)`

 SetFromCommentIdNil sets the value for FromCommentId to be an explicit nil

### UnsetFromCommentId
`func (o *RenderableUserNotification) UnsetFromCommentId()`

UnsetFromCommentId ensures that no value is present for FromCommentId, not even an explicit nil
### GetType

`func (o *RenderableUserNotification) GetType() NotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RenderableUserNotification) GetTypeOk() (*NotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RenderableUserNotification) SetType(v NotificationType)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *RenderableUserNotification) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RenderableUserNotification) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RenderableUserNotification) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetSent

`func (o *RenderableUserNotification) GetSent() string`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *RenderableUserNotification) GetSentOk() (*string, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *RenderableUserNotification) SetSent(v string)`

SetSent sets Sent field to given value.


### GetViewed

`func (o *RenderableUserNotification) GetViewed() string`

GetViewed returns the Viewed field if non-nil, zero value otherwise.

### GetViewedOk

`func (o *RenderableUserNotification) GetViewedOk() (*string, bool)`

GetViewedOk returns a tuple with the Viewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewed

`func (o *RenderableUserNotification) SetViewed(v string)`

SetViewed sets Viewed field to given value.


### GetRelatedObjectId

`func (o *RenderableUserNotification) GetRelatedObjectId() string`

GetRelatedObjectId returns the RelatedObjectId field if non-nil, zero value otherwise.

### GetRelatedObjectIdOk

`func (o *RenderableUserNotification) GetRelatedObjectIdOk() (*string, bool)`

GetRelatedObjectIdOk returns a tuple with the RelatedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectId

`func (o *RenderableUserNotification) SetRelatedObjectId(v string)`

SetRelatedObjectId sets RelatedObjectId field to given value.


### GetRelatedObjectType

`func (o *RenderableUserNotification) GetRelatedObjectType() NotificationObjectType`

GetRelatedObjectType returns the RelatedObjectType field if non-nil, zero value otherwise.

### GetRelatedObjectTypeOk

`func (o *RenderableUserNotification) GetRelatedObjectTypeOk() (*NotificationObjectType, bool)`

GetRelatedObjectTypeOk returns a tuple with the RelatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectType

`func (o *RenderableUserNotification) SetRelatedObjectType(v NotificationObjectType)`

SetRelatedObjectType sets RelatedObjectType field to given value.


### GetPageTitle

`func (o *RenderableUserNotification) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *RenderableUserNotification) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *RenderableUserNotification) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *RenderableUserNotification) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### SetPageTitleNil

`func (o *RenderableUserNotification) SetPageTitleNil(b bool)`

 SetPageTitleNil sets the value for PageTitle to be an explicit nil

### UnsetPageTitle
`func (o *RenderableUserNotification) UnsetPageTitle()`

UnsetPageTitle ensures that no value is present for PageTitle, not even an explicit nil
### GetUrl

`func (o *RenderableUserNotification) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RenderableUserNotification) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RenderableUserNotification) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUrlId

`func (o *RenderableUserNotification) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *RenderableUserNotification) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *RenderableUserNotification) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetId

`func (o *RenderableUserNotification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RenderableUserNotification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RenderableUserNotification) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


