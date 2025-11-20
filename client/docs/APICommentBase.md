# APICommentBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AiDeterminedSpam** | Pointer to **bool** |  | [optional] 
**AnonUserId** | Pointer to **NullableString** |  | [optional] 
**Approved** | **bool** |  | 
**AvatarSrc** | Pointer to **NullableString** |  | [optional] 
**Badges** | Pointer to [**[]CommentUserBadgeInfo**](CommentUserBadgeInfo.md) |  | [optional] 
**Comment** | **string** |  | 
**CommentHTML** | **string** |  | 
**CommenterEmail** | Pointer to **NullableString** |  | [optional] 
**CommenterLink** | Pointer to **NullableString** |  | [optional] 
**CommenterName** | **string** |  | 
**Date** | **NullableTime** |  | 
**DisplayLabel** | Pointer to **NullableString** |  | [optional] 
**Domain** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalParentId** | Pointer to **NullableString** |  | [optional] 
**ExpireAt** | Pointer to **NullableTime** |  | [optional] 
**FeedbackIds** | Pointer to **[]string** |  | [optional] 
**FlagCount** | Pointer to **NullableInt32** |  | [optional] 
**FromProductId** | Pointer to **int32** |  | [optional] 
**HasCode** | Pointer to **bool** |  | [optional] 
**HasImages** | Pointer to **bool** |  | [optional] 
**HasLinks** | Pointer to **bool** |  | [optional] 
**HashTags** | Pointer to [**[]CommentUserHashTagInfo**](CommentUserHashTagInfo.md) |  | [optional] 
**IsByAdmin** | Pointer to **bool** |  | [optional] 
**IsByModerator** | Pointer to **bool** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**IsDeletedUser** | Pointer to **bool** |  | [optional] 
**IsPinned** | Pointer to **NullableBool** |  | [optional] 
**IsLocked** | Pointer to **NullableBool** |  | [optional] 
**IsSpam** | Pointer to **bool** |  | [optional] 
**LocalDateHours** | Pointer to **NullableInt32** |  | [optional] 
**LocalDateString** | Pointer to **NullableString** |  | [optional] 
**Locale** | **NullableString** |  | 
**Mentions** | Pointer to [**[]CommentUserMentionInfo**](CommentUserMentionInfo.md) |  | [optional] 
**Meta** | Pointer to [**NullableFCommentMeta**](FCommentMeta.md) |  | [optional] 
**ModerationGroupIds** | Pointer to **[]string** |  | [optional] 
**NotificationSentForParent** | Pointer to **bool** |  | [optional] 
**NotificationSentForParentTenant** | Pointer to **bool** |  | [optional] 
**PageTitle** | Pointer to **NullableString** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Rating** | Pointer to **NullableFloat64** |  | [optional] 
**Reviewed** | Pointer to **bool** |  | [optional] 
**TenantId** | **string** |  | 
**Url** | **string** |  | 
**UrlId** | **string** |  | 
**UrlIdRaw** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Verified** | **bool** |  | 
**VerifiedDate** | Pointer to **NullableTime** |  | [optional] 
**Votes** | Pointer to **NullableInt32** |  | [optional] 
**VotesDown** | Pointer to **NullableInt32** |  | [optional] 
**VotesUp** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAPICommentBase

`func NewAPICommentBase(id string, approved bool, comment string, commentHTML string, commenterName string, date NullableTime, locale NullableString, tenantId string, url string, urlId string, verified bool, ) *APICommentBase`

NewAPICommentBase instantiates a new APICommentBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPICommentBaseWithDefaults

`func NewAPICommentBaseWithDefaults() *APICommentBase`

NewAPICommentBaseWithDefaults instantiates a new APICommentBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APICommentBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APICommentBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APICommentBase) SetId(v string)`

SetId sets Id field to given value.


### GetAiDeterminedSpam

`func (o *APICommentBase) GetAiDeterminedSpam() bool`

GetAiDeterminedSpam returns the AiDeterminedSpam field if non-nil, zero value otherwise.

### GetAiDeterminedSpamOk

`func (o *APICommentBase) GetAiDeterminedSpamOk() (*bool, bool)`

GetAiDeterminedSpamOk returns a tuple with the AiDeterminedSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiDeterminedSpam

`func (o *APICommentBase) SetAiDeterminedSpam(v bool)`

SetAiDeterminedSpam sets AiDeterminedSpam field to given value.

### HasAiDeterminedSpam

`func (o *APICommentBase) HasAiDeterminedSpam() bool`

HasAiDeterminedSpam returns a boolean if a field has been set.

### GetAnonUserId

`func (o *APICommentBase) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *APICommentBase) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *APICommentBase) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.

