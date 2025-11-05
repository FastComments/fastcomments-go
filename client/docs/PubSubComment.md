# PubSubComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**UrlId** | **string** |  | 
**Url** | **string** |  | 
**PageTitle** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**AnonUserId** | Pointer to **string** |  | [optional] 
**CommenterName** | **string** |  | 
**CommenterLink** | Pointer to **string** |  | [optional] 
**Comment** | **string** |  | 
**CommentHTML** | **string** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**Votes** | Pointer to **int32** |  | [optional] 
**VotesUp** | Pointer to **int32** |  | [optional] 
**VotesDown** | Pointer to **int32** |  | [optional] 
**ExpireAt** | Pointer to **time.Time** |  | [optional] 
**Verified** | **bool** |  | 
**Reviewed** | Pointer to **bool** |  | [optional] 
**AvatarSrc** | Pointer to **string** |  | [optional] 
**IsSpam** | Pointer to **bool** |  | [optional] 
**HasImages** | Pointer to **bool** |  | [optional] 
**HasLinks** | Pointer to **bool** |  | [optional] 
**HasCode** | Pointer to **bool** |  | [optional] 
**Approved** | **bool** |  | 
**Locale** | **string** |  | 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**IsDeletedUser** | Pointer to **bool** |  | [optional] 
**IsBannedUser** | Pointer to **bool** |  | [optional] 
**IsByAdmin** | Pointer to **bool** |  | [optional] 
**IsByModerator** | Pointer to **bool** |  | [optional] 
**IsPinned** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**FlagCount** | Pointer to **int32** |  | [optional] 
**Rating** | Pointer to **float64** |  | [optional] 
**DisplayLabel** | Pointer to **string** |  | [optional] 
**Badges** | Pointer to [**[]CommentUserBadgeInfo**](CommentUserBadgeInfo.md) |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**FeedbackIds** | Pointer to **[]string** |  | [optional] 
**GroupIds** | Pointer to **[]string** |  | [optional] 
**ViewCount** | Pointer to **int64** |  | [optional] 
**IsLive** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 

## Methods

### NewPubSubComment

`func NewPubSubComment(date string, id string, tenantId string, urlId string, url string, commenterName string, comment string, commentHTML string, verified bool, approved bool, locale string, ) *PubSubComment`

NewPubSubComment instantiates a new PubSubComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubSubCommentWithDefaults

`func NewPubSubCommentWithDefaults() *PubSubComment`

NewPubSubCommentWithDefaults instantiates a new PubSubComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *PubSubComment) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PubSubComment) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PubSubComment) SetDate(v string)`

SetDate sets Date field to given value.


### GetId

`func (o *PubSubComment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PubSubComment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PubSubComment) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *PubSubComment) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PubSubComment) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PubSubComment) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUrlId

`func (o *PubSubComment) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *PubSubComment) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *PubSubComment) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetUrl

`func (o *PubSubComment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PubSubComment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PubSubComment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPageTitle

`func (o *PubSubComment) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *PubSubComment) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *PubSubComment) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *PubSubComment) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### GetUserId

`func (o *PubSubComment) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PubSubComment) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PubSubComment) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PubSubComment) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAnonUserId

`func (o *PubSubComment) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *PubSubComment) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *PubSubComment) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.

### HasAnonUserId

`func (o *PubSubComment) HasAnonUserId() bool`

HasAnonUserId returns a boolean if a field has been set.

### GetCommenterName

`func (o *PubSubComment) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *PubSubComment) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *PubSubComment) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetCommenterLink

`func (o *PubSubComment) GetCommenterLink() string`

GetCommenterLink returns the CommenterLink field if non-nil, zero value otherwise.

### GetCommenterLinkOk

`func (o *PubSubComment) GetCommenterLinkOk() (*string, bool)`

GetCommenterLinkOk returns a tuple with the CommenterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterLink

`func (o *PubSubComment) SetCommenterLink(v string)`

SetCommenterLink sets CommenterLink field to given value.

### HasCommenterLink

`func (o *PubSubComment) HasCommenterLink() bool`

HasCommenterLink returns a boolean if a field has been set.

### GetComment

`func (o *PubSubComment) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PubSubComment) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PubSubComment) SetComment(v string)`

SetComment sets Comment field to given value.


### GetCommentHTML

`func (o *PubSubComment) GetCommentHTML() string`

GetCommentHTML returns the CommentHTML field if non-nil, zero value otherwise.

### GetCommentHTMLOk

`func (o *PubSubComment) GetCommentHTMLOk() (*string, bool)`

GetCommentHTMLOk returns a tuple with the CommentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentHTML

`func (o *PubSubComment) SetCommentHTML(v string)`

