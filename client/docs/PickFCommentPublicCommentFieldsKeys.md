# PickFCommentPublicCommentFieldsKeys

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

## Methods

### NewPickFCommentPublicCommentFieldsKeys

`func NewPickFCommentPublicCommentFieldsKeys(date time.Time, id string, commenterName string, commentHTML string, verified bool, ) *PickFCommentPublicCommentFieldsKeys`

NewPickFCommentPublicCommentFieldsKeys instantiates a new PickFCommentPublicCommentFieldsKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickFCommentPublicCommentFieldsKeysWithDefaults

`func NewPickFCommentPublicCommentFieldsKeysWithDefaults() *PickFCommentPublicCommentFieldsKeys`

NewPickFCommentPublicCommentFieldsKeysWithDefaults instantiates a new PickFCommentPublicCommentFieldsKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *PickFCommentPublicCommentFieldsKeys) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PickFCommentPublicCommentFieldsKeys) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetId

`func (o *PickFCommentPublicCommentFieldsKeys) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PickFCommentPublicCommentFieldsKeys) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *PickFCommentPublicCommentFieldsKeys) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PickFCommentPublicCommentFieldsKeys) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PickFCommentPublicCommentFieldsKeys) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAnonUserId

`func (o *PickFCommentPublicCommentFieldsKeys) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *PickFCommentPublicCommentFieldsKeys) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.

### HasAnonUserId

`func (o *PickFCommentPublicCommentFieldsKeys) HasAnonUserId() bool`

HasAnonUserId returns a boolean if a field has been set.

### GetCommenterName

`func (o *PickFCommentPublicCommentFieldsKeys) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *PickFCommentPublicCommentFieldsKeys) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetCommenterLink

`func (o *PickFCommentPublicCommentFieldsKeys) GetCommenterLink() string`

GetCommenterLink returns the CommenterLink field if non-nil, zero value otherwise.

### GetCommenterLinkOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetCommenterLinkOk() (*string, bool)`

GetCommenterLinkOk returns a tuple with the CommenterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterLink

`func (o *PickFCommentPublicCommentFieldsKeys) SetCommenterLink(v string)`

SetCommenterLink sets CommenterLink field to given value.

### HasCommenterLink

`func (o *PickFCommentPublicCommentFieldsKeys) HasCommenterLink() bool`

HasCommenterLink returns a boolean if a field has been set.

### GetCommentHTML

`func (o *PickFCommentPublicCommentFieldsKeys) GetCommentHTML() string`

GetCommentHTML returns the CommentHTML field if non-nil, zero value otherwise.

### GetCommentHTMLOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetCommentHTMLOk() (*string, bool)`

GetCommentHTMLOk returns a tuple with the CommentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentHTML

`func (o *PickFCommentPublicCommentFieldsKeys) SetCommentHTML(v string)`

SetCommentHTML sets CommentHTML field to given value.


### GetParentId

`func (o *PickFCommentPublicCommentFieldsKeys) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PickFCommentPublicCommentFieldsKeys) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PickFCommentPublicCommentFieldsKeys) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetVotes

`func (o *PickFCommentPublicCommentFieldsKeys) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *PickFCommentPublicCommentFieldsKeys) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *PickFCommentPublicCommentFieldsKeys) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetVotesUp

`func (o *PickFCommentPublicCommentFieldsKeys) GetVotesUp() int32`

GetVotesUp returns the VotesUp field if non-nil, zero value otherwise.

### GetVotesUpOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetVotesUpOk() (*int32, bool)`

GetVotesUpOk returns a tuple with the VotesUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUp

`func (o *PickFCommentPublicCommentFieldsKeys) SetVotesUp(v int32)`

SetVotesUp sets VotesUp field to given value.

### HasVotesUp

`func (o *PickFCommentPublicCommentFieldsKeys) HasVotesUp() bool`

HasVotesUp returns a boolean if a field has been set.

### GetVotesDown

`func (o *PickFCommentPublicCommentFieldsKeys) GetVotesDown() int32`

GetVotesDown returns the VotesDown field if non-nil, zero value otherwise.

### GetVotesDownOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetVotesDownOk() (*int32, bool)`

GetVotesDownOk returns a tuple with the VotesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDown

`func (o *PickFCommentPublicCommentFieldsKeys) SetVotesDown(v int32)`

SetVotesDown sets VotesDown field to given value.

### HasVotesDown

`func (o *PickFCommentPublicCommentFieldsKeys) HasVotesDown() bool`

HasVotesDown returns a boolean if a field has been set.

### GetVerified

`func (o *PickFCommentPublicCommentFieldsKeys) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PickFCommentPublicCommentFieldsKeys) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetAvatarSrc

`func (o *PickFCommentPublicCommentFieldsKeys) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *PickFCommentPublicCommentFieldsKeys) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *PickFCommentPublicCommentFieldsKeys) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### GetIsSpam

`func (o *PickFCommentPublicCommentFieldsKeys) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *PickFCommentPublicCommentFieldsKeys) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *PickFCommentPublicCommentFieldsKeys) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetHasImages

`func (o *PickFCommentPublicCommentFieldsKeys) GetHasImages() bool`

GetHasImages returns the HasImages field if non-nil, zero value otherwise.

### GetHasImagesOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetHasImagesOk() (*bool, bool)`

GetHasImagesOk returns a tuple with the HasImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImages

`func (o *PickFCommentPublicCommentFieldsKeys) SetHasImages(v bool)`

SetHasImages sets HasImages field to given value.

### HasHasImages

`func (o *PickFCommentPublicCommentFieldsKeys) HasHasImages() bool`

