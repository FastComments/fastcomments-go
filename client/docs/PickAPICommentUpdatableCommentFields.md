# PickAPICommentUpdatableCommentFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** |  | 
**UrlId** | **string** |  | 
**UrlIdRaw** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**PageTitle** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**CommenterEmail** | Pointer to **string** |  | [optional] 
**CommenterName** | **string** |  | 
**CommenterLink** | Pointer to **string** |  | [optional] 
**Comment** | **string** |  | 
**CommentHTML** | **string** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**LocalDateString** | Pointer to **string** |  | [optional] 
**LocalDateHours** | Pointer to **int32** |  | [optional] 
**Votes** | Pointer to **int32** |  | [optional] 
**VotesUp** | Pointer to **int32** |  | [optional] 
**VotesDown** | Pointer to **int32** |  | [optional] 
**ExpireAt** | Pointer to **time.Time** |  | [optional] 
**Verified** | **bool** |  | 
**VerifiedDate** | Pointer to **time.Time** |  | [optional] 
**NotificationSentForParent** | Pointer to **bool** |  | [optional] 
**NotificationSentForParentTenant** | Pointer to **bool** |  | [optional] 
**Reviewed** | Pointer to **bool** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalParentId** | Pointer to **string** |  | [optional] 
**AvatarSrc** | Pointer to **string** |  | [optional] 
**IsSpam** | Pointer to **bool** |  | [optional] 
**Approved** | **bool** |  | 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**IsDeletedUser** | Pointer to **bool** |  | [optional] 
**IsByAdmin** | Pointer to **bool** |  | [optional] 
**IsByModerator** | Pointer to **bool** |  | [optional] 
**IsPinned** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**FlagCount** | Pointer to **int32** |  | [optional] 
**DisplayLabel** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**PickFCommentAPICommentFieldsKeysMeta**](PickFCommentAPICommentFieldsKeysMeta.md) |  | [optional] 
**ModerationGroupIds** | Pointer to **[]string** |  | [optional] 
**FeedbackIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPickAPICommentUpdatableCommentFields

`func NewPickAPICommentUpdatableCommentFields(date time.Time, urlId string, url string, commenterName string, comment string, commentHTML string, verified bool, approved bool, ) *PickAPICommentUpdatableCommentFields`

NewPickAPICommentUpdatableCommentFields instantiates a new PickAPICommentUpdatableCommentFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickAPICommentUpdatableCommentFieldsWithDefaults

`func NewPickAPICommentUpdatableCommentFieldsWithDefaults() *PickAPICommentUpdatableCommentFields`

NewPickAPICommentUpdatableCommentFieldsWithDefaults instantiates a new PickAPICommentUpdatableCommentFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *PickAPICommentUpdatableCommentFields) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PickAPICommentUpdatableCommentFields) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PickAPICommentUpdatableCommentFields) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetUrlId

`func (o *PickAPICommentUpdatableCommentFields) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *PickAPICommentUpdatableCommentFields) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *PickAPICommentUpdatableCommentFields) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetUrlIdRaw

`func (o *PickAPICommentUpdatableCommentFields) GetUrlIdRaw() string`

GetUrlIdRaw returns the UrlIdRaw field if non-nil, zero value otherwise.

### GetUrlIdRawOk

`func (o *PickAPICommentUpdatableCommentFields) GetUrlIdRawOk() (*string, bool)`

GetUrlIdRawOk returns a tuple with the UrlIdRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdRaw

`func (o *PickAPICommentUpdatableCommentFields) SetUrlIdRaw(v string)`

SetUrlIdRaw sets UrlIdRaw field to given value.

### HasUrlIdRaw

`func (o *PickAPICommentUpdatableCommentFields) HasUrlIdRaw() bool`

HasUrlIdRaw returns a boolean if a field has been set.

### GetUrl

`func (o *PickAPICommentUpdatableCommentFields) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PickAPICommentUpdatableCommentFields) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PickAPICommentUpdatableCommentFields) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPageTitle

`func (o *PickAPICommentUpdatableCommentFields) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *PickAPICommentUpdatableCommentFields) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *PickAPICommentUpdatableCommentFields) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *PickAPICommentUpdatableCommentFields) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### GetUserId

`func (o *PickAPICommentUpdatableCommentFields) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PickAPICommentUpdatableCommentFields) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PickAPICommentUpdatableCommentFields) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PickAPICommentUpdatableCommentFields) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCommenterEmail

`func (o *PickAPICommentUpdatableCommentFields) GetCommenterEmail() string`

GetCommenterEmail returns the CommenterEmail field if non-nil, zero value otherwise.

### GetCommenterEmailOk

`func (o *PickAPICommentUpdatableCommentFields) GetCommenterEmailOk() (*string, bool)`

GetCommenterEmailOk returns a tuple with the CommenterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterEmail

`func (o *PickAPICommentUpdatableCommentFields) SetCommenterEmail(v string)`

SetCommenterEmail sets CommenterEmail field to given value.

### HasCommenterEmail

`func (o *PickAPICommentUpdatableCommentFields) HasCommenterEmail() bool`

HasCommenterEmail returns a boolean if a field has been set.

### GetCommenterName

`func (o *PickAPICommentUpdatableCommentFields) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *PickAPICommentUpdatableCommentFields) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *PickAPICommentUpdatableCommentFields) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetCommenterLink

`func (o *PickAPICommentUpdatableCommentFields) GetCommenterLink() string`

GetCommenterLink returns the CommenterLink field if non-nil, zero value otherwise.

### GetCommenterLinkOk

`func (o *PickAPICommentUpdatableCommentFields) GetCommenterLinkOk() (*string, bool)`

GetCommenterLinkOk returns a tuple with the CommenterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterLink

`func (o *PickAPICommentUpdatableCommentFields) SetCommenterLink(v string)`

SetCommenterLink sets CommenterLink field to given value.

### HasCommenterLink

`func (o *PickAPICommentUpdatableCommentFields) HasCommenterLink() bool`

HasCommenterLink returns a boolean if a field has been set.

### GetComment

`func (o *PickAPICommentUpdatableCommentFields) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PickAPICommentUpdatableCommentFields) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PickAPICommentUpdatableCommentFields) SetComment(v string)`

SetComment sets Comment field to given value.


### GetCommentHTML

`func (o *PickAPICommentUpdatableCommentFields) GetCommentHTML() string`

GetCommentHTML returns the CommentHTML field if non-nil, zero value otherwise.

### GetCommentHTMLOk

`func (o *PickAPICommentUpdatableCommentFields) GetCommentHTMLOk() (*string, bool)`

GetCommentHTMLOk returns a tuple with the CommentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentHTML

`func (o *PickAPICommentUpdatableCommentFields) SetCommentHTML(v string)`

SetCommentHTML sets CommentHTML field to given value.


### GetParentId

`func (o *PickAPICommentUpdatableCommentFields) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PickAPICommentUpdatableCommentFields) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PickAPICommentUpdatableCommentFields) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PickAPICommentUpdatableCommentFields) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetLocalDateString

`func (o *PickAPICommentUpdatableCommentFields) GetLocalDateString() string`

GetLocalDateString returns the LocalDateString field if non-nil, zero value otherwise.

### GetLocalDateStringOk

`func (o *PickAPICommentUpdatableCommentFields) GetLocalDateStringOk() (*string, bool)`

GetLocalDateStringOk returns a tuple with the LocalDateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateString

`func (o *PickAPICommentUpdatableCommentFields) SetLocalDateString(v string)`

SetLocalDateString sets LocalDateString field to given value.

### HasLocalDateString

`func (o *PickAPICommentUpdatableCommentFields) HasLocalDateString() bool`

HasLocalDateString returns a boolean if a field has been set.

### GetLocalDateHours

`func (o *PickAPICommentUpdatableCommentFields) GetLocalDateHours() int32`

GetLocalDateHours returns the LocalDateHours field if non-nil, zero value otherwise.

### GetLocalDateHoursOk

`func (o *PickAPICommentUpdatableCommentFields) GetLocalDateHoursOk() (*int32, bool)`

GetLocalDateHoursOk returns a tuple with the LocalDateHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateHours

`func (o *PickAPICommentUpdatableCommentFields) SetLocalDateHours(v int32)`

SetLocalDateHours sets LocalDateHours field to given value.

### HasLocalDateHours

`func (o *PickAPICommentUpdatableCommentFields) HasLocalDateHours() bool`

HasLocalDateHours returns a boolean if a field has been set.

### GetVotes

`func (o *PickAPICommentUpdatableCommentFields) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *PickAPICommentUpdatableCommentFields) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *PickAPICommentUpdatableCommentFields) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *PickAPICommentUpdatableCommentFields) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetVotesUp

`func (o *PickAPICommentUpdatableCommentFields) GetVotesUp() int32`

GetVotesUp returns the VotesUp field if non-nil, zero value otherwise.

### GetVotesUpOk

`func (o *PickAPICommentUpdatableCommentFields) GetVotesUpOk() (*int32, bool)`

GetVotesUpOk returns a tuple with the VotesUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUp

`func (o *PickAPICommentUpdatableCommentFields) SetVotesUp(v int32)`

SetVotesUp sets VotesUp field to given value.

### HasVotesUp

`func (o *PickAPICommentUpdatableCommentFields) HasVotesUp() bool`

HasVotesUp returns a boolean if a field has been set.

### GetVotesDown

`func (o *PickAPICommentUpdatableCommentFields) GetVotesDown() int32`

GetVotesDown returns the VotesDown field if non-nil, zero value otherwise.

### GetVotesDownOk

`func (o *PickAPICommentUpdatableCommentFields) GetVotesDownOk() (*int32, bool)`

GetVotesDownOk returns a tuple with the VotesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDown

`func (o *PickAPICommentUpdatableCommentFields) SetVotesDown(v int32)`

SetVotesDown sets VotesDown field to given value.

### HasVotesDown

`func (o *PickAPICommentUpdatableCommentFields) HasVotesDown() bool`

HasVotesDown returns a boolean if a field has been set.

### GetExpireAt

`func (o *PickAPICommentUpdatableCommentFields) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *PickAPICommentUpdatableCommentFields) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *PickAPICommentUpdatableCommentFields) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *PickAPICommentUpdatableCommentFields) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetVerified

`func (o *PickAPICommentUpdatableCommentFields) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PickAPICommentUpdatableCommentFields) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PickAPICommentUpdatableCommentFields) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetVerifiedDate

`func (o *PickAPICommentUpdatableCommentFields) GetVerifiedDate() time.Time`

GetVerifiedDate returns the VerifiedDate field if non-nil, zero value otherwise.

### GetVerifiedDateOk

`func (o *PickAPICommentUpdatableCommentFields) GetVerifiedDateOk() (*time.Time, bool)`

GetVerifiedDateOk returns a tuple with the VerifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedDate

`func (o *PickAPICommentUpdatableCommentFields) SetVerifiedDate(v time.Time)`

SetVerifiedDate sets VerifiedDate field to given value.

### HasVerifiedDate

`func (o *PickAPICommentUpdatableCommentFields) HasVerifiedDate() bool`

HasVerifiedDate returns a boolean if a field has been set.

### GetNotificationSentForParent

`func (o *PickAPICommentUpdatableCommentFields) GetNotificationSentForParent() bool`

GetNotificationSentForParent returns the NotificationSentForParent field if non-nil, zero value otherwise.

### GetNotificationSentForParentOk

`func (o *PickAPICommentUpdatableCommentFields) GetNotificationSentForParentOk() (*bool, bool)`

GetNotificationSentForParentOk returns a tuple with the NotificationSentForParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSentForParent

`func (o *PickAPICommentUpdatableCommentFields) SetNotificationSentForParent(v bool)`

SetNotificationSentForParent sets NotificationSentForParent field to given value.

### HasNotificationSentForParent

`func (o *PickAPICommentUpdatableCommentFields) HasNotificationSentForParent() bool`

HasNotificationSentForParent returns a boolean if a field has been set.

### GetNotificationSentForParentTenant

`func (o *PickAPICommentUpdatableCommentFields) GetNotificationSentForParentTenant() bool`

GetNotificationSentForParentTenant returns the NotificationSentForParentTenant field if non-nil, zero value otherwise.

### GetNotificationSentForParentTenantOk

`func (o *PickAPICommentUpdatableCommentFields) GetNotificationSentForParentTenantOk() (*bool, bool)`

GetNotificationSentForParentTenantOk returns a tuple with the NotificationSentForParentTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSentForParentTenant

`func (o *PickAPICommentUpdatableCommentFields) SetNotificationSentForParentTenant(v bool)`

SetNotificationSentForParentTenant sets NotificationSentForParentTenant field to given value.

### HasNotificationSentForParentTenant

`func (o *PickAPICommentUpdatableCommentFields) HasNotificationSentForParentTenant() bool`

HasNotificationSentForParentTenant returns a boolean if a field has been set.

### GetReviewed

`func (o *PickAPICommentUpdatableCommentFields) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *PickAPICommentUpdatableCommentFields) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *PickAPICommentUpdatableCommentFields) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *PickAPICommentUpdatableCommentFields) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetExternalId

`func (o *PickAPICommentUpdatableCommentFields) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PickAPICommentUpdatableCommentFields) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PickAPICommentUpdatableCommentFields) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PickAPICommentUpdatableCommentFields) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalParentId

`func (o *PickAPICommentUpdatableCommentFields) GetExternalParentId() string`

GetExternalParentId returns the ExternalParentId field if non-nil, zero value otherwise.

### GetExternalParentIdOk

`func (o *PickAPICommentUpdatableCommentFields) GetExternalParentIdOk() (*string, bool)`

GetExternalParentIdOk returns a tuple with the ExternalParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalParentId

`func (o *PickAPICommentUpdatableCommentFields) SetExternalParentId(v string)`

SetExternalParentId sets ExternalParentId field to given value.

### HasExternalParentId

`func (o *PickAPICommentUpdatableCommentFields) HasExternalParentId() bool`

HasExternalParentId returns a boolean if a field has been set.

### GetAvatarSrc

`func (o *PickAPICommentUpdatableCommentFields) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *PickAPICommentUpdatableCommentFields) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *PickAPICommentUpdatableCommentFields) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *PickAPICommentUpdatableCommentFields) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### GetIsSpam

`func (o *PickAPICommentUpdatableCommentFields) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *PickAPICommentUpdatableCommentFields) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *PickAPICommentUpdatableCommentFields) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *PickAPICommentUpdatableCommentFields) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetApproved

`func (o *PickAPICommentUpdatableCommentFields) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *PickAPICommentUpdatableCommentFields) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *PickAPICommentUpdatableCommentFields) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetIsDeleted

`func (o *PickAPICommentUpdatableCommentFields) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PickAPICommentUpdatableCommentFields) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PickAPICommentUpdatableCommentFields) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PickAPICommentUpdatableCommentFields) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsDeletedUser

`func (o *PickAPICommentUpdatableCommentFields) GetIsDeletedUser() bool`

GetIsDeletedUser returns the IsDeletedUser field if non-nil, zero value otherwise.

### GetIsDeletedUserOk

`func (o *PickAPICommentUpdatableCommentFields) GetIsDeletedUserOk() (*bool, bool)`

GetIsDeletedUserOk returns a tuple with the IsDeletedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletedUser

`func (o *PickAPICommentUpdatableCommentFields) SetIsDeletedUser(v bool)`

SetIsDeletedUser sets IsDeletedUser field to given value.

### HasIsDeletedUser

`func (o *PickAPICommentUpdatableCommentFields) HasIsDeletedUser() bool`

HasIsDeletedUser returns a boolean if a field has been set.

### GetIsByAdmin

`func (o *PickAPICommentUpdatableCommentFields) GetIsByAdmin() bool`

GetIsByAdmin returns the IsByAdmin field if non-nil, zero value otherwise.

### GetIsByAdminOk

`func (o *PickAPICommentUpdatableCommentFields) GetIsByAdminOk() (*bool, bool)`

GetIsByAdminOk returns a tuple with the IsByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByAdmin

`func (o *PickAPICommentUpdatableCommentFields) SetIsByAdmin(v bool)`

SetIsByAdmin sets IsByAdmin field to given value.

### HasIsByAdmin

`func (o *PickAPICommentUpdatableCommentFields) HasIsByAdmin() bool`

HasIsByAdmin returns a boolean if a field has been set.

### GetIsByModerator

`func (o *PickAPICommentUpdatableCommentFields) GetIsByModerator() bool`

GetIsByModerator returns the IsByModerator field if non-nil, zero value otherwise.

### GetIsByModeratorOk

`func (o *PickAPICommentUpdatableCommentFields) GetIsByModeratorOk() (*bool, bool)`

GetIsByModeratorOk returns a tuple with the IsByModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByModerator

`func (o *PickAPICommentUpdatableCommentFields) SetIsByModerator(v bool)`

SetIsByModerator sets IsByModerator field to given value.

### HasIsByModerator

`func (o *PickAPICommentUpdatableCommentFields) HasIsByModerator() bool`

HasIsByModerator returns a boolean if a field has been set.

### GetIsPinned

`func (o *PickAPICommentUpdatableCommentFields) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *PickAPICommentUpdatableCommentFields) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *PickAPICommentUpdatableCommentFields) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.

