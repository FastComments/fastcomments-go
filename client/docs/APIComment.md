# APIComment

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
**Date** | **NullableFloat64** |  | 
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

### NewAPIComment

`func NewAPIComment(id string, approved bool, comment string, commentHTML string, commenterName string, date NullableFloat64, locale NullableString, tenantId string, url string, urlId string, verified bool, ) *APIComment`

NewAPIComment instantiates a new APIComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPICommentWithDefaults

`func NewAPICommentWithDefaults() *APIComment`

NewAPICommentWithDefaults instantiates a new APIComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APIComment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIComment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIComment) SetId(v string)`

SetId sets Id field to given value.


### GetAiDeterminedSpam

`func (o *APIComment) GetAiDeterminedSpam() bool`

GetAiDeterminedSpam returns the AiDeterminedSpam field if non-nil, zero value otherwise.

### GetAiDeterminedSpamOk

`func (o *APIComment) GetAiDeterminedSpamOk() (*bool, bool)`

GetAiDeterminedSpamOk returns a tuple with the AiDeterminedSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiDeterminedSpam

`func (o *APIComment) SetAiDeterminedSpam(v bool)`

SetAiDeterminedSpam sets AiDeterminedSpam field to given value.

### HasAiDeterminedSpam

`func (o *APIComment) HasAiDeterminedSpam() bool`

HasAiDeterminedSpam returns a boolean if a field has been set.

### GetAnonUserId

`func (o *APIComment) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *APIComment) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *APIComment) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.

### HasAnonUserId

`func (o *APIComment) HasAnonUserId() bool`

HasAnonUserId returns a boolean if a field has been set.

### SetAnonUserIdNil

`func (o *APIComment) SetAnonUserIdNil(b bool)`

 SetAnonUserIdNil sets the value for AnonUserId to be an explicit nil

### UnsetAnonUserId
`func (o *APIComment) UnsetAnonUserId()`

UnsetAnonUserId ensures that no value is present for AnonUserId, not even an explicit nil
### GetApproved

`func (o *APIComment) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *APIComment) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *APIComment) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetAvatarSrc

`func (o *APIComment) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *APIComment) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *APIComment) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *APIComment) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### SetAvatarSrcNil

`func (o *APIComment) SetAvatarSrcNil(b bool)`

 SetAvatarSrcNil sets the value for AvatarSrc to be an explicit nil

### UnsetAvatarSrc
`func (o *APIComment) UnsetAvatarSrc()`

UnsetAvatarSrc ensures that no value is present for AvatarSrc, not even an explicit nil
### GetBadges

`func (o *APIComment) GetBadges() []CommentUserBadgeInfo`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *APIComment) GetBadgesOk() (*[]CommentUserBadgeInfo, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *APIComment) SetBadges(v []CommentUserBadgeInfo)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *APIComment) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### SetBadgesNil

`func (o *APIComment) SetBadgesNil(b bool)`

 SetBadgesNil sets the value for Badges to be an explicit nil

### UnsetBadges
`func (o *APIComment) UnsetBadges()`

UnsetBadges ensures that no value is present for Badges, not even an explicit nil
### GetComment

`func (o *APIComment) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *APIComment) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *APIComment) SetComment(v string)`

SetComment sets Comment field to given value.


### GetCommentHTML

`func (o *APIComment) GetCommentHTML() string`

GetCommentHTML returns the CommentHTML field if non-nil, zero value otherwise.

### GetCommentHTMLOk

`func (o *APIComment) GetCommentHTMLOk() (*string, bool)`

GetCommentHTMLOk returns a tuple with the CommentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentHTML

`func (o *APIComment) SetCommentHTML(v string)`

SetCommentHTML sets CommentHTML field to given value.


### GetCommenterEmail

`func (o *APIComment) GetCommenterEmail() string`

GetCommenterEmail returns the CommenterEmail field if non-nil, zero value otherwise.

### GetCommenterEmailOk

`func (o *APIComment) GetCommenterEmailOk() (*string, bool)`

GetCommenterEmailOk returns a tuple with the CommenterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterEmail

`func (o *APIComment) SetCommenterEmail(v string)`

SetCommenterEmail sets CommenterEmail field to given value.

### HasCommenterEmail

`func (o *APIComment) HasCommenterEmail() bool`