SetCommentHTML sets CommentHTML field to given value.


### GetParentId

`func (o *PubSubComment) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PubSubComment) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PubSubComment) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PubSubComment) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetVotes

`func (o *PubSubComment) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *PubSubComment) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *PubSubComment) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *PubSubComment) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetVotesUp

`func (o *PubSubComment) GetVotesUp() int32`

GetVotesUp returns the VotesUp field if non-nil, zero value otherwise.

### GetVotesUpOk

`func (o *PubSubComment) GetVotesUpOk() (*int32, bool)`

GetVotesUpOk returns a tuple with the VotesUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUp

`func (o *PubSubComment) SetVotesUp(v int32)`

SetVotesUp sets VotesUp field to given value.

### HasVotesUp

`func (o *PubSubComment) HasVotesUp() bool`

HasVotesUp returns a boolean if a field has been set.

### GetVotesDown

`func (o *PubSubComment) GetVotesDown() int32`

GetVotesDown returns the VotesDown field if non-nil, zero value otherwise.

### GetVotesDownOk

`func (o *PubSubComment) GetVotesDownOk() (*int32, bool)`

GetVotesDownOk returns a tuple with the VotesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDown

`func (o *PubSubComment) SetVotesDown(v int32)`

SetVotesDown sets VotesDown field to given value.

### HasVotesDown

`func (o *PubSubComment) HasVotesDown() bool`

HasVotesDown returns a boolean if a field has been set.

### GetExpireAt

`func (o *PubSubComment) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *PubSubComment) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *PubSubComment) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *PubSubComment) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetVerified

`func (o *PubSubComment) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PubSubComment) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PubSubComment) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetReviewed

`func (o *PubSubComment) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *PubSubComment) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *PubSubComment) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *PubSubComment) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetAvatarSrc

`func (o *PubSubComment) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *PubSubComment) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *PubSubComment) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *PubSubComment) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### GetIsSpam

`func (o *PubSubComment) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *PubSubComment) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *PubSubComment) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *PubSubComment) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetHasImages

`func (o *PubSubComment) GetHasImages() bool`

GetHasImages returns the HasImages field if non-nil, zero value otherwise.

### GetHasImagesOk

`func (o *PubSubComment) GetHasImagesOk() (*bool, bool)`

GetHasImagesOk returns a tuple with the HasImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImages

`func (o *PubSubComment) SetHasImages(v bool)`

SetHasImages sets HasImages field to given value.

### HasHasImages

`func (o *PubSubComment) HasHasImages() bool`

HasHasImages returns a boolean if a field has been set.

### GetHasLinks

`func (o *PubSubComment) GetHasLinks() bool`

GetHasLinks returns the HasLinks field if non-nil, zero value otherwise.

### GetHasLinksOk

`func (o *PubSubComment) GetHasLinksOk() (*bool, bool)`

GetHasLinksOk returns a tuple with the HasLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLinks

`func (o *PubSubComment) SetHasLinks(v bool)`

SetHasLinks sets HasLinks field to given value.

### HasHasLinks

`func (o *PubSubComment) HasHasLinks() bool`

HasHasLinks returns a boolean if a field has been set.

### GetHasCode

`func (o *PubSubComment) GetHasCode() bool`

GetHasCode returns the HasCode field if non-nil, zero value otherwise.

### GetHasCodeOk

`func (o *PubSubComment) GetHasCodeOk() (*bool, bool)`

GetHasCodeOk returns a tuple with the HasCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCode

`func (o *PubSubComment) SetHasCode(v bool)`

SetHasCode sets HasCode field to given value.

### HasHasCode

`func (o *PubSubComment) HasHasCode() bool`

HasHasCode returns a boolean if a field has been set.

### GetApproved

`func (o *PubSubComment) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *PubSubComment) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *PubSubComment) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetLocale

`func (o *PubSubComment) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PubSubComment) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PubSubComment) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetIsDeleted

`func (o *PubSubComment) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PubSubComment) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PubSubComment) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PubSubComment) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsDeletedUser

`func (o *PubSubComment) GetIsDeletedUser() bool`

GetIsDeletedUser returns the IsDeletedUser field if non-nil, zero value otherwise.

### GetIsDeletedUserOk

`func (o *PubSubComment) GetIsDeletedUserOk() (*bool, bool)`

GetIsDeletedUserOk returns a tuple with the IsDeletedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletedUser

`func (o *PubSubComment) SetIsDeletedUser(v bool)`

SetIsDeletedUser sets IsDeletedUser field to given value.

### HasIsDeletedUser

`func (o *PubSubComment) HasIsDeletedUser() bool`

HasIsDeletedUser returns a boolean if a field has been set.

### GetIsBannedUser

