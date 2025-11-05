# PublicComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** |  | 
**Id** | **string** |  | 
**UserId** | Pointer to **string** |  | [optional] 
**AnonUserId** | Pointer to **string** |  | [optional] 
**CommenterName** | **string** |  | 
**CommenterLink** | Pointer to **string** |  | [optional] 
**CommentHTML** | **string** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**Votes** | Pointer to **int32** |  | [optional] 
**VotesUp** | Pointer to **int32** |  | [optional] 
**VotesDown** | Pointer to **int32** |  | [optional] 
**Verified** | **bool** |  | 
**AvatarSrc** | Pointer to **string** |  | [optional] 
**IsSpam** | Pointer to **bool** |  | [optional] 
**HasImages** | Pointer to **bool** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**IsDeletedUser** | Pointer to **bool** |  | [optional] 
**IsByAdmin** | Pointer to **bool** |  | [optional] 
**IsByModerator** | Pointer to **bool** |  | [optional] 
**IsPinned** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**Rating** | Pointer to **float64** |  | [optional] 
**DisplayLabel** | Pointer to **string** |  | [optional] 
**Badges** | Pointer to [**[]CommentUserBadgeInfo**](CommentUserBadgeInfo.md) |  | [optional] 
**FeedbackIds** | Pointer to **[]string** |  | [optional] 
**ViewCount** | Pointer to **int64** |  | [optional] 
**RequiresVerification** | Pointer to **bool** |  | [optional] 
**EditKey** | Pointer to **string** |  | [optional] 
**IsUnread** | Pointer to **bool** |  | [optional] 
**MyVoteId** | Pointer to **string** |  | [optional] 
**IsVotedDown** | Pointer to **bool** |  | [optional] 
**IsVotedUp** | Pointer to **bool** |  | [optional] 
**HasChildren** | Pointer to **bool** | This is always set when asTree&#x3D;true | [optional] 
**NestedChildrenCount** | Pointer to **int32** | The total nested child count included in this response (may be more available w/ pagination) Only set with asTree&#x3D;true, otherwise this will be null. | [optional] 
**ChildCount** | Pointer to **int32** | You must ask the API to count children (with asTree&#x3D;true&amp;countChildren&#x3D;true), otherwise this will be null. This will be the complete direct child count, whereas children may only contain a subset based on pagination. | [optional] 
**Children** | Pointer to [**[]PublicComment**](PublicComment.md) |  | [optional] 
**IsFlagged** | Pointer to **bool** |  | [optional] 
**IsBlocked** | Pointer to **bool** |  | [optional] 

## Methods

### NewPublicComment

`func NewPublicComment(date time.Time, id string, commenterName string, commentHTML string, verified bool, ) *PublicComment`

NewPublicComment instantiates a new PublicComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCommentWithDefaults

`func NewPublicCommentWithDefaults() *PublicComment`

NewPublicCommentWithDefaults instantiates a new PublicComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *PublicComment) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PublicComment) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PublicComment) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetId

`func (o *PublicComment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicComment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicComment) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *PublicComment) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PublicComment) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PublicComment) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PublicComment) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAnonUserId

`func (o *PublicComment) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *PublicComment) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *PublicComment) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.

### HasAnonUserId

`func (o *PublicComment) HasAnonUserId() bool`

HasAnonUserId returns a boolean if a field has been set.

### GetCommenterName

`func (o *PublicComment) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *PublicComment) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *PublicComment) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetCommenterLink

`func (o *PublicComment) GetCommenterLink() string`

GetCommenterLink returns the CommenterLink field if non-nil, zero value otherwise.

### GetCommenterLinkOk

`func (o *PublicComment) GetCommenterLinkOk() (*string, bool)`

GetCommenterLinkOk returns a tuple with the CommenterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterLink

`func (o *PublicComment) SetCommenterLink(v string)`

SetCommenterLink sets CommenterLink field to given value.

### HasCommenterLink

`func (o *PublicComment) HasCommenterLink() bool`

HasCommenterLink returns a boolean if a field has been set.

### GetCommentHTML

`func (o *PublicComment) GetCommentHTML() string`

GetCommentHTML returns the CommentHTML field if non-nil, zero value otherwise.

### GetCommentHTMLOk

`func (o *PublicComment) GetCommentHTMLOk() (*string, bool)`

GetCommentHTMLOk returns a tuple with the CommentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentHTML

`func (o *PublicComment) SetCommentHTML(v string)`

SetCommentHTML sets CommentHTML field to given value.


### GetParentId

`func (o *PublicComment) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PublicComment) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PublicComment) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PublicComment) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetVotes

`func (o *PublicComment) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *PublicComment) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *PublicComment) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *PublicComment) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetVotesUp

`func (o *PublicComment) GetVotesUp() int32`

GetVotesUp returns the VotesUp field if non-nil, zero value otherwise.

### GetVotesUpOk

`func (o *PublicComment) GetVotesUpOk() (*int32, bool)`

GetVotesUpOk returns a tuple with the VotesUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUp

`func (o *PublicComment) SetVotesUp(v int32)`

SetVotesUp sets VotesUp field to given value.

### HasVotesUp

`func (o *PublicComment) HasVotesUp() bool`

HasVotesUp returns a boolean if a field has been set.

### GetVotesDown

`func (o *PublicComment) GetVotesDown() int32`

GetVotesDown returns the VotesDown field if non-nil, zero value otherwise.

### GetVotesDownOk

`func (o *PublicComment) GetVotesDownOk() (*int32, bool)`

GetVotesDownOk returns a tuple with the VotesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDown

`func (o *PublicComment) SetVotesDown(v int32)`

SetVotesDown sets VotesDown field to given value.

### HasVotesDown

`func (o *PublicComment) HasVotesDown() bool`

HasVotesDown returns a boolean if a field has been set.

### GetVerified

`func (o *PublicComment) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PublicComment) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PublicComment) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetAvatarSrc

`func (o *PublicComment) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *PublicComment) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *PublicComment) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *PublicComment) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### GetIsSpam

`func (o *PublicComment) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *PublicComment) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *PublicComment) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *PublicComment) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetHasImages

`func (o *PublicComment) GetHasImages() bool`

GetHasImages returns the HasImages field if non-nil, zero value otherwise.

### GetHasImagesOk

`func (o *PublicComment) GetHasImagesOk() (*bool, bool)`

GetHasImagesOk returns a tuple with the HasImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImages

`func (o *PublicComment) SetHasImages(v bool)`

SetHasImages sets HasImages field to given value.

### HasHasImages

`func (o *PublicComment) HasHasImages() bool`

HasHasImages returns a boolean if a field has been set.

### GetIsDeleted

`func (o *PublicComment) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PublicComment) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PublicComment) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PublicComment) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsDeletedUser

`func (o *PublicComment) GetIsDeletedUser() bool`

GetIsDeletedUser returns the IsDeletedUser field if non-nil, zero value otherwise.

### GetIsDeletedUserOk

`func (o *PublicComment) GetIsDeletedUserOk() (*bool, bool)`

GetIsDeletedUserOk returns a tuple with the IsDeletedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletedUser

`func (o *PublicComment) SetIsDeletedUser(v bool)`

SetIsDeletedUser sets IsDeletedUser field to given value.

### HasIsDeletedUser

`func (o *PublicComment) HasIsDeletedUser() bool`

HasIsDeletedUser returns a boolean if a field has been set.

### GetIsByAdmin

`func (o *PublicComment) GetIsByAdmin() bool`

GetIsByAdmin returns the IsByAdmin field if non-nil, zero value otherwise.

### GetIsByAdminOk

`func (o *PublicComment) GetIsByAdminOk() (*bool, bool)`

GetIsByAdminOk returns a tuple with the IsByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByAdmin

`func (o *PublicComment) SetIsByAdmin(v bool)`

SetIsByAdmin sets IsByAdmin field to given value.

### HasIsByAdmin

`func (o *PublicComment) HasIsByAdmin() bool`

HasIsByAdmin returns a boolean if a field has been set.

### GetIsByModerator

`func (o *PublicComment) GetIsByModerator() bool`

GetIsByModerator returns the IsByModerator field if non-nil, zero value otherwise.

### GetIsByModeratorOk

`func (o *PublicComment) GetIsByModeratorOk() (*bool, bool)`

GetIsByModeratorOk returns a tuple with the IsByModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByModerator

`func (o *PublicComment) SetIsByModerator(v bool)`

SetIsByModerator sets IsByModerator field to given value.

### HasIsByModerator

`func (o *PublicComment) HasIsByModerator() bool`

HasIsByModerator returns a boolean if a field has been set.

### GetIsPinned

`func (o *PublicComment) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *PublicComment) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *PublicComment) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.