HasCommenterEmail returns a boolean if a field has been set.

### SetCommenterEmailNil

`func (o *APIComment) SetCommenterEmailNil(b bool)`

 SetCommenterEmailNil sets the value for CommenterEmail to be an explicit nil

### UnsetCommenterEmail
`func (o *APIComment) UnsetCommenterEmail()`

UnsetCommenterEmail ensures that no value is present for CommenterEmail, not even an explicit nil
### GetCommenterLink

`func (o *APIComment) GetCommenterLink() string`

GetCommenterLink returns the CommenterLink field if non-nil, zero value otherwise.

### GetCommenterLinkOk

`func (o *APIComment) GetCommenterLinkOk() (*string, bool)`

GetCommenterLinkOk returns a tuple with the CommenterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterLink

`func (o *APIComment) SetCommenterLink(v string)`

SetCommenterLink sets CommenterLink field to given value.

### HasCommenterLink

`func (o *APIComment) HasCommenterLink() bool`

HasCommenterLink returns a boolean if a field has been set.

### SetCommenterLinkNil

`func (o *APIComment) SetCommenterLinkNil(b bool)`

 SetCommenterLinkNil sets the value for CommenterLink to be an explicit nil

### UnsetCommenterLink
`func (o *APIComment) UnsetCommenterLink()`

UnsetCommenterLink ensures that no value is present for CommenterLink, not even an explicit nil
### GetCommenterName

`func (o *APIComment) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *APIComment) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *APIComment) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetDate

`func (o *APIComment) GetDate() float64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *APIComment) GetDateOk() (*float64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *APIComment) SetDate(v float64)`

SetDate sets Date field to given value.


### SetDateNil

`func (o *APIComment) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *APIComment) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDisplayLabel

`func (o *APIComment) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *APIComment) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *APIComment) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *APIComment) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### SetDisplayLabelNil

`func (o *APIComment) SetDisplayLabelNil(b bool)`

 SetDisplayLabelNil sets the value for DisplayLabel to be an explicit nil

### UnsetDisplayLabel
`func (o *APIComment) UnsetDisplayLabel()`

UnsetDisplayLabel ensures that no value is present for DisplayLabel, not even an explicit nil
### GetDomain

`func (o *APIComment) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *APIComment) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *APIComment) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *APIComment) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *APIComment) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *APIComment) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetExternalId

`func (o *APIComment) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *APIComment) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *APIComment) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *APIComment) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalParentId

`func (o *APIComment) GetExternalParentId() string`

GetExternalParentId returns the ExternalParentId field if non-nil, zero value otherwise.

### GetExternalParentIdOk

`func (o *APIComment) GetExternalParentIdOk() (*string, bool)`

GetExternalParentIdOk returns a tuple with the ExternalParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalParentId

`func (o *APIComment) SetExternalParentId(v string)`

SetExternalParentId sets ExternalParentId field to given value.

### HasExternalParentId

`func (o *APIComment) HasExternalParentId() bool`

HasExternalParentId returns a boolean if a field has been set.

### SetExternalParentIdNil

`func (o *APIComment) SetExternalParentIdNil(b bool)`

 SetExternalParentIdNil sets the value for ExternalParentId to be an explicit nil

### UnsetExternalParentId
`func (o *APIComment) UnsetExternalParentId()`

UnsetExternalParentId ensures that no value is present for ExternalParentId, not even an explicit nil
### GetExpireAt

`func (o *APIComment) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *APIComment) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *APIComment) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *APIComment) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### SetExpireAtNil

`func (o *APIComment) SetExpireAtNil(b bool)`

 SetExpireAtNil sets the value for ExpireAt to be an explicit nil

### UnsetExpireAt
`func (o *APIComment) UnsetExpireAt()`

UnsetExpireAt ensures that no value is present for ExpireAt, not even an explicit nil
### GetFeedbackIds

`func (o *APIComment) GetFeedbackIds() []string`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *APIComment) GetFeedbackIdsOk() (*[]string, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *APIComment) SetFeedbackIds(v []string)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *APIComment) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.

### GetFlagCount

`func (o *APIComment) GetFlagCount() int32`

GetFlagCount returns the FlagCount field if non-nil, zero value otherwise.

### GetFlagCountOk

`func (o *APIComment) GetFlagCountOk() (*int32, bool)`

