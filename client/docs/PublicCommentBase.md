# PublicCommentBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 
**CommenterName** | **string** |  | 
**CommenterLink** | Pointer to **NullableString** |  | [optional] 
**CommentHTML** | **string** |  | 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Date** | **NullableTime** |  | 
**Votes** | Pointer to **NullableInt32** |  | [optional] 
**VotesUp** | Pointer to **NullableInt32** |  | [optional] 
**VotesDown** | Pointer to **NullableInt32** |  | [optional] 
**Verified** | **bool** |  | 
**AvatarSrc** | Pointer to **NullableString** |  | [optional] 
**HasImages** | Pointer to **bool** |  | [optional] 
**IsByAdmin** | Pointer to **bool** |  | [optional] 
**IsByModerator** | Pointer to **bool** |  | [optional] 
**IsPinned** | Pointer to **NullableBool** |  | [optional] 
**IsLocked** | Pointer to **NullableBool** |  | [optional] 
**DisplayLabel** | Pointer to **NullableString** |  | [optional] 
**Rating** | Pointer to **NullableFloat64** |  | [optional] 
**Badges** | Pointer to [**[]CommentUserBadgeInfo**](CommentUserBadgeInfo.md) |  | [optional] 
**ViewCount** | Pointer to **NullableInt64** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**IsDeletedUser** | Pointer to **bool** |  | [optional] 
**IsSpam** | Pointer to **bool** |  | [optional] 
**AnonUserId** | Pointer to **NullableString** |  | [optional] 
**FeedbackIds** | Pointer to **[]string** |  | [optional] 
**RequiresVerification** | Pointer to **bool** |  | [optional] 
**EditKey** | Pointer to **string** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 

## Methods

### NewPublicCommentBase

`func NewPublicCommentBase(id string, commenterName string, commentHTML string, date NullableTime, verified bool, ) *PublicCommentBase`

NewPublicCommentBase instantiates a new PublicCommentBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCommentBaseWithDefaults

`func NewPublicCommentBaseWithDefaults() *PublicCommentBase`

NewPublicCommentBaseWithDefaults instantiates a new PublicCommentBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicCommentBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicCommentBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicCommentBase) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *PublicCommentBase) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PublicCommentBase) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PublicCommentBase) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PublicCommentBase) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *PublicCommentBase) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *PublicCommentBase) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetCommenterName

`func (o *PublicCommentBase) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *PublicCommentBase) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *PublicCommentBase) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetCommenterLink

`func (o *PublicCommentBase) GetCommenterLink() string`

GetCommenterLink returns the CommenterLink field if non-nil, zero value otherwise.

### GetCommenterLinkOk

`func (o *PublicCommentBase) GetCommenterLinkOk() (*string, bool)`

GetCommenterLinkOk returns a tuple with the CommenterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterLink

`func (o *PublicCommentBase) SetCommenterLink(v string)`

SetCommenterLink sets CommenterLink field to given value.

### HasCommenterLink

`func (o *PublicCommentBase) HasCommenterLink() bool`

HasCommenterLink returns a boolean if a field has been set.

### SetCommenterLinkNil

`func (o *PublicCommentBase) SetCommenterLinkNil(b bool)`

 SetCommenterLinkNil sets the value for CommenterLink to be an explicit nil

### UnsetCommenterLink
`func (o *PublicCommentBase) UnsetCommenterLink()`

UnsetCommenterLink ensures that no value is present for CommenterLink, not even an explicit nil
### GetCommentHTML

`func (o *PublicCommentBase) GetCommentHTML() string`

GetCommentHTML returns the CommentHTML field if non-nil, zero value otherwise.

### GetCommentHTMLOk

`func (o *PublicCommentBase) GetCommentHTMLOk() (*string, bool)`

GetCommentHTMLOk returns a tuple with the CommentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentHTML

`func (o *PublicCommentBase) SetCommentHTML(v string)`

SetCommentHTML sets CommentHTML field to given value.


### GetParentId

`func (o *PublicCommentBase) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PublicCommentBase) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PublicCommentBase) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PublicCommentBase) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *PublicCommentBase) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *PublicCommentBase) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetDate

`func (o *PublicCommentBase) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PublicCommentBase) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PublicCommentBase) SetDate(v time.Time)`

SetDate sets Date field to given value.


### SetDateNil

`func (o *PublicCommentBase) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *PublicCommentBase) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetVotes

`func (o *PublicCommentBase) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *PublicCommentBase) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *PublicCommentBase) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *PublicCommentBase) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### SetVotesNil

`func (o *PublicCommentBase) SetVotesNil(b bool)`

 SetVotesNil sets the value for Votes to be an explicit nil

### UnsetVotes
`func (o *PublicCommentBase) UnsetVotes()`

UnsetVotes ensures that no value is present for Votes, not even an explicit nil
### GetVotesUp

`func (o *PublicCommentBase) GetVotesUp() int32`

GetVotesUp returns the VotesUp field if non-nil, zero value otherwise.

### GetVotesUpOk

`func (o *PublicCommentBase) GetVotesUpOk() (*int32, bool)`

GetVotesUpOk returns a tuple with the VotesUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUp

`func (o *PublicCommentBase) SetVotesUp(v int32)`

SetVotesUp sets VotesUp field to given value.

### HasVotesUp

`func (o *PublicCommentBase) HasVotesUp() bool`

HasVotesUp returns a boolean if a field has been set.

### SetVotesUpNil

`func (o *PublicCommentBase) SetVotesUpNil(b bool)`

 SetVotesUpNil sets the value for VotesUp to be an explicit nil

### UnsetVotesUp
`func (o *PublicCommentBase) UnsetVotesUp()`

UnsetVotesUp ensures that no value is present for VotesUp, not even an explicit nil
### GetVotesDown

`func (o *PublicCommentBase) GetVotesDown() int32`

GetVotesDown returns the VotesDown field if non-nil, zero value otherwise.

### GetVotesDownOk

`func (o *PublicCommentBase) GetVotesDownOk() (*int32, bool)`

GetVotesDownOk returns a tuple with the VotesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDown

`func (o *PublicCommentBase) SetVotesDown(v int32)`

SetVotesDown sets VotesDown field to given value.

### HasVotesDown

`func (o *PublicCommentBase) HasVotesDown() bool`

HasVotesDown returns a boolean if a field has been set.

### SetVotesDownNil

`func (o *PublicCommentBase) SetVotesDownNil(b bool)`

 SetVotesDownNil sets the value for VotesDown to be an explicit nil

### UnsetVotesDown
`func (o *PublicCommentBase) UnsetVotesDown()`

UnsetVotesDown ensures that no value is present for VotesDown, not even an explicit nil
### GetVerified

`func (o *PublicCommentBase) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PublicCommentBase) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PublicCommentBase) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetAvatarSrc

`func (o *PublicCommentBase) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *PublicCommentBase) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *PublicCommentBase) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *PublicCommentBase) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### SetAvatarSrcNil

`func (o *PublicCommentBase) SetAvatarSrcNil(b bool)`

 SetAvatarSrcNil sets the value for AvatarSrc to be an explicit nil

### UnsetAvatarSrc
`func (o *PublicCommentBase) UnsetAvatarSrc()`

UnsetAvatarSrc ensures that no value is present for AvatarSrc, not even an explicit nil
### GetHasImages

`func (o *PublicCommentBase) GetHasImages() bool`

GetHasImages returns the HasImages field if non-nil, zero value otherwise.

### GetHasImagesOk

`func (o *PublicCommentBase) GetHasImagesOk() (*bool, bool)`

GetHasImagesOk returns a tuple with the HasImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImages

`func (o *PublicCommentBase) SetHasImages(v bool)`

SetHasImages sets HasImages field to given value.

### HasHasImages

`func (o *PublicCommentBase) HasHasImages() bool`

HasHasImages returns a boolean if a field has been set.

### GetIsByAdmin