### HasAnonUserId

`func (o *APICommentBase) HasAnonUserId() bool`

HasAnonUserId returns a boolean if a field has been set.

### SetAnonUserIdNil

`func (o *APICommentBase) SetAnonUserIdNil(b bool)`

 SetAnonUserIdNil sets the value for AnonUserId to be an explicit nil

### UnsetAnonUserId
`func (o *APICommentBase) UnsetAnonUserId()`

UnsetAnonUserId ensures that no value is present for AnonUserId, not even an explicit nil
### GetApproved

`func (o *APICommentBase) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *APICommentBase) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *APICommentBase) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetAvatarSrc

`func (o *APICommentBase) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *APICommentBase) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *APICommentBase) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *APICommentBase) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### SetAvatarSrcNil

`func (o *APICommentBase) SetAvatarSrcNil(b bool)`

 SetAvatarSrcNil sets the value for AvatarSrc to be an explicit nil

### UnsetAvatarSrc
`func (o *APICommentBase) UnsetAvatarSrc()`

UnsetAvatarSrc ensures that no value is present for AvatarSrc, not even an explicit nil
### GetBadges

`func (o *APICommentBase) GetBadges() []CommentUserBadgeInfo`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *APICommentBase) GetBadgesOk() (*[]CommentUserBadgeInfo, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *APICommentBase) SetBadges(v []CommentUserBadgeInfo)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *APICommentBase) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### SetBadgesNil

`func (o *APICommentBase) SetBadgesNil(b bool)`

 SetBadgesNil sets the value for Badges to be an explicit nil

### UnsetBadges
`func (o *APICommentBase) UnsetBadges()`

UnsetBadges ensures that no value is present for Badges, not even an explicit nil
### GetComment

`func (o *APICommentBase) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *APICommentBase) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *APICommentBase) SetComment(v string)`

SetComment sets Comment field to given value.


### GetCommentHTML

`func (o *APICommentBase) GetCommentHTML() string`

GetCommentHTML returns the CommentHTML field if non-nil, zero value otherwise.

### GetCommentHTMLOk

`func (o *APICommentBase) GetCommentHTMLOk() (*string, bool)`

GetCommentHTMLOk returns a tuple with the CommentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentHTML

`func (o *APICommentBase) SetCommentHTML(v string)`

SetCommentHTML sets CommentHTML field to given value.


### GetCommenterEmail

`func (o *APICommentBase) GetCommenterEmail() string`

GetCommenterEmail returns the CommenterEmail field if non-nil, zero value otherwise.

### GetCommenterEmailOk

`func (o *APICommentBase) GetCommenterEmailOk() (*string, bool)`

GetCommenterEmailOk returns a tuple with the CommenterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterEmail

`func (o *APICommentBase) SetCommenterEmail(v string)`

SetCommenterEmail sets CommenterEmail field to given value.

### HasCommenterEmail

`func (o *APICommentBase) HasCommenterEmail() bool`

HasCommenterEmail returns a boolean if a field has been set.

### SetCommenterEmailNil

`func (o *APICommentBase) SetCommenterEmailNil(b bool)`

 SetCommenterEmailNil sets the value for CommenterEmail to be an explicit nil

### UnsetCommenterEmail
`func (o *APICommentBase) UnsetCommenterEmail()`

UnsetCommenterEmail ensures that no value is present for CommenterEmail, not even an explicit nil
### GetCommenterLink

`func (o *APICommentBase) GetCommenterLink() string`

GetCommenterLink returns the CommenterLink field if non-nil, zero value otherwise.

### GetCommenterLinkOk

`func (o *APICommentBase) GetCommenterLinkOk() (*string, bool)`

GetCommenterLinkOk returns a tuple with the CommenterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterLink

`func (o *APICommentBase) SetCommenterLink(v string)`

SetCommenterLink sets CommenterLink field to given value.

### HasCommenterLink

`func (o *APICommentBase) HasCommenterLink() bool`