GetFlagCountOk returns a tuple with the FlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCount

`func (o *APIComment) SetFlagCount(v int32)`

SetFlagCount sets FlagCount field to given value.

### HasFlagCount

`func (o *APIComment) HasFlagCount() bool`

HasFlagCount returns a boolean if a field has been set.

### SetFlagCountNil

`func (o *APIComment) SetFlagCountNil(b bool)`

 SetFlagCountNil sets the value for FlagCount to be an explicit nil

### UnsetFlagCount
`func (o *APIComment) UnsetFlagCount()`

UnsetFlagCount ensures that no value is present for FlagCount, not even an explicit nil
### GetFromProductId

`func (o *APIComment) GetFromProductId() int32`

GetFromProductId returns the FromProductId field if non-nil, zero value otherwise.

### GetFromProductIdOk

`func (o *APIComment) GetFromProductIdOk() (*int32, bool)`

GetFromProductIdOk returns a tuple with the FromProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromProductId

`func (o *APIComment) SetFromProductId(v int32)`

SetFromProductId sets FromProductId field to given value.

### HasFromProductId

`func (o *APIComment) HasFromProductId() bool`

HasFromProductId returns a boolean if a field has been set.

### GetHasCode

`func (o *APIComment) GetHasCode() bool`

GetHasCode returns the HasCode field if non-nil, zero value otherwise.

### GetHasCodeOk

`func (o *APIComment) GetHasCodeOk() (*bool, bool)`

GetHasCodeOk returns a tuple with the HasCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCode

`func (o *APIComment) SetHasCode(v bool)`

SetHasCode sets HasCode field to given value.

### HasHasCode

`func (o *APIComment) HasHasCode() bool`

HasHasCode returns a boolean if a field has been set.

### GetHasImages

`func (o *APIComment) GetHasImages() bool`

GetHasImages returns the HasImages field if non-nil, zero value otherwise.

### GetHasImagesOk

`func (o *APIComment) GetHasImagesOk() (*bool, bool)`

GetHasImagesOk returns a tuple with the HasImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImages

`func (o *APIComment) SetHasImages(v bool)`

SetHasImages sets HasImages field to given value.

### HasHasImages

`func (o *APIComment) HasHasImages() bool`

HasHasImages returns a boolean if a field has been set.

### GetHasLinks

`func (o *APIComment) GetHasLinks() bool`

GetHasLinks returns the HasLinks field if non-nil, zero value otherwise.

### GetHasLinksOk

`func (o *APIComment) GetHasLinksOk() (*bool, bool)`

GetHasLinksOk returns a tuple with the HasLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLinks

`func (o *APIComment) SetHasLinks(v bool)`

SetHasLinks sets HasLinks field to given value.

### HasHasLinks

`func (o *APIComment) HasHasLinks() bool`

HasHasLinks returns a boolean if a field has been set.

### GetHashTags

`func (o *APIComment) GetHashTags() []CommentUserHashTagInfo`

GetHashTags returns the HashTags field if non-nil, zero value otherwise.

### GetHashTagsOk

`func (o *APIComment) GetHashTagsOk() (*[]CommentUserHashTagInfo, bool)`

GetHashTagsOk returns a tuple with the HashTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashTags

`func (o *APIComment) SetHashTags(v []CommentUserHashTagInfo)`

SetHashTags sets HashTags field to given value.

### HasHashTags

`func (o *APIComment) HasHashTags() bool`

HasHashTags returns a boolean if a field has been set.

### GetIsByAdmin

`func (o *APIComment) GetIsByAdmin() bool`

GetIsByAdmin returns the IsByAdmin field if non-nil, zero value otherwise.

### GetIsByAdminOk

`func (o *APIComment) GetIsByAdminOk() (*bool, bool)`

GetIsByAdminOk returns a tuple with the IsByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByAdmin

`func (o *APIComment) SetIsByAdmin(v bool)`

SetIsByAdmin sets IsByAdmin field to given value.

### HasIsByAdmin

`func (o *APIComment) HasIsByAdmin() bool`

HasIsByAdmin returns a boolean if a field has been set.

### GetIsByModerator

`func (o *APIComment) GetIsByModerator() bool`

GetIsByModerator returns the IsByModerator field if non-nil, zero value otherwise.

### GetIsByModeratorOk