`func (o *PublicCommentBase) GetIsByAdmin() bool`

GetIsByAdmin returns the IsByAdmin field if non-nil, zero value otherwise.

### GetIsByAdminOk

`func (o *PublicCommentBase) GetIsByAdminOk() (*bool, bool)`

GetIsByAdminOk returns a tuple with the IsByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByAdmin

`func (o *PublicCommentBase) SetIsByAdmin(v bool)`

SetIsByAdmin sets IsByAdmin field to given value.

### HasIsByAdmin

`func (o *PublicCommentBase) HasIsByAdmin() bool`

HasIsByAdmin returns a boolean if a field has been set.

### GetIsByModerator

`func (o *PublicCommentBase) GetIsByModerator() bool`

GetIsByModerator returns the IsByModerator field if non-nil, zero value otherwise.

### GetIsByModeratorOk

`func (o *PublicCommentBase) GetIsByModeratorOk() (*bool, bool)`

GetIsByModeratorOk returns a tuple with the IsByModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByModerator

`func (o *PublicCommentBase) SetIsByModerator(v bool)`

SetIsByModerator sets IsByModerator field to given value.

### HasIsByModerator

`func (o *PublicCommentBase) HasIsByModerator() bool`

HasIsByModerator returns a boolean if a field has been set.

### GetIsPinned

`func (o *PublicCommentBase) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *PublicCommentBase) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *PublicCommentBase) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.

### HasIsPinned

`func (o *PublicCommentBase) HasIsPinned() bool`

HasIsPinned returns a boolean if a field has been set.

### SetIsPinnedNil

`func (o *PublicCommentBase) SetIsPinnedNil(b bool)`

 SetIsPinnedNil sets the value for IsPinned to be an explicit nil

### UnsetIsPinned
`func (o *PublicCommentBase) UnsetIsPinned()`

UnsetIsPinned ensures that no value is present for IsPinned, not even an explicit nil
### GetIsLocked

`func (o *PublicCommentBase) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *PublicCommentBase) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *PublicCommentBase) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *PublicCommentBase) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### SetIsLockedNil

`func (o *PublicCommentBase) SetIsLockedNil(b bool)`

 SetIsLockedNil sets the value for IsLocked to be an explicit nil

### UnsetIsLocked
`func (o *PublicCommentBase) UnsetIsLocked()`

UnsetIsLocked ensures that no value is present for IsLocked, not even an explicit nil
### GetDisplayLabel

`func (o *PublicCommentBase) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *PublicCommentBase) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *PublicCommentBase) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *PublicCommentBase) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### SetDisplayLabelNil

`func (o *PublicCommentBase) SetDisplayLabelNil(b bool)`

 SetDisplayLabelNil sets the value for DisplayLabel to be an explicit nil

### UnsetDisplayLabel
`func (o *PublicCommentBase) UnsetDisplayLabel()`

UnsetDisplayLabel ensures that no value is present for DisplayLabel, not even an explicit nil
### GetRating

`func (o *PublicCommentBase) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *PublicCommentBase) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *PublicCommentBase) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *PublicCommentBase) HasRating() bool`

HasRating returns a boolean if a field has been set.

### SetRatingNil

`func (o *PublicCommentBase) SetRatingNil(b bool)`

 SetRatingNil sets the value for Rating to be an explicit nil

### UnsetRating
`func (o *PublicCommentBase) UnsetRating()`

UnsetRating ensures that no value is present for Rating, not even an explicit nil
### GetBadges

`func (o *PublicCommentBase) GetBadges() []CommentUserBadgeInfo`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *PublicCommentBase) GetBadgesOk() (*[]CommentUserBadgeInfo, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *PublicCommentBase) SetBadges(v []CommentUserBadgeInfo)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *PublicCommentBase) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### SetBadgesNil

`func (o *PublicCommentBase) SetBadgesNil(b bool)`

 SetBadgesNil sets the value for Badges to be an explicit nil

### UnsetBadges
`func (o *PublicCommentBase) UnsetBadges()`

UnsetBadges ensures that no value is present for Badges, not even an explicit nil
### GetViewCount

`func (o *PublicCommentBase) GetViewCount() int64`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *PublicCommentBase) GetViewCountOk() (*int64, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *PublicCommentBase) SetViewCount(v int64)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *PublicCommentBase) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.

### SetViewCountNil

`func (o *PublicCommentBase) SetViewCountNil(b bool)`

 SetViewCountNil sets the value for ViewCount to be an explicit nil

### UnsetViewCount
`func (o *PublicCommentBase) UnsetViewCount()`

UnsetViewCount ensures that no value is present for ViewCount, not even an explicit nil
### GetIsDeleted

`func (o *PublicCommentBase) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PublicCommentBase) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PublicCommentBase) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PublicCommentBase) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsDeletedUser

`func (o *PublicCommentBase) GetIsDeletedUser() bool`

GetIsDeletedUser returns the IsDeletedUser field if non-nil, zero value otherwise.

### GetIsDeletedUserOk

`func (o *PublicCommentBase) GetIsDeletedUserOk() (*bool, bool)`

GetIsDeletedUserOk returns a tuple with the IsDeletedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletedUser

`func (o *PublicCommentBase) SetIsDeletedUser(v bool)`

SetIsDeletedUser sets IsDeletedUser field to given value.

### HasIsDeletedUser

`func (o *PublicCommentBase) HasIsDeletedUser() bool`

HasIsDeletedUser returns a boolean if a field has been set.

### GetIsSpam

`func (o *PublicCommentBase) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *PublicCommentBase) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *PublicCommentBase) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *PublicCommentBase) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetAnonUserId

`func (o *PublicCommentBase) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *PublicCommentBase) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *PublicCommentBase) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.

### HasAnonUserId

`func (o *PublicCommentBase) HasAnonUserId() bool`

HasAnonUserId returns a boolean if a field has been set.

### SetAnonUserIdNil

`func (o *PublicCommentBase) SetAnonUserIdNil(b bool)`

 SetAnonUserIdNil sets the value for AnonUserId to be an explicit nil

### UnsetAnonUserId
`func (o *PublicCommentBase) UnsetAnonUserId()`

UnsetAnonUserId ensures that no value is present for AnonUserId, not even an explicit nil
### GetFeedbackIds

`func (o *PublicCommentBase) GetFeedbackIds() []string`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *PublicCommentBase) GetFeedbackIdsOk() (*[]string, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *PublicCommentBase) SetFeedbackIds(v []string)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *PublicCommentBase) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.

### GetRequiresVerification

`func (o *PublicCommentBase) GetRequiresVerification() bool`

GetRequiresVerification returns the RequiresVerification field if non-nil, zero value otherwise.

### GetRequiresVerificationOk

`func (o *PublicCommentBase) GetRequiresVerificationOk() (*bool, bool)`

GetRequiresVerificationOk returns a tuple with the RequiresVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresVerification

`func (o *PublicCommentBase) SetRequiresVerification(v bool)`

SetRequiresVerification sets RequiresVerification field to given value.

### HasRequiresVerification

`func (o *PublicCommentBase) HasRequiresVerification() bool`

HasRequiresVerification returns a boolean if a field has been set.

### GetEditKey

`func (o *PublicCommentBase) GetEditKey() string`

GetEditKey returns the EditKey field if non-nil, zero value otherwise.

### GetEditKeyOk

`func (o *PublicCommentBase) GetEditKeyOk() (*string, bool)`

GetEditKeyOk returns a tuple with the EditKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditKey

`func (o *PublicCommentBase) SetEditKey(v string)`

SetEditKey sets EditKey field to given value.

### HasEditKey

`func (o *PublicCommentBase) HasEditKey() bool`

HasEditKey returns a boolean if a field has been set.

### GetApproved

`func (o *PublicCommentBase) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *PublicCommentBase) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *PublicCommentBase) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *PublicCommentBase) HasApproved() bool`

HasApproved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