### HasIsPinned

`func (o *PickAPICommentUpdatableCommentFields) HasIsPinned() bool`

HasIsPinned returns a boolean if a field has been set.

### GetIsLocked

`func (o *PickAPICommentUpdatableCommentFields) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *PickAPICommentUpdatableCommentFields) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *PickAPICommentUpdatableCommentFields) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *PickAPICommentUpdatableCommentFields) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetFlagCount

`func (o *PickAPICommentUpdatableCommentFields) GetFlagCount() int32`

GetFlagCount returns the FlagCount field if non-nil, zero value otherwise.

### GetFlagCountOk

`func (o *PickAPICommentUpdatableCommentFields) GetFlagCountOk() (*int32, bool)`

GetFlagCountOk returns a tuple with the FlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCount

`func (o *PickAPICommentUpdatableCommentFields) SetFlagCount(v int32)`

SetFlagCount sets FlagCount field to given value.

### HasFlagCount

`func (o *PickAPICommentUpdatableCommentFields) HasFlagCount() bool`

HasFlagCount returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *PickAPICommentUpdatableCommentFields) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *PickAPICommentUpdatableCommentFields) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *PickAPICommentUpdatableCommentFields) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *PickAPICommentUpdatableCommentFields) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetMeta

`func (o *PickAPICommentUpdatableCommentFields) GetMeta() PickFCommentAPICommentFieldsKeysMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PickAPICommentUpdatableCommentFields) GetMetaOk() (*PickFCommentAPICommentFieldsKeysMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PickAPICommentUpdatableCommentFields) SetMeta(v PickFCommentAPICommentFieldsKeysMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PickAPICommentUpdatableCommentFields) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetModerationGroupIds

`func (o *PickAPICommentUpdatableCommentFields) GetModerationGroupIds() []string`

GetModerationGroupIds returns the ModerationGroupIds field if non-nil, zero value otherwise.

### GetModerationGroupIdsOk

`func (o *PickAPICommentUpdatableCommentFields) GetModerationGroupIdsOk() (*[]string, bool)`

GetModerationGroupIdsOk returns a tuple with the ModerationGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationGroupIds

`func (o *PickAPICommentUpdatableCommentFields) SetModerationGroupIds(v []string)`

SetModerationGroupIds sets ModerationGroupIds field to given value.

### HasModerationGroupIds

`func (o *PickAPICommentUpdatableCommentFields) HasModerationGroupIds() bool`

HasModerationGroupIds returns a boolean if a field has been set.

### GetFeedbackIds

`func (o *PickAPICommentUpdatableCommentFields) GetFeedbackIds() []string`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *PickAPICommentUpdatableCommentFields) GetFeedbackIdsOk() (*[]string, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *PickAPICommentUpdatableCommentFields) SetFeedbackIds(v []string)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *PickAPICommentUpdatableCommentFields) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