HasCommenterLink returns a boolean if a field has been set.

### SetCommenterLinkNil

`func (o *APICommentBase) SetCommenterLinkNil(b bool)`

 SetCommenterLinkNil sets the value for CommenterLink to be an explicit nil

### UnsetCommenterLink
`func (o *APICommentBase) UnsetCommenterLink()`

UnsetCommenterLink ensures that no value is present for CommenterLink, not even an explicit nil
### GetCommenterName

`func (o *APICommentBase) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *APICommentBase) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *APICommentBase) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetDate

`func (o *APICommentBase) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *APICommentBase) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *APICommentBase) SetDate(v time.Time)`

SetDate sets Date field to given value.


### SetDateNil

`func (o *APICommentBase) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *APICommentBase) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDisplayLabel

`func (o *APICommentBase) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *APICommentBase) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *APICommentBase) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *APICommentBase) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### SetDisplayLabelNil

`func (o *APICommentBase) SetDisplayLabelNil(b bool)`

 SetDisplayLabelNil sets the value for DisplayLabel to be an explicit nil

### UnsetDisplayLabel
`func (o *APICommentBase) UnsetDisplayLabel()`

UnsetDisplayLabel ensures that no value is present for DisplayLabel, not even an explicit nil
### GetDomain

`func (o *APICommentBase) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *APICommentBase) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *APICommentBase) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *APICommentBase) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *APICommentBase) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *APICommentBase) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetExternalId

`func (o *APICommentBase) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *APICommentBase) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *APICommentBase) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *APICommentBase) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalParentId

`func (o *APICommentBase) GetExternalParentId() string`

GetExternalParentId returns the ExternalParentId field if non-nil, zero value otherwise.

### GetExternalParentIdOk

`func (o *APICommentBase) GetExternalParentIdOk() (*string, bool)`

GetExternalParentIdOk returns a tuple with the ExternalParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalParentId

`func (o *APICommentBase) SetExternalParentId(v string)`

SetExternalParentId sets ExternalParentId field to given value.

### HasExternalParentId

`func (o *APICommentBase) HasExternalParentId() bool`

HasExternalParentId returns a boolean if a field has been set.

### SetExternalParentIdNil

`func (o *APICommentBase) SetExternalParentIdNil(b bool)`

 SetExternalParentIdNil sets the value for ExternalParentId to be an explicit nil

### UnsetExternalParentId
`func (o *APICommentBase) UnsetExternalParentId()`

UnsetExternalParentId ensures that no value is present for ExternalParentId, not even an explicit nil
### GetExpireAt

`func (o *APICommentBase) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *APICommentBase) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *APICommentBase) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *APICommentBase) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### SetExpireAtNil

`func (o *APICommentBase) SetExpireAtNil(b bool)`

 SetExpireAtNil sets the value for ExpireAt to be an explicit nil

### UnsetExpireAt
`func (o *APICommentBase) UnsetExpireAt()`

UnsetExpireAt ensures that no value is present for ExpireAt, not even an explicit nil
### GetFeedbackIds

`func (o *APICommentBase) GetFeedbackIds() []string`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *APICommentBase) GetFeedbackIdsOk() (*[]string, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *APICommentBase) SetFeedbackIds(v []string)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *APICommentBase) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.

### GetFlagCount

`func (o *APICommentBase) GetFlagCount() int32`

GetFlagCount returns the FlagCount field if non-nil, zero value otherwise.

### GetFlagCountOk

`func (o *APICommentBase) GetFlagCountOk() (*int32, bool)`

GetFlagCountOk returns a tuple with the FlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCount

`func (o *APICommentBase) SetFlagCount(v int32)`

SetFlagCount sets FlagCount field to given value.

### HasFlagCount

`func (o *APICommentBase) HasFlagCount() bool`

HasFlagCount returns a boolean if a field has been set.

### SetFlagCountNil

`func (o *APICommentBase) SetFlagCountNil(b bool)`

 SetFlagCountNil sets the value for FlagCount to be an explicit nil

### UnsetFlagCount
`func (o *APICommentBase) UnsetFlagCount()`

UnsetFlagCount ensures that no value is present for FlagCount, not even an explicit nil
### GetFromProductId

`func (o *APICommentBase) GetFromProductId() int32`

GetFromProductId returns the FromProductId field if non-nil, zero value otherwise.