`func (o *APIComment) GetIsByModeratorOk() (*bool, bool)`

GetIsByModeratorOk returns a tuple with the IsByModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByModerator

`func (o *APIComment) SetIsByModerator(v bool)`

SetIsByModerator sets IsByModerator field to given value.

### HasIsByModerator

`func (o *APIComment) HasIsByModerator() bool`

HasIsByModerator returns a boolean if a field has been set.

### GetIsDeleted

`func (o *APIComment) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *APIComment) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *APIComment) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *APIComment) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsDeletedUser

`func (o *APIComment) GetIsDeletedUser() bool`

GetIsDeletedUser returns the IsDeletedUser field if non-nil, zero value otherwise.

### GetIsDeletedUserOk

`func (o *APIComment) GetIsDeletedUserOk() (*bool, bool)`

GetIsDeletedUserOk returns a tuple with the IsDeletedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletedUser

`func (o *APIComment) SetIsDeletedUser(v bool)`

SetIsDeletedUser sets IsDeletedUser field to given value.

### HasIsDeletedUser

`func (o *APIComment) HasIsDeletedUser() bool`

HasIsDeletedUser returns a boolean if a field has been set.

### GetIsPinned

`func (o *APIComment) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *APIComment) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *APIComment) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.

### HasIsPinned

`func (o *APIComment) HasIsPinned() bool`

HasIsPinned returns a boolean if a field has been set.

### SetIsPinnedNil

`func (o *APIComment) SetIsPinnedNil(b bool)`

 SetIsPinnedNil sets the value for IsPinned to be an explicit nil

### UnsetIsPinned
`func (o *APIComment) UnsetIsPinned()`

UnsetIsPinned ensures that no value is present for IsPinned, not even an explicit nil
### GetIsLocked

`func (o *APIComment) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *APIComment) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *APIComment) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *APIComment) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### SetIsLockedNil

`func (o *APIComment) SetIsLockedNil(b bool)`

 SetIsLockedNil sets the value for IsLocked to be an explicit nil

### UnsetIsLocked
`func (o *APIComment) UnsetIsLocked()`

UnsetIsLocked ensures that no value is present for IsLocked, not even an explicit nil
### GetIsSpam

`func (o *APIComment) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *APIComment) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *APIComment) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *APIComment) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetLocalDateHours

`func (o *APIComment) GetLocalDateHours() int32`

GetLocalDateHours returns the LocalDateHours field if non-nil, zero value otherwise.

### GetLocalDateHoursOk

`func (o *APIComment) GetLocalDateHoursOk() (*int32, bool)`

GetLocalDateHoursOk returns a tuple with the LocalDateHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateHours

`func (o *APIComment) SetLocalDateHours(v int32)`

SetLocalDateHours sets LocalDateHours field to given value.

### HasLocalDateHours

`func (o *APIComment) HasLocalDateHours() bool`

HasLocalDateHours returns a boolean if a field has been set.

### SetLocalDateHoursNil

`func (o *APIComment) SetLocalDateHoursNil(b bool)`

 SetLocalDateHoursNil sets the value for LocalDateHours to be an explicit nil

### UnsetLocalDateHours
`func (o *APIComment) UnsetLocalDateHours()`

UnsetLocalDateHours ensures that no value is present for LocalDateHours, not even an explicit nil
### GetLocalDateString

`func (o *APIComment) GetLocalDateString() string`

GetLocalDateString returns the LocalDateString field if non-nil, zero value otherwise.

### GetLocalDateStringOk

`func (o *APIComment) GetLocalDateStringOk() (*string, bool)`

GetLocalDateStringOk returns a tuple with the LocalDateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateString

`func (o *APIComment) SetLocalDateString(v string)`

SetLocalDateString sets LocalDateString field to given value.

### HasLocalDateString

`func (o *APIComment) HasLocalDateString() bool`

HasLocalDateString returns a boolean if a field has been set.

### SetLocalDateStringNil

`func (o *APIComment) SetLocalDateStringNil(b bool)`

 SetLocalDateStringNil sets the value for LocalDateString to be an explicit nil

### UnsetLocalDateString
`func (o *APIComment) UnsetLocalDateString()`

UnsetLocalDateString ensures that no value is present for LocalDateString, not even an explicit nil
### GetLocale

`func (o *APIComment) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *APIComment) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *APIComment) SetLocale(v string)`