HasHasImages returns a boolean if a field has been set.

### GetIsDeleted

`func (o *PickFCommentPublicCommentFieldsKeys) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PickFCommentPublicCommentFieldsKeys) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PickFCommentPublicCommentFieldsKeys) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsDeletedUser

`func (o *PickFCommentPublicCommentFieldsKeys) GetIsDeletedUser() bool`

GetIsDeletedUser returns the IsDeletedUser field if non-nil, zero value otherwise.

### GetIsDeletedUserOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetIsDeletedUserOk() (*bool, bool)`

GetIsDeletedUserOk returns a tuple with the IsDeletedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletedUser

`func (o *PickFCommentPublicCommentFieldsKeys) SetIsDeletedUser(v bool)`

SetIsDeletedUser sets IsDeletedUser field to given value.

### HasIsDeletedUser

`func (o *PickFCommentPublicCommentFieldsKeys) HasIsDeletedUser() bool`

HasIsDeletedUser returns a boolean if a field has been set.

### GetIsByAdmin

`func (o *PickFCommentPublicCommentFieldsKeys) GetIsByAdmin() bool`

GetIsByAdmin returns the IsByAdmin field if non-nil, zero value otherwise.

### GetIsByAdminOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetIsByAdminOk() (*bool, bool)`

GetIsByAdminOk returns a tuple with the IsByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByAdmin

`func (o *PickFCommentPublicCommentFieldsKeys) SetIsByAdmin(v bool)`

SetIsByAdmin sets IsByAdmin field to given value.

### HasIsByAdmin

`func (o *PickFCommentPublicCommentFieldsKeys) HasIsByAdmin() bool`

HasIsByAdmin returns a boolean if a field has been set.

### GetIsByModerator

`func (o *PickFCommentPublicCommentFieldsKeys) GetIsByModerator() bool`

GetIsByModerator returns the IsByModerator field if non-nil, zero value otherwise.

### GetIsByModeratorOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetIsByModeratorOk() (*bool, bool)`

GetIsByModeratorOk returns a tuple with the IsByModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByModerator

`func (o *PickFCommentPublicCommentFieldsKeys) SetIsByModerator(v bool)`

SetIsByModerator sets IsByModerator field to given value.

### HasIsByModerator

`func (o *PickFCommentPublicCommentFieldsKeys) HasIsByModerator() bool`

HasIsByModerator returns a boolean if a field has been set.

### GetIsPinned

`func (o *PickFCommentPublicCommentFieldsKeys) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *PickFCommentPublicCommentFieldsKeys) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.

### HasIsPinned

`func (o *PickFCommentPublicCommentFieldsKeys) HasIsPinned() bool`

HasIsPinned returns a boolean if a field has been set.

### GetIsLocked

`func (o *PickFCommentPublicCommentFieldsKeys) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *PickFCommentPublicCommentFieldsKeys) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *PickFCommentPublicCommentFieldsKeys) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetRating

`func (o *PickFCommentPublicCommentFieldsKeys) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *PickFCommentPublicCommentFieldsKeys) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *PickFCommentPublicCommentFieldsKeys) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *PickFCommentPublicCommentFieldsKeys) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *PickFCommentPublicCommentFieldsKeys) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *PickFCommentPublicCommentFieldsKeys) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetBadges

`func (o *PickFCommentPublicCommentFieldsKeys) GetBadges() []CommentUserBadgeInfo`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetBadgesOk() (*[]CommentUserBadgeInfo, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *PickFCommentPublicCommentFieldsKeys) SetBadges(v []CommentUserBadgeInfo)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *PickFCommentPublicCommentFieldsKeys) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetFeedbackIds

`func (o *PickFCommentPublicCommentFieldsKeys) GetFeedbackIds() []string`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetFeedbackIdsOk() (*[]string, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *PickFCommentPublicCommentFieldsKeys) SetFeedbackIds(v []string)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *PickFCommentPublicCommentFieldsKeys) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.

### GetViewCount

`func (o *PickFCommentPublicCommentFieldsKeys) GetViewCount() int64`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetViewCountOk() (*int64, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *PickFCommentPublicCommentFieldsKeys) SetViewCount(v int64)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *PickFCommentPublicCommentFieldsKeys) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.

### GetRequiresVerification

`func (o *PickFCommentPublicCommentFieldsKeys) GetRequiresVerification() bool`

GetRequiresVerification returns the RequiresVerification field if non-nil, zero value otherwise.

### GetRequiresVerificationOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetRequiresVerificationOk() (*bool, bool)`

GetRequiresVerificationOk returns a tuple with the RequiresVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresVerification

`func (o *PickFCommentPublicCommentFieldsKeys) SetRequiresVerification(v bool)`

SetRequiresVerification sets RequiresVerification field to given value.

### HasRequiresVerification

`func (o *PickFCommentPublicCommentFieldsKeys) HasRequiresVerification() bool`

HasRequiresVerification returns a boolean if a field has been set.

### GetEditKey

`func (o *PickFCommentPublicCommentFieldsKeys) GetEditKey() string`

GetEditKey returns the EditKey field if non-nil, zero value otherwise.

### GetEditKeyOk

`func (o *PickFCommentPublicCommentFieldsKeys) GetEditKeyOk() (*string, bool)`

GetEditKeyOk returns a tuple with the EditKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditKey

`func (o *PickFCommentPublicCommentFieldsKeys) SetEditKey(v string)`

SetEditKey sets EditKey field to given value.

### HasEditKey

`func (o *PickFCommentPublicCommentFieldsKeys) HasEditKey() bool`

HasEditKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