### HasIsPinned

`func (o *PublicComment) HasIsPinned() bool`

HasIsPinned returns a boolean if a field has been set.

### GetIsLocked

`func (o *PublicComment) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *PublicComment) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *PublicComment) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *PublicComment) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetRating

`func (o *PublicComment) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *PublicComment) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *PublicComment) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *PublicComment) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *PublicComment) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *PublicComment) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *PublicComment) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *PublicComment) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetBadges

`func (o *PublicComment) GetBadges() []CommentUserBadgeInfo`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *PublicComment) GetBadgesOk() (*[]CommentUserBadgeInfo, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *PublicComment) SetBadges(v []CommentUserBadgeInfo)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *PublicComment) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetFeedbackIds

`func (o *PublicComment) GetFeedbackIds() []string`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *PublicComment) GetFeedbackIdsOk() (*[]string, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *PublicComment) SetFeedbackIds(v []string)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *PublicComment) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.

### GetViewCount

`func (o *PublicComment) GetViewCount() int64`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *PublicComment) GetViewCountOk() (*int64, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *PublicComment) SetViewCount(v int64)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *PublicComment) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.

### GetRequiresVerification

`func (o *PublicComment) GetRequiresVerification() bool`

GetRequiresVerification returns the RequiresVerification field if non-nil, zero value otherwise.

### GetRequiresVerificationOk

`func (o *PublicComment) GetRequiresVerificationOk() (*bool, bool)`

GetRequiresVerificationOk returns a tuple with the RequiresVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresVerification

`func (o *PublicComment) SetRequiresVerification(v bool)`

SetRequiresVerification sets RequiresVerification field to given value.

### HasRequiresVerification

`func (o *PublicComment) HasRequiresVerification() bool`

HasRequiresVerification returns a boolean if a field has been set.

### GetEditKey

`func (o *PublicComment) GetEditKey() string`

GetEditKey returns the EditKey field if non-nil, zero value otherwise.

### GetEditKeyOk

`func (o *PublicComment) GetEditKeyOk() (*string, bool)`

GetEditKeyOk returns a tuple with the EditKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditKey

`func (o *PublicComment) SetEditKey(v string)`

SetEditKey sets EditKey field to given value.

### HasEditKey

`func (o *PublicComment) HasEditKey() bool`

HasEditKey returns a boolean if a field has been set.

### GetIsUnread

`func (o *PublicComment) GetIsUnread() bool`

GetIsUnread returns the IsUnread field if non-nil, zero value otherwise.

### GetIsUnreadOk

`func (o *PublicComment) GetIsUnreadOk() (*bool, bool)`

GetIsUnreadOk returns a tuple with the IsUnread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnread

`func (o *PublicComment) SetIsUnread(v bool)`

SetIsUnread sets IsUnread field to given value.

### HasIsUnread

`func (o *PublicComment) HasIsUnread() bool`

HasIsUnread returns a boolean if a field has been set.

### GetMyVoteId

`func (o *PublicComment) GetMyVoteId() string`

GetMyVoteId returns the MyVoteId field if non-nil, zero value otherwise.

### GetMyVoteIdOk

`func (o *PublicComment) GetMyVoteIdOk() (*string, bool)`

GetMyVoteIdOk returns a tuple with the MyVoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyVoteId

`func (o *PublicComment) SetMyVoteId(v string)`

SetMyVoteId sets MyVoteId field to given value.

### HasMyVoteId