### GetFromProductIdOk

`func (o *APICommentBase) GetFromProductIdOk() (*int32, bool)`

GetFromProductIdOk returns a tuple with the FromProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromProductId

`func (o *APICommentBase) SetFromProductId(v int32)`

SetFromProductId sets FromProductId field to given value.

### HasFromProductId

`func (o *APICommentBase) HasFromProductId() bool`

HasFromProductId returns a boolean if a field has been set.

### GetHasCode

`func (o *APICommentBase) GetHasCode() bool`

GetHasCode returns the HasCode field if non-nil, zero value otherwise.

### GetHasCodeOk

`func (o *APICommentBase) GetHasCodeOk() (*bool, bool)`

GetHasCodeOk returns a tuple with the HasCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCode

`func (o *APICommentBase) SetHasCode(v bool)`

SetHasCode sets HasCode field to given value.

### HasHasCode

`func (o *APICommentBase) HasHasCode() bool`

HasHasCode returns a boolean if a field has been set.

### GetHasImages

`func (o *APICommentBase) GetHasImages() bool`

GetHasImages returns the HasImages field if non-nil, zero value otherwise.

### GetHasImagesOk

`func (o *APICommentBase) GetHasImagesOk() (*bool, bool)`

GetHasImagesOk returns a tuple with the HasImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImages

`func (o *APICommentBase) SetHasImages(v bool)`

SetHasImages sets HasImages field to given value.

### HasHasImages

`func (o *APICommentBase) HasHasImages() bool`

HasHasImages returns a boolean if a field has been set.

### GetHasLinks

`func (o *APICommentBase) GetHasLinks() bool`

GetHasLinks returns the HasLinks field if non-nil, zero value otherwise.

### GetHasLinksOk

`func (o *APICommentBase) GetHasLinksOk() (*bool, bool)`

GetHasLinksOk returns a tuple with the HasLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLinks

`func (o *APICommentBase) SetHasLinks(v bool)`

SetHasLinks sets HasLinks field to given value.

### HasHasLinks

`func (o *APICommentBase) HasHasLinks() bool`

HasHasLinks returns a boolean if a field has been set.

### GetHashTags

`func (o *APICommentBase) GetHashTags() []CommentUserHashTagInfo`

GetHashTags returns the HashTags field if non-nil, zero value otherwise.

### GetHashTagsOk

`func (o *APICommentBase) GetHashTagsOk() (*[]CommentUserHashTagInfo, bool)`

GetHashTagsOk returns a tuple with the HashTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashTags

`func (o *APICommentBase) SetHashTags(v []CommentUserHashTagInfo)`

SetHashTags sets HashTags field to given value.

### HasHashTags

`func (o *APICommentBase) HasHashTags() bool`

HasHashTags returns a boolean if a field has been set.

### GetIsByAdmin

`func (o *APICommentBase) GetIsByAdmin() bool`

GetIsByAdmin returns the IsByAdmin field if non-nil, zero value otherwise.

### GetIsByAdminOk

`func (o *APICommentBase) GetIsByAdminOk() (*bool, bool)`

GetIsByAdminOk returns a tuple with the IsByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByAdmin

`func (o *APICommentBase) SetIsByAdmin(v bool)`

SetIsByAdmin sets IsByAdmin field to given value.

### HasIsByAdmin

`func (o *APICommentBase) HasIsByAdmin() bool`

HasIsByAdmin returns a boolean if a field has been set.

### GetIsByModerator

`func (o *APICommentBase) GetIsByModerator() bool`

GetIsByModerator returns the IsByModerator field if non-nil, zero value otherwise.

### GetIsByModeratorOk

`func (o *APICommentBase) GetIsByModeratorOk() (*bool, bool)`

GetIsByModeratorOk returns a tuple with the IsByModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByModerator

`func (o *APICommentBase) SetIsByModerator(v bool)`

SetIsByModerator sets IsByModerator field to given value.

### HasIsByModerator

`func (o *APICommentBase) HasIsByModerator() bool`

HasIsByModerator returns a boolean if a field has been set.

### GetIsDeleted