`func (o *PubSubComment) GetIsBannedUser() bool`

GetIsBannedUser returns the IsBannedUser field if non-nil, zero value otherwise.

### GetIsBannedUserOk

`func (o *PubSubComment) GetIsBannedUserOk() (*bool, bool)`

GetIsBannedUserOk returns a tuple with the IsBannedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBannedUser

`func (o *PubSubComment) SetIsBannedUser(v bool)`

SetIsBannedUser sets IsBannedUser field to given value.

### HasIsBannedUser

`func (o *PubSubComment) HasIsBannedUser() bool`

HasIsBannedUser returns a boolean if a field has been set.

### GetIsByAdmin

`func (o *PubSubComment) GetIsByAdmin() bool`

GetIsByAdmin returns the IsByAdmin field if non-nil, zero value otherwise.

### GetIsByAdminOk

`func (o *PubSubComment) GetIsByAdminOk() (*bool, bool)`

GetIsByAdminOk returns a tuple with the IsByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByAdmin

`func (o *PubSubComment) SetIsByAdmin(v bool)`

SetIsByAdmin sets IsByAdmin field to given value.

### HasIsByAdmin

`func (o *PubSubComment) HasIsByAdmin() bool`

HasIsByAdmin returns a boolean if a field has been set.

### GetIsByModerator

`func (o *PubSubComment) GetIsByModerator() bool`

GetIsByModerator returns the IsByModerator field if non-nil, zero value otherwise.

### GetIsByModeratorOk

`func (o *PubSubComment) GetIsByModeratorOk() (*bool, bool)`

GetIsByModeratorOk returns a tuple with the IsByModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByModerator

`func (o *PubSubComment) SetIsByModerator(v bool)`

SetIsByModerator sets IsByModerator field to given value.

### HasIsByModerator

`func (o *PubSubComment) HasIsByModerator() bool`

HasIsByModerator returns a boolean if a field has been set.

### GetIsPinned

`func (o *PubSubComment) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *PubSubComment) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *PubSubComment) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.

### HasIsPinned

`func (o *PubSubComment) HasIsPinned() bool`

HasIsPinned returns a boolean if a field has been set.

### GetIsLocked

`func (o *PubSubComment) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *PubSubComment) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *PubSubComment) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *PubSubComment) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetFlagCount

`func (o *PubSubComment) GetFlagCount() int32`

GetFlagCount returns the FlagCount field if non-nil, zero value otherwise.

### GetFlagCountOk

`func (o *PubSubComment) GetFlagCountOk() (*int32, bool)`

GetFlagCountOk returns a tuple with the FlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCount

`func (o *PubSubComment) SetFlagCount(v int32)`

SetFlagCount sets FlagCount field to given value.

### HasFlagCount

`func (o *PubSubComment) HasFlagCount() bool`

HasFlagCount returns a boolean if a field has been set.

### GetRating

`func (o *PubSubComment) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *PubSubComment) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *PubSubComment) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *PubSubComment) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *PubSubComment) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *PubSubComment) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *PubSubComment) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *PubSubComment) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetBadges

`func (o *PubSubComment) GetBadges() []CommentUserBadgeInfo`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *PubSubComment) GetBadgesOk() (*[]CommentUserBadgeInfo, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *PubSubComment) SetBadges(v []CommentUserBadgeInfo)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *PubSubComment) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetDomain

`func (o *PubSubComment) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PubSubComment) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PubSubComment) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PubSubComment) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetFeedbackIds

`func (o *PubSubComment) GetFeedbackIds() []string`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *PubSubComment) GetFeedbackIdsOk() (*[]string, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *PubSubComment) SetFeedbackIds(v []string)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *PubSubComment) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.

### GetGroupIds

`func (o *PubSubComment) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *PubSubComment) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *PubSubComment) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *PubSubComment) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetViewCount

`func (o *PubSubComment) GetViewCount() int64`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *PubSubComment) GetViewCountOk() (*int64, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *PubSubComment) SetViewCount(v int64)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *PubSubComment) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.

### GetIsLive

`func (o *PubSubComment) GetIsLive() bool`

GetIsLive returns the IsLive field if non-nil, zero value otherwise.

### GetIsLiveOk

`func (o *PubSubComment) GetIsLiveOk() (*bool, bool)`

GetIsLiveOk returns a tuple with the IsLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLive

`func (o *PubSubComment) SetIsLive(v bool)`

SetIsLive sets IsLive field to given value.

### HasIsLive

`func (o *PubSubComment) HasIsLive() bool`

HasIsLive returns a boolean if a field has been set.

### GetHidden

`func (o *PubSubComment) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *PubSubComment) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *PubSubComment) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *PubSubComment) HasHidden() bool`

HasHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