SetLocale sets Locale field to given value.


### SetLocaleNil

`func (o *APIComment) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *APIComment) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetMentions

`func (o *APIComment) GetMentions() []CommentUserMentionInfo`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *APIComment) GetMentionsOk() (*[]CommentUserMentionInfo, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *APIComment) SetMentions(v []CommentUserMentionInfo)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *APIComment) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetMeta

`func (o *APIComment) GetMeta() FCommentMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *APIComment) GetMetaOk() (*FCommentMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *APIComment) SetMeta(v FCommentMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *APIComment) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *APIComment) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *APIComment) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetModerationGroupIds

`func (o *APIComment) GetModerationGroupIds() []string`

GetModerationGroupIds returns the ModerationGroupIds field if non-nil, zero value otherwise.

### GetModerationGroupIdsOk

`func (o *APIComment) GetModerationGroupIdsOk() (*[]string, bool)`

GetModerationGroupIdsOk returns a tuple with the ModerationGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationGroupIds

`func (o *APIComment) SetModerationGroupIds(v []string)`

SetModerationGroupIds sets ModerationGroupIds field to given value.

### HasModerationGroupIds

`func (o *APIComment) HasModerationGroupIds() bool`

HasModerationGroupIds returns a boolean if a field has been set.

### SetModerationGroupIdsNil

`func (o *APIComment) SetModerationGroupIdsNil(b bool)`

 SetModerationGroupIdsNil sets the value for ModerationGroupIds to be an explicit nil

### UnsetModerationGroupIds
`func (o *APIComment) UnsetModerationGroupIds()`

UnsetModerationGroupIds ensures that no value is present for ModerationGroupIds, not even an explicit nil
### GetNotificationSentForParent

`func (o *APIComment) GetNotificationSentForParent() bool`

GetNotificationSentForParent returns the NotificationSentForParent field if non-nil, zero value otherwise.

### GetNotificationSentForParentOk

`func (o *APIComment) GetNotificationSentForParentOk() (*bool, bool)`

GetNotificationSentForParentOk returns a tuple with the NotificationSentForParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSentForParent

`func (o *APIComment) SetNotificationSentForParent(v bool)`

SetNotificationSentForParent sets NotificationSentForParent field to given value.

### HasNotificationSentForParent

`func (o *APIComment) HasNotificationSentForParent() bool`

HasNotificationSentForParent returns a boolean if a field has been set.

### GetNotificationSentForParentTenant

`func (o *APIComment) GetNotificationSentForParentTenant() bool`

GetNotificationSentForParentTenant returns the NotificationSentForParentTenant field if non-nil, zero value otherwise.

### GetNotificationSentForParentTenantOk

`func (o *APIComment) GetNotificationSentForParentTenantOk() (*bool, bool)`

GetNotificationSentForParentTenantOk returns a tuple with the NotificationSentForParentTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSentForParentTenant

`func (o *APIComment) SetNotificationSentForParentTenant(v bool)`

SetNotificationSentForParentTenant sets NotificationSentForParentTenant field to given value.

### HasNotificationSentForParentTenant

`func (o *APIComment) HasNotificationSentForParentTenant() bool`

HasNotificationSentForParentTenant returns a boolean if a field has been set.

### GetPageTitle

`func (o *APIComment) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *APIComment) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *APIComment) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *APIComment) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### SetPageTitleNil

`func (o *APIComment) SetPageTitleNil(b bool)`

 SetPageTitleNil sets the value for PageTitle to be an explicit nil

### UnsetPageTitle
`func (o *APIComment) UnsetPageTitle()`

UnsetPageTitle ensures that no value is present for PageTitle, not even an explicit nil
### GetParentId

`func (o *APIComment) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *APIComment) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *APIComment) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *APIComment) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *APIComment) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *APIComment) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetRating

`func (o *APIComment) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *APIComment) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *APIComment) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *APIComment) HasRating() bool`

HasRating returns a boolean if a field has been set.

### SetRatingNil

`func (o *APIComment) SetRatingNil(b bool)`

 SetRatingNil sets the value for Rating to be an explicit nil

### UnsetRating
`func (o *APIComment) UnsetRating()`

UnsetRating ensures that no value is present for Rating, not even an explicit nil
### GetReviewed