`func (o *APICommentBase) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *APICommentBase) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *APICommentBase) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *APICommentBase) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsDeletedUser

`func (o *APICommentBase) GetIsDeletedUser() bool`

GetIsDeletedUser returns the IsDeletedUser field if non-nil, zero value otherwise.

### GetIsDeletedUserOk

`func (o *APICommentBase) GetIsDeletedUserOk() (*bool, bool)`

GetIsDeletedUserOk returns a tuple with the IsDeletedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletedUser

`func (o *APICommentBase) SetIsDeletedUser(v bool)`

SetIsDeletedUser sets IsDeletedUser field to given value.

### HasIsDeletedUser

`func (o *APICommentBase) HasIsDeletedUser() bool`

HasIsDeletedUser returns a boolean if a field has been set.

### GetIsPinned

`func (o *APICommentBase) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *APICommentBase) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *APICommentBase) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.

### HasIsPinned

`func (o *APICommentBase) HasIsPinned() bool`

HasIsPinned returns a boolean if a field has been set.

### SetIsPinnedNil

`func (o *APICommentBase) SetIsPinnedNil(b bool)`

 SetIsPinnedNil sets the value for IsPinned to be an explicit nil

### UnsetIsPinned
`func (o *APICommentBase) UnsetIsPinned()`

UnsetIsPinned ensures that no value is present for IsPinned, not even an explicit nil
### GetIsLocked

`func (o *APICommentBase) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *APICommentBase) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *APICommentBase) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *APICommentBase) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### SetIsLockedNil

`func (o *APICommentBase) SetIsLockedNil(b bool)`

 SetIsLockedNil sets the value for IsLocked to be an explicit nil

### UnsetIsLocked
`func (o *APICommentBase) UnsetIsLocked()`

UnsetIsLocked ensures that no value is present for IsLocked, not even an explicit nil
### GetIsSpam

`func (o *APICommentBase) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *APICommentBase) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *APICommentBase) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *APICommentBase) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetLocalDateHours

`func (o *APICommentBase) GetLocalDateHours() int32`

GetLocalDateHours returns the LocalDateHours field if non-nil, zero value otherwise.

### GetLocalDateHoursOk

`func (o *APICommentBase) GetLocalDateHoursOk() (*int32, bool)`

GetLocalDateHoursOk returns a tuple with the LocalDateHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateHours

`func (o *APICommentBase) SetLocalDateHours(v int32)`

SetLocalDateHours sets LocalDateHours field to given value.

### HasLocalDateHours

`func (o *APICommentBase) HasLocalDateHours() bool`

HasLocalDateHours returns a boolean if a field has been set.

### SetLocalDateHoursNil

`func (o *APICommentBase) SetLocalDateHoursNil(b bool)`

 SetLocalDateHoursNil sets the value for LocalDateHours to be an explicit nil

### UnsetLocalDateHours
`func (o *APICommentBase) UnsetLocalDateHours()`

UnsetLocalDateHours ensures that no value is present for LocalDateHours, not even an explicit nil
### GetLocalDateString

`func (o *APICommentBase) GetLocalDateString() string`

GetLocalDateString returns the LocalDateString field if non-nil, zero value otherwise.

### GetLocalDateStringOk

`func (o *APICommentBase) GetLocalDateStringOk() (*string, bool)`

GetLocalDateStringOk returns a tuple with the LocalDateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateString

`func (o *APICommentBase) SetLocalDateString(v string)`

SetLocalDateString sets LocalDateString field to given value.

### HasLocalDateString

`func (o *APICommentBase) HasLocalDateString() bool`

HasLocalDateString returns a boolean if a field has been set.

### SetLocalDateStringNil

`func (o *APICommentBase) SetLocalDateStringNil(b bool)`

 SetLocalDateStringNil sets the value for LocalDateString to be an explicit nil

### UnsetLocalDateString
`func (o *APICommentBase) UnsetLocalDateString()`

UnsetLocalDateString ensures that no value is present for LocalDateString, not even an explicit nil
### GetLocale

`func (o *APICommentBase) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *APICommentBase) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *APICommentBase) SetLocale(v string)`

SetLocale sets Locale field to given value.


### SetLocaleNil

`func (o *APICommentBase) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *APICommentBase) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetMentions

`func (o *APICommentBase) GetMentions() []CommentUserMentionInfo`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *APICommentBase) GetMentionsOk() (*[]CommentUserMentionInfo, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *APICommentBase) SetMentions(v []CommentUserMentionInfo)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *APICommentBase) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetMeta

`func (o *APICommentBase) GetMeta() FCommentMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *APICommentBase) GetMetaOk() (*FCommentMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *APICommentBase) SetMeta(v FCommentMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *APICommentBase) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *APICommentBase) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *APICommentBase) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetModerationGroupIds

`func (o *APICommentBase) GetModerationGroupIds() []string`

GetModerationGroupIds returns the ModerationGroupIds field if non-nil, zero value otherwise.

### GetModerationGroupIdsOk

`func (o *APICommentBase) GetModerationGroupIdsOk() (*[]string, bool)`

GetModerationGroupIdsOk returns a tuple with the ModerationGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationGroupIds

`func (o *APICommentBase) SetModerationGroupIds(v []string)`

SetModerationGroupIds sets ModerationGroupIds field to given value.

### HasModerationGroupIds

`func (o *APICommentBase) HasModerationGroupIds() bool`

HasModerationGroupIds returns a boolean if a field has been set.

### SetModerationGroupIdsNil

`func (o *APICommentBase) SetModerationGroupIdsNil(b bool)`

 SetModerationGroupIdsNil sets the value for ModerationGroupIds to be an explicit nil

### UnsetModerationGroupIds
`func (o *APICommentBase) UnsetModerationGroupIds()`

UnsetModerationGroupIds ensures that no value is present for ModerationGroupIds, not even an explicit nil
### GetNotificationSentForParent

`func (o *APICommentBase) GetNotificationSentForParent() bool`

GetNotificationSentForParent returns the NotificationSentForParent field if non-nil, zero value otherwise.

### GetNotificationSentForParentOk

`func (o *APICommentBase) GetNotificationSentForParentOk() (*bool, bool)`

GetNotificationSentForParentOk returns a tuple with the NotificationSentForParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSentForParent

`func (o *APICommentBase) SetNotificationSentForParent(v bool)`

SetNotificationSentForParent sets NotificationSentForParent field to given value.

### HasNotificationSentForParent

`func (o *APICommentBase) HasNotificationSentForParent() bool`

HasNotificationSentForParent returns a boolean if a field has been set.

### GetNotificationSentForParentTenant

`func (o *APICommentBase) GetNotificationSentForParentTenant() bool`

GetNotificationSentForParentTenant returns the NotificationSentForParentTenant field if non-nil, zero value otherwise.

### GetNotificationSentForParentTenantOk

`func (o *APICommentBase) GetNotificationSentForParentTenantOk() (*bool, bool)`

GetNotificationSentForParentTenantOk returns a tuple with the NotificationSentForParentTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSentForParentTenant

`func (o *APICommentBase) SetNotificationSentForParentTenant(v bool)`

SetNotificationSentForParentTenant sets NotificationSentForParentTenant field to given value.

### HasNotificationSentForParentTenant

`func (o *APICommentBase) HasNotificationSentForParentTenant() bool`

HasNotificationSentForParentTenant returns a boolean if a field has been set.

### GetPageTitle

`func (o *APICommentBase) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *APICommentBase) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *APICommentBase) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *APICommentBase) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### SetPageTitleNil

`func (o *APICommentBase) SetPageTitleNil(b bool)`

 SetPageTitleNil sets the value for PageTitle to be an explicit nil

### UnsetPageTitle
`func (o *APICommentBase) UnsetPageTitle()`

UnsetPageTitle ensures that no value is present for PageTitle, not even an explicit nil
### GetParentId

`func (o *APICommentBase) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *APICommentBase) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *APICommentBase) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *APICommentBase) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *APICommentBase) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *APICommentBase) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetRating

`func (o *APICommentBase) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *APICommentBase) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *APICommentBase) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *APICommentBase) HasRating() bool`

HasRating returns a boolean if a field has been set.

### SetRatingNil

`func (o *APICommentBase) SetRatingNil(b bool)`

 SetRatingNil sets the value for Rating to be an explicit nil

### UnsetRating
`func (o *APICommentBase) UnsetRating()`

UnsetRating ensures that no value is present for Rating, not even an explicit nil
### GetReviewed

`func (o *APICommentBase) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *APICommentBase) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *APICommentBase) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *APICommentBase) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetTenantId

`func (o *APICommentBase) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *APICommentBase) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *APICommentBase) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUrl

`func (o *APICommentBase) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *APICommentBase) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *APICommentBase) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUrlId

`func (o *APICommentBase) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *APICommentBase) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *APICommentBase) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetUrlIdRaw

`func (o *APICommentBase) GetUrlIdRaw() string`

GetUrlIdRaw returns the UrlIdRaw field if non-nil, zero value otherwise.

### GetUrlIdRawOk

`func (o *APICommentBase) GetUrlIdRawOk() (*string, bool)`

GetUrlIdRawOk returns a tuple with the UrlIdRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdRaw

`func (o *APICommentBase) SetUrlIdRaw(v string)`

SetUrlIdRaw sets UrlIdRaw field to given value.

### HasUrlIdRaw

`func (o *APICommentBase) HasUrlIdRaw() bool`

HasUrlIdRaw returns a boolean if a field has been set.

### GetUserId

`func (o *APICommentBase) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *APICommentBase) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *APICommentBase) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *APICommentBase) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *APICommentBase) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *APICommentBase) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetVerified

`func (o *APICommentBase) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *APICommentBase) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *APICommentBase) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetVerifiedDate

`func (o *APICommentBase) GetVerifiedDate() time.Time`

GetVerifiedDate returns the VerifiedDate field if non-nil, zero value otherwise.

### GetVerifiedDateOk

`func (o *APICommentBase) GetVerifiedDateOk() (*time.Time, bool)`

GetVerifiedDateOk returns a tuple with the VerifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedDate

`func (o *APICommentBase) SetVerifiedDate(v time.Time)`

SetVerifiedDate sets VerifiedDate field to given value.

### HasVerifiedDate

`func (o *APICommentBase) HasVerifiedDate() bool`

HasVerifiedDate returns a boolean if a field has been set.

### SetVerifiedDateNil

`func (o *APICommentBase) SetVerifiedDateNil(b bool)`

 SetVerifiedDateNil sets the value for VerifiedDate to be an explicit nil

### UnsetVerifiedDate
`func (o *APICommentBase) UnsetVerifiedDate()`

UnsetVerifiedDate ensures that no value is present for VerifiedDate, not even an explicit nil
### GetVotes

`func (o *APICommentBase) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *APICommentBase) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *APICommentBase) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *APICommentBase) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### SetVotesNil

`func (o *APICommentBase) SetVotesNil(b bool)`

 SetVotesNil sets the value for Votes to be an explicit nil

### UnsetVotes
`func (o *APICommentBase) UnsetVotes()`

UnsetVotes ensures that no value is present for Votes, not even an explicit nil
### GetVotesDown

`func (o *APICommentBase) GetVotesDown() int32`

GetVotesDown returns the VotesDown field if non-nil, zero value otherwise.

### GetVotesDownOk

`func (o *APICommentBase) GetVotesDownOk() (*int32, bool)`

GetVotesDownOk returns a tuple with the VotesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDown

`func (o *APICommentBase) SetVotesDown(v int32)`

SetVotesDown sets VotesDown field to given value.

### HasVotesDown

`func (o *APICommentBase) HasVotesDown() bool`

HasVotesDown returns a boolean if a field has been set.

### SetVotesDownNil

`func (o *APICommentBase) SetVotesDownNil(b bool)`

 SetVotesDownNil sets the value for VotesDown to be an explicit nil

### UnsetVotesDown
`func (o *APICommentBase) UnsetVotesDown()`

UnsetVotesDown ensures that no value is present for VotesDown, not even an explicit nil
### GetVotesUp

`func (o *APICommentBase) GetVotesUp() int32`

GetVotesUp returns the VotesUp field if non-nil, zero value otherwise.

### GetVotesUpOk

`func (o *APICommentBase) GetVotesUpOk() (*int32, bool)`

GetVotesUpOk returns a tuple with the VotesUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUp

`func (o *APICommentBase) SetVotesUp(v int32)`

SetVotesUp sets VotesUp field to given value.

### HasVotesUp

`func (o *APICommentBase) HasVotesUp() bool`

HasVotesUp returns a boolean if a field has been set.

### SetVotesUpNil

`func (o *APICommentBase) SetVotesUpNil(b bool)`

 SetVotesUpNil sets the value for VotesUp to be an explicit nil

### UnsetVotesUp
`func (o *APICommentBase) UnsetVotesUp()`

UnsetVotesUp ensures that no value is present for VotesUp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