`func (o *PublicComment) HasMyVoteId() bool`

HasMyVoteId returns a boolean if a field has been set.

### GetIsVotedDown

`func (o *PublicComment) GetIsVotedDown() bool`

GetIsVotedDown returns the IsVotedDown field if non-nil, zero value otherwise.

### GetIsVotedDownOk

`func (o *PublicComment) GetIsVotedDownOk() (*bool, bool)`

GetIsVotedDownOk returns a tuple with the IsVotedDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVotedDown

`func (o *PublicComment) SetIsVotedDown(v bool)`

SetIsVotedDown sets IsVotedDown field to given value.

### HasIsVotedDown

`func (o *PublicComment) HasIsVotedDown() bool`

HasIsVotedDown returns a boolean if a field has been set.

### GetIsVotedUp

`func (o *PublicComment) GetIsVotedUp() bool`

GetIsVotedUp returns the IsVotedUp field if non-nil, zero value otherwise.

### GetIsVotedUpOk

`func (o *PublicComment) GetIsVotedUpOk() (*bool, bool)`

GetIsVotedUpOk returns a tuple with the IsVotedUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVotedUp

`func (o *PublicComment) SetIsVotedUp(v bool)`

SetIsVotedUp sets IsVotedUp field to given value.

### HasIsVotedUp

`func (o *PublicComment) HasIsVotedUp() bool`

HasIsVotedUp returns a boolean if a field has been set.

### GetHasChildren

`func (o *PublicComment) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *PublicComment) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *PublicComment) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.

### HasHasChildren

`func (o *PublicComment) HasHasChildren() bool`

HasHasChildren returns a boolean if a field has been set.

### GetNestedChildrenCount

`func (o *PublicComment) GetNestedChildrenCount() int32`

GetNestedChildrenCount returns the NestedChildrenCount field if non-nil, zero value otherwise.

### GetNestedChildrenCountOk

`func (o *PublicComment) GetNestedChildrenCountOk() (*int32, bool)`

GetNestedChildrenCountOk returns a tuple with the NestedChildrenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedChildrenCount

`func (o *PublicComment) SetNestedChildrenCount(v int32)`

SetNestedChildrenCount sets NestedChildrenCount field to given value.

### HasNestedChildrenCount

`func (o *PublicComment) HasNestedChildrenCount() bool`

HasNestedChildrenCount returns a boolean if a field has been set.

### GetChildCount

`func (o *PublicComment) GetChildCount() int32`

GetChildCount returns the ChildCount field if non-nil, zero value otherwise.

### GetChildCountOk

`func (o *PublicComment) GetChildCountOk() (*int32, bool)`

GetChildCountOk returns a tuple with the ChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCount

`func (o *PublicComment) SetChildCount(v int32)`

SetChildCount sets ChildCount field to given value.

### HasChildCount

`func (o *PublicComment) HasChildCount() bool`

HasChildCount returns a boolean if a field has been set.

### GetChildren

`func (o *PublicComment) GetChildren() []PublicComment`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *PublicComment) GetChildrenOk() (*[]PublicComment, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *PublicComment) SetChildren(v []PublicComment)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *PublicComment) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetIsFlagged

`func (o *PublicComment) GetIsFlagged() bool`

GetIsFlagged returns the IsFlagged field if non-nil, zero value otherwise.

### GetIsFlaggedOk

`func (o *PublicComment) GetIsFlaggedOk() (*bool, bool)`

GetIsFlaggedOk returns a tuple with the IsFlagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlagged

`func (o *PublicComment) SetIsFlagged(v bool)`

SetIsFlagged sets IsFlagged field to given value.

### HasIsFlagged

`func (o *PublicComment) HasIsFlagged() bool`

HasIsFlagged returns a boolean if a field has been set.

### GetIsBlocked

`func (o *PublicComment) GetIsBlocked() bool`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *PublicComment) GetIsBlockedOk() (*bool, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *PublicComment) SetIsBlocked(v bool)`

SetIsBlocked sets IsBlocked field to given value.

### HasIsBlocked

`func (o *PublicComment) HasIsBlocked() bool`

HasIsBlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