`func (o *APIComment) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *APIComment) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *APIComment) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *APIComment) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetTenantId

`func (o *APIComment) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *APIComment) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *APIComment) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUrl

`func (o *APIComment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *APIComment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *APIComment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUrlId

`func (o *APIComment) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *APIComment) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *APIComment) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetUrlIdRaw

`func (o *APIComment) GetUrlIdRaw() string`

GetUrlIdRaw returns the UrlIdRaw field if non-nil, zero value otherwise.

### GetUrlIdRawOk

`func (o *APIComment) GetUrlIdRawOk() (*string, bool)`

GetUrlIdRawOk returns a tuple with the UrlIdRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdRaw

`func (o *APIComment) SetUrlIdRaw(v string)`

SetUrlIdRaw sets UrlIdRaw field to given value.

### HasUrlIdRaw

`func (o *APIComment) HasUrlIdRaw() bool`

HasUrlIdRaw returns a boolean if a field has been set.

### GetUserId

`func (o *APIComment) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *APIComment) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *APIComment) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *APIComment) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *APIComment) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *APIComment) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetVerified

`func (o *APIComment) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *APIComment) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *APIComment) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetVerifiedDate

`func (o *APIComment) GetVerifiedDate() time.Time`

GetVerifiedDate returns the VerifiedDate field if non-nil, zero value otherwise.

### GetVerifiedDateOk

`func (o *APIComment) GetVerifiedDateOk() (*time.Time, bool)`

GetVerifiedDateOk returns a tuple with the VerifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedDate

`func (o *APIComment) SetVerifiedDate(v time.Time)`

SetVerifiedDate sets VerifiedDate field to given value.

### HasVerifiedDate

`func (o *APIComment) HasVerifiedDate() bool`

HasVerifiedDate returns a boolean if a field has been set.

### SetVerifiedDateNil

`func (o *APIComment) SetVerifiedDateNil(b bool)`

 SetVerifiedDateNil sets the value for VerifiedDate to be an explicit nil

### UnsetVerifiedDate
`func (o *APIComment) UnsetVerifiedDate()`

UnsetVerifiedDate ensures that no value is present for VerifiedDate, not even an explicit nil
### GetVotes

`func (o *APIComment) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *APIComment) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *APIComment) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *APIComment) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### SetVotesNil

`func (o *APIComment) SetVotesNil(b bool)`

 SetVotesNil sets the value for Votes to be an explicit nil

### UnsetVotes
`func (o *APIComment) UnsetVotes()`

UnsetVotes ensures that no value is present for Votes, not even an explicit nil
### GetVotesDown

`func (o *APIComment) GetVotesDown() int32`

GetVotesDown returns the VotesDown field if non-nil, zero value otherwise.

### GetVotesDownOk

`func (o *APIComment) GetVotesDownOk() (*int32, bool)`

GetVotesDownOk returns a tuple with the VotesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDown

`func (o *APIComment) SetVotesDown(v int32)`

SetVotesDown sets VotesDown field to given value.

### HasVotesDown

`func (o *APIComment) HasVotesDown() bool`

HasVotesDown returns a boolean if a field has been set.

### SetVotesDownNil

`func (o *APIComment) SetVotesDownNil(b bool)`

 SetVotesDownNil sets the value for VotesDown to be an explicit nil

### UnsetVotesDown
`func (o *APIComment) UnsetVotesDown()`

UnsetVotesDown ensures that no value is present for VotesDown, not even an explicit nil
### GetVotesUp

`func (o *APIComment) GetVotesUp() int32`

GetVotesUp returns the VotesUp field if non-nil, zero value otherwise.

### GetVotesUpOk

`func (o *APIComment) GetVotesUpOk() (*int32, bool)`

GetVotesUpOk returns a tuple with the VotesUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUp

`func (o *APIComment) SetVotesUp(v int32)`

SetVotesUp sets VotesUp field to given value.

### HasVotesUp

`func (o *APIComment) HasVotesUp() bool`

HasVotesUp returns a boolean if a field has been set.

### SetVotesUpNil

`func (o *APIComment) SetVotesUpNil(b bool)`

 SetVotesUpNil sets the value for VotesUp to be an explicit nil

### UnsetVotesUp
`func (o *APIComment) UnsetVotesUp()`

UnsetVotesUp ensures that no value is present for VotesUp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


